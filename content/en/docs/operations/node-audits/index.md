---
Title: "Organization Node Audits"
linkTitle: "Node Audits"
Tags: ["audits", "node"]
---

{{% pageinfo %}}
Node audits provide records when services on a node are invoked, or when upgrades are applied. This page shows node audits from all nodes in an organization.
{{% /pageinfo %}}

Node audits for the entire organization are available by navigating to **Operations** > **Node Audits**. {{<tgimg src="node-audits.png" width="100%">}}

To view audits for a specific node, navigate to that node and select **Audits** under the **History** section. See [Node Audits]({{<relref "docs/nodes/shared/audits" >}}) for more information.

{{< alert color="info" >}}Viewing node audits requires `audits::read:node` permissions.{{</ alert >}}

## Searching Node Audits

### Search Bar

The search bar at the top of the table performs a full text search across **Event**, **Node**, and **Value** fields. Type a search term and press Enter to filter the results. Searches match the full value of a field or substrings up to 16 characters.

### Date Range

The date range selector is always visible above the node audits table, showing the currently active time window. The default range is **Last 1w**. Click the date range button to change the range. The selector supports both relative ranges (e.g., Last 2h, Last 1w) and absolute date/time ranges. {{<tgimg src="node-audit-time-range.png" width="50%">}}

### Timezone

Use the **UTC/Local** toggle on the date range selector to switch between UTC and your local timezone. This affects all timestamps displayed in the table. CSV exports are always in UTC.

### Advanced Search

Click **Advanced Search** to filter node audits using any combination of the following fields:

{{<fields>}}
{{<field "Event" >}}Filter by the audit category (e.g., upgrade, service invocation){{</field>}}
{{<field "Node" >}}Filter by one or more specific nodes{{</field>}}
{{<field "Full Text Search" >}}Search across **Node**, **Event**, and **Value** for matching text{{</field>}}
{{</fields>}}

{{<tgimg src="node-audit-advanced-search.png" width="50%" caption="Advanced Search">}}


### Actions Menu

The gear icon above the table provides access to:

- **Refresh** - Refresh the table results
- **Export** - Download the current filtered results as a CSV file.
- **Column Selector** - Choose which columns are visible in the table. {{<tgimg src="node-audit-gear.png" width="25%">}}
