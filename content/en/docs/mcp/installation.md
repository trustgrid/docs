---
title: "Installation"
linkTitle: "Installation"
weight: 30
description: "How to configure the Trustgrid MCP server in Claude, Cursor, VS Code, Copilot, and other AI clients."
---

## Before you start

You need two things:

1. **Your MCP endpoint URL** — `https://mcp.<domain>.trustgrid.io/mcp` where `<domain>` matches your Trustgrid organization. The default is `https://mcp.trustgrid.io/mcp`. The path suffix selects the [tool group]({{<ref "docs/mcp/tools">}}): `/mcp` (codemode), `/mcp/read`, `/mcp/tools`, or a combination like `/mcp/codemode/read`.

2. **Credentials** — an API token (`clientId:clientSecret`) from **User Management** → **API Access** in the portal, or use OAuth if your client supports it (no manual token needed). See [Authentication]({{<ref "docs/mcp/authentication">}}).

Examples below use `https://mcp.trustgrid.io/mcp` as the URL and show API token auth. Replace both with your actual values.

---

## Claude Desktop

Edit the Claude Desktop config file:

- **macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

Claude Desktop supports OAuth. If you omit the `headers` block, it will open a browser login on first connection.

Restart Claude Desktop after saving the config. The Trustgrid tools will appear in the tool selector.

---

## Claude Code

Run once to register the server:

```bash
claude mcp add --transport http trustgrid https://mcp.trustgrid.io/mcp \
  --header "Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
```

Claude Code also supports OAuth. To use it, omit `--header` — Claude Code will prompt for browser authorization on first use:

```bash
claude mcp add --transport http trustgrid https://mcp.trustgrid.io/mcp
```

To configure directly in `~/.claude/settings.json`:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## GitHub Copilot (VS Code)

Open or create `.vscode/mcp.json` in your workspace (for project-scoped access) or use your VS Code user settings for global access.

**Workspace (`.vscode/mcp.json`):**

```json
{
  "servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

**User settings (`settings.json`):**

```json
{
  "mcp.servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

MCP tool use in VS Code requires Copilot Chat in agent mode. Open Copilot Chat, switch to **Agent** mode, and Trustgrid tools will be available.

---

## VS Code (without Copilot)

VS Code's built-in MCP support uses the same config format as Copilot. Add to your workspace `.vscode/mcp.json`:

```json
{
  "servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Cursor

Create or edit `.cursor/mcp.json` in your workspace, or `~/.cursor/mcp.json` for global access:

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

Cursor detects MCP servers automatically from this file. No restart required — open Cursor's chat panel and the Trustgrid tools will appear.

---

## Windsurf

Edit `~/.codeium/windsurf/mcp_config.json`:

```json
{
  "mcpServers": {
    "trustgrid": {
      "serverUrl": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

Restart Windsurf after saving. Trustgrid tools will be available in the Cascade panel.

---

## opencode

Edit `~/.config/opencode/opencode.json`:

```json
{
  "mcp": {
    "trustgrid": {
      "type": "remote",
      "enabled": true,
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Codex (OpenAI CLI)

Add to `~/.codex/config.toml` or pass via command line:

```toml
[[mcp_servers]]
name = "trustgrid"
uri = "https://mcp.trustgrid.io/mcp"

[mcp_servers.headers]
Authorization = "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
```

Or as a CLI flag:

```bash
codex --mcp-server 'trustgrid=https://mcp.trustgrid.io/mcp' \
  --mcp-header 'Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET' \
  "List all offline nodes"
```

---

## Kiro

Create or edit the MCP config in your Kiro workspace settings:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Antigravity

Add the server in Antigravity's MCP configuration:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Generic / other clients

Any MCP client that supports Streamable HTTP transport can connect to the Trustgrid MCP server. The minimal configuration you need:

| Setting | Value |
|---|---|
| Transport | `http` (Streamable HTTP) |
| URL | `https://mcp.<domain>.trustgrid.io/mcp` |
| Auth header | `Authorization: trustgrid-token clientId:clientSecret` |

If your client supports custom HTTP headers, pass the `Authorization` header directly. If it supports OAuth 2.0 discovery, point it at the URL and it will negotiate auth automatically via the `/.well-known/oauth-authorization-server` metadata endpoint.
