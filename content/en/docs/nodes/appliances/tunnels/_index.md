---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "Tunnels"
linkTitle: "Tunnels"
aliases: 
    - /docs/nodes/tunnels
description: Configure manual tunnels with non-Trustgrid devices via standard protocols  
---

### Tunnels Overview


Trustgrid Supports configuring various tunneling protocols to establish secure connectivity between any appliance that supports the underlying protocol.
The supported tunnels are listed below. 

- IPsec - Internet Protocol Security supporting IKEv1 and IKEv2 tunnels 
- GRE - Generic Routing Encapsulation used for setting up a direct point to point network connection
- Vnet - The virtual network tunnel is the Trustgrid Proprietary VPN tunneling protocol and therefore can only be used to establish connectivity to other Trustgrid nodes in the organization
- Wireguard - The latest open source tunneling protocol aiming to provide better performance than the traditional ipsec or openvpn protocols 

"Wireguard" is a registered trademark of Jason A. Donenfeld 


