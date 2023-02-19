---
Title: "Info Visor"
linkTitle: "Info Visor"
---

You can view at-a-glance node information from the info visor, either by clicking `Info` at the top right of the node view, or pressing the back tick key.

![img](visor.png)

{{<fields>}}
{{<field "Control Plane" >}}
Indicates whether the node is connected to the Trustgrid control plane
{{</field >}}

{{<field "Data Plane" >}}
Indicates whether the node is connected to its gateways
{{</field >}}

{{<field "Status" >}}
Enabled nodes will connect to gateways and the Trustgrid mesh network. If a node is disabled, it will connect only to the control plane and wait to be enabled. Disabled nodes do not pass traffic and do not support remote management.
{{</field >}}

{{<field "Last Heartbeat" >}}
The last time the node sent a heartbeat to the control plane
{{</field >}}

{{<field "Device Type" >}}
The type of hardware or virtualization of the node
{{</field >}}

{{<field "Upgrade Status" >}}
Information about the last upgrade request sent to the node
{{</field >}}

{{<field "Tags" >}}
Tags applied to the node
{{</field >}}

{{<field "Public IP" >}}
The WAN IP through which the node has connected to the control plane. You may toggle the lock icon to pin the IP address. Attempts to connect to the control plane from a different IP from the locked IP address will be denied
{{</field >}}

{{<field "Version" >}}
The node software version
{{</field >}}

{{<field "Name" >}}
The node name
{{</field >}}

{{<field "Node ID" >}}
The node's unique ID
{{</field >}}

{{<field "Package Version" >}}
The node software package version
{{</field >}}
{{</fields>}}

### Open Alerts

Unresolved alerts will be listed at the bottom of the visor. They can be directly resolved from the visor by clicking the `Resolve` button.
