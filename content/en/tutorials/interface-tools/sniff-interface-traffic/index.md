---
title: Sniff Traffic Interface Tool
linkTitle: Sniff Traffic
description: Monitor network traffic on a local interface in realtime
---

The Sniff Traffic tool allows you to monitor traffic on a local network interface in realtime within the Trustgrid portal. Only the header information such as source and destination IPs, TCP flags, sequence numbers, etc. are displayed. 

## Using Sniff Interface Traffic

1. Login to the Trustgrid portal and navigate to the Node from which you want to view traffic.
1. Select `Interfaces` under the `Network` section. Then from the dropdown select the interface whose traffic you want to monitor. {{<tgimg src="sniff-int-select.png" width="85%" caption="Selecting the desired interface">}}
1. Click the Sniff Traffic button. {{<tgimg src="sniff-int-button.png" width="85%">}}
1. Set the Interface and Filters to match the traffic you wish to monitor. {{<tgimg src="sniff-int-filter.png" width="60%">}}
    {{<fields>}}
        {{<field "Interface">}}The interface name. This should be auto-populated with the interface selected above. If not, either click the **Reset Default** button or erase the interface name and select from the list.{{</field>}}
        {{<field "Filter">}}The tool utilizes [TCPDump filtering syntax](https://www.tcpdump.org/manpages/pcap-filter.7.html), which can help isolate interesting traffic. The below filter would show only ICMP and TCP traffic. See the [useful filters](#useful-filters) below for a quick introduction. {{</field>}}
    {{</fields>}}
1. Click `Execute` to view the captured traffic {{<tgimg src="sniff-output.png" width="85%" caption="Example output of sniffing interface traffic">}}


## Useful Filters

{{<readfile "/tutorials/interface-tools/traffic-capture/useful-tcpdump-filter.md">}}