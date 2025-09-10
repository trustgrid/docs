---
title: Upgrading Trustgrid Appliance Nodes
linkTitle: Upgrade Appliance Nodes
date: 2023-05-02
description: 
weight: 50
---

## How an Appliance Node Upgrades
The upgrade process will:
1. Upgrade the Trustgrid application package. As part of this process the node will backup the current running version. 
1. Apply security updates to install operating system packages.
1. Reboot the Trustgrid node appliance. 

For details on what this looks like from a network level see the [Node Upgrade Process KB]({{<ref "/help-center/kb/upgrade-process/">}}) article.

## Frequently Asked Questions

* **Is the upgrade disruptive?** - Yes.  While the appliance is upgrading services will remain available, though performance may be degraded.  However, **the reboot step will make all services unavailable** until the appliance fully boots and starts the Trustgrid service. Typically this takes a few minutes depending on the underlying hardware. 
* **How long will the upgrade take?** - This is highly variable but typically this takes between 10 minutes and an hour.  Below are a list of things that influence how long the upgrade takes:
   * Internet bandwidth - This determines how quickly the node can download the updates.
   * Time since previous update - The longer a node goes without updating the more security packages are likely to be required.
   * CPU and disk performance - The upgrade process involves significant CPU processing and to a lesser extent disk I/O.  The upgrade process is single-threaded, so core speed is more important than number of cores.

## Best Practices

* Upgrade nodes regularly to ensure access to the latest features and security updates
* For clustered nodes:
   * Upgrade the standby node
   * Change the standby node to be the [configured master/active]({{<ref "/docs/clusters#configured-master">}}) and wait for it to claim this role.
   * Upgrade the new standby node in the cluster
* Utilize an [Alert Suppression Window]({{<ref "/docs/alarms/alert-suppression#define-alert-suppression-window">}}) to reduce the number of alarms generated if upgrading many nodes.

## Required Permissions
In order to perform the below actions a user must have a [policy]({{<ref "/docs/user-management/policies">}}) applied with the following permissions to the nodes that will be upgraded.
``` json
  "statements": [
    {
      "effect": "allow",
      "actions": [
        "nodes::read",
        "nodes::upgrade",
        "nodes::manage",
        "nodes::service:node-upgrade",
        "portal::access"

      ]
    }
  ],
```
{{<alert color="info">}}The above is only a partial policy. Use the policy creator to define a complete policy{{</alert>}}
## Upgrading a Single Node

1. Navigate to the node that will be upgraded. 
1. From the tool bar click the Upgrade button. {{<tgimg src="node-upgrade-button.png" width="40%" caption="Upgrade button" alt="Toolbar showing an Upgrade (surrounded by a red rectangle), Refresh and Info buttons.">}}
1. When prompted enter the node's name and click confirm. {{<tgimg src="node-upgrade-confirm.png" width="40%" caption="Node upgrade confirmation" alt="Dialogue asking to enter the node name to confirm">}}
1. A dialogue will indicate the command was successfully sent. {{<tgimg src="single-node-upgrade-report.png" width="40%" caption="Node upgrade response" alt="Dialogue stating 'Upgrade in progress">}}

## Upgrading Multiple Nodes
Trustgrid offers two ways to upgrade multiple nodes: selecting multiple nodes from the Nodes table or using the [Upgrade Manager]({{<relref "/docs/upgrade-manager">}}).  Each method has its advantages and disadvantages, but typically for larger or organization-wide upgrades the Upgrade Manager is recommended.

|   | [Nodes Table](#bulk-upgrade-from-the-nodes-table) | [Upgrade Manager]({{<relref "/docs/upgrade-manager">}}) |
|---|---|---|
|Number of nodes supported | Limited initiating upgrades on 10 nodes at a time| Unlimited |
| Feedback provided | Confirms if Node Appliance upgrade command was sent successfully. | Provides detailed monitoring of upgrade progress and error information. This includes the ability to export the completion status across the upgrade cohort. |
| Cluster upgrade coordination | Requires manually failing over nodes | Automatically handles cluster upgrades by upgrading the standby node first, then promoting it to active and upgrading the remaining member. |
| Handling unhealthy clusters | Can upgrade clusters that are not in a healthy state. Note that if the only online member is upgraded and runs into issues, the entire site could end up offline. | Automatically skips clusters in an unhealthy state to avoid outages. |


### Upgrade with Upgrade Manager

The [Upgrade Manager]({{<relref "/docs/upgrade-manager">}}) page provides full details, but the high level process is:
1. [Plan the upgrade]({{<relref "/docs/upgrade-manager#planning-an-upgrade">}}) by using include and/or exclude tag filters to determine which nodes and clusters will be upgraded. 
1. (optional, but recommended) [Perform a dry run]({{<relref "/docs/upgrade-manager#dry-runs">}}) to confirm which nodes and clusters would be upgraded vs which would be skipped either due to tag filters or failing to meet health conditions. 
1. Starting the upgrade job from the Actions drop-down menu.
1. Monitoring and troubleshooting issues. This include the ability to retry a workflow for a node or cluster.
1. Completing the job. This will happen automatically if all targeted nodes are upgraded, skipped or canceled. Or can be manually cancelled from the Actions drop-down menu.

### Bulk Upgrade From the Nodes Table
1. Navigate to the [Nodes table]({{<ref "/docs/nodes#node-list-view">}}).
1. (Optional) Use the search box or [apply a tag filter]({{<ref "/docs/nodes#applying-a-tag-filter-to-the-nodes-table">}}) to filter the displayed nodes. {{<alert color="info">}}Note: if you change the search or tag filter any selected nodes will be deselected.{{</alert>}}
1. Use the check boxes to select the desired nodes. {{<tgimg src="multi-node-select.png" width="40%" caption="Nodes table" alt="Nodes table showing two nodes selected" >}}
1. From the actions dropdown select Upgrade. {{<tgimg src="nodes-upgrade-action.png" width="40%" caption="Action > Upgrade" alt="Action dropdown menu with red box around Upgrade option">}}
1. When prompted, confirm you want to proceed with the bulk upgrade.{{<tgimg src="multi-node-upgrade-confirm.png" width="50%" caption="Confirm bulk upgrade">}}
1. A table will be presented showing the status of each node upgrade request. {{<tgimg src="multi-node-upgrade-report.png" width="70%" caption="Upgrade Operation Report" alt="Upgrade Operation Report table showing two nodes were requested to upgrade and both are in progress">}}

## Monitoring Upgrade Progress

You can monitor the upgrade status from a node's Info tab. 
{{<tgimg src="upgrade-status-inprog.png" width="75%" caption="Info tab, Upgrade in Progress">}}

There are two fields
{{<fields>}}
{{<field "Upgrade Status">}} Shows the most recent status of an upgrade operation on this node. Possibly values are:
* In Progress
* Completed
* Failed
{{</field>}}
{{<field "Completion Time">}} Shows a timestamp for the last time an upgrade operation completed. 
{{</field>}}
{{</fields>}}
{{<alert color="info">}}Note: if an upgrade is currently **In Progress**, then this will be the completion time of the prior upgrade operations{{</alert>}}

{{<tgimg src="upgrade-status-completed.png" width="80%" caption="Example of Completed upgrade status">}}