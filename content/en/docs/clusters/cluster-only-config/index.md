---
title: "Cluster-Only Configuration Items"
description: "Features and configuration options specific to clustered environments"
---

## Cluster IP

The Cluster IP is a virtual IP address used by the cluster to provide a consistent endpoint for clients, even when the active node changes.

### Availability

| Platform         | Supported |
|------------------|-----------|
| On-Premise       | ✅ Yes    |
| Azure            | ✅ Yes    |
| AWS              | ✅ Yes    |
| Google Cloud (GCP) | ❌ No     |

### On-Premise

The Cluster IP is supported in on-premise deployments. The active member of the cluster answers ARP requests for the Cluster IP. Only the active node will respond, ensuring traffic is directed correctly in failover scenarios.

### Azure

The Cluster IP is supported in Azure. It is managed through the Azure API by assigning it as an additional IP configuration on the network interface of the active node. This requires API-level permissions and is automatically handled by the control plane. See [Azure IP Failover]({{<relref "/tutorials/deployments/deploy-azure/ip-failover">}}) for setup details.

### AWS

The Cluster IP is supported in AWS. The active cluster member claims the configured IP as a secondary private IP on its data ENI via `ec2:AssignPrivateIpAddresses` (with `AllowReassignment=true`). On failover, the newly active member atomically migrates the secondary IP to its own ENI — no AWS route-table updates are required.

AWS-specific requirements:
- IAM instance profile must allow `ec2:AssignPrivateIpAddresses`, `ec2:UnassignPrivateIpAddresses`, `ec2:DescribeNetworkInterfaces`, and `ec2:DescribeInstances` on the cluster members' data ENIs.
- Source/destination check must be disabled on the data ENI of each cluster member.
- The cluster IP must be an unused secondary private IPv4 address within the data subnet — it is not an Elastic IP.

See [AWS Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws-ami/cluster-ip-failover-in-aws">}}) for the full setup walkthrough.

### Google Cloud (GCP)

The Cluster IP is not currently supported in GCP. GCP lacks the low-level IP address failover control necessary to support this feature reliably.

---
