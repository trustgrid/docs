---
categories: ["overview"]
tags: ["architecture"]
title: "Basic Architecture"
date: 2022-12-20
description: >
  Trustgrid basic architecture overview
---

## Cloud Components

- Portal: The cloud management UI.
- Gatekeeper: Provides configuration updates to edge nodes.
- Zuul: Maintains persistent connection with nodes for reporting, events, and updates.
- API: The management API that exposes 100% of UI elements to automation.
- Repo: The APT repository that stores all firmware, OS, and Node updates.

## Edge Components

- [Node]({{< ref "/docs/concepts/node" >}}): The software that provides core functionality in the edge including networking, security, compute, and management features. Edge Nodes build outbound connections to Gateway Nodes.

## Use Cases

- Software Defined Networking: Create a mesh network that connects cloud applications to edge data with load balancing, [clustering]({{< ref "/docs/concepts/cluster" >}}), and failover managed through a portal or API.
- Edge Compute: Deploy applications to the edge to access datasets not appropriate for replication to the cloud due to security or compliance concerns, latency, or cost.
- Device Management: Manage thousands of nodes with advanced tools to reduce the burden of operations at enterprise scale.
- Edge API: Integrate thousands of edge datasets with a single API interface and ETL functions executing at the edge.
- Security: Leverage Trustgrid's advanced security to protect against a wide range of threats.
