---
linkTitle: Client
Title: Gateway Client Settings
descriptions: Configure settings for the Trustgrid gateway client
aliases: 
    - /docs/nodes/gateway/gateway-client
description: Configure client settings 
---

Settings on this page determine how a node's client connection to gateway peers behave.

## Settings
{{<tgimg src="gateway-client-settings.png" width="50%" caption="Gateway Client settings" alt="screenshot of gateway client settings">}}
{{<fields>}}
{{<field "Connectivity to Public Gateways">}}
Options are `Allowed` or `Denied`. If set to `Denied` this will cause the node to not attempt connections to public gateways. This might be desired if you want the node to only connect to configured private gateways. Or if you have private gateways that do not need to connect to the public gateways in your organization.
{{</field>}}
{{</fields>}}



## Hop Monitoring Settings
{{<tgimg src="gateway-hop-monitor-settings.png" width="85%" caption="Gateway Hop Monitoring settings" alt="screenshot of gateway client hop monitor settings">}}
{{<alert color="info">}} The settings below require the node to be running the [April 2025 major appliance release]({{<ref "/release-notes/node/2025-04/index.md">}}) or later. Prior versions only support enabled or disabled.{{</alert>}}
{{<fields>}}
{{<field "Monitor Hops to Gateway Servers" >}}
Determines if the node will attempt to [monitor hops to gateway peers]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers">}}). The possible values are:
- **Always** - The node will attempt to monitor hops to all gateway peers.
- **Only when peer requests** - The node will only attempt to monitor hops to gateway peers that request it. This is the default state.
- **Never** - The node will never attempt to monitor hops to gateway peers, even if peers request it.
{{</field>}}
{{<field "Monitor Hops Protocol" >}}
The protocol used to probe each hop. Requires the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}) or later.
- **SYN** - Probes using TCP SYN packets, reusing the gateway's TCP port so traffic already allowed out to the gateway is allowed for monitoring. This is the default and is the method prior versions used.
- **ICMP** - Probes using ICMP. Both the edge and gateway must allow the traffic. ICMP makes it easier to detect a lower-than-expected MTU along the path.
- **SACK** - Builds a full TCP session and sends 1-byte TCP packets within it until it gets a response, then closes the session. This helps on paths where network devices do not handle bare SYN traces well.
{{</field>}}
{{<field "Monitor Hops Interval" >}} The interval time, in seconds, between gathering hop monitoring data. Default is 20s. {{</field>}}
{{<field "Support Monitor Hops Resets" >}}
{{<alert color="warning">}}Deprecated as of the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}). Hop monitoring no longer sends resets to intermediate hops, so this setting has no effect.{{</alert>}}
Determines if the node will send reset (RST) packets for the TCP connections it attempts.  Doing so reduces, but does not eliminate the number of resets seen on the WAN interface.
- **Enabled** - The node will send reset packets. This is the default state.
- **Disabled** - The node will not send reset packets.
{{</field>}}
{{<field "Monitor Hops SYN Payload Size" >}} Determines the size of the TCP SYN payload sent. By default the payload is 0. Can be set between 0 and 1440. Recommended max is the lower of 1440 or the WAN MTU minus 60 bytes. {{</field>}}
{{</fields>}}
## Gateway Latency Monitors
{{<alert color="info">}}Gateway Latency Monitors require the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}) or later.{{</alert>}}

A gateway latency monitor watches the round trip time (RTT) of the node's tunnel to a gateway and reacts when it exceeds a threshold. RTT is the time a packet takes to travel to the gateway and back. The purpose is to let a cluster member mark itself unhealthy when its tunnel latency to a gateway gets too high, triggering a failover to a healthier member. If the node is not in a cluster, the monitor still sends a [Gateway Latency Exceeded]({{<relref "/docs/alarms/event-types#gateway-latency-exceeded">}}) event, but traffic stays on the node.

{{<tgimg src="gateway-latency-monitors.png" width="85%" caption="Gateway Latency Monitors" alt="Gateway Latency Monitors settings with a Failure Mode and a table of monitors">}}

### Failure Mode
Determines how triggered monitors affect the node's health:
- **None** - The node is not marked unhealthy when monitors trigger, unless a monitor is marked as [Critical](#monitor-fields). This is the default.
- **Any** - The node is marked unhealthy when any single monitor triggers. All monitors must recover before the node becomes healthy again.
- **All** - The node is marked unhealthy only when every monitor is triggered. The node recovers as soon as any single monitor recovers.

### Monitor Fields
Use **Add Latency Monitor** to add a monitor.
{{<fields>}}
{{<field "Gateway Node">}}The gateway to monitor latency to.{{</field>}}
{{<field "Gateway Path">}}(Optional) A specific [gateway path](#gateway-paths) to use for the monitor. Configure this if you want to monitor latency over a second path.{{</field>}}
{{<field "Max Latency">}}The maximum acceptable latency, in milliseconds, for the monitor.{{</field>}}
{{<field "Trigger Count">}}The number of consecutive RTT values that must be above the Max Latency before the monitor triggers.{{</field>}}
{{<field "Recover Count">}}The number of consecutive RTT values that must be below the Max Latency before the monitor marks itself healthy again.{{</field>}}
{{<field "Critical">}}When enabled, the monitor is treated as essential. Its failure alone marks the node unhealthy, and the node cannot recover until this monitor recovers, regardless of the Failure Mode setting.{{</field>}}
{{</fields>}}

{{<alert color="info">}}RTT is tested every 20 seconds by default, so a Trigger Count of 6 requires latency above the threshold for roughly two minutes.{{</alert>}}

## Gateway Paths

Allows you to define alternate paths to a gateway server

{{<fields>}}
{{<field Status>}}
- Enabled - The node **will** attempt to build and utilize the additional path.
- Disabled - The node **will not** build the additional path.
{{</field>}}
{{<field "Name" >}}
A name for the path.
{{</field >}}

{{<field "Gateway Node" >}}
Gateway for which the path is applicable.
{{</field >}}

{{<field "Host IP" >}}
Destination IP address for the path.
{{</field >}}

{{<field "Host Port" >}}
Destination port for the path.
{{</field >}}

{{<field "Local IP" >}}
Use this local IP as the source IP for the connection to the gateway.
{{</field >}}

{{<field "Use as Default" >}}

- True - Will not attempt to connect to the configured Gateway Node using the WAN interface IP and Default Gateway path.
- False - Will attempt to connect to the Gateway node using both this defined path **and** the WAN Interface IP and Default Gateway path.
  {{</field >}}
  {{</fields>}}

{{<tgimg src="gateway-paths.png" width="85%" caption="Example Gateway Path" alt="Screenshot of the gateway paths table">}}

### Add A Gateway Path

1. Click the Add Path link
1. Fill in the fields as desired.
1. Click the green check mark to save the path.
1. Optionally, repeat with additional paths.
1. Click save.

### Delete a Gateway Path

1. Click the X to the right of the desired path.
1. Click save.
