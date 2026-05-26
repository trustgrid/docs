---
title: "Installation"
linkTitle: "Installation"
weight: 30
description: "How to configure the Trustgrid MCP server in Claude, Cursor, VS Code, Copilot, and other AI clients."
---

## Before you start

You need two things:

1. **Your MCP endpoint URL** — The default is `https://mcp.trustgrid.io/mcp`. If your organization uses a tenant-specific portal domain (for example, `acme.trustgrid.io`), use `https://mcp.acme.trustgrid.io/mcp` instead. The path suffix selects the [tool group]({{<relref "docs/mcp/tools" >}}): `/mcp/codemode` (codemode), `/mcp/read`, `/mcp/tools`, or a combination like `/mcp/codemode/read`.

2. **Credentials** — OAuth is the preferred method: if your client supports it, point it at the URL and it will handle login automatically — no token management required. For clients that don't support OAuth, generate an API token (`clientId:clientSecret`) from **User Management** → **API Access** in the portal. See [Authentication]({{<relref "docs/mcp/authentication" >}}).

Examples below use `https://mcp.trustgrid.io/mcp/codemode` as the URL. OAuth examples are shown first where supported; API token examples follow as a fallback. Replace URLs and credentials with your actual values.

---

## Claude Desktop

Claude Desktop users have two supported options:

### Option 1: Add Trustgrid as a custom connector

This is the preferred setup for Claude Desktop because it uses Anthropic's built-in remote MCP connector flow.

**For Pro and Max plans:**

1. Navigate to **Customize > Connectors**.
2. Click **+** then **Add custom connector**.
3. Add your connector's remote MCP server URL: `https://mcp.trustgrid.io/mcp/all`
4. Optionally, click **Advanced settings** to specify an OAuth Client ID and OAuth Client Secret for your server.
5. Finish configuring your connector by clicking **Add**.

**For Team and Enterprise plans:**

The user adding the connector must have permission to add a custom connector. In Anthropic's current documentation, that means an **Owner** or **Primary Owner** must first add the connector for the organization:

1. Navigate to **Organization settings > Connectors**.
2. Click the **Add** button.
3. Hover over **Custom**, then select **Web**.
4. Add your connector's remote MCP server URL: `https://mcp.trustgrid.io/mcp/all`
5. Optionally, click **Advanced settings** to specify an OAuth Client ID and OAuth Client Secret for your server.
6. Finish configuring your connector by clicking **Add**.

After the connector has been added, each member can enable it:

1. Navigate to **Customize > Connectors**.
2. Find the custom connector your Owner added in the list. It will have a **Custom** label.
3. Click **Connect** to authenticate and start using the connector with Claude.

After configuration, you can enable the connector in an individual conversation from the **+** button in the lower-left of the chat interface, then **Connectors**.

### Option 2: Use `mcp-remote` in `claude_desktop_config.json`

If you prefer Claude Desktop's local MCP config, use `mcp-remote` to bridge Claude Desktop to the Trustgrid remote MCP server.

- **macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "trustgrid": {
      "command": "npx",
      "args": ["-y", "mcp-remote", "https://mcp.trustgrid.io/mcp/all"]
    }
  }
}
```

Restart Claude Desktop after saving the config. On first connection, `mcp-remote` will open the OAuth login flow in your browser.

---

## Claude Code

Claude Code supports OAuth. This is the recommended setup — it will prompt for browser authorization on first use:

```bash
claude mcp add --transport http trustgrid https://mcp.trustgrid.io/mcp
```

If you prefer an API token, pass the header explicitly:

```bash
claude mcp add --transport http trustgrid https://mcp.trustgrid.io/mcp/codemode \
  --header "Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
```

Or add the API token directly in `~/.claude/settings.json`:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## GitHub Copilot (VS Code)

VS Code supports OAuth for MCP servers. Open or create `.vscode/mcp.json` in your workspace (for project-scoped access) or use your VS Code user settings for global access. Omitting the `headers` block will trigger OAuth login on first use.

**Workspace (`.vscode/mcp.json`) — OAuth:**

```json
{
  "servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

**User settings (`settings.json`) — OAuth:**

```json
{
  "mcp.servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

If you prefer an API token, add a `headers` block to either config:

```json
"headers": {
  "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
}
```

MCP tool use in VS Code requires Copilot Chat in agent mode. Open Copilot Chat, switch to **Agent** mode, and Trustgrid tools will be available.

---

## Cursor

Cursor supports OAuth for remote MCP servers. Create or edit `.cursor/mcp.json` in your workspace, or `~/.cursor/mcp.json` for global access. Omit `headers` and Cursor will handle OAuth automatically via dynamic client registration:

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

If you prefer an API token, add a `headers` block:

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp/codemode",
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

Windsurf supports OAuth for all MCP transport types. Edit `~/.codeium/windsurf/mcp_config.json` and omit `headers` to use OAuth:

```json
{
  "mcpServers": {
    "trustgrid": {
      "serverUrl": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

If you prefer an API token, add a `headers` block:

```json
{
  "mcpServers": {
    "trustgrid": {
      "serverUrl": "https://mcp.trustgrid.io/mcp/codemode",
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

opencode supports OAuth with automatic detection. Edit `~/.config/opencode/opencode.json` and omit `headers` — if the server requires authentication, opencode will prompt you to authorize in your browser:

```json
{
  "mcp": {
    "trustgrid": {
      "type": "remote",
      "enabled": true,
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

You can also trigger auth manually: `opencode mcp auth trustgrid`

If you prefer an API token, add a `headers` block:

```json
{
  "mcp": {
    "trustgrid": {
      "type": "remote",
      "enabled": true,
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Codex (OpenAI CLI)

Codex supports OAuth. Add the server URL to `~/.codex/config.toml`, then run `codex mcp login trustgrid` to complete the OAuth flow:

```toml
[mcp_servers.trustgrid]
url = "https://mcp.trustgrid.io/mcp/codemode"
```

If you prefer an API token, add the header to the config instead:

```toml
[mcp_servers.trustgrid]
url = "https://mcp.trustgrid.io/mcp/codemode"

[mcp_servers.trustgrid.http_headers]
Authorization = "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
```

Or pass via CLI flag:

```bash
codex --mcp-server 'trustgrid=https://mcp.trustgrid.io/mcp/codemode' \
  --mcp-header 'Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET' \
  "List all offline nodes"
```

---

## Kiro

Kiro does not currently support OAuth for MCP servers. Use an API token. Create or edit the MCP config in your Kiro workspace settings:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Antigravity

Add the server in Antigravity's MCP configuration. Check Antigravity's docs for OAuth support — if supported, omit the `headers` block. Otherwise use an API token:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET"
      }
    }
  }
}
```

---

## Generic / other clients

Any MCP client that supports Streamable HTTP transport can connect to the Trustgrid MCP server. If your client supports OAuth 2.0 discovery, that's the preferred approach — point it at the URL and it will negotiate auth automatically via the `/.well-known/oauth-authorization-server` metadata endpoint. For clients that don't support OAuth, pass the `Authorization` header directly.

| Setting | Value |
|---|---|
| Transport | `http` (Streamable HTTP) |
| URL | `https://mcp.<domain>.trustgrid.io/mcp/codemode` |
| Auth (OAuth) | Use `/.well-known/oauth-authorization-server` for discovery |
| Auth (API token) | `Authorization: trustgrid-token clientId:clientSecret` |
