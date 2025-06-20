---
Title: "Gateway"
linkTitle: "Gateway"
weight: 7
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

{{</field>}}
{{<field "Max Ingress">}} The maximum total amount of ingress traffic in megabits per second allowed on the node.{{</field>}}
{{<field "Max Egress">}} The maximum total amount of egress traffic in megabits per second allowed on the node.{{</field>}}
{{</fields>}} 

## Gateway Troubleshooting Tool

The Troubleshoot Gateway Traffic tool allows you to inspect live diagnostic messages about traffic between this node and its configured peers. 
{{<tgimg src="gateway-diag-troubleshoot.png" width="80%" caption="Gateway Troubleshooting Tool">}}

Clicking the "Troubleshoot Gateway Traffic" button will open the below diaglog that allows for filtering the output seen in the Troubleshoot Gateway tool by peer, local or service. Selecting what to filter on and clicking "Apply" will update the output seen in the tool. Accepting the default will display all gateway messages.

{{<tgimg src="launch-troubleshoot-gateway.png" width="50%" caption="Troubleshoot Gateway Traffic Filter Dialog">}}

The output can be useful in troubleshooting why a node and a peer are not connecting.

{{<tgimg src="troubleshoot-gateway-traffic.png" width="80%" caption="Example output of Troubleshoot Gateway tool">}}