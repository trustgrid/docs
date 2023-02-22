---
categories: ["node"]
tags: ["tunnels", "networking"]
title: "Tunnels"
linkTitle: "Tunnels"
weight: 22
---

### Tunnels Overview


Trustgrid Supports configuring various tunneling protocols in order to establish secure connectivity between any appliance that supports the underlying protocol.
The supported tunnels are listed below. 

- IPsec - Internet Protocol Security supporting ikev1 and ikev2 tunnels 
- GRE - Generic Routing Encapsulation used for setting up a direct point to point network connection
- Vnet - The virtual network tunnel is the Trustgrid Proprietary VPN tunneling protocol and therefore can only be used to establish connectivity to other Trustgrid nodes in the organization
- Wireguard - The latest open source tunneling protocol aiming to provide better performance than the traditional ipsec or openvpn protocols 

"Wireguard" is a registered trademark of Jason A. Donenfeld 


