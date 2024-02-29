---
categories: ["overview"]
weight: 10
tags: ["architecture"]
title: "Basic Architecture"
date: 2022-12-28
description: >
  Trustgrid basic architecture overview
---

## Nodes

The basic building block of a Trustgrid network is a [node]({{<ref "docs/nodes">}}). Nodes are [deployed]({{<relref "./deployment">}}) on-premises or in your cloud provider and connected to each other to build a [data plane]([{{<relref "#data-plane">}}]) on top of which additional services such as [virtual networking]({{<relref "#virtual-networks">}}), [ztna access]({{<relref "#ztna-access">}}), and [compute]({{<relref "#edge-compute">}}) can be provided.

Nodes can be [deployed](../deployment) either as an **agent**, a service running in an existing operating system, or as an **appliance** which bundles the operating system and Trustgrid software together. 

Additionally, appliance-based nodes can be [clustered]({{<relref "docs/clusters">}}) to provide high availability at a site. 

### Gateways
While all nodes can act as edge nodes (making only outbound connections), appliance-based nodes can also be configured to act as [gateway servers]({{<relref "/docs/nodes/appliances/gateway/gateway-server">}}) which listen for inbound connections from other edge nodes in the account.  

These connections create the [private data plane]({{<relref "#data-plane">}}) which enabls the Trustgrid network services.

## Data Plane

The data plane is built between edge nodes and their gateways nodes to allow private connectivity between sites. It provides a virtual network overlay on top of the public internet to securely connect distributed systems and applications.

### Virtual Networks

Nodes can be attached to [virtual networks]({{<ref "docs/domain/virtual-networks">}}), which provide a way to share network configuration like [routes]({{<ref "docs/domain/virtual-networks/routes">}}), NATs, and [ACLs]({{<ref "docs/domain/virtual-networks/access-policy">}}) at scale. 

#### Layer 4 Proxy
The data plane can also be used to enable Layer 4 proxy connections across two nodes.  In this setup, one node will act at the [connector]({{<relref "docs/nodes/shared/connectors">}}) listening on a specified port.  When traffic is received on that port it will be forwarded to the configured [service]({{<relref "docs/nodes/shared/services">}}). The node that hosts that service will then forward the traffic on to the configured IP and port using it's local interfaces.

### ZTNA Applications

[ZTNA Applications]({{<ref "docs/applications">}}) can be exposed through nodes or clusters. Access to an application can be restricted via [access policies]({{<ref "docs/applications/access-policy">}}), for example to only allow users from a specific country.


## Control plane

In addition to the data plane, all Trustgrid nodes build connections to the Trustgrid control plane. This allows centralized management and monitoring of all nodes from the Trustgrid portal and api including: 

* Software updates are provided through our apt repository that ensures nodes have security updates available and only run with tested software permutations
* Network, node, user, and application configuration is managed through our web portal
* Configuration changes are broadcast to nodes as needed, for example when adding routes or adding a gateway to your network

All services of the control plane are contained [Trustgrid's reserved address spaces]({{<relref "help-center/kb/site-requirements#trustgrid-control-plane">}}) to limit the required firewall rules required for an edge node to operate behind a firewall.


{{<tgimg src="basic-connectivity.png" width="80%" alt="Basic connectivity diagram" caption="Basic Connectivity Diagram">}}

