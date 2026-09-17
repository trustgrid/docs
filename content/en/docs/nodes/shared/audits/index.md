---
tags: ["node"]
title: "Audits"
linkTitle: "Audits"
aliases: 
  - /docs/nodes/audits
description: View audit logs of actions taken on a node
---

{{% pageinfo %}}
Node audits provide records when services on a node are invoked, or when upgrades are applied. This can be helpful to see the historical maintenance of a node.
{{% /pageinfo %}}

Node audits can be viewed for a specific node by navigating to **Audits** under the **History** section on the node view.

Users will need `audits::read:node` permissions to view node audits.

{{<tgimg src="node-list.png" width="100%">}}

Node audits can also be found for the entire organization by navigating to **Operations** > **Node Audits**.

## Searching Node Audits

### Date Range

The date range selector is always visible above the node audits table, showing the currently active time window. The default range is **Last 1w**. Click the date range button to change the range. The selector supports both relative ranges (e.g., Last 2h, Last 1w) and absolute date/time ranges. {{<tgimg src="node-audit-time-range.png" width="50%">}}

### Advanced Search

Click **Advanced Search** to filter node audits using any combination of the following fields:

{{<fields>}}
{{<field "Event" >}}Filter by the audit category (e.g., upgrade, service invocation){{</field>}}
{{<field "Full Text Search" >}}Search across **Event**, and **Value** for matching text{{</field>}}
{{</fields>}}

{{<tgimg src="node-audit-advanced-search.png" width="50%" caption="Advanced Search">}}


### Actions Menu

The gear icon above the table provides access to:

- **Refresh** - Refresh the table results
- **Export** - Download the current filtered results as a CSV file.
- **Column Selector** - Choose which columns are visible in the table. {{<tgimg src="node-audit-gear.png" width="25%">}}
