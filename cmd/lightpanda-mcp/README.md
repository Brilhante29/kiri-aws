# 🐼 Lightpanda MCP Server (Model Context Protocol)

[![MCP Protocol](https://img.shields.io/badge/MCP-2024--11--05-blue)](https://modelcontextprotocol.io)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://golang.org)
[![Lightpanda](https://img.shields.io/badge/Engine-Lightpanda_CDP-8b5cf6)](https://lightpanda.io)

An ultra-fast, standalone **Model Context Protocol (MCP)** server for **Lightpanda** — the AI-native headless browser built for machines (16x lower memory footprint than Chrome).

Plug this MCP module into **ANY AI agent framework** (Claude Code, Google Antigravity, Cursor, Windsurf, LangChain, LlamaIndex, AutoGPT) to grant the AI native, high-speed web browsing, JS evaluation, and Markdown/AX-Tree parsing capabilities.

---

## ⚡ Provided MCP Tools

| Tool Name | Description | Parameters |
| :--- | :--- | :--- |
| `lightpanda_fetch_html` | Fetches raw/clean HTML content from any web page using Lightpanda. | `url` (string) |
| `lightpanda_get_markdown` | Extracts clean Markdown & Accessibility Tree (AX Tree) from a web page. | `url` (string) |
| `lightpanda_execute_js` | Executes custom JavaScript inside Lightpanda headless browser over CDP. | `url` (string), `script` (string) |
| `lightpanda_status` | Returns the health and CDP connectivity status of the local Lightpanda daemon. | None |

---

## 🚀 How to Plug into Any AI Assistant

### 1. Build the Executable

```bash
go build -o bin/lightpanda-mcp ./cmd/lightpanda-mcp
```

---

### 2. Configure in your AI Assistant

#### 🤖 Claude Desktop / Claude Code (`claude_desktop_config.json`)

Add to `~/.config/Claude/claude_desktop_config.json` or `%APPDATA%\Claude\claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "path/to/bin/lightpanda-mcp",
      "env": {
        "LIGHTPANDA_HOST": "127.0.0.1",
        "LIGHTPANDA_PORT": "9222"
      }
    }
  }
}
```

---

#### 🪐 Google Antigravity (`antigravity.json`)

Add to `.gemini/config/mcp.json` or project MCP config:

```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "path/to/bin/lightpanda-mcp"
    }
  }
}
```

---

#### 💻 Cursor / Windsurf (`mcp.json`)

Add under `Settings` -> `Features` -> `MCP`:

```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "C:/Users/Guilherme/Downloads/kiro-main/kiro-main/bin/lightpanda-mcp.exe"
    }
  }
}
```

---

## ⚙️ Environment Variables

- `LIGHTPANDA_HOST` (default: `127.0.0.1`): IP/hostname of the Lightpanda CDP server.
- `LIGHTPANDA_PORT` (default: `9222`): Port of the Lightpanda CDP WebSocket endpoint.

---

## 📄 License

MIT License &copy; Kiro-AWS & Lightpanda Contributors.
