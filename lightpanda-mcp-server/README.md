# 🐼 Lightpanda MCP Server (100% Project-Agnostic)

[![MCP Protocol](https://img.shields.io/badge/MCP-Standard-blue)](https://modelcontextprotocol.io)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://golang.org)
[![Project Agnostic](https://img.shields.io/badge/Path-Project--Agnostic-10b981)](#)

A **100% self-contained, project-agnostic Model Context Protocol (MCP) server** for **Lightpanda**.

You can use this in **ANY project** or **ANY AI Assistant** (OpenCode, Claude Code, Codex, Antigravity, Cursor, Windsurf) without hardcoding user or project-specific paths!

---

## ⚡ Option 1: Global Agnostic PATH (Recommended for All Projects)

Install the binary globally to your system `PATH` once:

```bash
# Copy binary to your system PATH (e.g. /usr/local/bin or C:\Windows or %USERPROFILE%\go\bin)
go install .
```

Now, in **ANY project**, your MCP configuration is 100% clean & project-agnostic:

```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "lightpanda-mcp-server"
    }
  }
}
```

---

## 📁 Option 2: Project-Relative Path (Drop the folder into any repository)

If you copy the `lightpanda-mcp-server` folder directly into any project repo, use relative paths:

### On Windows / Linux / macOS:
```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "./lightpanda-mcp-server/lightpanda-mcp-server"
    }
  }
}
```

---

## 🚀 Option 3: Zero-Build `go run` (Runs dynamically in any project)

No pre-compilation needed! Runs dynamically inside any project:

```json
{
  "mcpServers": {
    "lightpanda": {
      "command": "go",
      "args": ["run", "./lightpanda-mcp-server"]
    }
  }
}
```

---

## ⚙️ Environment Variables

- `LIGHTPANDA_HOST` (default: `127.0.0.1`): Host IP of local Lightpanda daemon.
- `LIGHTPANDA_PORT` (default: `9222`): Port of local Lightpanda CDP server.

---

## 📄 License

MIT License &copy; Lightpanda & Open Source Contributors.
