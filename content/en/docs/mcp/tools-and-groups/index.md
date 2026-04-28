---
title: "Tools and groups"
linkTitle: "Tools and groups"
weight: 20
description: "Understand the Trustgrid MCP tool groups and what each surface exposes."
---

Trustgrid MCP exposes three public tool groups. You choose them by path.

{{% alert color="info" %}}
Start with the `read` group unless you already know you need something else. It is the safest entry point for exploring nodes, clusters, routes, alerts, and other Trustgrid resources before moving to more advanced or stateful workflows.
{{% /alert %}}

## Group: `read`

Path:

```text
https://mcp.trustgrid.io/mcp/read
```

This surface exposes one direct MCP tool per read-oriented Trustgrid API operation from a curated public allowlist. Tool names are normalized to `snake_case` from the API `operationId`.

Examples include tools such as `list_nodes`, `get_node`, `list_clusters`, `get_domain`, `list_virtual_networks`, `list_network_routes`, `list_alerts_v2`, and `tail_node_audit`.

### Why start with `read`

- safest place to explore what data is available
- direct mapping to common query operations
- easy for users to inspect results before attempting diagnostics or automation

### Read tool coverage

The public `read` scope is intentionally narrower than the full Trustgrid API. It currently exposes a curated subset focused on these read paths:

- alerts and event feeds
- clusters and cluster VPN details
- domains
- nodes and node VPN details
- virtual networks
- network access policies, auth groups, groups, objects, port forwardings, and routes
- config and node audit tail feeds

This is not a one-to-one mirror of every public API domain. If you need the exact current surface, connect to `/mcp/read` and inspect the advertised tools, or use `codemode` to search and describe available functions first.

For the broader Trustgrid API object model and endpoint details, see [apidocs.trustgrid.io](https://apidocs.trustgrid.io/).

## Group: `codemode`

Path:

```text
https://mcp.trustgrid.io/mcp/codemode
```

Use this when you want a compact programmable surface after you have already explored the domain with `read`, or when your client works best with search/describe/code style workflows.

### Tools in `codemode`

| Tool | Purpose |
| --- | --- |
| `search` | Search available Trustgrid API functions and documentation by keyword |
| `describe` | Return the full function signature and docs for a single Trustgrid API function |
| `code` | Execute JavaScript that chains Trustgrid API calls and returns a shaped result |
| `followUp` | Retrieve the next page of cached rows from a prior paginated result |

## Group: `tools`

Path:

```text
https://mcp.trustgrid.io/mcp/tools
```

This surface exposes live node diagnostic tools. These are mostly read-only diagnostics that call node services through the Trustgrid API.

Reach for `tools` after `read` when you need live diagnostics from a node, not just control-plane data. They are still safer than broad write APIs, but they are operational actions rather than plain queries.

### Tools in `tools`

| Tool | Purpose |
| --- | --- |
| `tg_node_runtime_status` | Live runtime status for a node |
| `tg_node_network_status` | Interfaces, addresses, and link state |
| `tg_node_dataplane_status` | Gateway route and dataplane status |
| `tg_node_errors_status` | Currently reported node errors |
| `tg_node_tcp_test` | Outbound TCP connectivity test from a node |
| `tg_node_repo_test` | Repository connectivity test |
| `tg_node_dns_test` | DNS resolution test against node DNS servers |
| `tg_node_gateway_latency_test` | Gateway latency trace |
| `tg_test_traffic` | Simulated packet path / policy decision test |
| `tg_node_sniff_traffic` | Bounded packet capture on a node interface |

## Combining groups

Groups are composable. These URLs are equivalent regardless of part order:

- `https://mcp.trustgrid.io/mcp/codemode/read`
- `https://mcp.trustgrid.io/mcp/read/codemode`

Use combined paths when your client benefits from both a compact codemode surface and direct read tools.

If you want the full per-tool inventory instead of the high-level group summary on this page, see [Tool reference]({{<relref "/docs/mcp/tool-reference" >}}).

Recommended progression for new users:

1. Start with `/mcp/read`
2. Add `/codemode` when you want search/describe/code workflows
3. Add `/tools` when you need live diagnostics from nodes
