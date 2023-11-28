---
categories: ["overview"]
weight: 1
tags: ["architecture"]
title: "Basic Architecture"
date: 2022-12-28
description: >
  Trustgrid basic architecture overview
---
## Concepts

### Nodes

The basic building block of a Trustgrid network is a [node]({{<ref "docs/nodes">}}). Nodes are deployed on-premises or in your cloud provider and are connected to each other with TLS tunnels. Nodes can be deployed in a variety of ways, including as virtual machines, containers (as [agents]({{<relref "/tutorials/agent-deploy">}})), or bare metal servers.

Nodes are used to enable connectivity and access for different use cases, including:

* ZTNA applications, allowing fine-grained network control for access to business applications or servers
* VPN-like functionality, allowing network address translation (NAT) and routing between nodes
* Remote container execution, allowing an admin to deploy and manage containers in edge networks
* FTP (remote or in a data center) and other layer 4 protocols

#### Clusters

Nodes can be grouped into [clusters]({{<ref "docs/clusters">}}) to share configuration and provide high-availability.


### Control Plane 

The Trustgrid control plane is used for all management and configuration of Trustgrid [nodes]({{<ref "docs/nodes" >}}). It consists of:

- Portal - The cloud management UI.
- API - The management API that exposes 100% of UI elements to automation.
- Gatekeeper - Provides configuration updates to edge nodes via a REST endpoint
- Zuul - Maintains persistent connection with nodes for [reporting]({{<relref "/docs/operations/">}}), [events]({{<ref "docs/alarms/events" >}}), and [updates]({{<relref "/tutorials/management-tasks/upgrade-nodes">}}).
- Repo - The APT repository that stores all firmware, OS, and node software updates.


## Use Cases

- Software Defined Networking - Create a mesh [network]({{<ref "getting-started/networking">}}) that connects cloud applications to edge data with load balancing, [clustering]({{<ref "/docs/clusters">}}), and failover managed through a portal or API.
- Edge Compute - Deploy applications to the edge to access datasets not appropriate for replication to the cloud due to [security]({{<ref "getting-started/security">}}) or compliance concerns, latency, or cost.
- Device Management - Manage thousands of [nodes]({{<ref "docs/nodes">}}) with advanced tools to reduce the burden of operations at enterprise scale.
- Edge API - Integrate thousands of edge datasets with a single API interface and ETL functions executing at the edge.
- Security - Leverage Trustgrid's advanced [security]({{<ref "getting-started/security" >}}) to protect against a wide range of threats.
