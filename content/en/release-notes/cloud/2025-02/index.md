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


## Data Plane Improvements
### Location and ISP in Data Plane Panel
This release adds the ability to add the Location and ISP information to the Data Plane panel.  This information is based off the public IP address the node uses to connect to the control plane, so accuracy can vary. But it can be useful for troubleshooting to determine if there is a pattern in the region or ISP of nodes experiencing data plane issues such as high latency. 

### UDP Data Plane Visibility
This release makes it easier to see if a node is enabled for [UDP data plane]({{<relref "/docs/nodes/appliances/gateway/gateway-server#udp-enabled" >}}) by adding a column to the Nodes table and a new section in the [Infovisor]({{<relref "/docs/nodes/shared/infovisor">}}) page.
{{<tgimg src="udp-column.png" alt="UDP Column option in the Nodes table" width="55%" caption="UDP Column option in the Nodes table" >}}

{{<tgimg src="udp-infovisor.png" alt="UDP mode field in Infovisor" width="85%" caption="UDP mode field in Infovisor" >}}




## Other Improvements and Fixes
- Resolve an issue that led to an incident on February 14, 2025 that caused a large number of nodes to show as disconnected from the control plane for a short period of time.
- Fixes an issue that would cause some gateways to report Data Plane status as "Degraded" when clients are not connected.  Data Plane status should be based solely on the connections the current node is the client for. 
- Changes the behavior when changing settings on the Gateway server and client panels so that each setting can be set individually. 
- Resolves an issue with how Data Plane health status was reported for private gateways.
- Routes for VLAN sub-interfaces now require a next hop address to be defined.