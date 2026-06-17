---
tags: ["aws"]
linkTitle: "Deploy to AWS"
title: "Deploy a Trustgrid Node in AWS"
no_list: true
aliases:
  - /tutorials/deployments/deploy-aws-ami/
---

Standing up a Trustgrid node in AWS uses a published Amazon Machine Image (AMI). Each node has two network interfaces — a WAN interface for control plane communication and TLS/UDP tunnel traffic, and a LAN interface for internal data traffic.

## Prerequisites

The Trustgrid AMI is published in us-east-1, us-east-2, us-west-1, and us-west-2. Deploying in other regions requires working with [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) to copy the AMI into the target region.

You will also need:

- A VPC with a public subnet for the WAN interface and a private subnet for the LAN interface.
- An available Elastic IP in the region — one EIP is associated with the WAN interface during provisioning.
- An SSH key pair in the target region for troubleshooting if required.

{{<alert>}}
Direct SSH access to the node is removed as part of the provisioning process. The SSH key pair is only used if troubleshooting is required during the initial deployment.
{{</alert>}}

### Instance Type

Trustgrid's AWS platform support is grouped into generations based on the EC2 instance types and their network interface naming. The generation a node uses is shown by its **Device Type** in the portal.

| Generation | Status | Portal Device Type | Interface Names | Example Instance Types | Architecture |
|------------|--------|--------------------|-----------------|------------------------|--------------|
| Gen1 | Deprecated | `AWS T2` | `eth0` (WAN) / `eth1` (LAN) | t2 | x86_64 only |
| Gen2 | Supported | `AWS T3` or `AWS C5` | `ens5` (WAN) / `ens6` (LAN) | t3, t3a, c5, c5n, c6i, c6in, c6a | x86_64 only |
| Gen3 | Supported | `AWS Gen3` | `enp39s0` (WAN) / `enp40s0` (LAN) | c7i, c7a, m8i, m8a | x86_64 only |

The gen3 platform requires the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}) or later and supports most current-generation x86_64 instance types. Deploy new nodes on Gen2 or Gen3 instance types; Gen1 is deprecated.

Additional x86_64 instance types may work but have not been tested. ARM-based instances (Graviton) are not supported. Contact [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) if a different type is believed necessary.

{{<alert>}}
If using a burstable performance instance type (T3, T3a):

- Set CPU Credits for all Gateway instances to unlimited so CPU can burst above the normal threshold. See [Unlimited mode for burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html).
- Configure monitoring of your CPU Credit Balance to alert if credits are being consumed or you are being charged for additional CPU usage. See [Monitor your CPU credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-monitoring-cpu-credits.html).
{{</alert>}}

## Networking

The WAN interface lives in the public subnet and must allow outbound access to the Trustgrid control plane. See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the full list of required IPs and ports.

{{<alert color="warning">}}
Inbound TCP/UDP 8443 on the WAN interface is only required if the node will act as a gateway (a node that terminates tunnels from remote edge nodes). Edge nodes do not require any inbound WAN security group rules.
{{</alert>}}

The LAN interface lives in the private subnet and is used for internal data traffic and, in clustered deployments, for inter-node communication on the cluster heartbeat port (typically TCP 9000).

The following table summarizes the security group rules required across both interfaces:

| Interface | Direction | Protocol | Ports     | Source/Destination | Purpose |
|-----------|-----------|----------|-----------|--------------------|---------|
| WAN       | Egress    | TCP      | 443, 8443 | Trustgrid control plane IPs — see [Network Requirements]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) | Control plane communication |
| WAN       | Egress    | TCP      | 443       | AWS API endpoints — see [AWS IP ranges](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html) | EC2 API calls for cluster failover |
| WAN       | Ingress   | TCP/UDP  | 8443      | 0.0.0.0/0 (or known edge IPs) | TLS/UDP tunnel traffic (gateway nodes only) |
| LAN       | Ingress/Egress | TCP | 9000      | LAN subnet CIDR | Cluster heartbeat (clustered nodes only) |

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

## Security

