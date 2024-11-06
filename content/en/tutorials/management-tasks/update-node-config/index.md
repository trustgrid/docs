---
linkTitle: Update Node Config
Title: "Force Node to Pull Down Configuration"
Date: 2024-11-06
weight: 40
description: Describes how to force a Trustgrid node to pull down the most up to date configuration
---

{{<tgimg src="update-node-config.png" alt="Update Node Config button" width="60%" caption="Update Node Config button" >}}

Trustgrid nodes are notified when configuration changes are saved in the Portal. However, if multiple changes are made in a short period of time, the notification will delay for a minute to reduce configuration churn and include all changes. This reduces the number of requests made by each node and the number of configuration changes the node has to process.  

If you have made multiple changes to a node's configuration (or its cluster configuration) and want to force the node to pull down the most recent configuration you can use the "Update Node Config" button. This will cause the node to pull down the configuration immediately.
