---
title: "Cluster"
date: 2022-2-13
---

{{<alert>}}Cluster configuration is only available for nodes that are members of a [cluster]({{<ref "docs/cluster">}}).{{</alert>}}

## Heartbeat

The cluster heartbeat defines a persistent connection between cluster members. The heartbeat is used to determine the health of the cluster and to determine which member is the active member.

![img](heartbeat.png)

{{<field-def "host">}}The IP (bound to an interface) on which heartbeat communication will occur{{</field-def>}}

{{<field-def "port">}}The port on which heartbeat communication will occur{{</field-def>}}

{{<alert>}}The same interface should be selected for all cluster members{{</alert>}}
