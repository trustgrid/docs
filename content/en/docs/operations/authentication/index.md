---
Title: "Authentication Logs"
linkTitle: "Authentication"
Tags: ["audits", "auth"]
---

{{% pageinfo %}}
Trustgrid records information for authentication attempts into the control plane. This information can be used to troubleshoot authentication issues or monitor for suspicious activity.
{{% /pageinfo %}}

Authentication logs are available by navigating to **Operations** > **Authentication Logs**. {{<tgimg src="auth-audits.png" width="100%" caption="Authentication Logs">}}

{{< alert color="info" >}}Viewing authentication logs requires `audits::read:user` permissions.{{</ alert >}}

{{<alert color="warning">}}
Successful logins from external identity providers are logged, but failed attempts are not. Check your IDP's documentation for information on how to view those.
{{</alert>}}

The following information is shown for each entry:

{{<fields>}}
{{<field "Date" >}}When the authentication event occurred{{</field>}}
{{<field "IP" >}}The IP address of the authentication attempt{{</field>}}
{{<field "Message" >}}A description of the authentication event{{</field>}}
{{<field "User ID" >}}The user associated with the authentication attempt{{</field>}}
{{</fields>}}

## Searching Authentication Logs

### Search Bar

The search bar at the top of the table performs a full text search across **User ID**, **IP**, and **Message** fields. Type a search term and press Enter to filter the results. Searches match the full value of a field or substrings up to 16 characters.

### Date Range

The date range selector is always visible above the table, showing the currently active time window. The default range is **Last 1w**. Click the date range button to change the range. The selector supports both relative ranges (e.g., Last 2h, Last 1w) and absolute date/time ranges. {{<tgimg src="auth-audit-time-range.png" width="50%">}}

### Timezone

Use the **UTC/Local** toggle on the date range selector to switch between UTC and your local timezone. This affects all timestamps displayed in the table. CSV exports are always in UTC.

### Advanced Search

Click **Advanced Search** to filter authentication logs using any combination of the following fields. All multi-select fields accept multiple values. 

{{<fields>}}
{{<field "IPs" >}}Filter by the originating IP address{{</field>}}
{{<field "User IDs" >}}Filter by user{{</field>}}
{{<field "Full Text Search" >}}Search across **User ID**, **IP**, and **Message** for matching text{{</field>}}
{{</fields>}}

{{<tgimg src="auth-audit-advanced-search.png" width="50%" caption="Advanced Search">}}

### Actions Menu

The gear icon above the table provides access to:

- **Refresh** - Refresh the table results
- **Export** - Download the current filtered results as a CSV file.
- **Column Selector** - Choose which columns are visible in the table. {{<tgimg src="auth-audit-gear.png" width="25%">}}
