---
tags: ["aws"]
title: "Configure HA L4 Cluster in AWS"
description: "Configure a high availability L4 proxy cluster in AWS using a stable cluster IP that survives member failover"
linkTitle: "HA L4 Cluster in AWS"
type: docs
---

This tutorial walks through setting up a two-member Trustgrid cluster in AWS for L4 proxy (connector/service) traffic, where clients and backends always see a consistent IP regardless of which cluster member is active.

The cluster IP mechanism is explained in [AWS Cluster IP Failover]({{<relref "cluster-ip-failover-in-aws">}}). This page focuses on the end-to-end L4 use case.

## Scenario

```
LAN clients
    │  (connect to cluster IP)
    ▼
[Active cluster member]  ──── TrustGrid tunnel ────  [Remote cluster]
[Standby cluster member]                               │
                                                       ▼
                                                 Backend service
                                          (sees cluster IP as source)
```

- A two-member Trustgrid cluster in AWS acts as the local cluster.
- A cluster IP is configured on the LAN interface and becomes a secondary private IP on the active member's data ENI.
- L4 connectors on the local cluster accept client connections on the cluster IP and forward them across the Trustgrid tunnel to the remote side.
- L4 services on the remote cluster forward traffic to backend services with Source from cluster IP enabled, so the backend always sees the same stable source IP regardless of which local cluster member is active.

## Prerequisites

- Two-member Trustgrid HA cluster in AWS with cluster heartbeat configured. See [Configure HA Gateway Cluster in AWS]({{<relref "configure-ha-gateway-cluster-in-aws">}}) for cluster setup.
- Cluster IP configured on the LAN interface. See [AWS Cluster IP Failover]({{<relref "cluster-ip-failover-in-aws">}}) for IAM prerequisites, source/dest check, and configuration steps.
- Trustgrid VPN tunnel established between the local cluster and the remote cluster.

## Configure L4 Connectors (Local Cluster)

L4 connectors define which traffic the local cluster intercepts and forwards across the tunnel.

1. Navigate to the local cluster in the Trustgrid portal.
1. Select L4 Proxy > Connectors.
1. Create a connector for each backend service. Set the Listen IP to the cluster IP so that the connector listens on the stable floating address. Set the Listen Port, Destination, and Destination Port for the service.

LAN clients connect to `<cluster-ip>:<listen-port>`. The active cluster member accepts the connection and proxies it through the tunnel to the remote side.

## Configure L4 Services (Remote Cluster)

L4 services on the remote cluster define how incoming tunneled connections are forwarded to backend hosts.

1. Navigate to the remote cluster in the Trustgrid portal.
1. Select L4 Proxy > Services.
1. Create a service for each backend. Enable Source from cluster IP so that the remote cluster sources its connection to the backend from the remote cluster's own cluster IP, not from the individual member's primary IP.

With Source from cluster IP enabled:
- Backends see a predictable, stable source IP across failover events on either side.
- Backend allowlists, firewall rules, and audit logs can pin to a single IP per cluster.

## Failover Behavior

Local cluster failover (active member stops or loses quorum):
- The standby member promotes itself and calls `ec2:AssignPrivateIpAddresses(AllowReassignment=true)` to migrate the secondary cluster IP to its own data ENI.
- New client connections to the cluster IP succeed within seconds.
- Existing TCP sessions are torn down; clients reconnect automatically if the application supports reconnection.
- The tunnel to the remote cluster re-establishes via the new active member.

Remote cluster failover:
- The remote cluster's active member also migrates its cluster IP to the new active member's ENI.
- Backends continue to see the same source IP (the remote cluster IP) as before.

What backends observe across any failover event: the source IP remains the remote cluster's cluster IP throughout. The backend cannot tell which physical member handled the connection.

## Verified Behavior

The following was confirmed on the Trustgrid test tenant (`20260504-220203.c086493`):

- Cluster IP changes are picked up live by L4 services — no service restart required.
- Cluster IP changes made while the configured-active member is stopped are reconciled cleanly on restart at both the AWS ENI and OS levels.
- Source IPs visible at the backend follow the current cluster IP across changes and failover events.

## Related

- [AWS Cluster IP Failover]({{<relref "cluster-ip-failover-in-aws">}}) — underlying mechanism, IAM requirements, and configuration reference
- [Configure HA Gateway Cluster in AWS (Route Failover)]({{<relref "configure-ha-gateway-cluster-in-aws">}}) — L3 route-based failover alternative
- [Cluster-Only Configuration Items]({{<relref "/docs/clusters/cluster-only-config">}}) — cluster IP reference
