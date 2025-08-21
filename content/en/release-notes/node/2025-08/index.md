---
title: August 2025 Major Appliance Release Notes
linkTitle: August 2025 Major
type: docs
date: 2025-04-21
description: "Release notes for the August 2025 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20250813-2291" core-version="20250813-201825.ef18f9a" release="n-2.23.0" >}}

## DNS Health
Appliances being able to successfully resolve DNS is important for their overall functionality and reliability. Ensuring that DNS queries are properly handled can prevent a range of issues, especially regarding connections to the Trustgrid control plane. 

To ensure newly provisioned nodes can connect, even if their provided DNS servers are not immediately reachable, the node has a DNS cache for critical control plane services. Previously this was only updated on startup if the DNS servers were accessible.  

With this release the DNS Health Check, located in the [Configuration Status section of Infovisor]({{<relref "/docs/nodes/shared/infovisor/#configuration-status---appliances-only">}}), has been enhanced to update the cache more frequently. If the DNS servers can be contacted successfully, the cache will be populated with the resolved IP addresses, improving the chances of successful connections if future DNS disruptions occur. The health check also runs daily, ensuring that the cache remains up-to-date.

## Improvements
- Update the APT security packages to those available on 2025-07-07
- The interfaces Flows tool now has `Any` option for protocol to allow for matching against ICMP, TCP, and UDP in the same search. {{<tgimg src="flows-any-protocol.png" width="60%" caption="Set Protocol to `Any`" >}}
- Activation codes for [remote registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) now display in a higher-contrast to make reading easier in some environments.

## Fixes
- Changing WAN IPs using the [Try method]({{<>}})
- Fixed an issue that prevented MTU changes on interfaces configured by DHCP.
- Removed a logging statement on AWS nodes about "ARP'ing for cluster IPs" since custer IPs aren't supported in AWS.
- 
