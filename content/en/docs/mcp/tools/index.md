---
title: "Available Tools"
linkTitle: "Tools"
description: "Tool groups exposed by the Trustgrid MCP server and what each tool does."
---

The Trustgrid MCP server organizes its tools into three groups. Each group is available at a distinct URL path, and paths can be combined to get the union of multiple groups.

## Tool groups

### codemode (`/mcp/codemode`)

The codemode group is the default when no path suffix is specified. It provides an AI code execution sandbox with full read-only access to the Trustgrid API, plus documentation search and structured resource inspection. This reduces token usage and agent turns, and allows your agent to perform complex multi-step reasoning and data retrieval in a deterministic single turn — without needing to switch back and forth between tools.

| Tool       | Description                                                                                                                           |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `search`   | Semantic and keyword search across Trustgrid documentation. Returns relevant excerpts with source links.                              |
| `describe` | Describe a codemode function. This provides the full signature to help your agent write valid code.                                   |
| `code`     | Execute sandboxed JavaScript with access to the full read-only Trustgrid API. Use for custom queries, aggregations, and bulk lookups. |
| `followUp` | Paginate through results from a prior `code` or `search` call.                                                                        |

The `code` tool gives the AI direct access to all Trustgrid API endpoints within a JavaScript sandbox. It can traverse relationships, aggregate data across nodes, or do anything the REST API supports — without risking writes or configuration changes.

Codemode also includes all [node diagnostic tools](#tools-mcptools) from the `tools` group.

The codemode scope requires broad read permissions. See [Authentication]({{<relref "docs/mcp/authentication" >}}) for credential setup.

### read (`/mcp/read`)

The read group exposes individual Trustgrid API operations as discrete MCP tools — one tool per API endpoint, roughly.

| Tool family      | Examples                                                                   | Description                                                                           |
| ---------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| Nodes            | `list_nodes`, `get_node`, `list_node_events`                               | Node inventory, full node records, and node event history.                            |
| Clusters         | `list_clusters`, `get_cluster`, `list_cluster_vpn_routes`                  | Cluster inventory plus cluster VPN topology, routes, interfaces, and services.        |
| Domains          | `get_domain`                                                               | Domain details and configuration.                                                     |
| Alerts           | `list_alerts_v2`, `list_node_alerts_v2`                                    | Active alert lists at the org or node level.                                          |
| Audit logs       | `tail_config_audit`, `tail_node_audit`                                     | Configuration change history and node audit trails.                                   |
| Events           | `list_events`                                                              | Platform event stream for the org.                                                    |
| VPN networks     | `list_node_vpn_networks`, `list_node_vpn_routes`, `list_node_vpn_services` | Node VPN topology, routes, interfaces, import/export routes, and services.            |
| Virtual networks | `list_virtual_networks`, `list_network_routes`, `list_network_objects`     | Overlay network configuration including routes, objects, groups, and port forwarding. |
| Flow logs        | `list_flow_logs`                                                           | Query network traffic flow logs with filters and pagination.                          |

The read group requires a smaller set of OAuth scopes than codemode, making it suitable for tightly scoped service integrations.

### tools (`/mcp/tools`)

The tools group exposes discrete, live diagnostic tools that execute directly on Trustgrid nodes. Each tool invokes a service on the node and returns the live result.

| Tool                     | Description                                                                                                    |
| ------------------------ | -------------------------------------------------------------------------------------------------------------- |
| `get_runtime_status`     | Node runtime status: running services, disk usage, memory, systemd unit states.                                |
| `get_network_status`     | Node networking: interface states, link status, assigned IP addresses as seen by the node right now.           |
| `get_dataplane_status`   | Gateway route connectivity: which gateway routes are up or down, including per-route status.                   |
| `get_errors_status`      | Currently reported errors from the node. Pass `startupOnly: true` to filter to startup-phase errors only.      |
| `test_tcp_connectivity`  | Attempt a TCP handshake from the node to a `host:port` target and report success or failure.                   |
| `test_repo_connectivity` | Verify the node can reach its configured Trustgrid apt repositories. Reports per-repo success.                 |
| `test_dns_health`        | Test DNS resolution against the node's configured resolvers. Returns per-server results.                       |
| `test_gateway_latency`   | Run a latency trace from the node to a named gateway. Returns per-hop latency measurements.                    |
| `test_packet_path`       | Simulate a TCP packet entering the node's virtual network and trace the path through routing and policy rules. |
| `get_packet_capture`     | Run a bounded tcpdump on a node interface and return captured packet lines. Accepts BPF filter expressions.    |

The tools group requires `nodes::read` plus service-specific scopes. The node does not need to be fully healthy to run most diagnostics — that's the point.

### Historical telemetry

These tools read the stored tunnel latency and [hop monitoring]({{<relref "tutorials/gateway-tools/monitoring-network-hops-to-peers" >}}) data shown in a node's **Data Plane** panel. They do not run anything on the node. They are available in the `codemode`, `read`, and `tools` groups.

| Tool                          | Description                                                                                                                                                   |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `get_hop_monitoring`          | Per-hop RTT history to each gateway peer: hop number, IP and FQDN variants, min/max/mean RTT, sample counts, and coverage across the returned minute buckets. |
| `get_gateway_latency_history` | Tunnel RTT history to each gateway and path: min/max/mean RTT, average per-minute standard deviation, sample counts, and coverage.                            |

Both tools accept the same inputs:

| Parameter               | Description                                                                                                                 |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `nodeID`                | Node UID, FQDN, or bare name. Required.                                                                                     |
| `window`                | Relative lookback: `1h` (default) or `24h`. Cannot be combined with `startTime` and `endTime`.                              |
| `startTime` / `endTime` | Absolute ISO-8601 timestamps with `Z` or a UTC offset. Both are required together. The range can be at most 24 hours.       |
| `gateway`               | Optional exact gateway peer name. Case-insensitive.                                                                         |
| `includeSeries`         | Set to `true` to include the per-minute series. Defaults to `false` so large ranges return summaries only.                  |

Timestamps in the response are UTC ISO-8601. Ranges longer than one hour are fetched in sequential one-hour requests and combined.

The source data is per-minute aggregates, not raw RTT samples. The tools do not calculate p95, jitter, or packet loss. Missing hop samples are reported as coverage gaps. The node must be an online appliance node.

## Combining groups

Paths are combinable and order-independent. The server serves the union of all named groups:

```
https://mcp.trustgrid.io/mcp/codemode/read
https://mcp.trustgrid.io/mcp/read/tools
https://mcp.trustgrid.io/mcp/codemode/read/tools
```

Use the combined paths when your workflow needs both the `code` sandbox and direct per-resource tools, or when you want a single endpoint for a client that doesn't let you configure multiple servers.

## Rate limits

Limits apply per token across all requests:

| Group    | Requests | Window     |
| -------- | -------- | ---------- |
| codemode | 10       | 60 seconds |
| read     | 30       | 60 seconds |
| tools    | 10       | 60 seconds |
| default  | 30       | 60 seconds |

When a limit is exceeded the server returns `429 Too Many Requests` with:

- `Retry-After: <seconds>` — how long to wait before retrying
- `X-RateLimit-Limit: <n>` — the limit for this bucket
- `X-RateLimit-Remaining: 0` — remaining requests in the window
- `X-RateLimit-Reset: <unix timestamp>` — when the window resets
