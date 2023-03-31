---
title: "Static Routing"
date: 2023-2-14
---

Static routing allows a node to route traffic to a remote network in addition to and [virtual network define routes]({{<ref "/docs/domain/virtual-networks/routes">}}). At this level routes can be created for destination CIDRs outside the [virtual network's Network CIDR]({{<ref "">}})

Find static route definitions and change them under the `Static Routing` section of the VPN configuration.

{{<tgimg src="add-modal.png" caption="Add Route dialog" >}}

{{<fields>}}
{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route. Only visible if [additional Gateway paths]({{<ref "/docs/nodes/gateway">}}) are defined on the node. {{</field>}} 
{{</fields>}}