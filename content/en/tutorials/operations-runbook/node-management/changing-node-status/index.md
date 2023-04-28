---
title: "Changing Node Status (Enabling / Disabling)"
linkTitle: "Changing Node Status"
date: 2023-04-28
description: "Describes the process of enabling or disabling a Trustgrid node and the impacts of such a change"
---

Trustgrid nodes can be in two states:
* Enabled - In this state the node will connected to the Trustgrid Control Plane and any Data Plane gateways the node is configured to use.
* Disabled - This is a reduced functionality state that cause the node to only attempt to connect to the Trustgrid control plane. It can be managed to a limited extent but many other Trustgrid features will not function in this state

Disabling a node allows for additional security if a device is in transport and can also be used as an easily reversible way of verifying a node is no longer in use before fully deleting the node. 

{{<alert color="warning">}}Changing the node status will trigger a restart of the Trustgrid node service.{{</alert>}}

## Disabling/Enabling from Info Tab
If disabling or enabling a single node you can:
1. Navigate to that node in the Trustgrid portal.
1. Open the **Info** tab by clicking the dropdown in the top-right, or hitting the `` ` `` key.
1. From the Status dropdown change to Disable or Enable {{<tgimg src="info-tab-status.png" width="70%" caption="Info tab with Status dropdown">}}

## Disabling/Enabling from Nodes Table
If disabling or enabling multiple nodes it is easier to do so from the Nodes table

1. Navigate to the nodes table.
1. (Optional) Use the search box to filter the displayed nodes. {{<alert color="warning">}} Note that if you change this any selected nodes will be de-selected. {{</alert>}}
1. Use the checkbox on the left to select the nodes you wish to disable or enable.
1. From the Actions dropdown select either **Enable** or **Disable** {{<tgimg src="nodes-table-status.png" width="60%" caption="Status actions on Nodes table">}}

## Disabling/Enabling via Terraform
Nodes can be enabled or disabled via the Trustgrid published Terraform provider using the <a href="https://registry.terraform.io/providers/trustgrid/tg/latest/docs/resources/node_state/" target="_blank">tg_node_state</a> resource.

{{<alert color="info">}}The Trustgrid API refers to this state via the `status` attribute.  The value `ACTIVE` means the node is **enabled** and the value `INACTIVE` indicates the node is currently **disabled** {{</alert>}}
