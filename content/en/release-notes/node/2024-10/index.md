---
title: October 2024 Major Appliance Release Notes
linkTitle: October 2024 Major
type: docs
date: 2024-10-07
description: "Release notes for the October 2024 Major Trustgrid Appliance release"
---
{{< node-release package-version="TBD" core-version="TBD" release="n-2.21.0" >}}


## Force Node Health Check Refresh
In the [July 2024 cloud release]({{<relref "/release-notes/cloud/2024-07#node-health-check-visibility">}}) Trustgrid exposed warnings in the portal if a node was unable to connect to its configured DNS server, the Trustgrid update repository (repo), or if SSH had been reconfigured to listen on an IP other than localhost.  These health checks run automatically at the following times:
- Repo Connectivity is checked every 24 hours
- DNS Resolution runs hourly 
- SSH Lockdown is checked when the node service first starts
With this release, rather than waiting for the health checks to rerun and update the state of each check you can now force the node to rerun each check.  
In the [infovisor]({{<relref "/docs/nodes/shared/infovisor">}}) refresh buttons have been added next to each health check.  Clicking this will force the node to verify if the check is healthy or nor.
{{<tgimg src="health-check-refresh.png" width="75%" caption="Infovisor with health check refresh buttons">}}

## UDP Tunnels
### Additional Path UDP Support
With this release, appliances with UDP enabled will attempt to build connections for any [additional paths]({{<relref "/docs/nodes/appliances/gateway/gateway-client#gateway-paths">}}) defined in the Gateway > Client settings. 

### UDP Tunnel Events and Logging
New `Gateway UDP Tunnel Error` events have been added to show when a UDP tunnel has disconnected and re-established.  After a future Cloud release the `re-established` event will automatically resolve the `timed out` alert. 
{{<tgimg src="udp-tunnel-events.png" width="75%" caption="Example UDP event messages">}}

## Additional New Event Types
The following events have been added or updated:
- Clustered AWS-based appliances will now send an event if they are unable to retrieve AWS credentials needed for updating AWS route table entries.
- The new `Networking Framework Memory Management` event type will alert if the appliance's java virtual machine experience out of memory errors. 
- This release adds new `SSH Lockdown` and `Repo Connectivity` events to give visibility on when a node transitioned between healthy and unhealthy states. 

## Other Improvements and Fixes
- Addresses an issue with [port forwards]({{<relref "/docs/nodes/appliances/vpn/port-forwarding">}}) that could cause failures under high load.
- Changing the listening interface or source block on a [connector]({{<relref "/docs/nodes/shared/connectors">}}) no longer requires a restart to be effective.
- Resolves an issue that caused the active cluster member to flap between devices. This issue would present in the logs with repeated messages about `Connection closed: node=<peer node>,reason=Session closed locally/remotely`
- Resolve an issue that prevented the [Traffic Capture]({{<relref "/tutorials/interface-tools/traffic-capture">}}) tool from working on appliances using the Ubuntu 22.04 operating system.