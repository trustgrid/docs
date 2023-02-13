---
title: "WireGuard ®"
date: 2023-2-14
---

Nodes may enable a WireGuard server to expose the virtual network to WireGuard clients.

## Configuration

![img](wireguard.png)

{{<field-def "Enabled">}}When enabled, this node will accept WireGuard client connections{{</field-def>}}
{{<field-def "Port">}}The port clients should connect to{{</field-def>}}
{{<field-def "Remote Traffic Forwarding">}}Whether the node should forward traffic to the virtual network{{</field-def>}}

## Key Management

Clients will need the server's public key to connect. This can be generated (using the `Regenerate` action) or can be derived from the private key (imported using the `Import` action).

{{<alert>}}Changing the server's private key will disconnect existing clients and require their reconfiguration.{{</alert>}}

## Clients

Clients must be added before they can connect.

![img](wg-client.png)

Each client will need to generate their own private and public key. Most WireGuard clients provide tools to do this.

{{<field-def "Name">}}The name of the client{{</field-def>}}
{{<field-def "Pre-shared Key">}}(Optional) The client's pre-shared key{{</field-def>}}
{{<field-def "Public Key">}}The client's public key{{</field-def>}}
{{<field-def "Allow Internet">}}Whether this client should be allowed to access the public network through this server{{</field-def>}}

An inside NAT should be created to map the client's local IP to an IP on the virtual network. The virtual network CIDR should not be a part of any existing routes in the domain. The IP should be included in an outside NAT network group, so that an ACL can allow the traffic.

An outside NAT should be created that includes the virtual network used on the gateway side.

_"WireGuard" is a registered trademark of Jason A. Donenfeld._
