---
title: "Cluster-Only Configuration Items"
description: "Features and configuration options specific to clustered environments"
---

## Cluster IP

The **Cluster IP** is a virtual IP address used by the cluster to provide a consistent endpoint for clients, even when the active node changes.

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

The Cluster IP is supported in Azure. It is managed through the Azure API by assigning it as an additional IP configuration on the network interface of the active node. This requires API-level permissions and is automatically handled by the control plane.

### AWS

The Cluster IP is supported in AWS as of the [June 2026 major appliance release]({{<ref "/release-notes/node/2026-06/index.md">}}). It is managed through the AWS API by assigning it as a secondary private IP on the network interface (ENI) of the active node, on both the WAN and LAN interfaces. This requires the instance's IAM role to allow managing private IP addresses on the ENIs. See [AWS IP Failover]({{<relref "/tutorials/deployments/deploy-aws-ami/ip-failover">}}) for setup and the required permissions.

### Google Cloud (GCP)

The Cluster IP is **not currently supported** in GCP. Similar to AWS, GCP lacks the low-level IP address failover control necessary to support this feature reliably.

---