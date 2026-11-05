---
title: "Historical Telemetry Tools"
linkTitle: "Historical Telemetry"
description: "MCP tools that read stored tunnel latency and hop monitoring data for a node."
---

Two tools read the stored tunnel latency and [hop monitoring]({{<relref "tutorials/gateway-tools/monitoring-network-hops-to-peers" >}}) data shown in a node's **Data Plane** panel. They do not run anything on the node. They are available through the `read` and `tools` groups, and through `codemode` because it includes the `tools` group. See [Available Tools]({{<relref "docs/mcp/tools" >}}) for the groups.

| Tool                          | Description                                                                                                                                                   |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `get_hop_monitoring`          | Per-hop RTT history to each gateway peer: hop number, IP and FQDN variants, min/max/mean RTT, sample counts, and coverage across the returned minute buckets. |
| `get_gateway_latency_history` | Tunnel RTT history to each gateway and path: min/max/mean RTT, average per-minute standard deviation, sample counts, and coverage.                            |

## Parameters

Both tools accept the same inputs:

| Parameter               | Description                                                                                                           |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `nodeID`                | Node UID, FQDN, or bare name. Required.                                                                               |
| `window`                | Relative lookback: `1h` (default) or `24h`. Cannot be combined with `startTime` and `endTime`.                        |
| `startTime` / `endTime` | Absolute ISO-8601 timestamps with `Z` or a UTC offset. Both are required together. The range can be at most 24 hours. |
| `gateway`               | Optional exact gateway peer name. Case-insensitive.                                                                   |
| `includeSeries`         | Set to `true` to include the per-minute series. Defaults to `false` so large ranges return summaries only.            |

## Output

Timestamps in the response are UTC ISO-8601. Ranges longer than one hour are fetched in sequential one-hour requests and combined.

`get_hop_monitoring` returns one entry per gateway peer. Each entry lists the hops on the path with the hop number, the IP and FQDN variants seen, the min, max, and mean RTT, the number of samples, and the percentage of returned minutes in which the hop reported. With `includeSeries: true` it also returns the per-minute hop data, the same data the hops table in the portal shows for a selected time point.

`get_gateway_latency_history` returns one summary per gateway and path. Distinct paths are never combined.

## Limits

- The source data is per-minute aggregates, not raw RTT samples. The tools do not calculate p95, jitter, or packet loss.
- Missing hop samples are reported as coverage gaps. A hop with few samples or low coverage did not respond to probes. See [Known Limitations]({{<relref "tutorials/gateway-tools/monitoring-network-hops-to-peers#known-limitations" >}}) for why this happens.
- Large requested series can exceed an MCP client's response limit. Leave `includeSeries` off for 24-hour ranges unless you need the per-minute data.
- The node must be an online appliance node.
