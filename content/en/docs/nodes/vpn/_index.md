---
title: "VPN"
date: 2023-2-14
weight: 22
---

{{% pageinfo %}}
Trustgrid nodes can be configured to provide VPN-like functionality, allowing network address translation (NAT) and routing between nodes.
{{% /pageinfo %}}

The VPN feature enables routing of IP packets by utilizing a virtual IP space that is configured on the nodes. In this configuration, the node can be used as the next hop for route destined for the virtual IP space, or the node can be used as the default gateway for a network. The virtual IP space is used as a transit network with NAT being utilized to translate IP addresses on to/from the virtual IP space on either side. It is also possible to preserve the original source and destination by having the virtual IP space be identical to the LAN network.

The Trustgrid VPN feature provides the capability to securely route IP packets between remote networks. In this configuration, Trustgrid nodes can operate as a distributed mesh virtual private network (VPN) that can allow applications to access remote data and services at layer 3 (L3) of the network OSI model. This is done by defining a virtual L3 network (similar to an Amazon VPC) and then selecting how local node networks are exposed and translated into the the virtual address space.

## Attaching a Virtual Network

To use VPN functionality, a node or cluster must be attached to a domain [virtual network]({{<ref "docs/domain/virtual-networks">}}).

Navigate to the `VPN` section for your node or cluster, and select `Actions`->`Attach`.

{{<tgimg src="attach-vnet.png" caption="Attach Virtual Network dialog on a Node" alt="Attach Network dialog with options to select network, Validation CIDR and Virtual Management IP" width="60%" >}}

{{<fields>}}
{{<field "Select Network">}}Drop down list to select [defined virtual networks]({{<ref "/docs/domain/virtual-networks/">}}). {{</field>}}
{{<field "Validation CIDR">}}
* Used to validate NATs defined on this node or cluster.  If the Virtual CIDR of [Inside NATs]({{<ref "docs/nodes/vpn/nats#inside-nats">}}) or the Network Group value of [Outside Nats]({{<ref "docs/nodes/vpn/nats#outside-nats">}}) are outside the Validation CIDR a warning will appear.
* Must be a subnet equal to or smaller than the Virtual Network's Network CIDR.
{{</field>}}
{{<field "Virtual Management IP">}} 
* This IP address is used by the node as the source IP for running [VPN troubleshooting tools]({{<ref "tutorials/remote-tools/">}}).
* Only visible in this dialog when attaching the network to a node.  If attaching to a cluster, you will need to navigate to each member node to set the Virtual Management IP.
{{</field>}}
{{</fields>}}


Once attached, navigate into the network to manage the VPN configuration.
