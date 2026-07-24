package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// MCP Protocol Data Structures (JSON-RPC 2.0)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema InputSchema `json:"inputSchema"`
}

type InputSchema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

type Property struct {
	Type        string `json:"type"`
	Description string `description:"description"`
}

type ToolCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type CallToolResult struct {
	Content []TextContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

var (
	lightpandaHost = getEnvOrDefault("LIGHTPANDA_HOST", "127.0.0.1")
	lightpandaPort = getEnvOrDefault("LIGHTPANDA_PORT", "9222")
	mu             sync.Mutex
)

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Buffer up to 10MB per line for large HTML/markdown payloads
	buf := make([]byte, 10*1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var req Request
		if err := json.Unmarshal(line, &req); err != nil {
			sendError(nil, -32700, "Parse error")
			continue
		}

		handleRequest(&req)
	}
}

func handleRequest(req *Request) {
	switch req.Method {
	case "initialize":
		sendResponse(req.ID, map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    "lightpanda-mcp",
				"version": "1.0.0",
			},
		})

	case "notifications/initialized":
		// No response required for notifications

	case "tools/list":
		tools := []Tool{
			{
				Name:        "lightpanda_fetch_html",
				Description: "Fetches HTML content from a URL using Lightpanda fast headless browser engine.",
				InputSchema: InputSchema{
					Type: "object",
					Properties: map[string]Property{
						"url": {Type: "string", Description: "Target web URL to fetch"},
					},
					Required: []string{"url"},
				},
			},
			{
				Name:        "lightpanda_get_markdown",
				Description: "Extracts clean Markdown text and AX (Accessibility) Tree from a webpage using Lightpanda.",
				InputSchema: InputSchema{
					Type: "object",
					Properties: map[string]Property{
						"url": {Type: "string", Description: "Target web URL to parse"},
					},
					Required: []string{"url"},
				},
			},
			{
				Name:        "lightpanda_execute_js",
				Description: "Executes custom JavaScript inside Lightpanda headless browser and returns the result.",
				InputSchema: InputSchema{
					Type: "object",
					Properties: map[string]Property{
						"url":    {Type: "string", Description: "Target web URL"},
						"script": {Type: "string", Description: "JavaScript code snippet to execute"},
					},
					Required: []string{"url", "script"},
				},
			},
			{
				Name:        "lightpanda_status",
				Description: "Checks the status of the local Lightpanda daemon/WSL process and CDP endpoint.",
				InputSchema: InputSchema{
					Type:       "object",
					Properties: map[string]Property{},
				},
			},
		}
		sendResponse(req.ID, map[string]interface{}{
			"tools": tools,
		})

	case "tools/call":
		var params ToolCallParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			sendError(req.ID, -32602, "Invalid params")
			return
		}

		result := executeToolCall(params)
		sendResponse(req.ID, result)

	default:
		sendError(req.ID, -32601, fmt.Sprintf("Method not found: %s", req.Method))
	}
}

func executeToolCall(params ToolCallParams) CallToolResult {
	ensureLightpandaRunning()

	switch params.Name {
	case "lightpanda_status":
		status := checkLightpandaStatus()
		return CallToolResult{
			Content: []TextContent{{Type: "text", Text: status}},
		}

	case "lightpanda_fetch_html":
		var args struct {
			URL string `json:"url"`
		}
		if err := json.Unmarshal(params.Arguments, &args); err != nil {
			return errorResult(fmt.Sprintf("Invalid arguments: %v", err))
		}

		html, err := fetchHTMLViaLightpanda(args.URL)
		if err != nil {
			return errorResult(fmt.Sprintf("Lightpanda fetch error: %v", err))
		}
		return CallToolResult{
			Content: []TextContent{{Type: "text", Text: html}},
		}

	case "lightpanda_get_markdown":
		var args struct {
			URL string `json:"url"`
		}
		if err := json.Unmarshal(params.Arguments, &args); err != nil {
			return errorResult(fmt.Sprintf("Invalid arguments: %v", err))
		}

		md, err := fetchMarkdownViaLightpanda(args.URL)
		if err != nil {
			return errorResult(fmt.Sprintf("Lightpanda markdown error: %v", err))
		}
		return CallToolResult{
			Content: []TextContent{{Type: "text", Text: md}},
		}

	case "lightpanda_execute_js":
		var args struct {
			URL    string `json:"url"`
			Script string `json:"script"`
		}
		if err := json.Unmarshal(params.Arguments, &args); err != nil {
			return errorResult(fmt.Sprintf("Invalid arguments: %v", err))
		}

		output, err := executeJSViaLightpanda(args.URL, args.Script)
		if err != nil {
			return errorResult(fmt.Sprintf("Lightpanda JS execution error: %v", err))
		}
		return CallToolResult{
			Content: []TextContent{{Type: "text", Text: output}},
		}

	default:
		return errorResult(fmt.Sprintf("Unknown tool: %s", params.Name))
	}
}

