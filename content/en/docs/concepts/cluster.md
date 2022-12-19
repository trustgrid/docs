---
categories: ["concepts"]
tags: ["cluster", "concepts"]
title: "Cluster"
date: 2022-12-19

---

{{% pageinfo %}}
A cluster is a pair of [nodes]({{< ref "node" >}}) that share configuration.
{{% /pageinfo %}}

A cluster is a pair of [nodes]({{< ref "node" >}}) at a single site that provides automated high-availability (HA) connectivity. An additional IP address is assigned as a Cluster Virtual IP address that can move between the [nodes]({{< ref "node" >}}) if failover occurs. 

Additionally, certain settings such as Network Services and VPN settings can be configured for the cluster and these settings will override the individual node's configuration.




