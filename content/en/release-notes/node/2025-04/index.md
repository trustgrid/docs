---
title: April 2025 Major Appliance Release Notes
linkTitle: April 2025 Major
type: docs
date: 2025-04-21
description: "Release notes for the April 2025 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20250415-2266" core-version="20250415-213243.4a224c3" release="n-2.22.0" >}}

## Hop Monitoring Improvements
This release makes some significant improvements to Hop Monitoring.  Previously the Hop Monitoring ran every 20 seconds providing only 3 samples per minute.  This was not enough to provide accurate data for some applications.  This release increases the sample rate to every 5 seconds.  This allows the Hop Monitoring to provide accurate data for applications that require a high sample rate.  

Additionally, the Hop Monitoring by default will attempt to send data payload in the SYN packets equal the 1440 bytes or the MTU of the WAN interface minus 60 byte.  The goal is to detect hops in the path with smaller than standard MTU that can lead to retransmissions and performance problems.  It is important to not that this check only runs in a single direction, from the edge node to the gateway node. As edge nodes do not listen a check in the other direction is not possible.

Finally, this release adds the ability for a gateway to request hop monitoring from all it's connected edge nodes.  This can make it easier to enable hop monitoring throughout an organization without having to enable it on each edge node individually.  As part of this change we have also changed the hop monitoring settings on the client page from disabled/enabled to:
- Always - This is equivalent to the previous enabled setting.  The node will attempt hop monitoring to all its connected gateways.
- Only when peer requests - This is the default setting.  The node will attempt hop monitoring to any gateway configured to request it. 
- Never - This is equivalent to the previous disabled setting.  The node will not attempt hop monitoring to any gateway, event if it is configured to request it.

## Cluster Improvements
### Cluster Communication Fixes
This release makes a number of changes to address issues with cluster communication.  These changes include:
- Azure Cluster IP addresses could end up in a state where neither member could set the IP configuration. Additionally any issue configuring the cluster IP logged every minute.  This release fixes both of these issues.
- Configuring a cluster IP on non-Azure appliances required a restart to be effective. This release fixes this issue.
- Nodes stuck in a dual active or flapping state could see direct memory (outside the JVM allocation) usage increase.  This release fixes this issue.
- Multiple race conditions that could lead to dual-active or no-active member clusters have been addressed. 

### Stateful Failover Beta 
Trustgrid is working to make upgrades of clustered nodes less disruptive. This release includes a beta feature that synchronizes VPN flows between nodes in the cluster. During a graceful failover, this can allow existing connections to continue to function without interruption.  This feature is disabled by default but can be enabled by contacting Trustgrid support.

## BGP Improvements
This release adds two new [event types]({{<relref "/docs/alarms/event-types">}}) for BGP connectivity under the category `BGP Peer Connectivity` that will be sent when a BGP disconnects or reconnects. It also rate limits the number of debug log entries when a BGP peer is disconnected. 

## Gateway Path Improvements
Prior to this release, the Control & Data Plane Health check did not include additional [gateway paths]({{<relref "docs/nodes/appliances/gateway/gateway-client#gateway-paths">}}) in determining the health of a cluster member. This could potentially lead to a failover when the node still had a healthy data plane path. This release fixes this issue by including all gateway paths in the health check.

## Other Improvements and Fixes
- Port Forwards were not adjusting the MSS value correctly when one side had larger than normal MTU, like AWS members.  This release fixes this issue.
- Fixes an issue where a disabled node could take up to 24 hours to re-enable if the first attempts to pull the updated configuration failed. Disabled nodes will now check every 10 minutes. 
- The node will now attempt to pull the configuration twice before sending the `Configuration Update Failure` event type.
- Fixes an issue causing Azure nodes to report startup errors. 
- Fixes an issue causing the CPU and memory usage stats to be incorrect.
- Update the APT security packages to those available on 2025-02-24.
- The Data Plane panels now accurately reports the UDP tunnel status. 
- Fixes an issue that could leave a process started in a terminal running after the terminal is closed.
