---
Title: "Event Types"
Tags: ["event types", "events", "alarms"]
Date: 2022-12-28
---

| Filter Type Name | Severity | Message | Description |
| ---------------- | -------- | ------- | ----------- |
| All Gateways Disconnected  | WARNING  |  | This [event]({{<ref "docs/alarms/events" >}}) will be triggered if an edge node loses connectivity to all it’s available gateways.  |
| All Peers Disconnected | WARNING  |  | This [event]({{<ref "docs/alarms/events" >}}) will be triggered if a node is configured as a gateway an it lose connectivity to all it’s configured edge nodes.   |
| BGP Peer Connectivity | ERROR  | BGP peer <peer IP> has disconnected	| This event will be triggered if a BGP peer disconnects.  |
| BGP Peer Connectivity | INFO  | BGP peer <peer IP> has re-connected	| This event will be triggered if a BGP peer reconnects.  |
| Certificate Expiring | WARNING  |  | Alerts when a certificate uploaded via Portal → Certificates is about to expire in less than three months.   |
| Cluster Failover | INFO | [Node]({{<ref "docs/nodes" >}}) is the active [cluster]({{<ref "docs/clusters" >}}) member | Sent by a [node]({{<ref "docs/nodes" >}}) when it claims the active role.    |
| Cluster Failover | INFO | [Node]({{<ref "docs/nodes" >}}) is no longer the active [cluster]({{<ref "docs/clusters" >}}) member | Sent by a [node]({{<ref "docs/nodes" >}}) when it releases the active role.    |
| Configuration Update Failure	| INFO | Unexpected error pulling the most recent configuration form the cloud endpoint | Indicate the node is unable to connect to the configuration REST API endpoint within the Trustgrid Control Plane. [Verify all required communication is allowed to Control Plane.]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}) |
| Connection Flapping  | WARNING  |  | Alerts when a [node]({{<ref "docs/nodes" >}}) disconnects and reconnects 10 times within 5 minutes. A follow up alert will be sent after each subsequent 120 disconnect/reconnects.  |
| Connection Flapping  | INFO |  | This alert will be sent when a [node’s]({{<ref "docs/nodes" >}}) "Connection Flapping" issue has been resolved.  |
| Connection Timeout | ERROR  |  | Alerts when a [node]({{<ref "docs/nodes" >}}) does not reconnect after a profile update has been pushed to the [node]({{<ref "docs/nodes" >}}).  |
| Data Plane Disruption  | WARNING  |  | This [event]({{<ref "docs/alarms/events" >}}) will be triggered if a data plane tunnel is terminated unexpectedly.  |
|Data and Control Plane Disruption| ERROR |  Data and Control plane connection is up/down | Generated when a node cannot establish a connection to the Trustgrid control plane and **all** its configured data plane gateways. If clustered this will mark the node as unhealthy triggering a cluster failover| 
| Data Plane Disruption  | WARNING  | Gateway _nodename_ removed from the domain | If a [node]({{<ref "docs/nodes" >}}) acting as a gateway (public, private or hub) is disabled or the gateway service is disabled all other [nodes]({{<ref "docs/nodes" >}}) in the domain will log this [event]({{<ref "docs/alarms/events" >}}). |
| DNS Resolution | ERROR | DNS resolution failed/re-established | Once per hour the node will attempt to resolve a trustgrid.io DNS address using the configured DNS servers. This event is triggered if the node is unable to resolve the requested address. 
| Gateway Connectivity Health Check  | CRITICAL |  | Alerts when an edge node is unable to communicate with a gateway node.   |
| Gateway Connectivity Health Check  | INFO |  | Alerts when an edge node reestablishes connectivity to a gateway node after a failure is reported.   |
| Gateway Ingress Limit Reached  | ERROR  |  | Alerts when a gateway node’s ingress limit is above 95 percent utilization for two minutes straight.   |
| Gateway UDP Tunnel Error | INFO | UDP Tunnel connection has been re-established for node=`<peer-node>` and endpoint=`<peer-ip>:<peer-port>` | Event generated when a previously disconnected UDP tunnel is re-established and traffic should flow through it again. |  
| Gateway UDP Tunnel Error | ERROR | UDP Tunnel has timed out for node=`<peer-node>` and endpoint=`<peer-ip>:<peer-port>` | Alerts when a UDP tunnel times out after not receiving the keep alive packet for 2 minutes (default gateway timeout). |  
| Metric Threshold Violation | ERROR  |  | Alerts when a [node]({{<ref "docs/nodes" >}}) cpu, ram, disk, or latency configured metric threshold is violated.  |
| Metric Threshold Violation | INFO |  | Alerts when a previously reported threshold violation has been cleared.    |
| Network Error  | CRITICAL | Stale ARP detected   | Alerts when the active node in a cluster detects another MAC address responding to ARPs for a configured [cluster IP address]({{<relref "/docs/clusters/cluster-only-config#cluster-ip-address">}}).  This can occur briefly during a failover if the standby node begins arping before the previous active node releases that role. Other causes include an IP conflict, proxy ARP configured on another device on that network, or the attached switch not updating its ARP cache.  |
| Network Error  | ERROR | Unable to create/update Azure IP configuration for nic=LAN. Error Code=&lt;Azure error code&gt; | Alerts when a [node]({{<ref "docs/nodes" >}}) is unable to create or update an Azure Cluster IP configuration. The error from the Azure API is included in the event. |
| Network Error  | WARNING  | Interface {OS interface name} is running with half-duplex  | Alerts if an interface has been detected running at half-duplex. This is almost always a result of a failure to auto-negotiate the speed/duplex and can result in poor performance.  |
| Networking Framework Memory Management | ERROR | The networking framework has exhausted all allocated memory. Please contact support@trustgrid.io to notify of this issue. | Alerts if the Java Virtual Machine the Trustgrid node service uses runs out of available memory. | 
| Network Health Check | ERROR | Interface {OS interface name} is down / All interfaces are up | Generated when link is lost on a configured interface. This is generated even if the interfaces is set with the "Ignore Health Check" setting enabled. Note: Interfaces with APIPA addresses are ignored. |
| [Node]({{<ref "docs/nodes" >}}) Connect  | INFO |  | Alerts when a [node]({{<ref "docs/nodes" >}}) connects to the control plane.   |
| [Node]({{<ref "docs/nodes" >}}) Disconnect | WARNING  |  | Alerts when a [node]({{<ref "docs/nodes" >}}) disconnects from the control plane.  |
| Node Delete | WARNING | Node deleted | Event generated whenever a node is deleted |
| Node Stop Error | ERROR | Failed to stop the Node service cleanly | Indicates that the Trustgrid service did not stop normally prior to this instance starting|
| Order Created  | INFO |  | Alerts when a new provisioning order has been created.   |
| Order Commented  | INFO |  | Alerts when a provisioning order case has been commented.    |
| Repo Connectivity | ERROR | Repo connectivity failed | Alerts when a node cannot connect to the Trustgrid update repository | 
| Repo Connectivity | INFO | Repo connectivity re-established | Alerts when a node re-establishes connectivty to the Trustgrid update repository. This event clears the Repo Connectivity error alert.| 
| SSH Lockdown | ERROR | SSH allowing connections from non local address and port | Alerts when SSH on an appliance-based node is configured to listen on any IP other than local host (127.0.0.1).|
| SSH Lockdown | INFO | SSH listening only on local address and port | Alerts when SSH on an appliance- based node is properly locked down. This event clears the SSH Lockdown error alert.| 
| Unauthorized IP  | WARNING  |  | Alerts when a [node's public IP has been locked]({{<ref "/tutorials/limit-node-functionality" >}}) but the connection to the control plane comes from a different IP.   |
| Deregister | INFO | Device was deregistered from the console | Event is sent if a user with console access runs the [deregistration process]({{<ref "/tutorials/local-console-utility/remote-registration#deregistration-process">}}) | 
---
