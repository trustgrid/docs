---
title: Node Deployment
linkTitle: Deployment
weight: 30
description: Intro to deploying Trustgrid Nodes
---


Nodes are available in two form factors: 
- Appliance - This is a combination of a fully managed operating system and Trustgrid software pre-installed. All configuration and management is done via the Trustgrid Control Plane. 
- Agent - This is a software package that can installed on supported operating systems. Users can install additional software and configure the operating system as needed. The Trustgrid Agent connects back to the Control Plane for management of only the Trustgrid node services. 

| Functionality | Agent | Appliance |
|-|-|-|
| Gateway Capabilities | None | Data Plane or ZTNA |
| Operating System | [Multiple Supported OS]({{<relref "/tutorials/agent-deploy#supported-operating-systems">}}), install additional software | Fully managed OS and Trustgrid service, no additional software permitted |
| Updates | Managed by user via OS native tools | OS and Trustgrid updates managed via Trustgrid |
| Interface IP Management | Managed by user via OS native tools | Configurable via Trustgrid portal or [local console]({{<relref "/tutorials/local-console-utility/">}})
| Layer 3 VPN | Single Virtual Network connectivity | Support for multiple Virtual Networks |
| Layer 4 Proxy | Full Support | Full Support |
| Compute | Commands only | Commands, Containers or VMs supported |

Nodes can be deployed in a variety of ways, including as virtual machines, containers (as [agents]({{<relref "/tutorials/agent-deploy">}})), or bare metal servers. Additionally, appliance-based nodes can be [clustered]({{<relref "/docs/clusters">}}) for high availability. 
