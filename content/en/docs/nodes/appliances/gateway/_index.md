---
Title: "Gateway"
linkTitle: "Gateway"
aliases: 
    - /docs/nodes/gateway
description: Configure gateway server and client settings to manage connections to peer nodes, and run related diagnostics
---

{{% pageinfo %}}
The Gateway panels contains settings determining how the node connects to peer nodes by either acting as a server or client.  The tunnels created by these connections are the foundation for the underlay network that enables overlay services like site-to-site VPN and L4 proxy services. 
Additionally, the Diagnostics page can be of use troubleshooting connectivity issues if layer 4 connectivity has already been confirmed.
{{% /pageinfo %}}

## Global Settings 
The header bar contains settings that apply to both the node's client and server functionality. 
{{<tgimg src="gateway-settings-header.png" alt="Gateway Settings Header" caption="Gateway Settings Header" width="95%">}}
{{<fields>}}
{{<field "Enable UDP">}} 
- On a server, this will determine if the server will listen for UDP tunnel connections on the configured [UDP port]({{<relref "./gateway-server/index.md#udp-port">}}).
- On a client, this will determine if the client will attempt to build UDP tunnels to any server with UDP enabled.

Firewalls between the client and server must permit UDP on the configured port end to end. See [Trustgrid Data Plane]({{<relref "help-center/kb/site-requirements#trustgrid-data-plane" >}}) for the requirements, including a note on [application-aware firewalls]({{<relref "help-center/kb/site-requirements#application-aware-layer-7-firewalls" >}}). See [Tunnels](#tunnels) for how the node decides a UDP tunnel is usable.

{{</field>}}
{{<field "Max Ingress">}} The maximum total amount of ingress traffic in megabits per second allowed on the node.{{</field>}}
{{<field "Max Egress">}} The maximum total amount of egress traffic in megabits per second allowed on the node.{{</field>}}
{{</fields>}} 

## Gateway Troubleshooting Tool

The Troubleshoot Gateway Traffic tool allows you to inspect live diagnostic messages about traffic between this node and its configured peers. 
{{<tgimg src="gateway-diag-troubleshoot.png" width="80%" caption="Gateway Troubleshooting Tool">}}

Clicking the "Troubleshoot Gateway Traffic" button will open the below dialog that allows for filtering the output seen in the Troubleshoot Gateway tool by peer, local or service. Selecting what to filter on and clicking "Apply" will update the output seen in the tool. Accepting the default will display all gateway messages.

{{<tgimg src="launch-troubleshoot-gateway.png" width="50%" caption="Troubleshoot Gateway Traffic Filter Dialog">}}

The output can be useful in troubleshooting why a node and a peer are not connecting.

{{<tgimg src="troubleshoot-gateway-traffic.png" width="80%" caption="Example output of Troubleshoot Gateway tool">}}

## Tunnels

A gateway client builds two kinds of tunnel to each gateway server it peers with. The portal shows a connection as **TLS** or **TLS/UDP** depending on which are in use.

### TLS Tunnels

The TLS tunnel is built first and is always present. It runs over TCP to the gateway's [Port]({{<relref "./gateway-server/index.md#port" >}}), carries traffic whenever a UDP tunnel is not available, and is what a node falls back to if one fails. A node with no usable UDP tunnel is fully functional.

Tunnels negotiate TLS 1.3 where both ends support it, falling back to TLS 1.2. See [Security]({{<relref "getting-started/security#tls-encryption" >}}).

The client pings the gateway over the tunnel every 20 seconds. If nothing arrives from the gateway for 120 seconds, six missed pings, the client closes the tunnel. It retries every 10 seconds. A network that silently stops passing traffic therefore takes up to two minutes to detect, where one that resets or refuses the connection is detected at once.

Each connection to a gateway is its own tunnel and times out on its own. This is the same whether the connection uses the default route or an additional [gateway path]({{<relref "./gateway-client/index.md#gateway-paths" >}}).

Closing the TLS tunnel takes its UDP tunnel with it, since the UDP keys are delivered over the TLS tunnel. Both rebuild together.

### UDP Tunnels

A UDP tunnel is built when **Enable UDP** is set on both the client and the server. It runs to the gateway's [UDP Port]({{<relref "./gateway-server/index.md#udp-port" >}}) and avoids the head-of-line blocking of a TCP tunnel, so it usually performs better on a lossy path. A blocked or failing UDP tunnel costs performance, not connectivity.

UDP tunnel traffic is encrypted with ChaCha20-Poly1305, independently of the TLS version the TLS tunnel negotiated. See [UDP Data Plane Encryption]({{<relref "getting-started/security#udp-data-plane-encryption" >}}).

#### Establishing a Tunnel

A UDP tunnel is never built on its own. It depends on the TLS tunnel throughout its life:

1. The client establishes the TLS tunnel to the gateway.
2. The gateway generates the UDP tunnel's encryption keys and sends them to the client over the TLS tunnel.
3. The client opens the UDP tunnel from a new source port to the gateway's UDP Port.
4. The node tests the tunnel. Traffic stays on the TLS tunnel until it passes.
5. Traffic moves to the UDP tunnel.

The gateway rotates the keys every 15 minutes, over the TLS tunnel, for the life of the tunnel. A rotation keeps the existing source port and does not repeat the test.

#### Testing a Tunnel

Step 4 confirms the path can carry the tunnel before any traffic depends on it:

- The node sends probe packets at the tunnel's full packet size, one second apart. The peer echoes each one.
- Three consecutive probes must return, within six attempts. A single dropped packet is tolerated.
- A tunnel that fails stays unused, and traffic stays on the TLS tunnel.

The test exists for paths crossing an application-aware firewall. Those firewalls pass the first packets of a flow while classifying it, then drop the rest once the traffic is marked unidentified. An untested tunnel looks healthy for that first moment and then silently stops carrying traffic. See [UDP Blocked by Layer 7]({{<relref "help-center/kb/startup-process/udp-blocked-by-app-inspection" >}}).

Both ends must run n-2.26.0 or later. Against an older peer the node activates a UDP tunnel as soon as it is built, without testing it.

#### Retries and Flapping

A UDP tunnel that fails its test, or that stops carrying traffic later, is retried on a widening interval:

| Attempt | Delay |
| --- | --- |
| 1 | 30 seconds |
| 2 | 1 minute |
| 3 | 2 minutes |
| 4 | 5 minutes |
| 5 and later | 10 minutes |

Each delay is varied slightly so a site full of nodes does not retry in lockstep. The interval resets after the tunnel has been healthy for 30 minutes.

A tunnel is treated as flapping after three healthy-to-failed cycles, or five consecutive failures, within an hour. While flapping, the node retries every 10 minutes but reports once a day, and the event states when it will report again. Flapping clears after an hour of health.

#### Events and Logging

A failed UDP tunnel raises a [Gateway UDP Tunnel Error]({{<relref "docs/alarms/event-types#gateway-udp-tunnel-error" >}}). The message distinguishes a tunnel that never established from one that established and then failed, and carries the current retry interval. Establishing a tunnel is logged on the node, so node logs show whether UDP was working before a problem started.
