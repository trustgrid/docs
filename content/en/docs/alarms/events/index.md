---
Title: "Events"
Tags: ["events", "alarms"]
---

{{% pageinfo %}}
Events are emitted from nodes and from the Trustgrid control plane when actionable things happen. Events are the basis for alarms and notifications.
{{% /pageinfo %}}

Events can be viewed for individual nodes by navigating to **Events** under the **History** section.

{{< alert color="info" >}}Viewing configuration changes requires `events::read` permissions.{{</ alert >}}

{{<tgimg src="node-events-table.png" width="100%" caption="">}}

{{< alert color="info" >}}Clicking the **Test** button will send the event through your configured alarms to help verify channels are configured as expected.{{< /alert >}}

## Event Times

{{<fields>}}
{{<field "Generated Time" >}}
The time the event was created.
{{</field >}}

{{<field "Received Time" >}}
The time the event was received by the control plane. This can be later than the generated time in the event of a network disruption, for eample.
{{</field >}}
{{</fields>}}

{{< alert color="info">}}
Events can also be viewed at the organization level:

- [Organization Events]({{<relref "docs/operations/events" >}}) — **Operations** > **Events**
{{< /alert >}}

## Searching Events

### Search Bar

The search bar at the top of the table performs a full text search across **Event Type**, **Event Level**, and **Message** fields. Type a search term and press Enter to filter the results. Searches match the full value of a field or substrings up to 16 characters.

### Date Range

{{<tgimg src="date-range-selector.png" width="60%" caption="Date range selector showing the active range">}}

The date range selector is always visible above the events table, showing the currently active time window. The default range is **Last 1w**. Click the date range button to change the range. The selector supports both relative ranges (e.g., Last 2h, Last 1w) and absolute date/time ranges. Use **Advanced Search** to filter by other criteria, and **Clear Advanced Search** to reset those filters.

### Timezone

Use the **UTC/Local** toggle on the date range selector to switch between UTC and your local timezone. This affects all timestamps displayed in the table. CSV exports are always in UTC.

### Advanced Search

Click **Advanced Search** to filter events using any combination of the following fields:

{{<fields>}}
{{<field "Event Type" >}}Filter by one or more specific event types (e.g., `Node Connect`, `Cluster Healthy`){{</field>}}
{{<field "Event Level" >}}Filter by one or more specific levels (e.g. `INFO`, `ERROR`){{</field>}}
{{<field "Full Text Search" >}}Search across **Event Type**, **Event Level**, and **Message** for matching text{{</field>}}
{{</fields>}}

{{<tgimg src="events-advanced-search.png" width="50%" caption="Advanced Search">}}

### Actions Menu

The gear icon above the table provides access to:

- **Refresh** - Refresh the table results
- **Export** - Download the current filtered results as a CSV file.
- **Column Selector** - Choose which columns are visible in the table. {{<tgimg src="events-gear.png" width="25%">}}
