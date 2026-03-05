---
title: "Route Monitors"
description: Configure health-check probes on VPN routes to enable automatic failover when a destination becomes unreachable
tags: ["vpn", "routing", "health-check", "failover"]
---

Route monitors are health-check probes attached to VPN routes. They periodically send probes to a destination IP and, when a configurable failure threshold is reached, mark the route as inactive. This enables automatic failover to an alternate route when a destination becomes unreachable or too slow.

Each VPN route can have one or more monitors. If any monitor on a route fails, the route is marked inactive.

{{<alert color="info">}}Route monitors are configured on [virtual network routes]({{<ref "/docs/domain/virtual-networks/routes">}}) at the domain level, not on node-level static routes.{{</alert>}}

{{<alert color="info">}}Cluster support for route monitors is coming soon. Route monitors are currently supported on standalone nodes only.{{</alert>}}

## Monitor Types

### ICMP

An ICMP monitor sends periodic ICMP echo requests (pings) to the destination IP and measures round-trip time (RTT). A probe is considered failed if:

- The destination returns an ICMP Destination Unreachable message.
- The RTT exceeds the configured **Max Latency**.
- No reply is received before the next probe cycle.

{{<alert color="warning">}}ICMP monitors are affected by a known issue where the monitor may not recover after a failure even when the destination becomes reachable again. Until this is resolved, TCP monitors are recommended for production use.{{</alert>}}

ICMP monitors also require a management IP to be configured on the node's VPN attachment (the **IP** field on the virtual network attachment). Without it, the node has no source address for ICMP probes and they will silently fail.

### TCP

A TCP monitor sends a TCP SYN packet to the destination IP and port, and waits for a SYN-ACK response. RTT is measured from when the SYN is sent to when the SYN-ACK is received. After evaluation, the node sends a RST to clean up the half-open connection. A probe is considered failed if:

- A TCP RST is received.
- An ICMP Destination Unreachable message is received.
- No SYN-ACK is received before the probe times out.
- The RTT exceeds the configured **Max Latency**.

TCP monitors are the recommended protocol. They work for destinations that filter ICMP and verify reachability to a specific port.

## Configuration Fields

{{<fields>}}
{{<field "Name">}}A human-readable label for the monitor.{{</field>}}

{{<field "Enabled">}}Whether the monitor is active. Disabling a monitor prevents it from running probes or affecting route state.{{</field>}}

{{<field "Protocol">}}The probe protocol: **ICMP** or **TCP**.{{</field>}}

{{<field "Destination">}}The IPv4 address to probe.{{</field>}}

{{<field "Port">}}(TCP only) The destination TCP port to probe (1–65535).{{</field>}}

{{<field "Interval">}}How often to send a probe, in seconds (1–86400).{{</field>}}

{{<field "Failures">}}The number of consecutive probe failures required before the monitor is marked as failed. Once the failure count drops back to zero the monitor recovers (1–86400). Must be at least 1.{{</field>}}

{{<field "Max Latency">}}The maximum acceptable round-trip time in milliseconds. Probes with an RTT above this value count as failures. Defaults to 1000 ms.{{</field>}}
{{</fields>}}

## Configuring Route Monitors

Route monitors are managed from the virtual network's **Routes** table in the Trustgrid Portal. The **Manage Monitors** action is available on each route row.

### Adding a Monitor

1. Navigate to **Domain** > **Virtual Networks** and select the network.
1. In the **Routes** table, find the route you want to monitor.
1. Click the **Manage Monitors** row action for that route.
1. In the Route Monitors dialog, click **Add**.
1. Fill in the monitor fields:
   - Enter a descriptive **Name**.
   - Set **Enabled** to on.
   - Select the **Protocol** (ICMP or TCP — see [Monitor Types](#monitor-types) for guidance).
   - Enter the **Destination** IP address to probe.
   - If using TCP, enter the **Port**.
   - Set the **Interval** (in seconds) between probes.
   - Set **Failures** — the consecutive failure count required to mark the monitor as failed (minimum 1).
   - Optionally adjust **Max Latency** (default: 1000 ms).
1. Click **Save**.
1. Repeat for any additional monitors on the same route.

### Editing or Disabling a Monitor

1. Navigate to the route's **Manage Monitors** dialog.
1. Click the edit icon on the monitor you want to change.
1. Update the desired fields and click **Save**.

To temporarily stop a monitor without deleting it, set **Enabled** to off.

### Deleting a Monitor

1. Navigate to the route's **Manage Monitors** dialog.
1. Select the monitor(s) to delete.
1. Click **Delete** and confirm.

## Relationship to Virtual Network Route Failover

Route monitors work alongside the [virtual network route]({{<ref "/docs/domain/virtual-networks/routes">}}) metric-based failover system. Virtual network routes with lower metrics are preferred; if the destination node goes offline, traffic falls back to a higher-metric route automatically.

Route monitors add a finer-grained layer of control: a route can be deactivated based on measured data-plane health (reachability and latency) even when the node itself remains online. This allows failover to trigger for scenarios like a downstream network becoming unreachable without the Trustgrid node going offline.

For example, you might configure:
- A primary route (lower metric) with a TCP monitor targeting a critical host behind the remote node.
- A secondary route with a higher metric as the fallback.

When the TCP monitor fails, the primary route deactivates and traffic shifts to the secondary route automatically. When the target recovers, the primary route becomes active again within one probe interval.

{{<alert color="info">}}A route must also have an active session to be considered active. Monitor state is evaluated on top of session state — a route with a healthy monitor but no active session will still be inactive.{{</alert>}}
