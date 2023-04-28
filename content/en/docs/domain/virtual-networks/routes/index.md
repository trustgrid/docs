---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Routes"
date: 2023-1-9
weight: 30
---

{{% pageinfo %}}
Routes configured under [Domains]({{<ref "docs/domain" >}}) determine to which [node]({{<ref "docs/nodes" >}}) or [cluster]({{<ref "docs/clusters" >}}) the [Trustgrid virtual network]({{<ref "docs/domain/virtual-networks" >}}) should route traffic for a specific subnet.
{{% /pageinfo %}}

## Virtual Network Routes
Routes defined a the virtual network level allow creating a global route table shared with all nodes and clusters attached to the virtual network. This table is then used to evaluate where to send VPN traffic for a specified destination CIDR.


![img](routes-list.png)

### Route Fields
A Route has the following fields:

{{<fields>}}
{{<field "Destination" >}}
This will be the name of the [node]({{<ref "docs/nodes" >}}) or [cluster]({{<ref "docs/clusters" >}}) that traffic will be routed to. This list is auto-populated based on the nodes and clusters in the selected [domain]({{<ref "docs/domain" >}}).
{{</field >}}

{{<field "Destination CIDR" >}}
This is the CIDR notation of the [virtual network]({{<ref "docs/domain/virtual-networks" >}}) that should be routed to the above destination [node]({{<ref "docs/nodes" >}}) or [cluster]({{<ref "docs/clusters" >}}).
{{</field >}}

{{<field "Metric" >}}
If there are multiple routes for the same [virtual network]({{<ref "docs/domain/virtual-networks" >}}) the metric will determine which route will be used. The lowest number is the highest priority. See the "Automatic Failover" text below.
{{</field >}}

{{<field "Description">}} (Optional) This field can be used to provide additional information about the purpose of the route.  e.g. it could label a route as a DR route indicating that in normal circumstances it would have a higher metric than a primary route.
{{</field>}}

{{</fields>}}
## Managing Virtual Network Routes

### Adding Virtual Network Routes

1. Navigate to Domain > Virtual Networks and select the desired Virtual Network.
1. The Routes table is selected by default{{<tgimg src="vnet-route-table.png" width="40%" caption="Virtual Network > Routes" alt="Virtual Network named vNet1 with the Routes option selected in the navigation menu">}} {{<alert color="info">}} Tip: Use the search field to filter the list of routes so you can see the routes you are adding. You could filter by the CIDR address or part of the destination node/cluster names. This will show any existing matching routes, and will make it easy to see the new routes you are adding.{{</alert>}}
1. Click the add route button.{{<tgimg src="add-route-button.png" width="40%" caption="Add route button">}}
1. Provide the desired route information in the [fields](#route-fields) <ol type="A"><li>Select the destination cluster or node</li><li>Enter the destination network in CIDR notation. For a single IP use /32.</li><li>Enter a metric between 1 and 200.</li><li>Optionally, provide a description.</li></ol>{{<tgimg src="new-route-entry.png" width="80%">}}
1. Repeat the above two steps for any additional routes you wish to add.
1. Click the Save button. {{<tgimg src="save-button.png" width="20%" caption="Save route button">}}
1. There should be a notification saying "Routes Updated" but you will need to [review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) before the changes are actually be published to nodes in your environment. {{<tgimg src="saved-route.png" width="50%" caption="Example of a saved route">}}
    
{{<alert color="warning">}} As stated above, the route changes are not actually implemented until you click the Apply button the [Review Changes]({{<ref "./review-changes">}}) panel. {{</alert>}}

### Deleting Virtual Network Routes

1. Navigate to Domain > Virtual Networks and select the desired Virtual Network.
1. The Routes table is selected by default{{<tgimg src="vnet-route-table.png" width="40%" caption="Virtual Network > Routes" alt="Virtual Network named vNet1 with the Routes option selected in the navigation menu">}} 
1. Use the search field to filter the list of routes. {{<alert color="info">}}You could filter by the CIDR address or part of the destination node/cluster names. This will show any existing matching routes, and will make it easy to see the new routes so you can delete them.{{</alert>}}
1. Click the red X on the far right of the route being removed. {{<tgimg src="delete-route.png">}}
1. Repeat the above two steps for any additional routes you wish to add.
1. Click the Save button. {{<tgimg src="save-button.png" width="20%" caption="Save route button">}}
1. There should be a notification saying "Routes Updated" but you will need to [review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) before the changes are actually be published to nodes in your environment. {{<tgimg src="saved-route.png" width="50%" caption="Example of a saved route">}}

## Route Failover 

Route failover allows a subnet to be routed to an alternate [node]({{<ref "/docs/nodes" >}}) or [cluster]({{<ref "docs/clusters" >}}) in the event of a failure.  This can be automatic or performed manually.

### Prerequisites 

In either configuration, the [virtual network]({{<ref "docs/domain/virtual-networks" >}}) settings under VPN settings for the primary and backup destination [nodes]({{<ref "docs/nodes" >}})/[clusters]({{<ref "docs/clusters" >}}) must match. Including:

- Network Virtual Route

  ![img](virtual-network-route.png)

- Network Group under Outside NAT Table

  ![img](outside-nat-table.png)

- Virtual CIDR under Inside NAT Table

  ![img](inside-nat-table.png)

### Automatic Failover

To have the route failover without manual intervention you must define two routes for the same subnet (Destination CIDR) with different metrics.  The lowest numerical metric will take precedence unless the destination [node]({{<ref "docs/nodes" >}}) or [cluster]({{<ref "docs/clusters" >}}) is offline.

In the below example configuration we want traffic for the 10.20.0.0/24 network to the `edge` [cluster]({{<ref "docs/clusters" >}}) first, and failover to the `edge-dr` [cluster]({{<ref "docs/clusters" >}}).

![img](automatic-failover.png)

### Preventing Automated Failback or Forcing a Failover

As mentioned above if multiple routes are configured traffic will route to destination with the lowest metric number. So if the primary destination fails but then comes back online traffic will be routed back automatically.

However, there are some circumstances where this is not desirable. For example, if the primary site is unstable you may wish to keep traffic at the backup site until the primary is stabilized. Alternately, you may wish to preemptively reroute traffic to the backup site in advance of planned maintenance at the primary site.

To preemptively reroute traffic to the backup site, update the backup route to have a lower metric than the primary route.

### Manual Failover

In some circumstances, it may be preferable for failover to only occur with manual intervention. In this situation, you will have a single route under the [domain]({{<ref "docs/domain" >}}).  To initiate a failover you'll need to update the destination [cluster]({{<ref "docs/clusters" >}})/[node]({{<ref "docs/nodes" >}}).

1. Login to the Portal and select your domain.
1. Click the link to the desired [domain]({{<ref "docs/domain" >}}) under the "Name" column.
1. Scroll down to the ["Virtual Networks"]({{<ref "docs/domain/virtual-networks" >}}) section and select the "Routes" tab.
1. Find the route you wish to failover.
1. Update the destination device. (e.g. in the below you would switch from edge to edge-dr).
  ![Update Destination](update-destination.png)
1. Click save.
