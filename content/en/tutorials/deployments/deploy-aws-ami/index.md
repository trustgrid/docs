---
tags: ["aws"]
title: "Deploy a Trustgrid Node AMI in AWS"
date: 2023-02-09
---

Standing up a Trustgrid node in AWS is easy using an Amazon AMI. Trustgrid nodes in AWS use two network interfaces - a management and a data interface. The management interface communicates with Trustgrid Cloud Management systems. The data interface is used to terminate TLS tunnels from Edge Nodes.

## Notes

- The cloudformation template below works with an AMI currently published in US-EAST-1/2 and US-WEST-1/2. Deploying in other regions requires working with Trustgrid Support
- Requires VPC and public subnet
- Does not create security groups or roles - those have to be managed separately (more below)

{{<alert>}}
If using a burstable performance instance types (T2, T3 and T3a) the following is advised:

- Set CPU Credits for all Gateway instances to unlimited to allow CPU to burst in the event there is a spike above the normal threshold. [Unlimited mode for burstable performance instances - Amazon Elastic Compute Cloud](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html)

- Configure monitoring of your CPU Credit Balance to alert if your credits are being consumed or you are being charged for additional CPU usage which might warrant resizing your devices. [Monitor your CPU credits - Amazon Elastic Compute Cloud](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-monitoring-cpu-credits.html)
{{</alert>}}

## Prerequisites

- VPC with public and private subnets - Management NIC goes in the public subnet, Data NIC goes in the private subnet
  - Note: If doing a multi-AZ cluster deployment the private subnets need to use the same route table for automated route management to work
- Security group for management NIC that allows the following traffic:

  - Inbound traffic on designated Trustgrid gateway port (typical TCP 8443) for remote nodes. Access to this port can be secured to only allow access from remote nodes if desired. This is only required if deploying a Trustgrid gateway. If the node is acting as an edge then no inbound access is required.
  - Outbound traffic to Trustgrid’s control plane IP (TCP 80/443 & 8443 to 35.171.100.16/28 & 34.223.12.192/28)
  - Outbound traffic to AWS API (TCP 443) https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html
  - Inbound & Outbound to/from management NIC security group on cluster port (typically TCP port 9000)
  - For the initial deployment outbound access for TCP 80/443 should be allowed. Upon successful registration with the Trustgrid Portal, this can be removed.

- IAM role for the instance with policies allowing changes to the routing table of the data NIC - See attached doc

- All Interfaces on the Trustgrid Gateway should have source/destination check disabled in AWS

- Security group for data NIC - No configuration for now

- An IP in the private subnet that will be used by the data NIC

- An SSH key-pair that can be used to SSH to the instance if necessary

- VPC must have unallocated public IP that will be claimed during provisioning

## Process

1. Create a new Node. When complete the Node license will copy to clipboard.

   - Note: The node will not be visible in the portal until the registration process is complete.
   - Download the license to local storage in case the clipboard is cleared. You cannot reissue a license without recreating the node.

1. Select the appropriate Cloud Formation Template based on the AWS region in which the Trustgrid node is being deployed
   - US-EAST-1: https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node-useast1.json
   - US-EAST-2: https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node-useast2.json
   - US-WEST-1: https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node-uswest1.json
   - US-WEST-2: https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node-uswest2.json

1. Fill out the fields in the CloudFormation form
## Parameters

### Stack Name
Unique name to describe this deployment

### Instance Configuration

{{<fields>}}
{{% field "Instance Type" %}}
Set the instance type of the EC2 instance to deploy (bigger instances cost more)
{{% /field %}}
{{% field "SSH Keypair" %}}
SSH keypair to SSH to the instance as ubuntu user if necessary

> SSH access requires a security group change allowing access. We strongly recommend that SSH is not allowed from anywhere (0.0.0.0/0).
{{% /field %}}


{{%field "Host IAM Role" %}}
An IAM role needs to be created with the permissions listed in the [IAM Role Requirements section below](#iam-role-requirements).
{{% /field %}}

{{</fields>}}

#### IAM Role Requirements
##### Encrypted EBS Volume 
> Required for all nodes

By default, the cloud formation template provided will configure an encrypted EBS volume on the Trustgrid Node.
The following permissions need to be applied to the associated IAM role to provide access to the default EBS key. 
Note you will need to input your applicable AWS account ID/region where this node is being deployed.

```json
 {
	"Effect": "Allow",
	"Action": [
        "kms:Decrypt",
        "kms:DescribeKey",
        "kms:ReEncrypt*",
        "kms:GenerateDataKey*"
        ],
        "Resource": "arn:aws:kms:us-east-1:$aws_accountid:alias/aws/ebs"        
}
```
##### Route Table
> Required for [clustered nodes]({{<relref "/docs/clusters">}})

If the node will be clustered the IAM role requires the following permissions (`ec2:DescribeRouteTables` for all resources and `ec2:CreateRoute` and `ec2:DeleteRoute` on the route table):

Route Table Policy

```json
{
	"Effect": "Allow",
	"Action": "ec2:DescribeRouteTables",
	"Resource": "*"
},
{
	"Effect": "Allow",
	"Action": [
		"ec2:CreateRoute",
		"ec2:DeleteRoute"
	],
	"Resource": "arn:aws:ec2:us-east-1:$aws_accountid:route-table/rtb-f428d58b"
}
```
**NOTE**: Set the Resource field to the ARN of the Routing Table associated with the data NICs of the instance. 


## Management Configuration
This section covers the configuration of the outside, internet-facing interface of the EC2 instance.
{{<fields>}}
{{% field "Security Group" %}}

- Needs to allow outbound traffic to other gateways and the [Trustgrid public IP range]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}), at a minimum.
- If it's a gateway node, needs to allow inbound access on the gateway port, typically TCP/UDP 8443.

{{% /field %}}
{{% field "Subnet" %}}
The VPC subnet for the public, internet-facing interface. The EIP that is created by the CloudFormation template will be associated with the interface on this subnet.
{{% /field %}}
{{</fields>}}

## Data Path Configuration
This section covers the configuration of the inward, private-facing interface of the EC2 instance.

{{<fields>}}
{{% field "Security Group" %}}
The security group for the data path interface. 

- Needs to allow communication between any private AWS network resources that need to access the Trustgrid EC2 node's private IP or any [virtual network]({{<relref "docs/domain/virtual-networks">}}) resources that will be accessed across the Trustgrid network.
- If the EC2 instance node will be clustered, the security group should allow communication between the private IPs of all the clustered nodes on the [cluster heartbeat port]({{<relref "/docs/nodes/appliances/cluster#heartbeat">}}), typically port TCP 9000.

{{% /field %}}

{{% field "Subnet" %}}
The VPC subnet for the data path interface.
{{% /field %}}

{{% field "Data IP" %}}
The private IP for the data path - must belong to the subnet and not already be allocated.
{{% /field %}}

{{</fields>}}


## Trustgrid Configuration
{{<fields>}}
{{% field "Trustgrid License" %}}
Copy/paste the license from the portal.

Note: It is critical that you copy/paste the license correctly.
{{% /field %}}

{{<alert>}}If you are not a direct Trustgrid customer please work with your vendor to get these licenses generated and sent to you.{{</alert>}}
{{</fields>}}
## Creating the Stack

1. Create the stack. 
    - Check the box acknowledging that AWS CloudFormation might create IAM resources. This is required because we create an instance profile for the to-be-run EC2 instance.
1. You can now manage the node as you would any other in the Portal UI.
