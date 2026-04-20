---
title: "Static Routing"
aliases: 
    - /docs/nodes/vpn/static-routing
description: Configure static routes for the virtual network for the selected node or cluster
---

Static routing allows a node to route traffic to a remote network in addition to virtual network-defined routes({{<ref "/docs/domain/virtual-networks/routes">}}). Routes defined at this level can be created for destination CIDRs outside the [virtual network's Network CIDR]({{<relref "docs/domain/virtual-networks#network-cidr">}})

Find static route definitions and change them under the **Static Routing** section of the VPN configuration.

{{<tgimg src="add-modal.png" caption="Add Route dialog" >}}

{{<fields>}}
{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route. Only visible if [additional Gateway paths]({{<ref "/docs/nodes/appliances/gateway">}}) are defined on the node. {{</field>}} 
{{</fields>}}

## Route Monitors

Route monitors let Trustgrid test a route destination and mark the route unavailable when the configured monitor checks fail. They are configured per static route.

{{<tgimg src="route-monitors-route-table.png" caption="Static Routing table with route monitor management" width="80%">}}

### Prerequisites

- The node or cluster must already be attached to a [virtual network]({{<ref "/docs/domain/virtual-networks">}}).
- The route must already exist.
- The node should have a **Virtual Management IP** configured for that virtual network because route monitor traffic is sourced from that address. See [Attaching a Virtual Network]({{<ref "/docs/nodes/appliances/vpn">}}).
- The destination IP and protocol must be allowed by the remote host and any firewalls in the path.

### Manage Route Monitors

1. Navigate to the desired node or cluster VPN network.
1. Open **Static Routing**.
1. Find the route you want to monitor.
1. Select **Manage Monitors** on that route.

{{<tgimg src="route-monitors-modal.png" caption="Route Monitors dialog" width="80%">}}

1. Select **Add Monitor**.
1. Enter the monitor settings.
1. Select **Save**.

{{<tgimg src="route-monitors-add-modal.png" caption="Add Route Monitor dialog" width="55%">}}

### Route Monitor Fields

{{<fields>}}
{{<field "Name">}}A unique name for the monitor on that route.{{</field>}}
{{<field "Status">}}Enabled monitors run on schedule. Disabled monitors are kept in configuration but are not executed.{{</field>}}
{{<field "Protocol">}}Supported values are **ICMP** and **TCP**.{{</field>}}
{{<field "Destination IP">}}IPv4 address to test through the selected route.{{</field>}}
{{<field "Destination Port">}}Required for TCP monitors. Not used for ICMP monitors.{{</field>}}
{{<field "Monitor Interval">}}How often the monitor runs, in seconds.{{</field>}}
{{<field "Failures Count">}}How many failed checks must accumulate before the monitor is considered failed.{{</field>}}
{{<field "Max Latency">}}Optional latency threshold in milliseconds. If omitted, the node uses a default of 1000 ms.{{</field>}}
{{</fields>}}

### How Route Monitors Work

- Route monitors run from the node that owns the route and use the virtual network path for that route.
- ICMP monitors send echo requests.
- TCP monitors send a TCP SYN to the configured destination IP and port.
- A monitor can fail because the destination does not respond, returns an unreachable response, returns a TCP reset, or exceeds the configured maximum latency.
- When a monitor reaches the configured **Failures Count**, the route is treated as failed.
- When successful checks resume, the route can recover automatically.

{{<alert color="info">}}Successful checks reduce the accumulated failure count one step at a time. Recovery is not an instant full reset after a single success.{{</alert>}}

### Failover and Recovery Behavior

Route monitors affect whether a route is considered available. If you have multiple routes for the same destination CIDR, Trustgrid will still prefer the lowest metric route that is currently available.

If a failed route starts passing its monitors again, it can become eligible for traffic again automatically. Trustgrid reuses existing monitor instances during route updates to reduce unnecessary monitor flapping, but unstable targets can still cause route changes if the monitor repeatedly fails and recovers.

For guidance on choosing good monitor targets and avoiding false failures, see [Route Monitor Best Practices]({{<relref "tutorials/monitoring/route-monitors">}}).
