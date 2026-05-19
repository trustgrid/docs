---
tags: ["aws"]
linkTitle: "Remote Registration"
title: "Deploy a Trustgrid Node in AWS via Remote Registration"
description: "Step-by-step walkthrough for deploying a Trustgrid node in AWS by launching the AMI manually and registering via the EC2 Serial Console."
---

This walkthrough deploys a Trustgrid node in AWS by launching the AMI manually from the EC2 console, then registering the node via the EC2 Serial Console using the Trustgrid local console utility. The node generates an activation code after launch that is redeemed to license it.

## Step 1: Confirm Prerequisites

You will need:

- AWS console (or CLI) access with permission to launch EC2 instances, create network interfaces, allocate EIPs, and (for HA) create IAM instance profiles in the target region.
- A VPC with a public subnet for the WAN interface and a private subnet for the LAN interface.
- A security group on each interface that allows the traffic in [Step 3](#step-3-configure-networking) below. See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the IPs and ports the WAN interface must reach.
- An available Elastic IP in the region for the WAN interface, if the node will act as a gateway (terminate inbound tunnels from remote edges). Edge nodes can run without a public IP as long as outbound NAT to the Trustgrid control plane is available.
- An SSH key pair in the target region. Direct SSH access is removed by the provisioning process — the key pair is only used if troubleshooting is required during initial deployment.
- Someone with Trustgrid portal access who can redeem the activation code generated in Step 5. If you are a direct Trustgrid customer, this is you. If you are not a direct Trustgrid customer, this is your vendor or the Trustgrid provisioning team — pass the activation code to them when prompted.

The Trustgrid AMI is published in `us-east-1`, `us-east-2`, `us-west-1`, and `us-west-2`. Deploying in other regions requires working with [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) to copy the AMI into the target region.

## Step 2: Pick an Instance Type and AMI

The following x86_64 EC2 instance families are supported. The AMI you launch depends on the family you choose:

| Instance families | AMI name prefix |
|-------------------|-----------------|
| t3, t3a, c5, c5n, c5a, c6i, c6in, c6a | `trustgrid-node-2204-` |
| c7a, c7i, c8a, c8i | `trustgrid-node-gen3-2204-` |

ARM-based instances (Graviton) are not supported. Contact [Trustgrid Support]({{<relref "/help-center/trustgrid-support">}}) if a different type is believed necessary.

In the EC2 console, search the AMI catalog (Public images, owned by Trustgrid) by the prefix matching your chosen instance family and select the most recent image.

{{<alert>}}
If using a burstable performance instance type (T3, T3a):

- Set CPU Credits for all Gateway instances to unlimited so CPU can burst above the normal threshold. See [Unlimited mode for burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html).
- Configure monitoring of your CPU Credit Balance to alert if credits are being consumed or you are being charged for additional CPU usage. See [Monitor your CPU credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-monitoring-cpu-credits.html).
{{</alert>}}

## Step 3: Configure Networking

Create the network interfaces and security group rules before launching the instance. The Trustgrid node uses two ENIs.

WAN interface (public subnet):

| Direction | Protocol | Ports | Source/Destination | Purpose |
|-----------|----------|-------|--------------------|---------|
| Egress    | TCP      | 443, 8443 | Trustgrid control plane IPs — see [Network Requirements]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) | Control plane communication |
| Egress    | TCP      | 443   | AWS API endpoints — see [AWS IP ranges](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html) | EC2 API calls for cluster failover |
| Ingress   | TCP/UDP  | 8443  | `0.0.0.0/0` (or known edge IPs) | TLS/UDP tunnel traffic (gateway nodes only) |

LAN interface (private subnet):

| Direction | Protocol | Ports | Source/Destination | Purpose |
|-----------|----------|-------|--------------------|---------|
| Ingress/Egress | TCP | 9000 | LAN subnet CIDR | Cluster heartbeat (clustered nodes only) |

Source/destination check must be disabled on both ENIs after they are created — Remote Registration does not configure this for you.

## Step 4: Launch the EC2 Instance

In the EC2 console, launch a new instance from the AMI selected in Step 2 with:

- The instance type from Step 2.
- A primary network interface in the public subnet with the WAN security group attached. For gateway nodes, associate the Elastic IP from your prerequisites with this interface; edge nodes can skip this if outbound NAT is available.
- A secondary network interface in the private subnet with the LAN security group attached.
- Source/destination check disabled on both interfaces.
- The SSH key pair for the region.
- An IAM instance profile attached if the node will participate in an HA cluster — see [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}) for the policy each mechanism needs.

## Step 5: Register via the EC2 Serial Console

Once the instance is running, open the [EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-serial-console.html) for the instance from the AWS console and log in to the Trustgrid local console utility. From there, initiate the remote registration process — the console will generate a short activation code.

Redeeming the activation code requires Trustgrid portal access. If you are a direct Trustgrid customer, redeem it yourself. Otherwise, pass the code to your vendor or the Trustgrid provisioning team and they will license the node on your behalf.

See [Remote Registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) for full instructions on the activation code workflow.

## Step 6: Verify

If you are not a direct Trustgrid customer, contact the Trustgrid provisioning team to confirm registration status.

If you are a direct Trustgrid customer, open the [Nodes page]({{<relref "/docs/nodes">}}) in the Trustgrid portal once registration is complete — the node shows as online when its control plane connection is established and is managed like any other Trustgrid node.

## Next Steps

- Deploy a second node and set up HA — see [Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/cluster-ip-failover">}}) or [VPC Route Failover]({{<relref "/tutorials/deployments/deploy-aws/vpc-route-failover">}}).
- Review the [AWS Network Firewall and UDP Tunnels]({{<relref "/tutorials/deployments/deploy-aws#aws-network-firewall-and-udp-tunnels">}}) caveat if the node sits behind an AWS Network Firewall.
