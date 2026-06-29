---
tags: ["aws"]
title: "Configure HA L4 Cluster in AWS"
description: "Use the cluster IP on L4 connectors or services to make a Trustgrid HA cluster present a stable IP across member failover"
linkTitle: "HA L4 Cluster in AWS"
type: docs
aliases:
  - /tutorials/deployments/deploy-aws-ami/configure-ha-l4-in-aws/
  - /tutorials/deployments/deploy-aws/configure-ha-l4-in-aws/
---

In a high-availability (HA) cluster, the cluster IP gives you a stable address that survives member failover. L4 connectors and services can both leverage it, depending on which side of the traffic flow the cluster sits on. Use one or both — they are independent.

- Connectors give clients a stable destination IP:port to connect to.
- Services give backends a stable source IP for connections originated from the cluster.

## Prerequisites

- A two-member [Trustgrid HA cluster]({{<ref "/docs/clusters">}}) deployed in AWS.
- A cluster IP configured on the LAN interface. See [Choose a failover mechanism](#choose-a-failover-mechanism) below.

## Choose a failover mechanism

Both modes use a cluster IP on the LAN interface for the stable address. What differs is where that IP lives and how it moves to the new active member on failover, based on whether your two members share an Availability Zone (AZ). The connector and service steps later in this guide are the same for both.

### Same AZ

The cluster IP is an unused private IP inside the data-NIC subnet. On failover it is reassigned as a secondary private IP to the new active member. This is the simpler option, with nothing to manage in the route table, but both members must sit in the same AZ. Configure it per [AWS Cluster IP Failover]({{<relref "ip-failover" >}}).

### Different AZs

A secondary private IP cannot move across AZs, because it is bound to one subnet and a subnet lives in one AZ. To get a stable IP that survives the loss of an entire AZ, the cluster IP is a VIP outside every VPC CIDR, and failover is handled by the route table. This mode combines two things:

1. A cluster IP set to the VIP on the LAN interface. Set it the same way as a normal [cluster IP]({{<relref "ip-failover#configuration" >}}), but use an address outside every VPC CIDR. No IAM permissions are needed for the cluster IP itself: the active member only binds it locally on the interface, nothing is changed in AWS.
1. [AWS Cluster Route Failover]({{<relref "route-failover" >}}), which keeps a `/32` route to the VIP pointing at the active member and moves it on failover. Follow that page for the route-table setup and the IAM permissions, which are only needed for managing the route.

Set both, then return here to configure the L4 connector or service.

## Connectors — Stable Destination IP for Clients

Use this pattern when clients should reach the cluster on a fixed IP:port. Set the connector's Listen Interface to **All** so it listens on every address on the member, including the cluster IP. On failover the cluster IP moves to the new active member, which is also listening, so clients reconnect to the same address.

1. In the Trustgrid portal, navigate to your cluster and select Connectors from the Networking menu.
1. Click Add Connector and configure:
   - Listen Interface — set to **All**.
   - Listen Port, Destination Node, Destination Service — set per the service being fronted.

{{<tgimg src="l4-add-connector.png" alt="Connector configuration dialog with Listen Interface set to All, listening on port 8080 and forwarding to a remote cluster" caption="Connector configured with Listen Interface = All" >}}

The configured connector appears in the cluster's Connectors list:

{{<tgimg src="l4-connectors-list.png" alt="Connectors list showing two TCP connectors" caption="Cluster Connectors list">}}

## Services — Stable Source IP Toward Backends

Use this pattern when the cluster originates connections to backends and the backend must see a single stable source IP (allowlists, return routing, audit logging). With Source Interface set to the LAN interface and Use Cluster IP selected, outbound connections are sourced from the cluster IP regardless of which member is currently active.

1. In the Trustgrid portal, navigate to your cluster and select Services from the Networking menu.
1. Click Add Service and configure:
   - Source Interface — set the interface dropdown to the LAN interface (typically `eth1`), then set the second dropdown to Use Cluster IP.
   - Host, Port, Protocol — set per the backend being reached.

{{<tgimg src="l4-add-service.png" alt="Service configuration dialog with Source Interface set to eth1 and Use Cluster IP selected, forwarding to a backend on TCP 8080" caption="Service configured with Source Interface = eth1 and Use Cluster IP">}}

The configured service appears in the cluster's Services list with the Source Interface column showing `eth1 (Use Cluster IP)`:

{{<tgimg src="l4-services-list.png" alt="Services list showing one service with Source Interface eth1 (Use Cluster IP)" caption="Cluster Services list">}}
