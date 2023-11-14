---
title: "Getting Started"
linkTitle: "Getting Started"
menu:
  main:
    weight: 10
cascade:
  type: docs
---

Trustgrid is a platform that provides secure, reliable and high-performance connectivity between distributed systems and applications. The platform is designed to simplify networking for distributed enterprises by providing a software-defined network (SDN) that operates as a virtual overlay on top of existing networks, enabling organizations to build and manage their own private networks.


### Virtual Networks

Nodes can be attached to [virtual networks]({{<ref "docs/domain/virtual-networks">}}), which provide a way to share network configuration like [routes]({{<ref "docs/domain/virtual-networks/routes">}}), NATs, and [ACLs]({{<ref "docs/domain/virtual-networks/access-policy">}}) at scale. 

### Applications

[Applications]({{<ref "docs/applications">}}) can be exposed through nodes or clusters. Access to an application can be restricted via [access policies]({{<ref "docs/applications/access-policy">}}), for example to only allow users from a specific country.

## Management

All Trustgrid nodes are entirely managed through our control plane:

* Software updates are provided through our apt repository that ensures nodes have security updates available and only run with tested software permutations
* Network, node, user, and application configuration is managed through our web portal
* Configuration changes are broadcast to nodes as needed, for example when adding routes or adding a gateway to your network
