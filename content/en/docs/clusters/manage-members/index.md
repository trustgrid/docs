---
title: "Manage Cluster Members"
linkTitle: "Manage Members"
---

## Add Member
This will add the selected node to the cluster and gain any configuration defined at the cluster level.

1. Click Actions > Add Node.
{{<tgimg src="add-node.png" width="40%" caption="Action > Add Node" alt="Action > Add Node" >}}
2. Select the desired node from the list. This is a type-ahead field that will filter the list as you type. Click Add Node.
{{<tgimg src="add-node-list.png" width="50%" caption="Member node selection list" alt="Add node dialog with list of nodes filtered to show those matching 'edge'">}}
3. You are prompted if you want to configure cluster heartbeat settings.
{{<tgimg src="config-heart1.png" width="60%" caption="Prompt to configure heartbeat settings" alt="Prompt to configure heartbeat settings">}}
   1. No - The node is added to the cluster and you will manually need to navigate to the [node's cluster settings]({{<ref "/docs/nodes/cluster">}}) and configure the heartbeat before it will be healthy.
   2. Yes - The configuration prompt allows you to set the [Heartbeat]({{<ref "/docs/nodes/cluster#heartbeat">}}) and, optionally, the [Status Endpoint]({{<ref "/docs/nodes/cluster#status-endpoint">}}) settings.
   {{<tgimg src="config-heart2.png" width="60%" caption="Heartbeat and Status Endpoint settings" alt="Dialog to enter heartbeat IP and port, and status endpoint port">}}



## Remove Member
Removing a node from a cluster will remove the cluster heartbeat and endpoint settings defined on the node and any cluster-level configuration.

1. Select the desired node from the member table.
2. Click Action > Remove Node.
{{<tgimg src="remove-node.png" width="30%" caption="Remove Node option from Actions">}}
3. Confirm you want to remove the node by clicking Yes.