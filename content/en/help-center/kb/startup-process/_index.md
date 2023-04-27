---
title: Edge Node Startup
linkTitle: Node Startup
weight: 30
description: A step by step description of how a healthy Edge Node starts up and connects to the Trustgrid system.
---

## Startup Process

1. ARP for the default gateway
1. DNS request for A record for gatekeeper.trustgrid.io & zuul.trustgrid.io
    1. The response will be IPs in the Trustgrid AWS networks, `35.171.100.16/28` (us-east-1) or `34.223.12.192/28` (us-west-2)
1. The node will initiate a connection to port 8443 on the IP addresses returned and start a TLS handshake.
1. After a successfully TLS handshake, the node will show as connected in the Trustgrid Portal

