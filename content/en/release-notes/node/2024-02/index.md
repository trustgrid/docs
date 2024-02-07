---
title: February 2024 Appliance Release Notes
linkTitle: February 2024
type: docs
date: 2024-02-01
description: "Release notes for the February 2024 Trustgrid Appliance release focused on support for agents"
---
{{< node-release package-version="1.5.20240131-1959" core-version="20240131-005233.4838689" release="n-2.18.0" >}}

## Removal of Local UI System
The local UI functionality allowed users with direct physical access to an appliance network port to connect to a lightweight web server on the appliance to change the WAN interface IP and see the connectivity status. Several years ago, this solution was augmented with the [Console Utility]({{<relref "/tutorials/local-console-utility">}}) which has proved easier to use and has been expanded to provide much more functionality. The local UI has now been removed to simplify the appliance and focus efforts on the Console Utility.

## VPN Sniff Improvements
This release improves the [Virtual Network Sniff]({{<relref "/tutorials/remote-tools/sniff-virtual-traffic">}}) tool to allow more flexibility matching based on ports.  We now support port matching based on:
- equal (=)
- not equal (!=) to exclude traffic for the specified port
- greater than (>) and less than (<)
- a port range (<>) which will show all traffic with ports between the two values
{{<tgimg src="vpn-sniff-port-opts.png">}}

## Agent Support
Gateways running this version will be the first to support the new lightweight Trustgrid agent-based nodes. More details will be announced once the agent has been officially released. 

## Other Improvements
- The [JVM Heap metric]({{<relref "/docs/nodes/appliances/metrics#jvm-heap">}}) is now exposed via the [Trustgrid SNMP OID]({{<relref "/docs/nodes/appliances/snmp/oids">}})
- The `dig` and `nslookup` commands are now available from the [Console Network Tools Shell]({{<relref "/tutorials/local-console-utility/troubleshooting#network-tools-shell">}})
- [Remote Console Registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) no longer forces upgrading to the latest version. This makes the process of registration quicker, though it is still recommended to [upgrade to the latest version]({{<relref "/tutorials/management-tasks/upgrade-nodes">}}) after the node is online. 

## Other Fixes
- If a [BGP]({{<relref "docs/nodes/appliances/bgp">}}) server received a route removal update, it would remove the OS route but sometimes not update the BGP table. Then if the same route was added back, the BGP server would fail to recreate the OS route. This release resolves this issue. 
- Resolves an issue where traffic egressing a Virtual Network would not evaluate the specificity of interface routes when determining the next hop. For example, if you had two routes (172.16.22.0/24 and 172.16.22.100/32) with different next hops, the more specific /32 route next hop was not utilized. 
- Resolves an issue with [Access Policies]({{<relref "/docs/domain/virtual-networks/access-policy">}}) using the reject action.  The reject response (e.g. RESET for TCP) generated had the source and destination IPs transposed. This release fixes the source/destination IP placement in reject responses.
