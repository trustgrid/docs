---
title: Virtual Network IP Pools
linkTitle: IP Pools
description: Define pools of IPs that are used for automatically assigning virtual management addresses to nodes
---

Virtual Network IP Pools are used for automatically assigning virtual management to nodes.  This is optional but makes configuration and management easier.

{{<alert color="warning">}}At this time only agent-based nodes support automatic IP assignment{{</alert>}}

{{<tgimg src="ip-pools.png" width="90%">}}

## Address Assignment
Addresses are assigned sequentially from the pool by Trustgrid. When a node is attached to a virtual network with an IP pool configured, it will automatically be assigned the next available IP from the pool. When a node is detached its IP is returned to the pool for reuse. However, the released addresses are not used until all addresses in the pool have cycled through once.

{{<alert color="info">}}You will still need to create [domain routes]({{<relref "/docs/domain/virtual-networks/routes">}}) to point the assigned IP to the targeted node. {{</alert>}}

## Adding Pools
You can add one or more virtual pool ranges to a virtual network. It is recommended to size the initial pool to a realistic estimate of your expected nodes plus some overhead. Additional ranges can be added later if needed.

1. Click the **Add IP Pool** button.
1. Provide the requested information: {{<fields>}}
{{<field CIDR>}}Provide the CIDR of the network you want to make available in the pool.  For example, 10.1.1.0/24{{</field>}}
{{<field Description>}}(optional)Provide a user-friendly description of the pool.{{</field>}}
{{</fields>}}
1. Click Save

{{<tgimg src="add-ip-pool.png" width="60%" caption="Add IP Pool">}}