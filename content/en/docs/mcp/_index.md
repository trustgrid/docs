---
title: "Trustgrid MCP"
linkTitle: "MCP"
weight: 5
description: "Connect AI clients to Trustgrid with the public remote MCP server."
---

{{% pageinfo %}}
Trustgrid provides a public [Model Context Protocol](https://modelcontextprotocol.io/) server for AI clients that support remote MCP over HTTPS. Use it to search the Trustgrid API surface, run scoped read tools, and trigger node diagnostics without hand-writing API calls.
{{% /pageinfo %}}

## Public endpoint

The Trustgrid remote MCP endpoint is:

```text
https://mcp.trustgrid.io/mcp
```

The public hosted server supports **remote Streamable HTTP over HTTPS**.

## Recommended tool group

For most clients, start with:

```text
https://mcp.trustgrid.io/mcp/codemode
```

That gives the model the smaller codemode tool surface:

- `search`
- `describe`
- `code`
- `followUp`

If you want other tool groups, append scope parts to the URL:

| URL | Purpose |
| --- | --- |
| `https://mcp.trustgrid.io/mcp/codemode` | Recommended default for most agentic clients |
| `https://mcp.trustgrid.io/mcp/read` | One direct MCP tool per read API operation |
| `https://mcp.trustgrid.io/mcp/tools` | Live node diagnostic tools |
| `https://mcp.trustgrid.io/mcp/codemode/read` | Codemode plus direct read tools |
| `https://mcp.trustgrid.io/mcp/read/tools` | Direct read tools plus node diagnostics |
| `https://mcp.trustgrid.io/mcp/codemode/read/tools` | Everything |

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
