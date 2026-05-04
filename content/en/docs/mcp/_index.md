---
title: "MCP Server"
linkTitle: "MCP Server"
weight: 5
description: "Connect AI assistants and coding agents to your Trustgrid infrastructure using the Model Context Protocol."
---

The Trustgrid MCP server lets AI assistants — Claude, Copilot, Cursor, and others — interact with your Trustgrid environment using natural language. Ask questions about node status, run network diagnostics, query alerts, and explore your topology without leaving your editor or chat interface.

## What you can do

- **Query infrastructure** — list nodes, inspect configuration, check VPN routes, review certificates and policies
- **Run live diagnostics** — test TCP connectivity, DNS resolution, gateway latency, and packet paths directly on nodes
- **Explore topology** — understand how your virtual networks, clusters, and gateways are connected
- **Search documentation** — semantic search across Trustgrid docs returns relevant excerpts alongside live API data
- **Execute code** (codemode) — run sandboxed JavaScript with full API access for custom analysis and bulk queries
- **Review audit history** — query config changes, node events, and user activity logs

## Endpoint URL

The MCP server is hosted at `https://mcp.<domain>.trustgrid.io` where `<domain>` matches your organization's Trustgrid domain. For example, if your portal is at `portal.acme.trustgrid.io`, your MCP endpoint is `https://mcp.acme.trustgrid.io`.

The default production endpoint is:

```
https://mcp.trustgrid.io
```

The server exposes three tool groups, each at its own path:

| Path | Group | Purpose |
|---|---|---|
| `/mcp` or `/mcp/codemode` | codemode | AI sandbox with full read API access |
| `/mcp/read` | read | Direct read-only API tools |
| `/mcp/tools` | tools | Live node diagnostic tools |
| `/mcp/codemode/read` | codemode + read | Both tool sets combined |
| `/mcp/codemode/read/tools` | all | Everything |

See [Tools]({{<ref "docs/mcp/tools">}}) for details on what each group provides.

## Transport

The Trustgrid MCP server uses **Streamable HTTP transport**. This is the only transport available for external integrations — there is no stdio or WebSocket option. Your MCP client must support HTTP-based MCP connections.

## Authentication

The server accepts two credential types:

- **API token** — a `clientId:clientSecret` pair generated in the Trustgrid portal. Recommended for most integrations.
- **Bearer JWT** — a portal JWT extracted from the `portaljwt` session cookie. Useful for interactive sessions.

The server also supports **OAuth 2.0**. Clients that implement the MCP OAuth handshake (Claude Desktop, Claude Code, and others) will trigger a browser-based login automatically — no manual token needed.

See [Authentication]({{<ref "docs/mcp/authentication">}}) for setup details.

## Rate limits

Requests are rate-limited per token:

| Group | Limit |
|---|---|
| codemode | 10 requests / 60 seconds |
| read | 30 requests / 60 seconds |
| tools | 10 requests / 60 seconds |

The server returns `429 Too Many Requests` with `Retry-After` and `X-RateLimit-*` headers when limits are exceeded.

## Next steps

- [Authentication]({{<ref "docs/mcp/authentication">}}) — get credentials and understand auth options
- [Tools]({{<ref "docs/mcp/tools">}}) — explore available tool groups
- [Installation]({{<ref "docs/mcp/installation">}}) — connect your AI client
- [Troubleshooting]({{<ref "docs/mcp/troubleshooting">}}) — common setup issues
