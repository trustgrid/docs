---
title: "Dynamic Routing"
date: 2023-2-14
---

Dynamic routing allows a node to export and import routes from other nodes in the virtual network. Find dynamic route definitions and change them under the `Dynamic Routing` section of the VPN configuration.

{{<tgimg src="list.png" caption="Dynamic Routing table" alt="A table for Export Routes and a table for Imports" width="85%">}}

## Exports

Export routes for other nodes in the virtual network to import.

{{<tgimg src="add-export.png" caption="Add Export Route dialog" alt="Dialog to add export route with options of destination, destination CIDR, metric and optional gateway path">}}

{{<fields>}}
{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route.{{</field>}}
{{</fields>}}

## Imports

Import routes from other nodes in the virtual network.

{{<tgimg src="add-import.png" caption="Add Import Route dialog" alt="Dialog to add import route with options of destination, destination CIDR, metric and optional gateway path">}}

{{<fields>}}
{{<field "Destination">}}The node to route traffic to.{{</field>}}
{{<field "Destination CIDR">}}The network to route traffic to.{{</field>}}
{{<field "Metric">}}The route metric. Lower metrics are processed first.{{</field>}}
{{<field "Gateway Path">}}An optional gateway path to use for the route.{{</field>}}
{{</fields>}}