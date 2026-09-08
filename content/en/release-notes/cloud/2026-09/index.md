---
title: "September 2026 Cloud Release Notes"
linkTitle: "September 2026"
date: 2026-09-03
description: "September 2026 Cloud Release Notes"
type: docs
---

## September 3, 2026 - Minor Release

### Track and Reply to Support Requests in the Portal
[Support requests]({{<relref "docs/support/support-request" >}}) submitted from the portal now open a tracked ticket with the Trustgrid Support Team instead of sending an email. Submitting a request no longer leaves you without a receipt or a way to check on it.

The **Support Requests** section now lists requests in two columns:

- **Open Support Requests** - requests Trustgrid is still working.
- **Past Support Requests** - requests that have been closed.

Each entry shows the subject, its current status, and when it was last updated. You see the requests you opened.

Select a request to read the full conversation, including messages from Trustgrid Support and any status changes. Reply from the request itself and click **Send**, or click **Mark Ticket Resolved** once the issue is addressed. Replies are attributed to your portal user account.

The request form is unchanged. It now opens from the **Create Request** button rather than sitting on the page.

## September 11, 2026 - Minor Release

### Search Flow Logs by Partial IP Address or CIDR

Flow Log search now supports partial IP address matching and CIDR searches. Search for part of an address to find matching flow records, or use a CIDR range to search an entire network.

### Improved Audit Search

The **Changes**, **Node Audits**, and **Authentication Logs** tables now provide improved search performance with expanded filters, full-text search, and CSV export. You can filter by audit type, item type, IP address, user, node, item ID, category, and time range.

### Timezone-Aware Audit and Flow Log Timestamps

Flow Log and audit table timestamps now follow the **Local/UTC** selection in the time-range picker.
