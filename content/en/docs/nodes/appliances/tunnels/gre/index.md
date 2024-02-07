---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "GRE"
linkTitle: "GRE"
aliases: 
    - /docs/nodes/tunnels/gre
description: Configure GRE tunnels to establish connectivity between nodes and external devices
---


{{% pageinfo %}}
Trustgrid supports configuring GRE tunnels that can be used to establish connectivity to any appliance that supports the GRE protocol.
A GRE tunnel interface is created which can then be attached to an interface on a node. [VRF]({{<relref "/docs/nodes/appliances/vrfs" >}})s can then be used to control the flow of traffic and how it should appear on the network. 
{{% /pageinfo %}}

## Configure GRE Tunnel 
On a Trustgrid Node navigate to Tunnels under the Network Menu, select Add Tunnel, and then select GRE

![img](add_tunnel.png)

![img](add_gre.png)

![img](gre.png)

### GRE Tunnel Configuration Parameters

{{<fields>}}
{{<field Name>}} The name of the GRE tunnel interface that will be created on the Trustgrid node {{</field>}}
{{<field Description>}} (optional) Used to document the purpose of the tunnel {{</field>}}
{{<field IP>}} the IP address assigned to the node's tunnel interface in CIDR notation {{</field>}}
{{<field Peer>}} the remote IP address of the device the tunnel is being established with {{</field>}}
{{<field VRF>}} existing [VRF]({{<relref "/docs/nodes/appliances/vrfs" >}}) on the node the tunnel is being associated with{{</field>}}
{{<field MTU>}} the maximum transmission unit (MTU) of the tunnel interface. Typically should be set to 1476 or lower. {{</field>}}
{{</fields>}}