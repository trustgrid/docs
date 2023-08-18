---
title: "DNS"
date: 2023-2-14
---

Nodes can act as DNS servers for requests ingressing the the virtual network on the node if [DNS is configured at the virtual network]({{<ref "/docs/domain/virtual-networks/dns">}}). 

![img](list.png)

## DNS Status
This setting must be set to enabled for the node to process requests targeting the configured [virtual network DNS server IP]({{<ref "/docs/domain/virtual-networks/dns#dns-settings">}}). 

## Upstream Servers
If configured, the node will use these servers to resolve DNS requests for [zones not configured at the virtual network level]({{<ref "/docs/domain/virtual-networks/dns#dns-zones">}}). 

{{<fields>}}
{{<field "Host IP Address">}}The upstream DNS server's IP address{{</field>}}
{{<field "Host Port">}}The upstream DNS server's port (usually 53){{</field>}}
{{<field "Description">}}(Optional) a user friends description of the server{{</field>}}
{{</fields>}}