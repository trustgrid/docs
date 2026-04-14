---
title: "View Virtual Route Table"
description: View the routing table on a node
---

## Usage

1. Login to the Trustgrid portal and navigate to the Node from which you want to test connectivity.
1. Select **VPN** under the **Network** section.
1. Click the **Tools** button and click the **View Virtual Routes** button

	![Virtual Network Tools](network-tools.png)

1. A table will appear with routing information for the node.

	![Virtual Route Table](virtual-route-table.png)

## Analyzing the Route Table

### Route Availability

If a route is available, its `Active` column will show a green up arrow (<i class="fa-solid fa-arrow-up" style="color: green;"></i>). If a route is not available, its `Active` column will show a red down arrow (<i class="fa-solid fa-arrow-down" style="color: red;"></i>).

### Route Preference

Routes will be selected based on the following order:

1. In overlapping routes, the most specific route wins. For example, 10.20.0.51/32 takes precedence over 10.20.0.0/24.
1. If multiple matching routes exist, the route with the lowest metric wins. For example, a metric of 1 takes precedence over a metric of 50.