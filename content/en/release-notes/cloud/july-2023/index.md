---
title: July 2023 Release Notes
linkTitle: 'July 2023 Release'
type: docs
date: 2023-07-31
description: July Release focusing on general improvements and fixes
---

## Network Interface Panel Changes

### Interface Details Copy All Button
Frequently Trustgrid users need to share the interface configuration information with people who do not have access to the Trustgrid portal.  There is now a **Copy All** button that will load all of these details into your clipboard.

{{<tgimg src="interface-copy-all.png" caption="Copy All button" alt="Button to copy all interface details" width="60%" >}}

Example WAN interface output:
```
IP: 172.16.11.31/24
Hardware Address: 00:50:56:8e:03:54
Gateway: 172.16.11.1
DNS1: 172.16.11.4
DNS2: 4.2.2.2

```
### Public IP Removed from WAN Interface
The WAN interface used to include a field labeled "Public IP" that was the IP address the Trustgrid control plane saw as the source IP for traffic from the node.  This would be the IP address used on the internet after any NATs where applied.  

Some users were confused by this IP and believed it was the assigned IP address for the WAN interface which is not always true.  Therefore we have removed this from the WAN interface panel.  

The observed Public IP is still visible in the [Infovisor]({{<ref "/docs/nodes/infovisor">}}).

## Data Plane Panel
### Export Data Plane Peers

You can now export a CSV-formatted list of peers from the [Data Plane panel]({{<ref "/docs/nodes/data-plane">}}). 
{{<tgimg src="/docs/nodes/data-plane/data-plane-export.png" caption="Export button" width="25%">}}

### List Gateway Ports
The [Data Plane Panel]({{<ref "/docs/nodes/data-plane">}}) now lists the Gateway Port for each tunnel. This is the port the client node uses to connect to the listed gateway. 

While this is typically port 8443,, it can be configured differently and having that info visible aids in troubleshooting.

## BGP Status Tool
A new [BGP Status tool]({{<ref "/docs/nodes/bgp#bgp-status">}}) has been added to the BGP panel. This tool displays the current BGP peer status including peer connection status, uptime, routes received and advertised, and more.
{{<tgimg src="/docs/nodes/bgp/bgp-status-output.png" caption="BGP Status" width="80%">}}

## Flow Logs Advanced Search Improvement
The [flow logs advanced search]({{<ref "/help-center/flow-logs#advanced-search">}}) now allows selecting clusters as the source or destination nodes.  This enables search for flows that were sent to or received from a cluster, rather than individual nodes.

## Trustgrid Resource Names (TGRN) in Infovisor
When creating [Resource Scoped Policies]({{<ref "/docs/user-management/policies#resource-scoped-policies">}}) to limit a user's visibility to specific nodes or clusters you specify resources using Trustgrid Resource Names (TGRNs). Prior to this release it was necessary to generate the TGRN for by combining the common prefix with the node's UID.  With this release the TGRN is available in the node's [Infovisor]({{<ref "/docs/nodes/infovisor">}}) panel with a button to easily copy the TGRN into your clipboard. 

{{<tgimg src="infovisor-tgrn.png" caption="Copy TGRN field on Infovisor" width="70%" >}}

## Other Notable Fixes
The below issues have been resolved in this release:

* Adding and saving multiple tags use to frequently lead to an error that caused an endless status spinner and would fail to save one or more of the tags. 
* Overview graphs were not aggregating statistics correctly for view of 6 hours or longer.* 