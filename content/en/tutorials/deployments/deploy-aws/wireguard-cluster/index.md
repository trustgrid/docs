---
title: "AWS HA Wireguard Cluster"
description: "Front a Trustgrid HA cluster with an AWS Network Load Balancer to terminate Wireguard connections"
linkTitle: "AWS HA Wireguard Cluster"
aliases:
  - /tutorials/deployments/deploy-aws-ami/aws-ha-cluster/
  - /tutorials/deployments/deploy-aws/aws-ha-cluster/
---

This tutorial covers fronting a clustered Trustgrid node in AWS with a Network Load Balancer (NLB) so that Wireguard clients connect to a single stable endpoint and the NLB directs traffic to the active cluster member.

Only the active member of a Trustgrid cluster responds healthy to the load balancer health check, so the NLB always forwards traffic to the active node.

## Prerequisites

- A two-member [Trustgrid HA cluster]({{<ref "/docs/clusters">}}) deployed in AWS.
- A configured [Wireguard Gateway]({{<ref "/docs/nodes/appliances/wireguard-gateway">}}) on the cluster.

## Configuration

### 1. Create the EC2 Target Group

Create a target group containing both Trustgrid cluster members. Configure the target group as follows:

- Target port: the Wireguard server port (default UDP 51820).
- Health check protocol: HTTP.
- Health check path: `/status`.
- Health check port override: 80.

{{<tgimg src="wg-ha.png" alt="EC2 target group with both cluster members on UDP 51820" caption="Target group with both cluster members targeted on the Wireguard port">}}

{{<tgimg src="health-checks.png" alt="Target group health check configured for HTTP /status on port 80" caption="Health check — HTTP /status on port 80">}}

### 2. Create the Network Load Balancer

Create an internet-facing IPv4 Network Load Balancer mapped to the public subnets of the cluster members. Add a UDP listener on the Wireguard server port (default 51820) that forwards to the target group created above.

{{<tgimg src="edit-listener.png" alt="NLB UDP listener forwarding to the Wireguard target group" caption="NLB UDP listener forwarding to the target group">}}

Wireguard clients should be configured to connect to the NLB DNS name (or the Elastic IPs assigned to it). On cluster failover, the NLB stops receiving healthy responses from the previously active member and begins forwarding traffic to the new active member with no client-side change required.
