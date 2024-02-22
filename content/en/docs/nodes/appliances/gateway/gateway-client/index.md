---
linkTitle: Client
Title: Gateway Client Settings
descriptions: Configure settings for the Trustgrid gateway client
weight: 20
aliases: 
    - /docs/nodes/gateway/gateway-client
description: Configure client settings 
---

Settings on this page determine how a node's client connection to gateway peers behave.

## Settings

{{<fields>}}
{{<field "UDP Enabled" >}}
This settings can be configured on both the [gateway server](../gateway-server) and [gateway client](../gateway-client) sub-panel.

- On a server, this will determine if the server will listen for UDP tunnel connections on the configured [UDP port](#udp-port).
- On a client, this will determine if the client will attempt to build UDP tunnels to any server with UDP enabled.

**UDP tunnels are only attempted if both the client and the server have the UDP Enabled field set to Enabled**

{{</field >}}
{{<field "Max Egress Mbps" >}}
The egress bandwidth limit for the gateway. Connections will be throttled when this limit is reached.
{{</field >}}
{{<field "Monitor Network Hops to Peers" >}}
Whether to [monitor latency to peers]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers">}}) through this gateway. This can have a performance impact and is not recommended for high-traffic gateways.
{{</field >}}
{{<field "Connectivity to Public Gateways">}}
Options are `Allowed` or `Denied`. If set to `Denied` this will cause the node to not attempt connections to public gateways. This might be desired if you want the node to only connect to configured private gateways. Or if you have private gateways that do not need to connect to the public gateways in your organization.
{{</field>}}
{{</fields>}}

{{<tgimg src="gateway-client-settings.png" width="85%" caption="Gateway Client settings" alt="screenshot of gateway client settings">}}

#### Gateway Paths

Allows you to define alternate paths to a gateway server

{{<fields>}}
{{<field "Name" >}}
A name for the path.
{{</field >}}

{{<field "Gateway Node" >}}
Gateway for which the path is applicable.
{{</field >}}

{{<field "Host IP" >}}
Destination IP address for the path.
{{</field >}}

{{<field "Host Port" >}}
Destination port for the path.
{{</field >}}

{{<field "Local IP" >}}
Use this local IP as the source IP for the connection to the gateway.
{{</field >}}

{{<field "Use as Default" >}}

- True - Will not attempt to connect to the configured Gateway Node using the WAN interface IP and Default Gateway path.
- False - Will attempt to connect to the Gateway node using both this defined path **and** the WAN Interface IP and Default Gateway path.
  {{</field >}}
  {{</fields>}}

{{<tgimg src="gateway-paths.png" width="85%" caption="Example Gateway Path" alt="Screenshot of the gateway paths table">}}

### Add A Gateway Path

1. Click the Add Path link
1. Fill in the fields as desired.
1. Click the green check mark to save the path.
1. Optionally, repeat with additional paths.
1. Click save.

### Delete a Gateway Path

1. Click the X to the right of the desired path.
1. Click save.
