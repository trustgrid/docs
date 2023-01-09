---
Title: "Investigate Flow Logs" 
Date: 2023-1-9
---

{{% pageinfo %}}
Flow Logs can assist in troubleshooting at a basic level. Is traffic flowing? Is it in within bandwidth expectations?
{{% /pageinfo %}}

To examine flow logs for a specific node, select the node from the node list:

Navigate to the `Nodes` page, find the desired node and click on it, then scroll down to `Flow Logs` to locate the page below.

![img](flow-logs-node.png)

This view on the `Node` page shows only source and destination Nodes with bandwidth. A session will only log when it's completed so some long running sessions may not appear for some time. Traffic is flowing between the edge node and the gateway node in this example. Excess bandwidth may indicate other problems and potentially a security breach.

