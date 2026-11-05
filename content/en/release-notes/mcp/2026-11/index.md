---
title: November 2026 MCP Server Release Notes
linkTitle: November 2026
type: docs
date: 2026-11-05
description: "November 2026 MCP Server Release Notes"
---

## November 5, 2026 - Minor Release

### Historical Hop and Gateway Latency Tools

Agents can now read the historical tunnel latency and hop monitoring data shown in a node's **Data Plane** panel. Two read-only tools have been added:

- `get_hop_monitoring` returns per-hop RTT summaries for each gateway peer: hop number, IP and FQDN variants, min/max/mean RTT, sample counts, and coverage.
- `get_gateway_latency_history` returns tunnel RTT summaries for each gateway and path: min/max/mean RTT, sample counts, and coverage.

Both tools accept a node UID, FQDN, or bare name, a relative window of `1h` (default) or `24h`, or an absolute `startTime` and `endTime` range of up to 24 hours. An optional `gateway` filter limits the result to one peer. Set `includeSeries: true` to include the per-minute series.

See [Available Tools]({{<relref "docs/mcp/tools" >}}) for the tool list and [Monitoring Network Hops to Peers]({{<relref "tutorials/gateway-tools/monitoring-network-hops-to-peers#viewing-hop-data-with-mcp" >}}) for an example.