Source/destination check must be disabled on both interfaces of every Trustgrid node. The CloudFormation and Terraform deployment paths configure this automatically; for [Remote Registration](#remote-registration) deployments it must be disabled manually after the instance is launched.

An IAM role is only required if the node will be deployed as part of an HA cluster using one of the failover mechanisms below. Each tutorial documents the specific permissions required:

- [IP Failover]({{<relref "ip-failover">}}) — `ec2:AssignPrivateIpAddresses` on the LAN ENI.
- [Route Failover]({{<relref "route-failover">}}) — `ec2:DescribeRouteTables`, `ec2:CreateRoute`, `ec2:DeleteRoute` on the LAN route table.

---

## Deployment Methods

Choose one of the following methods to deploy and register the node.

### Terraform

The [trustgrid-infra-as-code](https://github.com/trustgrid/trustgrid-infra-as-code) repository provides purpose-built Terraform modules for deploying Trustgrid nodes in AWS. These modules encapsulate the EC2 instance, two ENIs, EIP, source/destination check, and IAM wiring.

| Module | Path | Purpose |
|--------|------|---------|
| `trustgrid_single_node_auto_reg` | `aws/terraform/modules/compute/trustgrid_single_node_auto_reg` | Deploys a node with two ENIs and automatic registration via license key |
| `trustgrid_single_node_manual_reg` | `aws/terraform/modules/compute/trustgrid_single_node_manual_reg` | Deploys a node with two ENIs without a license key, to be registered after launch via [Remote Registration](#remote-registration) |
| `trustgrid_cluster_route_role` | `aws/terraform/modules/iam/trustgrid_cluster_route_role` | Creates and binds the IAM role required for HA cluster route failover |

### CloudFormation

In this path, you create the node in the Trustgrid portal first to obtain a license key, then pass the key as a parameter to the CloudFormation stack. The node registers automatically with the Trustgrid control plane on first boot.

#### Step 1: Add the Node and Obtain a License Key

In the Trustgrid portal, go to the [Nodes page]({{<relref "/docs/nodes">}}), click **Add Node**, enter a name for the node, and click **Create License**. The portal generates a license key — copy it to your clipboard or click **Download License** to save it locally. See [Adding Node Appliances]({{<relref "/docs/nodes#adding-node-appliances---generating-licenses">}}) for the full walkthrough with screenshots.

{{<alert>}}
Generating a license requires a Trustgrid portal account. If you are not a direct Trustgrid customer, work with your vendor or contact [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) to have a license generated for you.
{{</alert>}}

{{<alert>}}
The node will not appear in the portal until the EC2 instance successfully completes registration with the Trustgrid control plane. You cannot reissue a license without recreating the node, so download the license to local storage in case the clipboard is cleared.
{{</alert>}}

#### Step 2: Launch the CloudFormation Stack

Open the CloudFormation template in the AWS console:

[https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node.json](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node.json)

Change the `region=` query parameter in the URL to the region you are deploying in. All template fields are required.

{{<fields>}}
{{% field "Stack Name" %}}
A unique name describing this deployment.
{{% /field %}}
{{% field "Instance Type" %}}
The EC2 instance type. See [Instance Type](#instance-type) above.
{{% /field %}}
{{% field "SSH Keypair" %}}
SSH keypair used to SSH to the instance as the `ubuntu` user if necessary.

> SSH access requires a security group rule allowing inbound port 22. We strongly recommend SSH not be allowed from `0.0.0.0/0`.
{{% /field %}}
{{% field "Host IAM Role" %}}
Optional. Only required if the node will participate in an HA cluster — see [Security](#security).
{{% /field %}}
{{% field "WAN — Security Group / Subnet" %}}
The security group and public VPC subnet for the WAN interface. The EIP created by the template is associated with the interface in this subnet.
{{% /field %}}
{{% field "LAN — Security Group / Subnet" %}}
The security group and private VPC subnet for the LAN interface.
{{% /field %}}
{{% field "Trustgrid License" %}}
Paste the license key from the portal (see Step 1 above). It is critical that the license is copied exactly.
{{% /field %}}
{{</fields>}}

When creating the stack, check the box acknowledging that AWS CloudFormation may create IAM resources — this is required because the template creates an instance profile for the EC2 instance.

The CloudFormation template configures an encrypted EBS volume. If you supply a custom IAM role, it must include access to the default EBS KMS key:

```json
{
  "Effect": "Allow",
  "Action": [
    "kms:Decrypt",
    "kms:DescribeKey",
    "kms:ReEncrypt*",
    "kms:GenerateDataKey*"
  ],
  "Resource": "arn:aws:kms:REGION:ACCOUNT_ID:alias/aws/ebs"
}
```

Once registration is complete, the node appears as online in the portal and is ready to manage like any other.

### Remote Registration

In this path, you launch an EC2 instance from the Trustgrid AMI without a license key and then register the node using the Trustgrid remote registration utility over SSH. Use this approach if you need to register a node without pre-generating a license key from the portal.

#### Step 1: Launch the EC2 Instance

Launch an EC2 instance from the Trustgrid AMI in the target region. The AMI is published with a name prefixed `trustgrid-node-prod` — search the AMI catalog by that prefix and select the most recent version. Configure the instance with:

- The instance type from [Instance Type](#instance-type) above.
- A primary network interface in the public subnet with the WAN security group attached and an Elastic IP associated.
- A secondary network interface in the private subnet with the LAN security group attached.
- Source/destination check disabled on both interfaces.
- An IAM instance profile attached if the node will participate in an HA cluster — see [Security](#security) for the permissions required by IP Failover or Route Failover.
- The SSH key pair for the region.

#### Step 2: Register via the EC2 Serial Console

Once the instance is running, open the [EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-serial-console.html) for the instance from the AWS console and log in to the Trustgrid local console utility. From there, initiate the remote registration process — the console will generate a short activation code that someone with Trustgrid portal access can use to license the node.

{{<alert color="warning">}}
Completing remote registration requires access to the Trustgrid portal. This step must be performed by someone with portal access — either the end customer or [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}).
{{</alert>}}

See [Remote Registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) for full instructions.

---

## High Availability

High availability for a Trustgrid cluster in AWS is achieved using one of two failover mechanisms. Either mechanism can support L3 or L4 traffic patterns.

- [IP Failover]({{<relref "ip-failover">}}) — Claims a secondary private IP on the active member's LAN ENI. Does not require route-table changes.
- [Route Failover]({{<relref "route-failover">}}) — Updates AWS route-table entries to point overlay CIDRs at the active member's ENI.

### Use Case Tutorials

- [HA L4 Cluster in AWS]({{<relref "l4-cluster">}}) — Using the cluster IP on L4 connectors and services so the cluster presents a stable IP across member failover.
- [HA Wireguard Cluster]({{<relref "wireguard-cluster">}}) — Fronting Wireguard listeners with an AWS Network Load Balancer.
