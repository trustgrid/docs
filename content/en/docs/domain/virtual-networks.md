---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Virtual Networks"
date: 2022-12-28
---

{{% pageinfo %}}
Trustgrid networks are assigned to [domains]({{< ref "docs/domain" >}}). They define an overlay network and routes for layer 3 connectivity. Trustgrid [nodes]({{< ref "docs/node" >}}) are then assigned to [networks]({{< ref "docs/overview/networking" >}}). Virtual Network Overlays work similarly to Amazon's VPCs but can span between clouds, datacenters, and between the cloud and the edge.
{{% /pageinfo %}}

### Virtual Network Overlay

##### Virtual Networks - Layer 3

- Trustgrid uses a virtual network to traffic between [nodes]({{< ref "docs/node" >}}). This avoids common
  challenges like conflicting subnets in large, non-centralized networks.

- Inside and outside NATs are used to expose hosts for traffic

- [Routes]({{< ref "docs/domain/routes" >}}) may be configured on the [nodes]({{< ref "docs/node" >}})

- ACLs can be applied

##### Simple Host Communication

- An inside Nat on the edge node that maps 10.0.5.250 to 172.16.3.250 would allow the data center host (10.0.1.150) to communicate with 172.16.3.250 host using the address 10.0.5.250

- An outside Nat on the edge node that maps 10.0.1.150 to 10.0.5.150 would allow the edge host (172.16.3.250) to communicate the data center host at 10.0.5.150.

- **NOTE:** A [route]({{< ref "/docs/domain/routes" >}}) would have to be added to the default gateway of 172.16.3.0/24 for 10.0.5.150 to route traffic for 10.0.5.0/24 through the gw of the edge node (172.16.0.5) data interface.

![img](/docs/domain/virtual-networks-yay.png)

### Virtual IP Addresses

Through the use of inside and outside NAT rules, local IP addresses to the Trustgrid [node]({{< ref "docs/node" >}}) may be exposed on the [network]({{< ref "docs/overview/networking" >}}) through a Virtual IP Address. For instance, a local device may have a local IP address of 192.168.1.100, but you may want it to appear to other devices or applications on the virtual network as IP address 10.0.20.100. Virtual IP Subnets are also supported where each IP in a subnet is mapped to the corresponding IP in a virtual subnet.

### Inside NAT

Inside NATs are added to a network to expose hosts in the private network to authorized traffic in the virtual network. For example a host at 192.168.1.184/32 could be exposed as 172.16.1.184/32. Traffic from authorized edge nodes could then communicate with 172.16.1.184.

### Outside NAT

The use of an outside NAT can eliminate the need to create [routes]({{< ref "/docs/domain/routes" >}}) on internal gateways. Outsides NATs translate address in edge (remote) private networks to a privately addressable local IP. For example a host at 172.16.10.184/32 could be translated to 192.168.1.184/32.

### Routes

[Routes]({{< ref "docs/domain/routes" >}}) are used to specify the network interface to be used when accessing specific subnets.

### Domains

All Trustgrid [nodes]({{< ref "docs/node" >}}) are grouped into [domains]({{< ref "docs/domain" >}}). Communication is only possible within the [domain]({{< ref "docs/domain" >}}) a [node]({{< ref "docs/node" >}}) is deployed into. [Domains]({{< ref "docs/domain" >}}) operate as a security group for connectivity and a management group for bulk updates or changes. Virtual networks are assigned to individual [domains]({{< ref "docs/domain" >}}).
