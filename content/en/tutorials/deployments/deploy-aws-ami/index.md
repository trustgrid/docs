---
tags: ["aws"]
title: "Deploy a Trustgrid Node AMI in AWS"
date: 2023-2-9
---

Standing up a Trustgrid node in AWS is easy using an Amazon AMI. Trustgrid nodes in AWS use two network interfaces - a management and data interface. The management interface communicates with Trustgrid Cloud Management systems. The data interface is used to terminate TLS tunnels from Edge Nodes.

## Notes

- The cloudformation template below works with an AMI currently published in US-EAST1. Deploying in other regions requires working with Trustgrid Support
- Logs /var/log/syslog, /var/log/secure, and /var/log/trustgrid/tg-default.log to CloudWatch
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
  - For the initial deployment outbound access for TCP 80/443 should be allowed. Upon successful registration with the Trustgrid Portal this can be removed.

- IAM role for the instance with policies allowing publishing to Cloudwatch logs and changes to the routing table of the data NIC - See attached doc

- All Interfaces on the Trustgrid Gateway should have source / destination check disabled in AWS

- Security group for data NIC - No configuration for now

- An IP in the private subnet that will be used by the data NIC

- An SSH keypair that can be used to SSH to the instance if necessary

- VPC must have unallocated public IP that will be claimed during provisioning

## Process

1. Create a new Node. When complete the Node license will copy to clipboard.

   - Note: The node will not be visible in the portal until the registration process is complete.
   - Download the license to local storage in case the clipboard is cleared. You cannot reissue a license without recreating the node.

1. Click this link to start the AMI provisioning process. https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://s3.amazonaws.com/tg-dev-public/cf-trustgrid-node.json

1. Fill out the fields in the CloudFormation form

## Instance Configuration

{{<field "Stack Name">}}
Unique name to describe this deployment
{{</field>}}

{{<field "Instance Type">}}
Set the instance type of the EC2 instance to deploy (bigger instances cost more)
{{</field>}}

{{<field "Host IAM Role">}}
An EC2 associated role that allows creating and writing to CloudWatch logs. Only the role name itself is required.

The role needs these IAM privileges: `logs:CreateLogGroup`, `logs:CreateLogStream`, `logs:PutLogEvents`, `logs:DescribeLogStreams` on resource `arn:aws:logs:*:*:*`

Cloudwatch logs policy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": ["*"]
    }
  ]
}
```

If EC2 Nodes will be clustered (Layer 3) the IAM role needs `ec2:DescribeRouteTables` for any resource and `ec2:CreateRoute` & `ec2:DeleteRoute` on the route table

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
	"Resource": "arn:aws:ec2:us-east-1:079972220921:route-table/rtb-f428d58b"
}
```

**NOTE**: Set the Resource field to the ARN of the Routing Table associated with the data NICs of the instance.

**CloudWatch Note**: We will create log groups named /trustgrid/var/log/syslog and /trustgrid/var/log/trustgrid/tg-default.log.
{{</field>}}

{{<field "SSH Keypair">}}
SSH keypair to SSH to the instance as ubuntu user if necessary

**NOTE**: SSH access requires a security group change allowing access. We strongly recommend that SSH is not allowed from anywhere (0.0.0.0/0).
{{</field>}}

## Management Configuration

{{<field "Security Group">}}
Needs to allow outbound traffic to other gateways and our public IP range, at a minimum. If it's a gateway node, needs to allow inbound access on desired gateway port.
{{</field>}}

{{<field "Subnet">}}
The VPC subnet to put the EC2 instance in. This needs to be a subnet with public IP enables (the instance will automatically claim one; the Auto-Assign Public IP does not need to be enabled)
{{</field>}}

## Data Path Configuration

{{<field "Security Group">}}
The security group for the data path - needs to allow outbound communication to other gateways, and inbound communication on its gateway port
{{</field>}}

{{<field "Subnet">}}
The VPC subnet to put the data interface in - if it's a cloud-accessible gateway, should be a public subnet, if it's only for internal AWS traffic, can be a private subnet. Will need outbound access either through IGW or NAT GW.
{{</field>}}

{{<field "Data IP">}}
The private IP for the data path - must belong to the subnet
{{</field>}}

## Trustgrid Configuration

{{<field "Security Group">}}
Copy/paste the license from the portal.

Note: It is critical that you copy/paste the license correctly.
{{</field>}}

{{<alert>}}If you are not a direct Trustgrid customer please work with your vendor to get these licenses generated and sent to you.{{</alert>}}

## Creating the Stack

1. Create the stack. Check the box acknowleding that AWS CloudFormation might create IAM resources. This is required because we create an instance profile for the to-be-run EC2 instance.
1. When the node appears in the Portal, activate it.
1. In the EC2 console, reboot the node (it will be named trustgrid-node)
1. You can now manage the node as you would any other in the Portal UI.
