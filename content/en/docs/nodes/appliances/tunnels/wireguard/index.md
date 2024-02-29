---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "WireGuard"
linkTitle: "WireGuard"
aliases: 
    - /docs/nodes/tunnels/wireguard
description: Configure WireGuard tunnels to allow client connectivity via the WireGuard protocol
---


{{% pageinfo %}}
A WireGuard tunnel interface can be configured to allow a WireGuard client to connect to a Trustgrid Node running as a WireGuard server. 
The tunnel interface is associated with a [VRF]({{<relref "/docs/nodes/appliances/vrfs" >}}) which is used to define and control what traffic is allowed to pass and how it should appear on the network.
{{% /pageinfo %}}

## Configure WireGuard Tunnel
On a Trustgrid Node navigate to Tunnels under the Network Menu, select Add Tunnel, and then select WireGuard

![img](add_tunnel.png)

![img](add_wireguard.png)

![img](wireguard.png)

### WireGuard Configuration Parameters

{{<fields>}}
{{<field "Name">}}
The name of the WireGuard tunnel interface created on the Trustgrid node or cluster.
{{</field>}}

{{<field "Description">}}
Optional field used to document what clients are utilizing the tunnel.
{{</field>}}

{{<field "Interface IP">}}
The IP address in CIDR notation that will be assigned to the node's WireGuard tunnel interface.
{{</field>}}

{{<field "Public Key">}} 
The public key of the remote WireGuard client.
{{</field>}}

{{<field "Pre-shared Key">}}
Optional 256-bit to provide an additional level of security.
{{</field>}}

{{<field "VRF">}}
The Virtual Routing and Forwarding ([VRF]({{<relref "/docs/nodes/appliances/vrfs" >}})) table of the WireGuard tunnel will be attached to
{{</field>}}

{{<field "MTU">}}
The maximum transmission unit (MTU) of the WireGuard tunnel interface.  Defaults of 1430 or lower are recommended.
{{</field>}}

{{</fields>}}

## Example Client Config
As you configure a WireGuard tunnel interface the portal UI will automatically generate an example WireGuard configuration for the peer/client based on the information provided.

{{<tgimg src="wireguard-tunnel-example-config.png" width="90%" caption="Example WireGuard client config">}}

{{<alert color="info">}} Note: this is only a partial config. You will likely need to adjust the `AllowedIps` and other fields to meet the needs of the configured tunnel.
{{</alert>}}

