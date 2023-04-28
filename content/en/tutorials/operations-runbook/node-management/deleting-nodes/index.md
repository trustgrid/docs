---
title: Deleting Trustgrid Nodes
linkTitle: Deleting Nodes
description: Describes the process of deleting Trustgrid nodes along with best practices
date: 2023-04-28
--- 
{{<alert color="danger">}}**Be advised:** Deleting a node is an irreversible action. Delete nodes only after you are certain it is no longer needed.{{</alert>}}

## Best Practices
* [Disable the node(s)]({{<ref "/tutorials/operations-runbook/node-management/changing-node-status">}}) for several days prior to deleting. This way if it is discovered that it is still needed it is simple to re-enable.
* Before deleting, [remove any virtual network routes]({{<ref "/docs/domain/virtual-networks/routes#deleting-virtual-network-routes">}}) with the node as a destination. While these can be removed later, they will be easier to find while the node is still listed as the destination.

## Deleting Nodes

1. If the node(s) is not already disabled, [disable the node(s)]({{<ref "/tutorials/operations-runbook/node-management/changing-node-status">}}) **See best practice above before proceeding**
1. If the node(s) is a member of a cluster, [remove the node(s) from the cluster]({{<ref "/docs/clusters/manage-members#remove-member">}}). 
1. Navigate to the Nodes table and select the target node(s). {{<alert color="info">}}Optionally, Use the search to filter the displayed nodes. {{</alert>}}
1. From the Actions menu select the Delete option. {{<tgimg src="delete-node-action.png" width="50%" caption="Delete node action">}}
   1. If deleting a single node, you will be prompted to enter the name of the node about to be deleted to confirm.
   1. If deleting multiple nodes, you will be prompted to enter `DELETE` to confirm.