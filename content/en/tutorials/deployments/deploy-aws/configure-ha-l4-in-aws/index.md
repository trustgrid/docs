---
tags: ["aws"]
title: "Configure HA L4 Cluster in AWS"
description: "Use the cluster IP on L4 connectors or services to make a Trustgrid HA cluster present a stable IP across member failover"
linkTitle: "HA L4 Cluster in AWS"
type: docs
aliases:
  - /tutorials/deployments/deploy-aws-ami/configure-ha-l4-in-aws/
---

In an HA cluster, the cluster IP gives you a stable address that survives member failover. L4 connectors and services can both leverage it, depending on which side of the traffic flow the cluster sits on. Use one or both — they are independent.

- Connectors give clients a stable destination IP:port to connect to.
- Services give backends a stable source IP for connections originated from the cluster.

## Prerequisites

- A two-member [Trustgrid HA cluster]({{<ref "/docs/clusters">}}) deployed in AWS.
- A cluster IP configured on the LAN interface. See [AWS Cluster IP Failover]({{<ref "ip-failover">}}) for IAM permissions, source/destination check, and configuration steps.

## Connectors — Stable Destination IP for Clients

Use this pattern when clients should reach the cluster on a fixed IP:port. When a cluster IP is configured on the LAN interface, a connector with Listen Interface set to the LAN interface binds only to the cluster IP — not to the member's primary IP. On failover, the cluster IP migrates to the new active member and clients reconnect to the same address.

1. In the Trustgrid portal, navigate to your cluster and select Connectors from the Networking menu.
1. Click Add Connector and configure:
   - Listen Interface — the LAN interface where the cluster IP is configured (typically `eth1`).
   - Listen Port, Destination Node, Destination Service — set per the service being fronted.

{{<tgimg src="l4-add-connector.png" alt="Add Connector dialog with Listen Interface set to eth1" caption="Add Connector — Listen Interface = eth1 binds the connector to the cluster IP only">}}

The configured connector appears in the cluster's Connectors list:

{{<tgimg src="l4-connectors-list.png" alt="Connectors list showing two TCP connectors" caption="Cluster Connectors list">}}

## Services — Stable Source IP Toward Backends

Use this pattern when the cluster originates connections to backends and the backend must see a single stable source IP (allowlists, return routing, audit logging). With Source Interface set to the LAN interface and Use Cluster IP selected, outbound connections are sourced from the cluster IP regardless of which member is currently active.

1. In the Trustgrid portal, navigate to your cluster and select Services from the Networking menu.
1. Click Add Service and configure:
   - Source Interface — set the interface dropdown to the LAN interface (typically `eth1`), then set the second dropdown to Use Cluster IP.
   - Host, Port, Protocol — set per the backend being reached.

{{<tgimg src="l4-add-service.png" alt="Add Service dialog with Source Interface set to eth1 and Use Cluster IP selected" caption="Add Service — Source Interface = eth1, Use Cluster IP">}}

The configured service appears in the cluster's Services list with the Source Interface column showing `eth1 (Use Cluster IP)`:

{{<tgimg src="l4-services-list.png" alt="Services list showing one service with Source Interface eth1 (Use Cluster IP)" caption="Cluster Services list">}}
