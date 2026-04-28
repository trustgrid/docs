---
title: "Trustgrid MCP"
linkTitle: "MCP"
weight: 5
description: "Connect AI clients to Trustgrid with the public remote MCP server."
---

{{% pageinfo %}}
Trustgrid provides a public [Model Context Protocol](https://modelcontextprotocol.io/) server for AI clients that support remote **Streamable HTTP over HTTPS**. Start with the scoped `read` tools to safely explore what is available in your org, then move on to broader or state-changing workflows only when you know what you need.
{{% /pageinfo %}}

## Public endpoint

The Trustgrid remote MCP endpoint is:

```text
https://mcp.trustgrid.io/mcp
```

The public hosted server supports **remote Streamable HTTP over HTTPS**, not the legacy HTTP SSE transport.

## Start here: `read`

For most users, start with:

```text
https://mcp.trustgrid.io/mcp/read
```

This gives your client a direct, read-only oriented tool surface for querying Trustgrid resources without jumping straight into broader action-capable workflows.

Use `read` first to inspect objects such as nodes, clusters, domains, routes, alerts, and audit feeds. Once you understand the shape of the data you need, move on to `codemode` or other combined scopes for more advanced workflows.

## Tool groups

The hosted MCP server exposes these scoped paths:

| URL | Purpose |
| --- | --- |
| `https://mcp.trustgrid.io/mcp/read` | Recommended starting point for direct read/query operations |
| `https://mcp.trustgrid.io/mcp/codemode` | Compact agent-oriented tool surface for search, describe, and scripted workflows |
| `https://mcp.trustgrid.io/mcp/tools` | Live node diagnostic tools |
| `https://mcp.trustgrid.io/mcp/codemode/read` | Codemode plus direct read tools |
| `https://mcp.trustgrid.io/mcp/read/tools` | Direct read tools plus node diagnostics |
| `https://mcp.trustgrid.io/mcp/codemode/read/tools` | Everything |

### `codemode` tools

If you want the smaller programmable codemode surface, use:

```text
https://mcp.trustgrid.io/mcp/codemode
```

That scope provides:

- `search`
- `describe`
- `code`
- `followUp`

If you are unsure where to begin, do the boring safe thing first: connect to `/mcp/read`, look around, and only then graduate to workflows that can do more than query state. Saves everyone from exciting 3 a.m. surprises.

## Choose an auth model

- **OAuth** is best for interactive use in clients that support remote MCP OAuth.
- **JWT Bearer tokens** are good for advanced setups and temporary interactive testing.
- **API tokens** (`trustgrid-token client-id:client-secret`) are best for service accounts and automation.

See [Authentication]({{<relref "/docs/mcp/authentication" >}}) for the exact header formats and client behavior.

## In this section

- [Authentication]({{<relref "/docs/mcp/authentication" >}})
- [Tools and groups]({{<relref "/docs/mcp/tools-and-groups" >}})
- [Transport, endpoints, errors, and rate limits]({{<relref "/docs/mcp/reference" >}})
- [Client installation guides]({{<relref "/docs/mcp/clients" >}})
