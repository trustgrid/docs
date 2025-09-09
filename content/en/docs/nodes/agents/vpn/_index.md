---
title: Virtual Network(VPN)
linkTitle: VPN
description: Agent Virtual Network configuration
---
The VPN section allows for viewing and configuring an agent's virtual network settings and access tools for troubleshooting.

{{<tgimg src="agent-vpn.png" width="95%">}}

## VPN Toolbar
At the top of all agent VPN pages is a bar that contains information about the virtual network attachment and tools.
{{<fields>}}
{{<field Network>}}Displays the name of the virtual network the agent is attached to and its configured address space.{{</field>}}
{{<field "Validation CIDR">}}(Optional) Can be defined to validate that the Virtual CIDR address in Network Address Translations falls within an expected network.{{</field>}}
{{<field "Virtual Management IP">}}The IP address assigned to the `trustgrid0` tunnel interface. This IP is either assigned manually or automatically from an [IP Pool]({{<relref "/docs/domain/virtual-networks/ip-pool">}}) {{</field>}}
{{<field "Sniff Virtual Traffic">}}This button launches the [Interface Sniff Tool]({{<relref "/tutorials/interface-tools/sniff-interface-traffic">}}) scoped to the `trustgrid0` interface. {{</field>}}
{{<field "View Virtual Routes">}}This button launches the [View Virtual Route Table]({{<relref "/tutorials/remote-tools/view-virtual-route-table">}}) tool.{{</field>}}
{{</fields>}}

## Virtual Network Attachment

Agents are attached to a virtual network during creation but can be detached and attached as needed. At this time agents only support a single virtual network attachment.

### Attach Agent to Virtual Network 
1. Navigate to the VPN page of the agent
1. Under the Actions dropdown select Attach
1. Select the desired virtual network.
    - If the virtual network has [IP Pools]({{<relref "/docs/domain/virtual-networks/ip-pool">}}) configured, the agent will automatically be assigned an IP from the pool. Click Save to complete.
    - If not, you will be provided the following fields: {{<fields>}}
    {{<field "Validation CIDR">}}(optional) if provided this will be used to confirm all NAT rule virtual CIDRs fall within the validation CIDR.{{</field>}}
    {{<field "Virtual Management IP">}}Assigned to the `trustgrid0` tunnel interface on the agent for communication on the virtual network. {{</field>}}
    {{</fields>}}
1. Click Save

### Detach Agent from Virtual Network
1. Navigate to the VPN page of the agent. You will be automatically sent to the Address Translation panel for the currently attached network. 
1. Click the "Manage Networks" link in the left navigation bar.
1. Click the radio button to the left of the attached network.
1. From the Actions dropdown select Detach.
1. Confirm you want to detach the agent from the virtual network when prompted.