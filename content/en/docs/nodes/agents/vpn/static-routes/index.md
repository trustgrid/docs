---
title: Agent Virtual Network Static Routes
linkTitle: Static Routes
description: Define static routes for an agent's virtual network
---
Agents can define static routes for their virtual network to augment or override routes defined in the global [virtual network route table]({{<relref "/docs/domain/virtual-networks/routes">}}). 
{{<tgimg src="agent-static-route.png" width="90%">}}
{{<fields>}}
{{<field Destination>}}The destination node or cluster where matching traffic will be routed.{{</field>}}
{{<field "Destination CIDR">}}The CIDR notation of the IP address or network traffic that this route should manage.{{</field>}}
{{<field Metric>}}Value between 1 and 200. The available route with the lowest metric will be selected when a flow is created. {{</field>}}
{{<field Description>}}(optional) A user-friendly description for the route and its purpose. {{</field>}}
{{</fields>}}

## Manage Agent Static VPN Routes
### Add Agent Static VPN Route

1. Navigate to the Static Routes page of the desired agent.
1. Click the **Add Route** button at the top of the table.
1. Provide the Destination, Destination CIDR, Metric, and optional description for the new route.
1. Optionally, repeat with additional routes.
1. Click Save


### Delete Agent Static VPN Routes
1. Navigate to the Static Routes page of the desired agent.
1. Select the check box to the left of the route you wish to delete.
1. From the Actions dropdown above the table select Delete.
1. When prompted confirm you wish to delete the rule.
