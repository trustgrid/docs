---
title: "September 2026 Cloud Release Notes"
linkTitle: "September 2026"
date: 2026-09-03
description: "September 2026 Cloud Release Notes"
type: docs
---

## September 21, 2026 - Minor Release

### SAML Authentication Supports Alternate Methods

SAML authentication now supports alternate authentication methods, including passkeys. Trustgrid no longer sends restrictive password-oriented authentication requirements by default, so identity providers can use the authentication method selected by the user. This applies to SAML identity providers generally, including Microsoft Entra ID.

## September 17, 2026 - Major Release

### Enhanced Search and Filtering for Operations Pages

The [Flow Logs]({{<relref "help-center/flow-logs" >}}), [Configuration Changes]({{<relref "docs/operations/changes" >}}), [Node Audits]({{<relref "docs/nodes/shared/audits" >}}), and [Authentication Logs]({{<relref "docs/operations/authentication" >}}) pages have been upgraded with improved search, filtering, and export capabilities.

#### Flow Logs
- **Source IP** and **Dest IP** filters now accept multiple values and support wildcard patterns (e.g., `192.168.*`, `10.*.1.*`).
- New **Source IP CIDR** and **Dest IP CIDR** fields allow filtering by CIDR range (e.g., `172.16.0.0/16`).

#### Configuration Changes
- The default time range is now **Last 4w** (previously 2 hours).
- The **Clear** button in the date range selector removes the time filter to show all available records.
- A new **Advanced Search** dialog provides multi-select filters for events, item types, IPs, user names, item IDs, and full text search.
- The **Item Type** column now links to the changed item in the portal when a link is available.
- Results can be exported to CSV with all matching records, not just the current page.

#### Node Audits
- A date range selector is now available (default **Last 1w**).
- A new **Advanced Search** dialog provides filters for event category, node, and full text search.
- Results can be exported to CSV.

#### Authentication Logs
- A date range selector is now available (default **Last 1w**).
- A new **Advanced Search** dialog provides multi-select filters for IPs, user IDs, and full text search.
- Results can be exported to CSV.

#### All Operations Pages
- A **UTC/Local** toggle on the date range selector lets you switch between UTC and local timezone for all displayed timestamps.
- CSV exports now use UTC timestamps for consistency.
- The date range selector now supports month-based ranges (1, 3, 6, and 12 months).

## September 3, 2026 - Minor Release

### Track and Reply to Support Requests in the Portal
[Support requests]({{<relref "docs/support/support-request" >}}) submitted from the portal now open a tracked ticket with the Trustgrid Support Team instead of sending an email. Submitting a request no longer leaves you without a receipt or a way to check on it.

The **Support Requests** section now lists requests in two columns:

- **Open Support Requests** - requests Trustgrid is still working.
- **Past Support Requests** - requests that have been closed.

Each entry shows the subject, its current status, and when it was last updated. You see the requests you opened.

Select a request to read the full conversation, including messages from Trustgrid Support and any status changes. Reply from the request itself and click **Send**, or click **Mark Ticket Resolved** once the issue is addressed. Replies are attributed to your portal user account.

The request form is unchanged. It now opens from the **Create Request** button rather than sitting on the page.
