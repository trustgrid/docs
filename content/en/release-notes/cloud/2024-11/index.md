---
title:  November 2024 Release Notes
linkTitle:  November 2024 Release
date: 2024-11-06
description: " November 2024 cloud release"
type: docs
---

## Upgrade Manager
This release makes available the [Upgrade Manager]({{<relref "/docs/upgrade-manager">}}) to orchestrate upgrades across large numbers of devices. The Upgrade Manager can target specific nodes and clusters based on [tags]({{<relref "/docs/nodes/shared/tags">}}) and provides a single pane of glass for monitoring the upgrade process. 

Upgrade Manager also automates Cluster upgrades by starting with the standby node in the cluster, triggering a failover once it that node has completed the upgrade, and then upgrading the other member of the cluster. 

By default, all users with the [tg-builtin-admin]({{<relref "/docs/user-management/policies#builtin-tg-admin">}}) policy will have **read-only** access to the Upgrade Manager.  The new [tg-builtin]({{<relref "/docs/user-management/policies#builtin-tg-upgrade-admin">}}) policy has been created to easily grant users permissions to run jobs in the Upgrade Manager. 

## NAT Hit Counter
This release when combined nodes running the [October 2024 Trustgrid Appliance Release]({{<relref "/release-notes/node/2024-10">}}) and later, will add a **Hits** button to all NATs defined on the [VPN Address Translation]({{<relref "/docs/nodes/appliances/vpn/nats/">}}) page. When clicked the Hits button will display:
- The number of times the NAT has been hit (or used) since the last reset or reboot.
- The first (since reset) and most recent times the NAT has been hit.
- Provides a button to reset the hit counter for that specific NAT.
{{<tgimg src="nat-hit-counter-example.png" alt="Inside NAT hit counter" width="70%" caption="Inside NAT hit counter" >}}


## Google Chat Channel 
This release add the ability to create a [Google Chat (gChat) channel]({{<relref "/docs/alarms/channels#google-chat-channel">}}) for notifications.

## Cluster Overview Changes
### Update to Cluster Terminology
Over the past few years we've been moving away from the term "master" to "active" in regards to cluster status. With this release we've update the labels in the Cluster Overview table to use the new terminology.
### Make Active Button
This release adds a new button to the Cluster Overview table that allows you to make a node the configured active node in the cluster.  Previously this required selecting the node from the table, selecting the Actions menu and then selecting the "Set as Master" option.  And there was no confirmation of this action. 

Now each row in the Cluster overview table has a new button to make the node the active node.
{{<tgimg src="make-active-button.png" alt="Make Active Button" width="70%" caption="Make Active Button" >}}
And you will now be prompted before the action is taken.
{{<tgimg src="make-active-prompt.png" alt="Make Active Prompt" width="50%" caption="Make Active Prompt" >}}

## Update Node Config
This release adds a new button a node's tool bar called "Update Node Config." This button will send a message to the node to pull down it's most recent configuration.  By default if multiple changes are made to a node's configuration (or its cluster configuration) the control plane will only send the signal to pull the configuration once per minute.  This is done to reduce the number of requests made by each node.  The "Update Node Config" button will bypass this restriction and force the node to pull down the most recent configuration.

### Bulk Change Limits
Prior to this release selecting multiple nodes from the Nodes table and then performing a bulk action like Disable/Enable provided insufficient prompting that prevented the user from knowing exactly which nodes the actions would be applied.  This release changes the behavior to:
- Prevent any action from being applied to more than 10 nodes at once.
- List the nodes that are selected for the chosen action and request confirmation. 

## Other Fixes and Improvements
- `node::read` permissions now includes viewing Thresholds at the Node and Cluster level.
- Fixes an issue that prevented [resource scoped]({{<relref "/docs/user-management/policies/#resource-scoped-policies">}}) users without `user::read` permissions from creating and viewing the API keys. 
- Fixes an issue were nodes that are members of a cluster that has a Virtual Network attached to multiple interfaces could only see the VPN config of one of the interfaces. 
- Fixes an issue with Container configurations that prevented removing the Command and Stop Time values. 
- Fixes an issue preventing setting flags on newly created agent-based nodes.
- Reduces the number of domain update notifications generated when a node is enabled or disabled. Now only nodes that need to know the change should be updated. 
- Fixes an issue preventing [application]({{<relref "/docs/applications">}}) session history from loading. This release also adds the ability to export the session and acces history to CSV.
- Prior to this release if you attempted to export [flow logs]({{<relref "/docs/operations/flow-logs">}}) after changing to a different page of results the export would fail.  This release fixes this issue.
- Introduces the [tg-builtin-provisioning-admin]({{<relref "/docs/user-management/policies#builtin-tg-provisioning-admin">}})
