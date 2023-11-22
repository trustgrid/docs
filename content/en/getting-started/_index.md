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
- Deploy nodes worldwide to establish private network connectivity between sites, data centers, and clouds
- Provide simplified, secured, and compliance-friendly edge connectivity for your remote sites, customers, vendors, or partners **requiring only outbound connectivity**
- Manage fleets of edge devices from a centralized cloud-based console for configuration, monitoring, updates, and more
- Gain visibility into traffic flows across your network with integrated flow logs

## Core Concepts

### Nodes

The basic building block of a Trustgrid network is a [node]({{<ref "docs/nodes">}}). Nodes are [deployed]({{<relref "/tutorials/deployments/">}}) on-premises or in your cloud provider and connected to each other to build a [data plane]([{{<relref "#data-plane">}}]) on top of which additional services such as [virtual networking]({{<relref "#virtual-networks">}}), [ztna access]({{<relref "#ztna-access">}}), and [compute]({{<relref "#edge-compute">}}) can be provided.

Nodes can be deployed in a variety of ways, including as virtual machines, containers (as [agents]({{<relref "/tutorials/agent-deploy">}})), or bare metal servers.


### Virtual Networks

Nodes can be attached to [virtual networks]({{<ref "docs/domain/virtual-networks">}}), which provide a way to share network configuration like [routes]({{<ref "docs/domain/virtual-networks/routes">}}), NATs, and [ACLs]({{<ref "docs/domain/virtual-networks/access-policy">}}) at scale. 

### Applications

[Applications]({{<ref "docs/applications">}}) can be exposed through nodes or clusters. Access to an application can be restricted via [access policies]({{<ref "docs/applications/access-policy">}}), for example to only allow users from a specific country.

## Management

All Trustgrid nodes are entirely managed through our control plane:

* Software updates are provided through our apt repository that ensures nodes have security updates available and only run with tested software permutations
* Network, node, user, and application configuration is managed through our web portal
* Configuration changes are broadcast to nodes as needed, for example when adding routes or adding a gateway to your network
