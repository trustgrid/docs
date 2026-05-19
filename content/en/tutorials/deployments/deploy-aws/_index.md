---
tags: ["aws"]
linkTitle: "Deploy to AWS"
title: "Deploy a Trustgrid Node in AWS"
no_list: true
aliases:
  - /tutorials/deployments/deploy-aws-ami/
---

Standing up a Trustgrid node in AWS uses a published Amazon Machine Image (AMI). Each node has two network interfaces — a WAN interface for control plane communication (and, on gateway nodes, inbound tunnel traffic from remote edges), and a LAN interface for internal data traffic.

## Pick a Deployment Method

Pick whichever fits your tooling and workflow:

- [**Terraform**]({{<relref "terraform">}}) — Apply the official Trustgrid Terraform modules. (Requires a license key up front — direct Trustgrid customers generate one via the portal or API; otherwise, request one from the Trustgrid provisioning team.)
- [**CloudFormation**]({{<relref "cloudformation">}}) — Launch a published CloudFormation template from the AWS console. (Requires a license key up front — direct Trustgrid customers generate one via the portal or API; otherwise, request one from the Trustgrid provisioning team.)
- [**Remote Registration**]({{<relref "remote-registration">}}) — Launch the AMI manually and register the node via the EC2 Serial Console. (No license key required up front; if you are not a direct Trustgrid customer, requires working with the Trustgrid provisioning team to register the node.)

Each walkthrough covers prerequisites, networking, the deployment steps, and verification.

## AWS Network Firewall and UDP Tunnels

If nodes are deployed behind an [AWS Network Firewall](https://aws.amazon.com/network-firewall/) and UDP tunnels are used, explicit rules must be added in both directions. Unlike TCP, AWS Network Firewall does not maintain UDP connection state during maintenance events. If the first packet the firewall sees after a maintenance event arrives from the opposite direction of the original flow (e.g., a gateway initiating a keepalive back toward an edge node), it will not recognize the tuple as an established session and will block it.

{{<alert color="warning">}}
This bidirectional rule requirement applies only when using UDP tunnels. TCP tunnels are not affected because AWS Network Firewall tracks TCP state normally.
{{</alert>}}

Gateway node — required rules:

| Direction | Source IP | Source Port | Destination IP | Destination Port |
|-----------|-----------|-------------|----------------|------------------|
| Inbound   | Remote edge node IPs (or `any`) | Any | Gateway IP | UDP 8443 (or configured gateway port) |
| Outbound  | Gateway IP | Any (ephemeral UDP source port) | Remote edge node IPs (or `any`) | Any |

Edge node — required rules:

| Direction | Source IP | Source Port | Destination IP | Destination Port |
|-----------|-----------|-------------|----------------|------------------|
| Outbound  | Edge node IP | Any | Known gateway IPs | UDP 8443 (or configured gateway port) |
| Inbound   | Known gateway IPs | UDP 8443 (or configured gateway port) | Edge node IP | Any |

## High Availability

Once a pair of nodes is deployed, you can cluster them using one of two failover mechanisms. Either supports L3 or L4 traffic patterns.

- [Cluster IP Failover]({{<relref "cluster-ip-failover">}}) — Claims a secondary private IP on the active member's LAN ENI. Does not require route-table changes.
- [VPC Route Failover]({{<relref "vpc-route-failover">}}) — Updates AWS route-table entries to point overlay CIDRs at the active member's ENI.

### Use Case Tutorials

- [HA L4 Cluster in AWS]({{<relref "l4-cluster">}}) — Using the cluster IP on L4 connectors and services so the cluster presents a stable IP across member failover.
- [HA Wireguard Cluster]({{<relref "wireguard-cluster">}}) — Fronting Wireguard listeners with an AWS Network Load Balancer.
