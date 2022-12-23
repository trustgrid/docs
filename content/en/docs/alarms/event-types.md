---
Title: "Event Types"
Tags: ["event types", "events", "alarms"]
Date: 2022-12-23
---

| Filter Type Name | Severity | Message | Description |
|------------------|----------|---------|-------------|
| All Gateways Disconnected | WARNING | | This event will be triggered if an edge node loses connectivity to all it’s available gateways. |
| All Peers Disconnected | WARNING | | This even will be triggered if a node is configured as a gateway an it lose connectivity to all it’s configured edge nodes. |
| Certificate Expiring | WARNING | | Alerts when a certificate uploaded via Portal → Certificates is about to expire in less than 3 months. |
| Cluster Failover | INFO | [Node]({{< ref "docs/concepts/node" >}}) is the active [cluster]({{< ref "docs/concepts/cluster" >}}) member | Sent by a [node]({{< ref "docs/concepts/node" >}}) when it claims the Active role. |
| Cluster Failover | Info | [Node]({{< ref "docs/concepts/node" >}}) is no longer the active [cluster]({{< ref "docs/concepts/cluster" >}}) member | Sent by a [node]({{< ref "docs/concepts/node" >}}) when it releases the active role. |
| Connection Flapping | WARNING | | Alerts when a [node]({{< ref "docs/concepts/node" >}}) disconnects and reconnects 10 times within 5 minutes. A follow up alert will be sent after each subsequent 120 disconnect/reconnects. |
| Connection Flapping | INFO | | This alert will be sent when a [node’s]({{< ref "docs/concepts/node" >}}) Connection Flapping issue has been resolved. |
| Connection Timeout | ERROR | | Alerts when a [node]({{< ref "docs/concepts/node" >}}) does not reconnect after a profile update has been pushed to the [node]({{< ref "docs/concepts/node" >}}). |
| Data Plane Disruption | WARNING | | This event will be triggered if a data plane tunnel is terminated unexpectedly. |
| Date Plane Disruption | WARNING | Gateway *nodename* removed from the domain | If a [node]({{< ref "docs/concepts/node" >}}) acting as a gateway (public, private or hub) is disabled or the gateway service is disabled all other [nodes]({{< ref "docs/concepts/node" >}}) in the domain will log this event. |
| Gateway Connectivity Health Check | CRITICAL | | Alerts when an edge node is unable to communicate with a gateway node. |
| Gateway Connectivity Health Check | INFO | | Alerts when an edge node reestablishes connectivity to a gateway node after a failure is reported. |
| Gateway Ingress Limit Reached | ERROR | | Alerts when a gateway node’s ingress limit is above 95 percent utilization for 2 minutes straight
| Metric Threshold Violation | ERROR | | Alerts when a [node]({{< ref "docs/concepts/node" >}}) cpu, ram, disk, or latency configured metric threshold is violated. 
| Metric Threshold Violation | INFO | | Alerts when a previously reported threshold violation has been cleared. |
| Network Error | CRITICAL | Stale ARP Detected | Alerts when a stale ARP entry is detected for a [node’s]({{< ref "docs/concepts/node" >}}) [cluster]({{< ref "docs/concepts/cluster" >}}) IP. |
| Network Error | WARNING | Interface {OS interface name} is running with half-duplex | Alerts if an interface has been detected running at half-duplex. This is almost always a result of a failure to auto-negotiate the speed/duplex and can result in poor performance. |
| [Node]({{< ref "docs/concepts/node" >}}) Connect | INFO | | Alerts when a [node]({{< ref "docs/concepts/node" >}}) connects to the control plane. |
| [Node]({{< ref "docs/concepts/node" >}}) Disconnect | WARNING | | Alerts when a [node]({{< ref "docs/concepts/node" >}}) disconnects from the control plane
| Order Created | INFO | | Alerts when a new provisioning order has been created. |
| Order Commented | INFO | | Alerts when a provisioning order case has been commented.
| Unauthorized IP | WARNING | | Alerts when a [node's]({{< ref "docs/concepts/node" >}}) public IP has been locked but the connection to the control plane comes from a different IP.



 


