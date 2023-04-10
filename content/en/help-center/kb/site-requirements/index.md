---
title: "Site Requirements"
date: 2017-01-05
description: >
  Network access requirements
---

{{% pageinfo %}}
Trustgrid’s architecture consists of Nodes deployed at the Edge (on-premise) and management infrastructure built in AWS. Nodes create tunnels to other nodes for data transfer, and to the management infrastructure for control. All tunnels are TLS connections and must be permitted by any firewall between
{{% /pageinfo %}}

### Network Requirements for All Nodes

In order to connect to the Trustgrid Control Plane, the following outbound traffic must be allowed from the node’s configured primary interface IP address

- TCP Port 443 and TCP/UDP 8443 to:
  - 35.171.100.16/28
  - 34.223.12.192/28
- TCP/UDP Port 53 to the configured DNS servers. These DNS servers must be accessible from the WAN/outside interface IP and be able to resolve DNS requests for the trustgrid.io domain

{{<alert color="info">}} Trustgrid nodes will attempt connect to 169.254.169.254 on port 80. This is a standard instance metadata server in public clouds such as AWS and Azure. This connectivity is not required outside these providers. {{</alert>}}

### Recommendations

At least 10 Mbps download speed is recommended for reasonable performance during upgrades. Actual bandwidth requirement will be specific to the workloads running across the device.

Hardware devices are recommended to be connected to ports configured for 1000 Mbps and full duplex with auto negotiation on.

## Gateway Node Network Requirements

In addition to being able to connect to the Control Plane resource outlined above, Gateway Nodes must be allowed to receive inbound traffic on their configured public IP and port (typically TCP/UDP port 8443). Both the advertised public IP and port are configurable when enabling a node as a gateway.
