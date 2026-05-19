---
tags: ["aws", "cloudformation"]
linkTitle: "CloudFormation"
title: "Deploy a Trustgrid Node in AWS via CloudFormation"
description: "Step-by-step CloudFormation walkthrough for deploying a Trustgrid node in AWS using the published Trustgrid templates."
---

This walkthrough deploys a Trustgrid node in AWS using a published CloudFormation template. The template provisions the EC2 instance, management and data ENIs, EIP, and instance profile, and looks up the latest Trustgrid AMI automatically.

## Step 1: Confirm Prerequisites

You will need:

- AWS console (or CLI) access with permission to create EC2, EIP, security group, IAM, and Lambda resources in the target region.
- A VPC with a public subnet for the WAN interface and a private subnet for the LAN interface.
- A security group for each interface. The template attaches the security groups you provide to the ENIs as-is — make sure they allow egress to the Trustgrid control plane and any inbound traffic the node needs (e.g., TCP/UDP 8443 for gateway nodes). See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the IPs and ports.
- An SSH key pair in the target region. Direct SSH access is removed by the provisioning process — the key pair is only used if troubleshooting is required during initial deployment.

The Trustgrid AMI is published in `us-east-1`, `us-east-2`, `us-west-1`, and `us-west-2`. Deploying in other regions requires working with [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) to copy the AMI into the target region.

## Step 2: Obtain a License Key

The CloudFormation template takes the Trustgrid license key as a parameter, so you need it before launching the stack.

- **Non-Trustgrid customers** — Contact the Trustgrid provisioning team to obtain a license key.
- **Direct Trustgrid customers** — Generate the key via the portal/API. For the portal flow, go to the [Nodes page]({{<relref "/docs/nodes">}}), click **Add Node**, enter a name, and click **Create License**. Copy the key or click **Download License** to save it locally. See [Adding Node Appliances]({{<relref "/docs/nodes#adding-node-appliances---generating-licenses">}}) for the full walkthrough.

{{<alert>}}
You cannot reissue a license without recreating the node, so save it to local storage in case the clipboard is cleared. The node will not appear in the portal until the EC2 instance successfully completes registration with the Trustgrid control plane.
{{</alert>}}

## Step 3: Pick a Template

Trustgrid publishes two CloudFormation templates. Pick the one that matches the instance family you want to deploy — each template restricts the **Instance Type** dropdown to the families it supports and selects the matching AMI automatically.

| Supported instance families | Template URL |
|----------------------------|--------------|
| t3, t3a, c5, c5n, c5a, c6i, c6in, c6a | [Launch Trustgrid Node — Gen2](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node.json) |
| c7a, c7i, c8a, c8i | [Launch Trustgrid Node — Gen3](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node-gen3.json) |

ARM-based instances (Graviton) are not supported. Change the `region=` query parameter in the URL to the region you are deploying in.

{{<alert>}}
If using a burstable performance instance type (T3, T3a):

- Set CPU Credits for all Gateway instances to unlimited so CPU can burst above the normal threshold. See [Unlimited mode for burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html).
- Configure monitoring of your CPU Credit Balance to alert if credits are being consumed or you are being charged for additional CPU usage. See [Monitor your CPU credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-monitoring-cpu-credits.html).
{{</alert>}}

## Step 4: Fill in the Stack Parameters

All template fields are required.

{{<fields>}}
{{% field "Stack Name" %}}
A unique name describing this deployment.
{{% /field %}}
{{% field "Instance Type" %}}
The EC2 instance type, picked from the template's allowed list (see Step 3).
{{% /field %}}
{{% field "SSH Keypair" %}}
SSH keypair used to SSH to the instance as the `ubuntu` user if necessary.

> SSH access requires a security group rule allowing inbound port 22. We strongly recommend SSH not be allowed from `0.0.0.0/0`.
{{% /field %}}
{{% field "Host IAM Role" %}}
The IAM role attached to the instance profile the template creates. Only required if the node will participate in an HA cluster — see [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}) for the policy each mechanism needs.
{{% /field %}}
{{% field "WAN — Security Group / Subnet" %}}
The security group and public VPC subnet for the WAN interface. The EIP created by the template is associated with the interface in this subnet.
{{% /field %}}
{{% field "LAN — Security Group / Subnet" %}}
The security group and private VPC subnet for the LAN interface.
{{% /field %}}
{{% field "Trustgrid License" %}}
Paste the license key from Step 2. It is critical that the license is copied exactly.
{{% /field %}}
{{</fields>}}

## Step 5: Launch the Stack

When creating the stack, check the box acknowledging that AWS CloudFormation may create IAM resources — the template creates an instance profile for the EC2 instance.

The template configures an encrypted EBS volume. If you supply a custom IAM role, it must include access to the default EBS KMS key:

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

## Step 6: Verify

If you are not a direct Trustgrid customer, contact the Trustgrid provisioning team to confirm registration status.

If you are a direct Trustgrid customer, open the [Nodes page]({{<relref "/docs/nodes">}}) in the Trustgrid portal — the node appears once registration completes and shows as online when its control plane connection is established. From there it is managed like any other Trustgrid node.

## Next Steps

- Deploy a second node from the same template and set up HA — see [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}).
- Review the [AWS Network Firewall and UDP Tunnels]({{<relref "/tutorials/deployments/deploy-aws#aws-network-firewall-and-udp-tunnels">}}) caveat if the node sits behind an AWS Network Firewall.
