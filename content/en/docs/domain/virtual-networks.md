---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Virtual Networks"
date: 2022-12-22
---

{{% pageinfo %}}
Trustgrid networks are assigned to [Domains]({{< ref "docs/domain" >}}). They define an overlay network and routes for layer 3 connectivity. Trustgrid [Nodes]({{< ref "docs/concepts/node" >}}) are then assigned to Networks. Virtual Network Overlays work similarly to Amazon's VPCs but can span between clouds, datacenters, and between the cloud and the edge.
{{% /pageinfo %}}

## Virtual Network Overlay

Virtual Networks - Layer 3 

 - Trustgrid uses a Virtual Network to traffic between Nodes. This avoids common 
challenges like conflicting subnets in large, non-centralized networks. 

- Inside and Outside NATs are used to expose hosts for traffic

- Routes may configured on the Nodes

- ACLs can be applied

Simple Host Communication

- An Inside Nat on the Edge Node that maps 10.0.5.250 to 172.16.3.250  would allow the Data Center host (10.0.1.150) to communicate with 172.16.3.250 host using the address 10.0.5.250

- An Outside Nat on the Edge Node that maps 10.0.1.150 to 10.0.5.150 would allow the Edge host (172.16.3.250) to communicate the Data Center host at 10.0.5.150.

- **NOTE:** A route would have to be added to the default gateway of 172.16.3.0/24 for 10.0.5.150 to route traffic for 10.0.5.0/24 through the gw of the Edge Node (172.16.0.5) data interface.

![img](/docs/domain/virtual-networks-yay.png)



## Virtual IP Addresses

Through the use of Inside and Outside NAT rules, local IP addresses to the Trustgrid [Node]({{< ref "docs/concepts/node" >}}) may be exposed on the network through a Virtual IP Address.  For instance, a local device may have a local IP address of 192.168.1.100, but you may want it to appear to other devices or applications on the virtual network as IP address 10.0.20.100. Virtual IP Subnets are also supported where each IP in a subnet is mapped to the corresponding IP in a virtual subnet.

## Inside NAT

Inside NATs are added to a network to expose hosts in the private network to authorized traffic in the virtual network. For example a host at 192.168.1.184/32 could be exposed as 172.16.1.184/32. Traffic from authorized Edge Nodes could then communicate with 172.16.1.184.

## Outside NAT

The use of an Outside NAT can eliminate the need to create [routes]({{< ref "docs/domain/routes" >}}) on internal gateways. Outsides NATs translate address in Edge (remote) private networks to a privately addressable local IP. For example a host at 172.16.10.184/32 could be translated to 192.168.1.184/32.

## Routes

[Routes]({{< ref "docs/domain/routes" >}}) are used to specify the network interface to be used when accessing specific subnets.

## Domains

All Trustgrid [Nodes]({{< ref "docs/concepts/node" >}})  are grouped into [domains]({{< ref "docs/domain" >}}). Communication is only possible within the [domain]({{< ref "docs/domain" >}}) a [node]({{< ref "docs/concepts/node" >}})  is deployed into. [Domains]({{< ref "docs/domain" >}}) operate as a security group for connectivity and a management group for bulk updates or changes. Virtual Networks are assigned to individual [domains]({{< ref "docs/domain" >}}).


