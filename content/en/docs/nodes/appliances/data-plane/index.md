---
tags: ["node"]
title: "Data Plane"
linkTitle: "Data Plane"
description: Listing of peers and connectivity testing tools for appliance-based nodes
aliases: 
    - /docs/nodes/data-plane
---

{{% pageinfo %}}
The Data Plane panel shows a list of all statistics known peers, their connection status, provides tools for testing connectivity and can display statistics for the connection between the node and the peer.
{{% /pageinfo %}}



## Select Data Plane Columns
Select **Columns** to choose which selectable columns appear in the peers table. All options are enabled by default except **TLS Version** and **TLS Cipher**.

| Column | Description | Enabled by default |
| --- | --- | --- |
| Path | Network path used to reach the peer, when available. | Yes |
| RTT | Round-trip time observed when the panel is opened or refreshed. | Yes |
| Mode | Connection mode: TLS or TLS/UDP. | Yes |
| TLS Version | TLS version negotiated for this peer connection. | No |
| TLS Cipher | Cipher suite negotiated for this peer connection. | No |
| Type | Peer role: `SERVER` for gateway or `CLIENT` for edge. | Yes |
| Port(s) | Ports used to establish the peer connection. | Yes |
| ISP | Internet service provider associated with the peer's public IP address. | Yes |
| Location | Geographic location associated with the peer's public IP address. | Yes |

{{<alert color="info">}}The TLS columns are available only on nodes running the [September 2026 Major Appliance release]({{<relref "release-notes/node/2026-09" >}}) or newer. The TLS columns appear only when at least one route reports TLS information.{{</alert>}}

## Export List of Peers
You can export a CSV formatted list of all peers by clicking **Actions** > **Export**.
{{<tgimg src="data-plane-export.png" alt="Export option in the Data Plane Actions menu" caption="Export button" width="25%">}}

The CSV file will include columns for:
* Connection status (true/false)
* Peer name
* Peer IP
* Path (if applicable)
* RTT - Return Trip Time or latency observed when the panel was opened or refreshed
* Mode - TLS or TLS/UDP
* Type - The role of the node to peer. SERVER = gateway, CLIENT = edge device
* Ports - The ports used to establish the connection
* TLS Version - The negotiated TLS version for the peer, when reported and selected
* TLS Cipher - The negotiated cipher suite for the peer, when reported and selected

## Tools
The far right column as links to launch the tools listed below:
{{<tgimg src="data-plane-tools.png" alt="Tool actions for a Data Plane peer" caption="Data Plane Tools" width="25%">}}
### Data Plane Ping

{{<tgimg src="gateway-ping-modal.png" alt="Data Plane Ping dialog for a selected peer" caption="Data Plane Ping dialog" width="400px">}}

The gateway ping tool will launch a ping inside the data plane tunnel between the node and the gateway, similar to ICMP ping.  This can provide a more realtime view of latency between node and gateway compared to the RTT column (only loaded on refresh) and the historical data which provides a minute by minute view.

{{<tgimg src="ping-output.png" alt="Ping results between the node and selected peer" caption="Ping Output" width="50%">}}

### Data Plane Performance Tool

The gateway performance tool attempts to estimate the available **upload bandwidth** between current node and the selected peer. 

{{<tgimg src="gateway-performance-modal.png" alt="Data Plane performance test dialog for a selected peer" caption="Data Plane Performance dialog" width="400px">}}

By default this tool will attempt to send 10MB of data to the destination device. This then used to calculate the available bandwidth.  Sending more data provides a more accurate estimate but be aware that other devices using the same internet circuit may experience degraded performance while the test is running. 

{{<tgimg src="performance-output.png" alt="Gateway performance test results" caption="Gateway Performance Output" width="50%">}}

The periods are posted for about every 1MB of data sent. If you see exclamation marks this can indicate lost packets.

### Trace Route and MTR
Both of the links here will populate the [interface tool of the same name]({{<ref "/docs/nodes/appliances/interfaces">}}) with the selected peer's IP address and TCP port. This allows you to trace the route or run MTR diagnostics directly between the current node and the selected peer.

{{<alert color="info">}}This is most effective if run from the client/edge node to a gateway since the gateway is the only one actually listening on that port.  If run in the other direction (gateway -> client) it is likely any firewall between will drop the connections and the client will definitely not respond on that port. {{</alert>}}

### Peer Disconnect
The disconnect button will disconnect the node from the selected peer. The tunnel should be reestablished automatically but there will be a period of disruption. 
After clicking the link, you will be prompted to confirm the action.
{{<tgimg src="disconnect-confirm.png" alt="Confirmation dialog for disconnecting from a selected peer" caption="Disconnect prompt" width="45%">}}

{{<alert color="warning">}} If the peer node is the primary site and a DR site with a higher metric exists, some traffic may be routed to the DR site.{{</alert>}}
## View Latency to Peers

To view latency data between two nodes, select either the edge or gateway node, and then from the peers table, select the node to view.

{{<alert color="info">}}Gateway nodes will list edge nodes in their peers table{{</alert>}}

{{<tgimg src="peer-list.png" alt="Data Plane peers table with a peer selected" caption="Data Plane peer list" width="50%">}}

Once a peer is selected, the monitoring section will populate with relevant data.

{{<tgimg src="monitoring.png" alt="Monitoring details for the selected Data Plane peer" caption="Data Plane peer monitoring" width="50%">}}

Hop data is only available for nodes that have hop monitoring enabled. See [Monitoring Network Hops to Peers]({{<ref "/tutorials/gateway-tools/monitoring-network-hops-to-peers" >}}).
