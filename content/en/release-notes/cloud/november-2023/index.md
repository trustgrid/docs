---
title: November 2023 Release Notes
linkTitle: 'November 2023 Release'
type: docs
date: 2023-11-01
description: "November release focused on various UI improvements"
---

## Long Term Statistics
Starting with this release the Overview graphs for nodes running the [July 2023 release]({{<ref "/release-notes/node/july-2023">}}) or newer will display 1 week and 1 month statistic views. Previously only 24 hours was visible in the portal. This allows viewing historical usage trends over longer periods of time directly from the node detail page.  
- 1 week stats are aggregated into hourly data points
- 1 months stats are aggregated into 6 hour data points

{{<tgimg src="long-term-stats.png" alt="Screenshot showing 1 week and 1 month statistic views" width="80%">}}

## Remote Registration
The [October 2023 node release]({{<ref "/release-notes/node/oct-2023#preview-feature-node-registration-from-the-trustgrid-console">}}) introduced a new remote registration feature to allow nodes to be registered directly from the console. At the time this was a preview feature that required Trustgrid support to grant users the ability to complete the [portal activation process]({{<ref "/tutorials/local-console-utility/remote-registration#portal-activation-process">}}).  

With this release this permission can now be granted via [policies]({{<ref "/docs/user-management/policies">}}) by granting the `nodes::remote-activation` permission.  This permission was also added to the [tg-builtin-admin]({{<ref "/docs/user-management/policies#tg-builtin-admin">}}) policy.

## Copy NATs Button
It is common to need to share the information in the VPN > Address Translation panel with people that do not have access to the Trustgrid portal.  With this release two buttons have been added to "Copy outside NATs" and "Copy inside NATs" that will copy the respective NAT rules to the clipboard in comma seperated (CSV) format that can be easily shared directly or imported into a spreadsheet.

{{<tgimg src="copy-nats.png" alt="Screenshot showing new Copy outside NATs and Copy inside NATs buttons" width="80%">}}

Example CSV output:
```csv
Virtual CIDR,Local CIDR,Description
10.210.200.0/24,10.0.2.0/24,
```
## Gateway Panel Redesign

The [System > Gateway]({{<relref "/docs/nodes/gateway">}}) panel has been redesigned to separate server and client settings and make it more clear which settings apply to each role.

## Request Support
We have added a way to [request support directly from the portal]({{<relref "/docs/support/support-request">}}). This eliminates confusion about our support email address (it is not .com) and allows you to list the actual impacted nodes or clusters.

Navigate to the [Support]({{<relref "/docs/support">}})" and fill out the **Support Request** section with the relevant information. We will respond as quickly as possible based on the urgency selected and established SLAs for your organization. 

## Other Fixes and Improvements
- Resolved an issue where the S3 Bucket Policy on the Operations > Flow Logs table was listing the incorrect ARN
- Restored the ability to list the virtual management IP as a column on the Nodes table
- Fixed an issue that set all Host Ports to 0 (zero) when a container definition was imported
- Resolved an issue causing VPN Flow stats to sometimes not accurately display active vs new flows