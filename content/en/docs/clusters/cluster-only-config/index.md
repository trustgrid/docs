---
title: "Cluster-Only Configuration Items"
weight: 10
description: "Features and configuration options specific to clustered environments"
---

## Cluster IP

The **Cluster IP** is a virtual IP address used by the cluster to provide a consistent endpoint for clients, even when the active node changes.

### Availability

| Platform         | Supported |
|------------------|-----------|
| On-Premise       | ✅ Yes    |
| Azure            | ✅ Yes    |
| AWS              | ❌ No     |
| Google Cloud (GCP) | ❌ No     |

### On-Premise

The Cluster IP is supported in on-premise deployments. The active member of the cluster answers ARP requests for the Cluster IP. Only the active node will respond, ensuring traffic is directed correctly in failover scenarios.

### Azure

The Cluster IP is supported in Azure. It is managed through the Azure API by assigning it as an additional IP configuration on the network interface of the active node. This requires API-level permissions and is automatically handled by the control plane.

### AWS

The Cluster IP is **not currently supported** in AWS due to limitations in IP failover and reassignment capabilities in the VPC networking model.

### Google Cloud (GCP)

The Cluster IP is **not currently supported** in GCP. Similar to AWS, GCP lacks the low-level IP address failover control necessary to support this feature reliably.

---