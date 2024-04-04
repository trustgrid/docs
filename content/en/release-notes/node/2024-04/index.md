---
title: April 2024 Minor Appliance Release Notes
linkTitle: April 2024 Minor
type: docs
date: 2024-04-04
description: "Release notes for the April 2024 Minor Trustgrid Appliance release"
---
{{< node-release package-version="TBD" core-version="TBD" release="n-2.19.1" >}}
## Abnormal Disconnect Messages limited to Clients
Before this release, both the appliance acting as the gateway server and the edge node client. This change limits abnormal disconnect messages to only be sent from the client to avoid excessive alert notifications on the gateway server.

## Other Fixes and Changes
- The reboot service will not respond allowing [portal notification]({{<relref "/release-notes/cloud/2024-03#portal-notification-improvements">}}) if the request fails.
- Resolves an issue that prevented the [Trustgrid SNMP OID]({{<relref "/docs/nodes/appliances/snmp/oids">}}) from providing VPN stats if attached to a VLAN interface. 
- Fixes an issue that sometimes causes client peers to be removed from a gateway's data plane panel table on disconnect and not be added back when they reconnect.
- Resolves an issue related to VPN traffic to agent-based nodes running in a container.
- Fixes a leak related to DNS requests made by containers running on appliance-based nodes.
