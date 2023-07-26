---
tags: ["node"]
title: "Data Plane"
linkTitle: "Data Plane"
weight: 4
---

{{% pageinfo %}}
Data plane statistics between a node and its peers can help troubleshoot connectivity or performance issues.
{{% /pageinfo %}}

## Export List of Peers
You can export a CSV formatted list of all peers by clicking Actions -> Export
{{<tgimg src="data-plane-export.png" caption="Export button" width="25%">}}

The CSV file will include columns for:
* Connection status (true/false)
* Peer name
* Peer IP
* Path (if applicable)
* RTT - Return Trip Time or latency observed when the panel was opened or refreshed
* Mode - TLS or TLS/UDP
* Type - The role of the node to peer. SERVER = gateway, CLIENT = edge device
* Ports - The ports used to establish the connection


## View Latency to Peers

To view latency data between two nodes, select either the edge or gateway node, and then from the peers table, select the node to view.

{{<alert color="info">}}Gateway nodes will list edge nodes in their peers table{{</alert>}}

![img](peer-list.png)

Once a peer is selected, the monitoring section will populate with reelvant data.

![img](monitoring.png)

Hop data is only available for nodes that have hop monitoring enabled. See [Monitoring Network Hops to Peers]({{<ref "/tutorials/gateway-tools/monitoring-network-hops-to-peers" >}}).
