---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "WireGuard"
linkTitle: "WireGuard"
---


{{% pageinfo %}}
A WireGuard tunnel interface can be configured to allow a WireGuard client to connect to a Trustgrid Node running as a WireGuard server. 
The tunnel interface is associated to a vrf which is used to define and control what traffic is allowed to pass and how it should appear on the network.
{{% /pageinfo %}}

## Configure WireGuard Tunnel
On a Trustgrid Node navigate to Tunnels under the Network Menu and select Add Tunnel and then select WireGuard

![img](add_tunnel.png)

![img](add_WireGuard.png)

![img](WireGuard.png)

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
Optional 256 bit to provide an additonal level of security.
{{</field>}}

{{<field "VRF">}}
The Virtual Routing and Fowarding (VRF) table the WireGuard tunnel will be attached to
{{</field>}}

{{<field "MTU">}}
The maximum transmission unit (MTU) of the WireGuard tunnel interface.  Defaults of 1430 or lower recommended.
{{</field>}}

{{</fields>}}

## Example Client Config
As you configure a WireGuard tunnel interface the portal UI will automatically generate an example WireGuard configuration for the peer/client based on the information provided.

{{<tgimg src="WireGuard-tunnel-example-config.png" width="90%" caption="Example WireGuard client config">}}

{{<alert color="info">}} Note: this is only a partial config. You will likely need to adjust to the `AllowedIps` and other fields to meet the needs of the configured tunnel.
{{</alert>}}

