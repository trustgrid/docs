---
Title: "Find the Public IP Address of a Node"
linkTitle: "Finding Public IP"
description: "How to find the public IP address of a Trustgrid node"
Date: 2023-01-03
---

The public IP address refers to the source IP address observed by the Trustgrid cloud when the control plane connections are established. This is the IP address after any NATs are applied to outbound internet traffic on the node's WAN interface.

## Nodes Table
The public IP address is a default column in the Nodes table list. You can use search to filter the list to see the public IP of the desired node.

{{<tgimg src="nodes-table-public-ip.png" caption="Example Nodes table with public IP column" width="80%">}}

## Node Infovisor
The public IP address is also listed in the [infovisor]({{<ref "/docs/nodes/shared/infovisor">}}) panel accessible from the node detail page.

{{<tgimg src="node-infovisor-public-ip.png" caption="Public IP listed in infovisor" width="80%">}}