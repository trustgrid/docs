---
title: ZTNA Gateway
description: Configure a node to allow end-user connectivity for ZTNA applications
aliases: 
    - /docs/nodes/ztna-gateway
description: Configure the node to act as a Zero Trust Network Access server
---

Navigate to a node, then select `ZTNA Gateway` under the `Network` section.

![ztna-gateway](ztna-gateway.png)

## HTTPS Endpoint

{{<fields>}}
{{<field "Enabled">}}When enabled, this node will listen for ZTNA traffic.{{</field>}}
{{<field "Public FQDN">}}The IP or FQDN of the node or the load balancer in front of the node. If using a domain name without a load balancer, be sure to select the appropriate [certificate]({{<ref "certificates">}}).{{</field>}}
{{<field "Port">}}The port to listen on. Users will always connect on port 443, so only change this if the node is behind a load balancer.{{</field>}}
{{<field "Certificate">}}The certificate to use for TLS traffic. Only needed if the node is to be directly accessed by users (i.e., without a load balancer) and the FQDN is a domain name.{{</field>}}
{{</fields>}}

## WireGuard Endpoint

The WireGuard endpoint is used to provide connectivity for WireGuard Tunnels. 

{{<fields>}}
{{<field "Enabled">}}When enabled, this node will listen for WireGuard traffic.{{</field>}}
{{<field "Public FQDN">}}The IP or FQDN of the node or the load balancer in front of the node.{{</field>}}
{{<field "Port">}}The port to listen on.{{</field>}}
{{<field "Public Key">}}The node's WireGuard public key. This can be generated or imported using the actions dropdown. Note that regenerated the key will disconnect existing clients and require users to reconfigure their WireGuard connection.{{</field>}}
{{</fields>}}

*“WireGuard” is a registered trademark of Jason A. Donenfeld.*
