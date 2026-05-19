---
tags: ["aws", "terraform"]
linkTitle: "Terraform"
title: "Deploy a Trustgrid Node in AWS via Terraform"
description: "Step-by-step Terraform walkthrough for deploying a Trustgrid node in AWS using the official trustgrid-infra-as-code modules."
---

This walkthrough deploys a Trustgrid node in AWS using the official [trustgrid-infra-as-code](https://github.com/trustgrid/trustgrid-infra-as-code) Terraform modules. The modules encapsulate the EC2 instance, two ENIs, EIP, source/destination check, and IAM wiring, and select the correct Trustgrid AMI based on the instance type.

## Step 1: Confirm Prerequisites

You will need:

- Terraform 1.5 or newer and the AWS provider 6.0 or newer.
- AWS credentials with permission to create EC2, EIP, security group, and IAM resources in the target region.
- A VPC with a public subnet for the WAN interface and a private subnet for the LAN interface.
- An SSH key pair in the target region. Direct SSH access is removed by the provisioning process — the key pair is only used if troubleshooting is required during initial deployment.
- A Trustgrid license key for each node you plan to deploy with the `trustgrid_single_node_auto_reg` module. If you are a direct Trustgrid customer, generate one via the portal or the Trustgrid API — see [Adding Node Appliances]({{<relref "/docs/nodes#adding-node-appliances---generating-licenses">}}) for the portal walkthrough. If you are not a direct Trustgrid customer, request the key from your vendor or the Trustgrid provisioning team. Save it to local storage — you cannot reissue a license without recreating the node.

The Trustgrid AMI is published in `us-east-1`, `us-east-2`, `us-west-1`, and `us-west-2`. Deploying in other regions requires working with [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) to copy the AMI into the target region.

## Step 2: Pick an Instance Type

The following x86_64 EC2 instance families are supported:

- t3, t3a, c5, c5n, c5a, c6i, c6in, c6a
- c7a, c7i, c8a, c8i

ARM-based instances (Graviton) are not supported. Contact [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) if a different type is believed necessary. The module reads `var.instance_type` and selects the matching AMI variant automatically — no extra configuration is required.

{{<alert>}}
If using a burstable performance instance type (T3, T3a):

- Set CPU Credits for all Gateway instances to unlimited so CPU can burst above the normal threshold. See [Unlimited mode for burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html).
- Configure monitoring of your CPU Credit Balance to alert if credits are being consumed or you are being charged for additional CPU usage. See [Monitor your CPU credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-monitoring-cpu-credits.html).
{{</alert>}}

## Step 3: Pick a Module

| Module | Path | What it does |
|--------|------|--------------|
| `trustgrid_single_node_auto_reg` | `aws/terraform/modules/trustgrid_single_node_auto_reg` | Provisions the EC2 instance, two ENIs, EIP, and source/destination-check settings, and registers the node with the Trustgrid control plane on first boot using a license key passed in as a variable. |
| `trustgrid_single_node_manual_reg` | `aws/terraform/modules/trustgrid_single_node_manual_reg` | Provisions the same AWS resources as `auto_reg` but does not register the node. After apply, register it via the EC2 Serial Console — see [Remote Registration]({{<relref "/tutorials/deployments/deploy-aws/remote-registration">}}) for the activation-code workflow. |
| `trustgrid_cluster_route_role` | `aws/terraform/modules/iam/trustgrid_cluster_route_role` | Creates the IAM role and policy required for HA [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}). Attach to either single-node module via `instance_profile_name`. |

## Step 4: Configure the Module

Reference the module from your Terraform configuration. For `trustgrid_single_node_auto_reg`:

```hcl
module "trustgrid_node" {
  source = "git::https://github.com/trustgrid/trustgrid-infra-as-code.git//aws/terraform/modules/trustgrid_single_node_auto_reg?ref=main"

  name                          = "tg-node-1"
  instance_type                 = "c7a.large"
  key_pair_name                 = "my-keypair"
  license                       = var.trustgrid_license

  management_subnet_id          = aws_subnet.public.id
  management_security_group_ids = [aws_security_group.wan.id]

  data_subnet_id                = aws_subnet.private.id
  data_security_group_ids       = [aws_security_group.lan.id]

  # Default is edge — set to true only if this node will terminate
  # inbound tunnels from remote edges (i.e., act as a gateway).
  # is_tggateway                = true
}
```

The module creates its own management security group that holds the optional gateway ingress rule (`is_tggateway`). The security groups you supply via `management_security_group_ids` and `data_security_group_ids` are attached alongside it and should provide egress to the Trustgrid control plane and any other rules required by your environment. See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the IPs and ports the WAN interface must reach.

For an HA pair, instantiate the module twice and add the `trustgrid_cluster_route_role` module (for VPC Route Failover) or attach an instance profile with `ec2:AssignPrivateIpAddresses` on the LAN ENI (for Cluster IP Failover). Pass the resulting instance profile name through the `instance_profile_name` variable. See [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}) for the full IAM policy.

## Step 5: Apply

```shell
terraform init
terraform plan
terraform apply
```

Terraform creates the management and data ENIs, allocates and associates an EIP on the management ENI, disables source/destination check on both ENIs, and launches the EC2 instance with the matching Trustgrid AMI. With `trustgrid_single_node_auto_reg`, the bootstrap script writes the license and calls the registration endpoint on first boot.

## Step 6: Verify

If you are not a direct Trustgrid customer, contact the Trustgrid provisioning team to confirm registration status.

If you are a direct Trustgrid customer, open the [Nodes page]({{<relref "/docs/nodes">}}) in the Trustgrid portal — the node appears once registration completes and shows as online when its control plane connection is established. From there it is managed like any other Trustgrid node.

## Next Steps

- Deploy a second node and set up HA — see [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}).
- Review the [AWS Network Firewall and UDP Tunnels]({{<relref "/tutorials/deployments/deploy-aws#aws-network-firewall-and-udp-tunnels">}}) caveat if the node sits behind an AWS Network Firewall.
