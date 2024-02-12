---
title: Virtual Network(VPN)
linkTitle: VPN
description: Agent Virtual Network configuration
weight: 25
---
The VPN section allows for viewing and configuring an agent's virtual network settings and access tools for troubleshooting.

{{<tgimg src="agent-vpn.png" width="95%">}}

## VPN Toolbar
At the top of all agent VPN pages is a bar that contains information about the virtual network attachment and tools.
{{<fields>}}
{{<field Network>}}Displays the name of the virtual network the agent is attached to and its configured address space.{{</field>}}
{{<field "Validation CIDR">}}(Optional) Can be defined to validate that the Virtual CIDR address in Network Address Translations falls within an expected network.{{</field>}}
{{<field "Virtual Management IP">}}The IP address assigned to the `trustgrid0` tunnel interface. This IP is either assigned manually or automatically from an [IP Pool]() {{</field>}}
{{<field "Sniff Virtual Traffic">}}This button launches the [Interface Sniff Tool]({{<relref "/tutorials/interface-tools/sniff-interface-traffic">}}) scoped to the `trustgrid0` interface. {{</field>}}
{{<field "View Virtual Routes">}}This button launches the [View Virtual Route Table]({{<relref "/tutorials/remote-tools/view-virtual-route-table">}}) tool.{{</field>}}
{{</fields>}}

## Virtual Network Attachment

