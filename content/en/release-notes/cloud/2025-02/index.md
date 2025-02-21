---
title:  February 2025 Release Notes
linkTitle:  February 2025 Release 
date: 2025-02-18
description: "February 2025 Cloud Release Notes"
type: docs
---

## Permissions and Policies
### Condition Scoped Policies
With this release we expand our permissions model to allow for using [tags to scope the policy]({{<relref "/docs/user-management/policies#condition-scoped-policies" >}}). This allows for more granular control of permissions, or as an alternative to specifying TGRNs in the policy. 

For example, if you wanted to limit a group of users to only see the nodes and cluster for a specific client, you could create a policy that looks like this:
{{<tgimg src="client-tag-scope.png" width="70%" caption="Tag Scoped Policy based on ClientID tag">}}

Any cluster or node with this tag would have the permissions granted by this policy.

## Gateway Configuration Changes
Several changes have been made to the [gateway configuration]({{<relref "/docs/nodes/appliances/gateway" >}}) panels to improve the usability.
1. The Gateway Server "Enabled" and "Enable UDP" buttons have been replaced with dropdown menus to make it more clear that saving is required.
1. Many fields can now be left unset. Previously, these fields might end up set based off of default values that could cause unwanted behavior.
1. The UDP Server port now defaults to 8443 which is the most common port used. 

## Data Plane Improvements
### Location and ISP in Data Plane Panel
This release adds the Location and ISP information as a column in the Data Plane panel.  This information is based off the public IP address the node uses to connect to the control plane, so accuracy can vary. But it can be useful for troubleshooting to determine if there is a pattern in the region or ISP of nodes experiencing data plane issues such as high latency. 

{{<tgimg src="data-plane-isp.png" width="65%" caption="Location and ISP columns in the Data Plane panel">}}


### UDP Data Plane Visibility
This release makes it easier to see if a node is enabled for [UDP data plane]({{<relref "/docs/nodes/appliances/gateway/gateway-server#udp-enabled" >}}) by adding a column to the Nodes table and a new section in the [Infovisor]({{<relref "/docs/nodes/shared/infovisor">}}) page.
{{<tgimg src="udp-column.png" alt="UDP Column option in the Nodes table" width="55%" caption="UDP Column option in the Nodes table" >}}

{{<tgimg src="udp-infovisor.png" alt="UDP mode field in Infovisor" width="85%" caption="UDP mode field in Infovisor" >}}

## Azure Cluster IP Support
This cloud release, combined with the [January 2025 Minor Appliance release]({{<relref "/release-notes/node/2025-01" >}})) adds support for using a floating, local [cluster IP address for Azure-based appliances]({{<relref "/tutorials/deployments/deploy-azure/ip-failover">}}). This provides an alternative method of providing highly available network connectivity when working with Azure-based appliances that behaves similarly to on-premises appliances.  


## ARP Ping Interface Tool
This release adds an [ARP ping tool]({{<relref "/tutorials/interface-tools/arping" >}}) to the interface tools section of the portal.  This tool allows you to send an ARP request to a specific IP address on a local interface. This tool is useful for troubleshooting connectivity issues such as confirming if the interface is in the same layer 2 broadcast domain as the target IP address and confirming their is not an IP address conflict on the network.


## Other Improvements and Fixes
- Resolve an issue that led to an incident on February 14, 2025 that caused a large number of nodes to show as disconnected from the control plane for a short period of time.
- Fixes an issue that would cause some gateways to report Data Plane status as "Degraded" when clients are not connected.  Data Plane status should be based solely on the connections where the current node is the initiating client and does not cover the connections where the current node is the gateway server.
- Changes the behavior when changing settings on the Gateway server and client panels so that each setting can be set individually. 
- Resolves an issue with how Data Plane health status was reported for private gateways.
- Routes for VLAN sub-interfaces now require a next hop address to be defined. 
- [Slack alarm channels with formatting enabled]({{<relref "/docs/alarms/channels#slack-channel">}}) now specify that the time stamp is in UTC.