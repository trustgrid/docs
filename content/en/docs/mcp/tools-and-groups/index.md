---
title: "Tools and groups"
linkTitle: "Tools and groups"
weight: 20
description: "Understand the Trustgrid MCP tool groups and what each surface exposes."
---

Trustgrid MCP exposes three public tool groups. You choose them by path.

## Group: `codemode`

Path:

```text
https://mcp.trustgrid.io/mcp/codemode
```

This is the recommended default because it gives the model a small, high-signal tool surface.

### Tools in `codemode`

| Tool | Purpose |
| --- | --- |
| `search` | Search available Trustgrid API functions and documentation by keyword |
| `describe` | Return the full function signature and docs for a single Trustgrid API function |
| `code` | Execute JavaScript that chains Trustgrid API calls and returns a shaped result |
| `followUp` | Retrieve the next page of cached rows from a prior paginated result |
| `switchOrg` | Switch the session to a different organization when using JWT/Bearer auth |

## Group: `read`

Path:

```text
https://mcp.trustgrid.io/mcp/read
```

This surface exposes one direct MCP tool per read-oriented Trustgrid API operation. Tool names are normalized to `snake_case` from the API `operationId`.

Examples include tools such as `list_nodes`, `get_node`, `list_groups`, `get_domain`, and many more.

### Read tool coverage

The read surface is grouped by the same high-level API domains documented in the public API:

- **Alerts**
  - alarms
  - alerts
- **Node**
  - access applications
  - agents
  - appliances
  - certificates
  - clusters
  - domains
  - tags
  - upgrade manager
  - virtual networks
- **Virtual Network**
  - access policies
  - change management
  - DNS
  - DNS zones
  - DNS records
  - network groups
  - network objects
  - port forwarding
  - routes
- **Compute**
  - edge compute
  - repositories
- **Management**
  - audits
  - auth groups
  - groups
  - identity providers
  - observability exporters
  - organizations
  - orders
  - permissions
  - service users
  - users

For the full object model and endpoint details, see [apidocs.trustgrid.io](https://apidocs.trustgrid.io/).

## Group: `tools`

Path:

```text
https://mcp.trustgrid.io/mcp/tools
```

This surface exposes live node diagnostic tools. These are mostly read-only diagnostics that call node services through the Trustgrid API.

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
