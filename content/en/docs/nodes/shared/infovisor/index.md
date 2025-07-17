---
title: Infovisor
linkTitle: Infovisor
description: The infovisor displays details about the Trustgrid node
weight: 8
aliases:
    - /docs/nodes/infovisor
---

The infovisor displays details about the Trustgrid node and can be displayed by clicking the Info button in the top right or hitting the backtick `` ` `` key.
{{<tgimg src="info-button.png" caption="Button to open Infovisor" width="25%">}} 

It displays the following information in a card layout broken up into a number of sections that can be expanded and collapsed, or removed from view.
{{<tgimg src="infovisor.png" caption="Example infovisor" width="80%">}}

It contains the below information but can vary between appliance and agent-based nodes. The table below lists the fields and if they are displayed depending on node type.


## Sections

The node name is displayed at the top, center of the infovisor. This is the one section that cannot be removed from view.

### Essentials
|Field | Agent | Appliance | Description|
|---|---|---|---|
|Status|✅|✅| The status of the node, either ["Enabled" or "Disabled"]({{<relref "/tutorials/management-tasks/changing-node-status/">}}).|
|Node/Agent ID|✅ |✅|The unique ID of the node.|
|Location|✅|✅|The location of the node as determined by geo-location of the public IP address of the node. |
|ISP|✅|✅|The ISP of the node as determined by the public IP address of the node.|
|TGRN | ✅ |✅| [Trustgrid Resource Name]({{<ref "/docs/user-management/policies#trustgrid-resource-names-or-tgrn">}}) of the node. |
|Version| ✅ |✅|
|Package Version| ❌ |✅|

### Connection
|Field | Agent | Appliance | Description|
|---|---|---|---|
|Control and Data Plane Status| ✅ |✅| Status of the control and data plane connection to the Trustgrid control plane. |
|Public IP | ✅ |✅| The public IP address of the node as observed by the Trustgrid control plane. This includes a lock icon to indicate if the node is [limited to a specific public IP address.]{{<relref "/tutorials/management-tasks/limit-node-functionality">}} |

### Interfaces
This panel includes a table listing information about the interfaces on the node. 

Copy buttons are included next to each cell. And there is a Copy All button that will copy the entire table to the clipboard. 

Additionally, under actions you can export the table to a CSV file.

### Device
|Field | Agent | Appliance | Description|
|---|---|---|---|
|Device Type | ✅ | ✅ | The hardware type of an appliance or will list "Agent" as the type for agent-based nodes. |
|Role | ❌ | ✅ | Will list the [Gateway Server role]({{<relref "/docs/nodes/appliances/gateway/gateway-server#gateway-server-types">}}) if acting as a Gateway Server. Or `Edge` if acting as an Edge node. |
|Operating System (OS) Version| ✅ |✅| The OS version of the node. |
|UDP Mode | ✅ | ✅ | Displays if the node is [enabled to attempt using UDP tunnels]({{<relref "/docs/nodes/appliances/gateway#enable-udp">}})|

### Configuration Status - Appliances Only
|Field | Description|
|---|---|
|SSH Lockdown | Displays if the SSH is restricted properly. Clicking the refresh button will recheck the status. |
|DNS Resolution | Displays if the node can successfully resolve DNS using the configured DNS servers. Clicking the refresh button will recheck the status. |
|Repo Connectivity | Displays if the node can connect to the Trustgrid update repository. Clicking the refresh button will recheck the status. |
|Last Node Config Update | The recent date and time the node last downloaded a new node or cluster configuration **and applied it**. </br></br> If this is older than the last known node (or cluster) configuration change, this could indicate the node is having issues connecting to the Trustgrid control plane REST API for configuration updates. |
|Last Domain Config Update | The most recent date and time the node last downloaded a new domain configuration (e.g. virtual network route changes) **and applied it**.  </br></br> If this is older than the last known domain configuration change, this could indicate the node is having issues connecting to the Trustgrid control plane REST API for configuration updates.|
|Upgrade Status | Displays if an upgrade is either: <ul><li>In Progress</li><li>Complete</li><li>Failed</li></ul> |
|Upgrade Completion Time | The most recent date and time the upgrade was completed or failed. |

### AWS Metadata (AWS Appliances only)
|Field | Description|
|---|---|
|Availability Zone | The Availability Zone of the node. |
|Instance ID | The Instance ID of the node. |
|Instance Type | The Instance Type of the node. |
| AMI ID | The AMI ID of the node. |

### Azure Metadata (Azure Appliances only)
|Field | Description|
|---|---|
|VM ID | The VM ID of the node. |
|VM Size | The VM Size of the node. |
|Subscription ID | The Subscription ID of the node. |
|Location | The Location of the node. |


### Tags
This section displays a table of all [Trustgrid tags]({{<ref "/docs/nodes/shared/tags">}}) applied to the node. 

Under actions, you can export the table to a CSV file.

