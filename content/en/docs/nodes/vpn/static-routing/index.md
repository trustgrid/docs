---
title: "Static Routing"
date: 2023-2-14
---

Static routing allows a node to route traffic to a remote network. Find static route definitions and change them under the `Static Routing` section of the VPN configuration.

![img](add-modal.png)

{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route.{{</field>}}
