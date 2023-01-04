---
tags: ["applications", "ztna"]
title: "RDP App"
date: 2023-01-05
---

A RDP app is a ZTNA application that allows remote access to an internal RDP server. RDP servers can be hosted internally but exposed to authorized users.

Fields (convert these to field defs):
section: general

- Name - app name
- Description - app description
- Icon - the application's icon (optional) to show in the application dashboard

section: Connectivity

- Connectivity type:
- - local to gateway - the application is hosted on the same network as the gateway
- - remote node - the application is hosted on an edge node's network
- - virtual network - the application is accessible over the Trustgrid virtual network from the ZTNA gateway
- ZTNA Gateway - the ZTNA gateway node that will be used to connect to the application
- Destination Node - only available if connectivity type is Remote Node. The edge node with access to the application
- VRF - only available if connectivity type is Remote Node. The VRF used to connect to the application.

- Internal server hostname or IP - the internal hostname or IP address of the SSH server and port number

- Virtual server URL - the internal URL of the application
- Virtual Network - only available if connectivity type is Virtual Network. The virtual network that will be used to connect to the application
- Client Virtual IP - only available if connectivity type is Virtual Network. The virtual IP address that will be used to connect to the application

section: Security

- Identity Provider - the Identity Provider [link] to authenticate users

SHOW HOW TO CREATE RDPAPP
