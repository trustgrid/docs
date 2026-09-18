---
Title: "Configuration Changes"
linkTitle: "Changes"
Tags: ["audits", "changes"]
---

{{% pageinfo %}}
Trustgrid records information about changes users make to all configuration. This can be helpful when troubleshooting recent changes.
{{% /pageinfo %}}

Configuration changes for the entire organization are available by navigating to **Operations** > **Changes**.

{{<tgimg src="changes-list.png" width="100%" caption="Configuration changes table showing recent changes across the organization">}}

{{< alert color="info" >}}Viewing configuration changes requires `audits::read:config` permissions.{{</ alert >}}

The following information is shown for each change:

{{<fields>}}
{{<field "Date" >}}When the change was made{{</field>}}
{{<field "IP" >}}The IP address from which the change originated{{</field>}}
{{<field "Event" >}}The type of change (delete/create/change/action). Note that for some entities where the entire entity is re-sent to the API, a change may show `create` instead of `change`.{{</field>}}
{{<field "Details" >}}Information about the change. Typically includes the item's ID and a brief summary of the changes made.{{</field>}}
{{<field "User Name" >}}The user who made the change{{</field>}}
{{<field "Item Type" >}}The type of item changed. When a link is available, clicking the item type navigates to the changed item in the portal.{{</field>}}
{{</fields>}}

{{< alert color="info" >}}For the results that support linking to the item that generated the audit, there will be an external link to navigate to it.{{</ alert >}}

{{< alert color="warning" >}}Configuration changes are stored for 90 days.{{</ alert >}}

Configuration changes can also be viewed at a narrower scope:

- [Node Changes]({{<relref "docs/nodes/shared/changes" >}}) — **History** > **Changes** on a node
- [Cluster Changes]({{<relref "docs/clusters/shared/changes" >}}) — **Changes** on a cluster
- [Domain Changes]({{<relref "docs/domain/shared/changes" >}}) — **Changes** on the domain

## Searching for Changes

### Search Bar

The search bar at the top of the table performs a full text search across **Event**, **Item**, **IP**, **User Name**, **Item ID**, and **Item Type** fields. Type a search term and press Enter to filter the results. Searches match the full value of a field or substrings up to 16 characters.

### Date Range

The date range selector is always visible above the changes table, showing the currently active time window. The default range is **Last 4w**. Click the date range button to change the range. The selector supports both relative ranges (e.g., Last 2h, Last 1w) and absolute date/time ranges.

To remove the time filter entirely, click **Clear** inside the date range picker. This shows all available changes regardless of when they occurred.

{{<tgimg src="date-range-selector.png" width="50%" caption="Date range selector showing the active range">}}

### Timezone

Use the **UTC/Local** toggle on the date range selector to switch between UTC and your local timezone. This affects all timestamps displayed in the table. CSV exports are always in UTC.

### Advanced Search

Click **Advanced Search** to filter changes using any combination of the following fields. All multi-select fields accept multiple values.

{{<fields>}}
{{<field "Events" >}}Filter by change type (e.g., create, change, delete, action){{</field>}}
{{<field "Item Types" >}}Filter by the type of item that was changed (e.g., Node, Cluster, Domain){{</field>}}
{{<field "IPs" >}}Filter by the originating IP address{{</field>}}
{{<field "User Names" >}}Filter by the user who made the change{{</field>}}
{{<field "Item IDs" >}}Filter by the specific item identifier{{</field>}}
{{<field "Full Text Search" >}}Search across all fields for matching text{{</field>}}
{{</fields>}}

{{<tgimg src="advanced-search.png" width="50%" caption="Advanced Search dialog for configuration changes">}}


### Actions Menu

The gear icon above the table provides access to:

- **Refresh** - Refresh the table results
- **Export** - Download the current filtered results as a CSV file. The export includes all records matching the active filters, not just the current page.
- **Column Selector** - Choose which columns are visible in the table. {{<tgimg src="change-audit-gear.png" width="25%" caption="Advanced Search dialog for configuration changes">}}