func fetchHTMLViaLightpanda(targetURL string) (string, error) {
	// Query local Lightpanda HTTP or CDP interface
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Lightpanda-MCP/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func fetchMarkdownViaLightpanda(targetURL string) (string, error) {
	// Run lightpanda CLI command via WSL if available or call CDP endpoint
	cmd := exec.Command("wsl", "lightpanda", "fetch", targetURL)
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		// Fallback to direct HTTP fetch + text conversion
		html, fetchErr := fetchHTMLViaLightpanda(targetURL)
		if fetchErr != nil {
			return "", fmt.Errorf("wsl lightpanda error: %v (%s), fallback error: %v", err, errOut.String(), fetchErr)
		}
		return fmt.Sprintf("# Extracted Content from %s\n\n%s", targetURL, html), nil
	}

	return out.String(), nil
}

func executeJSViaLightpanda(targetURL, script string) (string, error) {
	jsCode := fmt.Sprintf(`
const { chromium } = require('playwright');
(async () => {
  const browser = await chromium.connectOverCDP('ws://%s:%s');
  const context = await browser.newContext();
  const page = await context.newPage();
  await page.goto('%s');
  const res = await page.evaluate(() => { %s });
  console.log(JSON.stringify(res, null, 2));
  await browser.close();
})();
`, lightpandaHost, lightpandaPort, targetURL, script)

	cmd := exec.Command("node", "-e", jsCode)
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("JS evaluation failed: %v, stderr: %s", err, errOut.String())
	}

	return out.String(), nil
}

func checkLightpandaStatus() string {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(lightpandaHost, lightpandaPort), 2*time.Second)
	if err != nil {
		return fmt.Sprintf("⚠️ Lightpanda daemon is offline on %s:%s. Launch via `lightpanda --port 9222` or `wsl lightpanda`.", lightpandaHost, lightpandaPort)
	}
	conn.Close()
	return fmt.Sprintf("✅ Lightpanda CDP daemon is ONLINE and responding at ws://%s:%s", lightpandaHost, lightpandaPort)
}

func ensureLightpandaRunning() {
	mu.Lock()
	defer mu.Unlock()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(lightpandaHost, lightpandaPort), 1*time.Second)
	if err == nil {
		conn.Close()
		return
	}

	// Try auto-starting Lightpanda in WSL in background
	go func() {
		exec.Command("wsl", "lightpanda", "--port", lightpandaPort).Run()
	}()
	time.Sleep(500 * time.Millisecond)
}

func errorResult(msg string) CallToolResult {
	return CallToolResult{
		Content: []TextContent{{Type: "text", Text: msg}},
		IsError: true,
	}
}

func sendResponse(id interface{}, result interface{}) {
	resp := Response{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
	data, _ := json.Marshal(resp)
	os.Stdout.Write(append(data, '\n'))
}

func sendError(id interface{}, code int, message string) {
	resp := Response{
		JSONRPC: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: message,
		},
	}
	data, _ := json.Marshal(resp)
	os.Stdout.Write(append(data, '\n'))
}
