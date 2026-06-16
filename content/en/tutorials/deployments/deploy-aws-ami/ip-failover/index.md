---
title: "AWS IP Failover"
linkTitle: "IP Failover"
description: "Cluster IP failover for AWS appliances"
---
{{<alert color="info">}}AWS Cluster IP failover requires the [June 2026 major appliance release]({{<ref "/release-notes/node/2026-06/index.md">}}) or later.{{</alert>}}

## How it works
Trustgrid provides the ability for a floating [Cluster IP]({{<relref "/docs/clusters/cluster-only-config#cluster-ip">}}) to be assigned to a cluster of AWS appliances. The Cluster IP is assigned as a secondary private IP on the network interface (ENI) of the active appliance, and is supported on both the WAN and LAN interfaces. When failover occurs, the appliance taking the active role claims the secondary private IP on its own ENI through the AWS API, which moves it from the previously active appliance.

### Graceful Failover
This describes a graceful failover, where the active member changes while both members are online and working normally.
1. The appliance relinquishing the active role unassigns the Cluster IP from its ENI.
1. The appliance gaining the active role assigns the Cluster IP to its ENI.

### Ungraceful Failover
This describes an ungraceful failover, where [cluster health conditions]({{<relref "/docs/clusters#cluster-member-health-conditions">}}) prevent the active member from functioning.
1. After the configured [Cluster Timeout]({{<relref "/docs/clusters#cluster-timeout">}}) elapses, the appliance taking the active role assigns the Cluster IP to its own ENI through the AWS API, which moves it off the previously active ENI.

## Requirements for HA IP Failover
- An unused private IP address within the subnet CIDR of each interface where the Cluster IP is defined.
- Source/destination check disabled on the appliance interfaces, which is already required for Trustgrid nodes in AWS.
- An IAM role attached to the instances with permission to manage private IP addresses on the ENIs.

{{<alert color="warning">}}On startup, the appliance reconciles the secondary private IPs on its ENIs and removes any that do not match its current interface IP or Cluster IP. Do not assign additional secondary private IPs to a Trustgrid appliance's ENIs directly in AWS, as they will be removed on the next restart.{{</alert>}}

### Permissions Required for Cluster IP Failover
The appliances need permission to describe their network interfaces and to assign and unassign private IP addresses on them. Add the following policy to the [IAM role attached to the instances]({{<relref "/tutorials/deployments/deploy-aws-ami#iam-role-requirements">}}). `AssignPrivateIpAddresses` and `UnassignPrivateIpAddresses` can be scoped to the appliance ENIs, while `DescribeNetworkInterfaces` does not support resource-level permissions and must use `*`.

> Set the region and `$aws_accountid` to match where the appliances are deployed. You can optionally replace the `network-interface/*` wildcard with the specific ENI ARNs of the appliance interfaces.

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ec2:DescribeNetworkInterfaces",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AssignPrivateIpAddresses",
                "ec2:UnassignPrivateIpAddresses"
            ],
            "Resource": "arn:aws:ec2:us-east-1:$aws_accountid:network-interface/*"
        }
    ]
}
```

#### Add the policy via the AWS Console
1. Open the [IAM console](https://console.aws.amazon.com/iam/) and select **Roles**.
1. Select the role attached to your Trustgrid appliances (the **Host IAM Role** used during [deployment]({{<relref "/tutorials/deployments/deploy-aws-ami#iam-role-requirements">}})).
1. On the **Permissions** tab, choose **Add permissions**, then **Create inline policy**.
1. Select the **JSON** tab and paste the policy above. Update the region and account ID, and optionally narrow the `network-interface` ARN to the specific ENIs.
1. Choose **Next**.
1. Give the policy a name, such as `tg-cluster-ip-failover`, and choose **Create policy**.
1. If your cluster members use separate IAM roles, repeat these steps for each role.

{{<alert color="info">}}These permissions can be combined with the existing [route table permissions]({{<relref "/tutorials/deployments/deploy-aws-ami#route-table">}}) on the same IAM role rather than created as a separate policy.{{</alert>}}
