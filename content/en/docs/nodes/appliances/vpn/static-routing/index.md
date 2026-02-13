---
title: "Static Routing"
aliases: 
    - /docs/nodes/vpn/static-routing
description: Configure static routes for the virtual network for the selected node or cluster
---

Static routing allows a node to route traffic to a remote network in addition to virtual network-defined routes({{<ref "/docs/domain/virtual-networks/routes">}}). Routes defined at this level can be created for destination CIDRs outside the [virtual network's Network CIDR]({{<relref "/docs/domain/virtual-networks#network-cidr">}})

Find static route definitions and change them under the **Static Routing** section of the VPN configuration.

{{<tgimg src="add-modal.png" caption="Add Route dialog" >}}

{{<fields>}}
{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route. Only visible if [additional Gateway paths]({{<ref "/docs/nodes/appliances/gateway">}}) are defined on the node. {{</field>}} 
{{</fields>}}