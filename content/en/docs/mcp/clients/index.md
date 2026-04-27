---
title: "Client installation guides"
linkTitle: "Client installation"
weight: 40
description: "Install the Trustgrid MCP server in popular AI clients and editors."
---

All examples below use the recommended production scoped endpoint:

```text
https://mcp.trustgrid.io/mcp/codemode
```

If you want direct read tools or node diagnostics instead, replace the URL with another scoped path such as `/mcp/read`, `/mcp/tools`, or `/mcp/codemode/read`.

{{% alert color="info" %}}
Auth placeholder syntax is client-specific. Examples using `${...}` rely on that client's native variable interpolation. Examples using values like `YOUR_TG_JWT` are literal placeholders that you must replace manually before saving the config.
{{% /alert %}}

{{% alert color="warning" %}}
The snippets below are example configurations based on each client's documented MCP config shape at the time of writing. These clients change fast, because apparently shipping stable config formats would be too fucking easy, so if a field name has moved check the client's current MCP docs and keep the Trustgrid URL and auth header format the same.
{{% /alert %}}

## Claude Desktop

Claude Desktop remote MCP uses **custom connectors**.

1. Open **Customize → Connectors**.
2. Add a custom connector.
3. Enter the server URL: `https://mcp.trustgrid.io/mcp/codemode`
4. If Claude Desktop offers OAuth, complete the browser login flow.
5. If you are using manual auth instead, configure the Trustgrid auth header in the connector settings if your deployment exposes that option.

## Claude Code

### OAuth

```bash
claude mcp add trustgrid --transport http https://mcp.trustgrid.io/mcp/codemode
```

Then run `/mcp` in Claude Code and complete the browser login.

### JWT or API token

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer ${TG_JWT}"
      }
    }
  }
}
```

You can also use `trustgrid-token ${TG_CLIENT_ID}:${TG_CLIENT_SECRET}` in the `Authorization` header.

## GitHub Copilot CLI

Add Trustgrid to `~/.copilot/mcp-config.json`:

```json
{
  "mcpServers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer YOUR_TG_JWT"
      },
      "tools": ["*"]
    }
  }
}
```

If your Copilot CLI build supports remote MCP OAuth in your environment, you can try the same URL without static headers and follow the interactive auth flow.

## OpenCode

Add Trustgrid to `~/.config/opencode/opencode.json`:

### OAuth

```json
{
  "mcp": {
    "trustgrid": {
      "type": "remote",
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

### JWT or API token

```json
{
  "mcp": {
    "trustgrid": {
      "type": "remote",
      "enabled": true,
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "oauth": false,
      "headers": {
        "Authorization": "Bearer YOUR_TG_JWT"
      }
    }
  }
}
```

## Windsurf

Add Trustgrid to `~/.codeium/windsurf/mcp_config.json`:

```json
{
  "mcpServers": {
    "trustgrid": {
      "serverUrl": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer ${env:TG_JWT}"
      }
    }
  }
}
```

Windsurf documents OAuth support for remote HTTP MCP servers. If you prefer OAuth, add the remote server through the MCP UI and complete the browser flow when prompted.

## Kiro

Add Trustgrid to `.kiro/settings/mcp.json` in your workspace or `~/.kiro/settings/mcp.json` for user scope:

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer YOUR_TG_JWT"
      },
      "disabled": false,
      "autoApprove": [],
      "disabledTools": []
    }
  }
}
```

## Antigravity

Antigravity supports remote MCP servers, but current public guidance indicates MCP OAuth support is limited. Use a JWT or API token header.

Raw config shape:

```json
{
  "mcpServers": {
    "trustgrid": {
      "serverUrl": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer YOUR_TG_JWT"
      }
    }
  }
}
```

## VS Code

Add Trustgrid to `.vscode/mcp.json` or your user MCP config:

```json
{
  "servers": {
    "trustgrid": {
      "type": "http",
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer ${input:tg-jwt}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "tg-jwt",
      "description": "Trustgrid JWT",
      "password": true
    }
  ]
}
```

If your VS Code build supports remote MCP OAuth for this server, you can omit the header and authenticate through the VS Code MCP flow.

## Cursor

Add Trustgrid to `~/.cursor/mcp.json` or `.cursor/mcp.json`:

### OAuth

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp/codemode"
    }
  }
}
```

### JWT

```json
{
  "mcpServers": {
    "trustgrid": {
      "url": "https://mcp.trustgrid.io/mcp/codemode",
      "headers": {
        "Authorization": "Bearer ${env:TG_JWT}"
      }
    }
  }
}
```

## Codex

Add Trustgrid to `~/.codex/config.toml`:

### OAuth

```toml
[mcp_servers.trustgrid]
url = "https://mcp.trustgrid.io/mcp/codemode"
```

Then authenticate with:

```bash
codex mcp login trustgrid
```

### JWT

```toml
[mcp_servers.trustgrid]
url = "https://mcp.trustgrid.io/mcp/codemode"
bearer_token_env_var = "TG_JWT"
```

## Using API tokens instead of JWTs

Where a client supports custom headers, replace the `Authorization` value with:

```text
trustgrid-token <client-id>:<client-secret>
```

Examples:

```text
Authorization: trustgrid-token 12345-abcde:super-secret-value
```

## OAuth vs manual headers

| Client | Best option |
| --- | --- |
| Claude Desktop | OAuth |
| Claude Code | OAuth |
| Copilot CLI | Manual headers unless your build supports OAuth for remote MCP |
| OpenCode | OAuth |
| Windsurf | OAuth or manual headers |
| Kiro | Manual headers |
| Antigravity | Manual headers |
| VS Code | OAuth or manual headers |
| Cursor | OAuth or manual headers |
| Codex | OAuth or `bearer_token_env_var` |
