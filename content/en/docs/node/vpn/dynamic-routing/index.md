---
title: "Dynamic Routing"
date: 2023-2-14
---

Dynamic routing allows a node to export and import routes from other nodes in the virtual network. Find dynamic route definitions and change them under the `Dynamic Routing` section of the VPN configuration.

![img](list.png)

## Exports

Export routes for other nodes in the virtual network to import.

![img](add-export.png)

{{<field-def "Destination">}}The node to route traffic to.{{</field-def>}}
{{<field-def "Destination CIDR">}}The network to route traffic to.{{</field-def>}}
{{<field-def "Metric">}}The route metric. Lower metrics are processed first.{{</field-def>}}
{{<field-def "Gateway Path">}}An optional gateway path to use for the route.{{</field-def>}}

## Imports

Import routes from other nodes in the virtual network.

![img](add-import.png)

{{<field-def "Destination">}}The node to route traffic to.{{</field-def>}}
{{<field-def "Destination CIDR">}}The network to route traffic to.{{</field-def>}}
{{<field-def "Metric">}}The route metric. Lower metrics are processed first.{{</field-def>}}
{{<field-def "Gateway Path">}}An optional gateway path to use for the route.{{</field-def>}}
