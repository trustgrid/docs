---
title: "Cluster"
date: 2022-2-13
---

{{<alert>}}Cluster configuration is only available for nodes that are members of a [cluster]({{<ref "docs/cluster">}}).{{</alert>}}

## Heartbeat

The cluster heartbeat defines a persistent connection between cluster members. The heartbeat is used to determine the health of the cluster and to determine which member is the active member.

![img](heartbeat.png)

{{<field "host">}}The IP (bound to an interface) on which heartbeat communication will occur{{</field>}}

{{<field "port">}}The port on which heartbeat communication will occur{{</field>}}

{{<alert>}}The same interface should be selected for all cluster members{{</alert>}}
