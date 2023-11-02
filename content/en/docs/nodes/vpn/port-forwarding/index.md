---
title: "Port Forwarding"
date: 2023-5-8
---

Port forwarding allows a node to expose an [L4 TCP service]({{<ref "services">}}) on the virtual network.

{{<tgimg src="list.png" caption="Port Forward listing" width="90%">}}


{{<fields>}}
{{<field "Virtual IP Address">}}The IP to expose the service on the virtual network.{{</field>}}
{{<field "Virtual Port">}}The port to expose the service on the virtual network.{{</field>}}
{{<field "Local Service">}}The [service]({{<relref "/docs/nodes/services">}}) to expose. Alternately, you can specify an IP and port.{{</field>}}
{{</fields>}}