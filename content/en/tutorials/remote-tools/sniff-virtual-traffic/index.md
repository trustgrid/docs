---
title: "Sniff Virtual Traffic"
linkTitle: "Sniff VPN Traffic"
description: "Monitor traffic inside a virtual network tunnel"
---

## Summary
This tool allows you to monitor the traffic flowing over a VPN tunnel by "sniffing" inside the tunnel.  This allows traffic to be viewed after NATs have been applied and confirm traffic is routed to or from the expected node.

## Usage

This tool is accessed on a node's detail page under the VPN panel. Select the desired attached virtual network (note: if only on virtual network is present it will be selected by default) and click the "Tools" button in the top right.
{{<tgimg src="sniff-virtual-traffic-button.png" caption="Sniff Virtual Traffic Button" width="70%">}}

Once launched, you will be prompted to provide any filters to limit the traffic displayed. The default is to show all traffic on that virtual network tunnel.

{{<tgimg src="sniff-virtual-traffic-prompt.png" caption="Filter Prompt" width="50%">}}

{{<fields>}}
{{<field "Peer(s)">}}Use this field to filter for traffic only between the current node and the listed peer.{{</field>}}
{{<field "Protocol(s)">}} Filter to specified network protocol(s): any, tcp, udp, or icmp {{</field>}}
{{<field "Host(s)">}} Filter to see traffic matching the specified **virtual** IP address(es){{</field>}}
{{<field "Port(s)">}} Filter to see traffic matching the specified ports. This will match either a destination or source port. 
The operator on the left allows for matching based on:
- equal (=)
- not equal (!=) to exclude traffic for the specified port
- greater than (>) and less than (<)
- a port range (<>) which will show all traffic with ports between the two values

For ICMP, this field does not apply. {{</field>}}
{{<field "TCP Flags">}} Only applies to TCP traffic. It can filter based on the following flags:
- S = SYN - The start of a new TCP connection
- . = ACK - Acknowledges received data
- P = PSH - Data is ready to be read by the application
- F = FIN - Closing the connection
- R = RST - Resetting the connection

More than one flag can be specified, such as "S." to show only traffic with SYN **or** ACK flags set.
{{</field>}}
{{</fields>}}

{{<alert color="info">}}For the peer, protocol, host, and ports fields additional values can be provided to filter for more than one item.  When multiple values are listed they are combined in using an **OR** operator. Packets matching any of the listed values will be shown. {{</alert>}}

## Output

Below shows sample output from using this tool to monitor traffic on a virtual network between two nodes:

{{<tgimg src="sniff-virtual-traffic-output.png" caption="Sample Output" width="80%">}}

The output includes the following data for each packet:
{{<fields>}}
{{<field "Timestamp">}}The time the packet was seen{{</field>}}
{{<field "Peer">}}The source or destination node of the packet{{</field>}}
{{<field "Remote Virtual IP:Port">}} The virtual IP address and port of the source or destination remote endpoint{{</field>}}
{{<field "Direction Indicator">}} The symbols indicating if the packet was sent or received by the local node:
- `<` indicates the local node sent the packet to the peer
- `>` indicates the remote peer sent the packet to the local node

{{</field>}}
{{<field "Local Virtual IP:Port">}} The virtual IP address and port of the local endpoint{{</field>}}
{{<field "Packet Header Information">}}Additional metadata about the packet such as:
- protocol
- TCP flags
- sequence and acknowledgement numbers 
- data length{{</field>}}
{{</fields>}}