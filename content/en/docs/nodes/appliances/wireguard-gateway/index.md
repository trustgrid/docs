---
title: WireGuard Gateway
aliases: 
    - /docs/nodes/ztna-gateway
    - /docs/nodes/appliances/ztna-gateway
description: Configure the node to act as a Wireguard server for Wireguard tunnels.
---

## WireGuard Endpoint

The WireGuard endpoint is used to provide connectivity for [WireGuard Tunnels]({{<ref "docs/nodes/appliances/tunnels/wireguard" >}}). 

{{<fields>}}
{{<field "Enabled">}}When enabled, this node will listen for WireGuard traffic.{{</field>}}
{{<field "Public FQDN">}}The IP or FQDN of the node or the load balancer in front of the node.{{</field>}}
{{<field "Port">}}The port to listen on.{{</field>}}
{{<field "Public Key">}}The node's WireGuard public key. This can be generated or imported using the actions dropdown. Note that regenerated the key will disconnect existing clients and require users to reconfigure their WireGuard connection.{{</field>}}
{{</fields>}}

*“WireGuard” is a registered trademark of Jason A. Donenfeld.*
