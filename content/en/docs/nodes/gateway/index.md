---
Title: "Gateway"
linkTitle: "Gateway"
weight: 7
---

{{% pageinfo %}}
Gateway nodes accept TLS tunnel connections from other nodes in the domain. They are responsible for routing traffic between nodes.
{{% /pageinfo %}}

Gateways are only configurable at the node level. Configuring a gateway requires `nodes::configure:gateway` permissions.

## Gateway Types

A gateway can be one of three types:

- `public` - allows all connections from edge nodes in the domain
- `private` - allows only allow-listed clients to connect
- `hub` - allows connections from edge nodes and other gateway nodes

A public gateway should be secured with a firewall to prevent unauthorized access.

## Configuration

Navigate to a node and select `Gateway` under the `System` section.

### Server Settings

{{<tgimg src="public-config.png" caption="Gateway Server settings" alt="input table with all Trustgrid gateway server options" width="90%">}}

{{<fields>}}
{{<field "Status" >}}
Either `Enabled` or `Disabled` - when enabled, the node will listen and accept connections from other nodes.
{{</field >}}

{{<field "UDP Enabled" >}}
Whether to allow UDP tunnels to be established through this gateway.
{{</field >}}

{{<field "Public IP or DNS" >}}
The public IP or DNS name of the gateway. This is the address that other nodes will use to connect to this gateway.
{{</field >}}

{{<field "Port" >}}
The port that the gateway will listen on. This is the port that other nodes will use to connect to this gateway.
{{</field >}}

{{<field "UDP Port" >}}
The port that the gateway will listen on for UDP tunnels. This is the port that other nodes will use to connect to this gateway.
{{</field >}}

{{<field "Max Ingress Mbps" >}}
The ingress bandwidth limit for the gateway. Connections will be throttled when this limit is reached.
{{</field >}}

{{<field "Certificate" >}}
When using DNS for gateway connections, you may specify an uploaded [certificate]({{ ref "/docs/certificate" }}) to use for TLS connections.
{{</field >}}

{{<field "Gateway Type" >}}
The type of gateway. See [Gateway Types](#gateway-types) for more information.
{{</field >}}

{{<field "Monitor Network Hops to Peers" >}}
Whether to monitor latency to peers through this gateway. This can have a performance impact and is not recommended for high-traffic gateways.
{{</field >}}
{{</fields>}}

#### Gateway Clients
Private gateways only allow connectivity from listed and enabled clients. To add a client, use the typeahead textbox at the bottom of the clients table and select the desired node.
{{<fields>}}
{{<field Client>}}Name of a node that should connect to this private gateway{{</field>}}
{{<field Enabled>}}Values are Enabled _(default)_ or Disabled. If set to Enabled the client node will attempt to connect.{{</field>}}
{{</fields>}}


### Client Settings
Settings in this section define how the node connects to gateway servers as a client

{{<tgimg src="clients.png" caption="Gateway client settings" width="90%" alt="table of gateway client settings">}}


{{<fields>}}
{{<field "Max Egress Mbps" >}}
The egress bandwidth limit for the gateway. Connections will be throttled when this limit is reached.
{{</field >}}
{{</fields>}}

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
* True - Will not attempt to connect to the configured Gateway Node using the WAN interface IP and Default Gateway path.
* False - Will attempt to connect to the Gateway node using both this defined path **and** the WAN Interface IP and Default Gateway path.
{{</field >}}
{{</fields>}}

## Gateway Tools

To see detailed messages about gateway traffic, select the `Troubleshoot Gateway Traffic` option from the services menu under `Gateway Tools`.

{{<tgimg src="launch-troubleshoot-gateway-traffic.png" caption="Troubleshoot Gateway Traffic dialog" alt="Troubleshoot gateway Traffic dialog with options of Local, Peer and Service" width="50%">}}

This will open a new window with live diagnostic messages about traffic

{{<tgimg src="troubleshoot-gateway-traffic.png" caption="Output of Troubleshoot Gateway tool" alt="terminal output showing gateway log messages between two Trustgrid nodes" width="80%">}}
