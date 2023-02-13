---
title: "DNS"
date: 2023-2-14
---

Nodes can act as DNS servers for the virtual network. Find DNS definitions and change them under the `DNS` section of the VPN configuration.

![img](list.png)

## Upstream Servers

Upstream servers will be used to answer DNS queries for domains that are not in the virtual network.

{{<field-def "Host IP Address">}}The upstream DNS server's IP address{{</field-def>}}
{{<field-def "Host Port">}}The upstream DNS server's port (usually 53){{</field-def>}}
