---
title: Deleting Trustgrid Clusters
linkTitle: Deleting Clusters
description: Describes the process of deleting Trustgrid Cluster along with best practices
date: 2023-05-02
weight: 90
--- 
{{<alert color="danger">}}**Be advised:** Deleting a cluster is an irreversible action. Delete clusters only after you are certain they are no longer needed.{{</alert>}}

## Best Practices
* [Remove all members]({{<ref "/docs/clusters/manage-members#remove-member">}}) for a period of time to ensure the configuration is no longer needed.
* [Update any status tag]({{<ref "/docs/nodes/tags/prod-status-tag">}}) to indicate the cluster(s) is being decommissioned.
* Before deleting, [remove any virtual network routes]({{<ref "/docs/domain/virtual-networks/routes#deleting-virtual-network-routes">}}) with the cluster as a destination. While these can be removed later, they will be easier to find while the cluster is still listed as the destination.

## Deleting Clusters
1. [Remove the node(s) from the cluster]({{<ref "/docs/clusters/manage-members#remove-member">}}) and [remove any virtual network routes]({{<ref "/docs/domain/virtual-networks/routes#deleting-virtual-network-routes">}}) with the cluster as a destination. **See best practices above**
1. Navigate to the Cluster table and select the target cluster(s). {{<alert color="info">}}Optionally, Use the search to filter the displayed clusters. {{</alert>}}
1. From the Actions menu select the Delete option. {{<tgimg src="delete-cluster-action.png" width="30%" caption="Delete cluster action">}}
   1. If deleting a single cluster, you will be prompted to enter the fully-qualified domain name (fqdn) of the cluster about to be deleted to confirm.
   1. If deleting multiple clusters, you will be prompted to enter `DELETE` to confirm.