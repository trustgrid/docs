---
title: August 2025 Major Appliance Release Notes
linkTitle: August 2025 Major
type: docs
date: 2025-08-21
description: "Release notes for the August 2025 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20250813-2291" core-version="20250813-201825.ef18f9a" release="n-2.23.0" >}}

## DNS Health
Appliances being able to successfully resolve DNS is important for their overall functionality and reliability. Ensuring that DNS queries are properly handled can prevent a range of issues, especially regarding connections to the Trustgrid control plane. 

To ensure newly provisioned nodes can connect, even if their provided DNS servers are not immediately reachable, the node has a DNS cache for critical control plane services. Previously this was only updated on startup if the DNS servers were accessible.  

With this release the DNS Health Check, located in the [Configuration Status section of Infovisor]({{<relref "/docs/nodes/shared/infovisor/#configuration-status---appliances-only">}}), has been enhanced to update the cache more frequently. If the DNS servers can be contacted successfully, the cache will be populated with the resolved IP addresses, improving the chances of successful connections if future DNS disruptions occur. The health check also runs daily, ensuring that the cache remains up-to-date.

## Hop Monitoring 
The [previous release]({{<relref "/release-notes/node/2025-04/#hop-monitoring-improvements">}}) attempted to increase data gathered via [Hop Monitoring]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers">}}) by increasing the frequency of checks and attempting to send data in the SYN packet to look for MTU issues in the network path.  This was found to cause issues on some firewalls. With this release we are reverting those changes back to every 20 seconds and a 0 byte payload.  These settings can been [changed in the portal]({{<relref "/docs/nodes/appliances/gateway/gateway-client#hop-monitoring-settings">}}) if more detailed data is desired.


## Improvements
- Update the APT security packages to those available on 2025-07-07
- The interfaces Flows tool now has `Any` option for protocol to allow for matching against ICMP, TCP, and UDP in the same search. {{<tgimg src="flows-any-protocol.png" width="60%" caption="Set Protocol to `Any`" >}}
- Activation codes for [remote registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) now display in a higher-contrast to make reading easier in some environments.
- Prior to this release, nodes could rarely end up in a state where they were not connected to the Trustgrid control plane but did not attempt to reconnect. This has been addressed to ensure better connectivity.
- The `Repo Health` connectivity check has been improved to reduce false positives. 

## Fixes
- Addresses an issue where some connections across [Port Forwards]({{<relref "/docs/nodes/appliances/vpn/port-forwarding">}}) could end up leaving connections stuck in a `CLOSE-WAIT` state. This could lead to a conflict if the ephemeral source port of that connection was attempted to be used again. This release also resolve a null pointer exception that could occur in the same scenario.
- Changing WAN IPs using the [Try method]({{<relref "/tutorials/wan-interface-ip">}}) now adds an audit entry in the [Changes](<{{<relref "/docs/nodes/shared/changes">}}>) section of the appliance. Additional changes were made to improve the consistency of the `try` experience.
- Fixes an issue that prevented MTU changes on interfaces configured by DHCP.
- Removed a logging statement on AWS nodes about "ARP'ing for cluster IPs" since custer IPs aren't supported in AWS.
- Azure appliances that were not part of a cluster would attempt to load the Azure API credentials required for managing [route]({{<relref "/tutorials/deployments/deploy-azure/route-failover">}}) and [ip]({{<relref "/tutorials/deployments/deploy-azure/ip-failover">}}) failover methods, causing an error in the debug logs.  This has been addressed.
