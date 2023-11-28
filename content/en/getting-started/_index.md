---
title: "Getting Started"
linkTitle: "Getting Started"
menu:
  main:
    weight: 10
cascade:
  type: docs
---

Trustgrid is an edge platform that provides [secure]({{<relref "./security">}}), reliable, and high-performance connectivity between distributed systems and applications. The platform is designed to simplify networking for distributed enterprises by providing a software-defined network (SDN) that operates as a virtual overlay on top of existing networks, enabling organizations to build and manage private virtual networks with ease.

With Trustgrid you can:
- Deploy nodes worldwide to establish private network connectivity between sites, data centers, and public clouds
- Provide simplified, secured, and compliance-friendly edge connectivity for your remote sites, customers, vendors, or partners **requiring only outbound connectivity**
- Manage fleets of edge devices from a centralized cloud-based console for configuration, monitoring, updates, and more
- Gain visibility into traffic across your network with integrated flow logs and live troubleshooting tools

## Core Concepts

### Nodes

The basic building block of a Trustgrid network is a [node]({{<ref "docs/nodes">}}). Nodes are [deployed]({{<relref "./deployment">}}) on-premises or in your cloud provider and connected to each other to build a [data plane]([{{<relref "#data-plane">}}]) on top of which additional services such as [virtual networking]({{<relref "#virtual-networks">}}), [ztna access]({{<relref "#ztna-access">}}), and [compute]({{<relref "#edge-compute">}}) can be provided.


#### Gateways
While all nodes can act as edge nodes (making only outbound connections), appliance-based nodes can also be configured to act as [gateway servers]({{<relref "/docs/nodes/gateway/gateway-server">}}) which listen for inbound connections from other edge nodes in the account.  

These connections create the [private data plane]({{<relref "#data-plane">}}) which enable the Trustgrid network services.

### Data Plane

#### Virtual Networks

Nodes can be attached to [virtual networks]({{<ref "docs/domain/virtual-networks">}}), which provide a way to share network configuration like [routes]({{<ref "docs/domain/virtual-networks/routes">}}), NATs, and [ACLs]({{<ref "docs/domain/virtual-networks/access-policy">}}) at scale. 

#### Layer 4 Proxy
The data plane can also be used to enable Layer 4 proxy connections across two nodes.  In this setup, one node will act at the [connector]({{<relref "docs/nodes/connectors">}}) listening on a specified port.  When traffic is received on that port it will be forwarded to the configured [service]({{<relref "docs/nodes/services">}}). The node that hosts that service will then forward the traffic on to the configured IP and port using it's local interfaces.

#### Applications

[Applications]({{<ref "docs/applications">}}) can be exposed through nodes or clusters. Access to an application can be restricted via [access policies]({{<ref "docs/applications/access-policy">}}), for example to only allow users from a specific country.


### Control plane

All Trustgrid nodes are entirely managed through our control plane:

* Software updates are provided through our apt repository that ensures nodes have security updates available and only run with tested software permutations
* Network, node, user, and application configuration is managed through our web portal
* Configuration changes are broadcast to nodes as needed, for example when adding routes or adding a gateway to your network




