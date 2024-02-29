---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "VNET"
linkTitle: "VNET"
aliases: 
    - /docs/nodes/tunnels/vnet
description: Configure a tunnel interface representing the Trustgrid virtual network to allow for connections to VRFs
---


{{% pageinfo %}}
A virtual network tunnel interface can be configured on a Trustgrid node or cluster to allow the forwarding of traffic to a remote Trustgrid Node or Cluster that is attached to the same virtual network. The tunnel interface is associated with a [VRF]({{<relref "/docs/nodes/appliances/vrfs" >}}) which is used to define and control what traffic is allowed to pass and how it should appear on the network. There is no attachment of the vnet interface to a node interface required but the virtual network does need to be attached to the node or cluster before the tunnel interface is created. 
{{% /pageinfo %}}

## Configure VNET Tunnel
On a Trustgrid Node navigate to Tunnels under the Network Menu, select Add Tunnel, and then select VNET

![img](add_tunnel.png)

![img](add_vnet.png)

![img](vnet.png)


#### Configuration Parameters
{{<fields>}}
{{<field Name>}}the name of the tunnel interface created on the Trustgrid Node or Cluster{{</field>}}
{{<field Description>}} (optional) descriptive parameters related to the tunnel {{</field>}}
{{<field "Virtual Network">}} the Trustgrid virtual network the tunnel is being associated with. The virtual network should already be attached to the node or cluster to be selectable. {{</field>}}
{{<field VRF>}}The [VRF]({{<relref "/docs/nodes/appliances/vrfs" >}}) the tunnel is being associated with. Only VRF's which exist on the Trustgrid Node or Cluster can be selected.{{</field>}}
{{<field MTU>}} the maximum transmission unit (MTU) size of the tunnel interface. Typically the default of 1430 should be used. {{</field>}}
{{</fields>}}
