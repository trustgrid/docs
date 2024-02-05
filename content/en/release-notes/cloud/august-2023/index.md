---
title: August 2023 Release Notes
linkTitle: 'August 2023 Release'
type: docs
date: 2023-08-18
description: August Release focusing on general improvements and fixes
---

## Cluster Site Health Improvements
Since introducing our [cluster site health state]({{<ref "/docs/clusters/cluster-health">}}) feature there have been several instances where the state was being reported incorrectly.  This release has taken steps to prevent that and has also introduced a new tool to [force validating the health state]({{<ref "/docs/clusters/cluster-health#force-health-validation">}}) in the event the state does appear incorrect.
{{<tgimg src="/docs/clusters/cluster-health/cluster-validate-health.png" caption="Validate Health action" width="80%">}}


## Data Plane Panel Improvements
This release includes two improvements to the [Data Plane panel]({{<ref "/docs/nodes/appliances/data-plane">}}):
1. The protocol field which used to only display either TCP or TCP/UDP is now called Ports and it lists the port in use for each protocol.
1. Links have been added to the far right column for MTR and Trace Route. When used these will populate the appropriate fields in each tool to do a Trace Route or MTR to a gateway peer on the TCP port in use. This makes it easier to troubleshoot connectivity issues. 
{{<tgimg src="data-plane-new.png" caption="Data Plane panel with ports and Trace Route/MTR links" width="40%">}}


## Other Improvements
- Nodes now display their operating system (OS) version information in the [infovisor]({{<ref "/docs/nodes/shared/infovisor">}}) and it can also be added as a column to the Nodes table
- We now display the physical interface names assigned by the OS on the [interfaces panel]({{<ref "/docs/nodes/appliances/interfaces">}}). This name is sometimes needed when running tools such as the Sniff Traffic.
- It is now possibly for customers to generate and download [Debug Logs]({{<ref "/help-center/ops-logs/debug-logs">}}) from a Trustgrid node. This can be handy when working with Trustgrid support to troubleshoot issues.
- [Users can be added to one or more groups when viewing the user]({{<ref "/docs/user-management/users#adding-a-user-to-groups">}}). Previously you had to navigate to each group and add the user there.
- You can now select multiple protocols, IPs or ports when using the [Sniff Virtual Traffic]({{<ref "/tutorials/remote-tools/sniff-virtual-traffic">}}) VPN tool.
- [Alarm Filter Tag Matches]({{<ref "/docs/alarms/alarm-filters#tag-matches">}}) can now be set to Any or All.  `Any` would cover a scenario where you want to match say `Environment=Prod` **OR** `Environment=Test`.  `All` would cover if you wanted a filter to match something like `Environment=Prod` **AND** `Region=EAST`