---
title: Data Plane
linkTitle: Data Plane
description: Listing of peers and connectivity information for agent-based nodes
weight: 15
---

The Data Plane panel provides information about the tunnels between the agent and its gateway peers that allow for data plane communication.  
{{<tgimg src="agent-peers.png" width="90%" caption="Table of connected gateway peers and stats">}}

Each peer will show the following data:
{{<fields>}}
{{<field Status>}}This field will display a green circle if connected and a red circle if disconnected. {{</field>}}
{{<field Name>}}The name of the peer node.{{</field>}}
{{<field IP>}}The IP address the peer node is advertising for connectivity.{{</field>}}
{{<field Path>}}Name of additional path if one is defined. When blank the default WAN/internet path is implied.{{</field>}}
{{<field RTT>}}Return trip time. The agent sends a ping every 20 seconds and records the time for a response. This field shows the most recent result for each peer.{{</field>}}
{{<field Mode>}}<ul>
<li>TLS - Connections using only the Trustgrid TLS tunnel method</li>
<li>TLS/UPD - Connections using both Trustgrid TLS and UDP tunnels</li>
 </ul>
UDP is something that needs to be enabled on the gateway node. {{</field>}}
{{<field Type>}}For an agent this will always show **client** since it only makes outbound connections. Appliance-based nodes could show **server** (gateway) or **client**.{{</field>}}
{{<field Ports>}}The destination port(s) used for creating the tunnel to the peer. {{</field>}}
{{</fields>}}


{{<alert color="info">}}Organization using the Trustgrid hosted gateway for trial will not see any peers listed at this time.{{</alert>}}