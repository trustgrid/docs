---
tags: ["node"]
title: "Data Plane"
linkTitle: "Data Plane"
---

{{% pageinfo %}}
Data plane statistics between a node and its peers can help troubleshoot connectivity or performance issues.
{{% /pageinfo %}}

To view latency data between two nodes, select either the edge or gateway node, and then from the peers table, select the node to view.

{{<alert color="info">}}Gateway nodes will list edge nodes in their peers table{{</alert>}}

![img](peer-list.png)

Once a peer is selected, the monitoring section will populate with reelvant data.

![img](monitoring.png)

Hop data is only available for nodes that have hop monitoring enabled. See [Monitoring Network Hops to Peers]({{< ref "/tutorials/gateway-tools/monitoring-network-hops-to-peers" >}}).
