---
title: "DNS"
date: 2023-02-14
aliases: 
    - /docs/nodes/vpn
description: Configure a DNS 
---

Nodes can act as domain name servers (DNS) for requests ingressing the virtual network on the node if [DNS is configured at the virtual network]({{<ref "/docs/domain/virtual-networks/dns">}}).

{{<tgimg src="list.png" width="90%" caption="VPN DNS Settings page">}}

## DNS Status
This setting must be set to `Enabled` on the node to process requests targeting the configured [virtual network DNS server IP]({{<ref "/docs/domain/virtual-networks/dns#dns-settings">}}). 

## Upstream Servers
If configured, the node will use these servers to resolve DNS requests for [zones not configured at the virtual network level]({{<ref "/docs/domain/virtual-networks/dns#dns-zones">}}). 

{{<fields>}}
{{<field "Host IP Address">}}The upstream DNS server's IP address{{</field>}}
{{<field "Host Port">}}The upstream DNS server's port (usually 53){{</field>}}
{{<field "Description">}}(Optional) a user friends description of the server{{</field>}}
