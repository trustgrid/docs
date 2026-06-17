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

The Cluster IP is supported in AWS. It is managed through the AWS API by assigning it as a secondary private IP address on the network interface of the active node. This requires API-level permissions and is automatically handled by the control plane. See [AWS Cluster IP Failover]({{<relref "/tutorials/deployments/deploy-aws/ip-failover">}}) for setup details.

### Google Cloud (GCP)

The Cluster IP is not currently supported in GCP. GCP lacks the low-level IP address failover control necessary to support this feature reliably.

---
