---
categories: ["concepts"]
tags: ["cluster", "concepts"]
title: "Cluster"
date: 2022-12-07
description: >
  Trustgrid node clusters
---

{{% pageinfo %}}
Clusters are pairs of [nodes]({{< ref "node" >}}) that share configuration.
{{% /pageinfo %}}

A cluster is a pair of nodes at a single site that provides automated high-availability (HA) connectivity. An additional IP address is assigned as a Cluster Virtual IP address that can move between the nodes if failover. Additionally, certain settings such as Network Services and VPN settings can be configured one for the cluster and these settings will override the individual node's configuration.



