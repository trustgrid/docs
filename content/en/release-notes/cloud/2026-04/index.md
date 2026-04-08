---
title: "April 2026 Cloud Release Notes"
linkTitle: "April 2026"
date: 2026-04-08
description: "April 2026 Cloud Release Notes"
type: docs
---

## April 8, 2026 - Major Release

### Observability HTTP Exporter Authentication
The [HTTP exporter]({{<relref "/docs/observability#http-exporter-settings">}}) now supports authentication when sending telemetry data to external endpoints. A new **Authentication** section has been added with three **Auth Type** options:
- **None** - No authentication (previous behavior)
- **Bearer Token** - Sends a bearer token in the `Authorization` header
- **Basic Auth** - Sends a username and password

### Date Range Selector for Events, Flow Logs, and Changes
The date range selector has been moved out of Advanced Search and is now always visible above the table in the [Events]({{<relref "/docs/alarms/events">}}), [Flow Logs]({{<relref "/help-center/flow-logs">}}), and [Changes]({{<relref "/docs/operations/changes">}}) sections. The active range is shown at all times, eliminating confusion about what time window is being displayed. Default ranges are **Last 1w** for Events and **Last 2h** for Flow Logs and Changes.

### Other Improvements and Fixes
- Virtual Network route destinations are now validated with case-sensitive matching against the actual node or cluster name. Previously, a destination like `node-DC1` would be accepted even if the target was named `node-dc1`, creating a route to a non-existent device. The save will now be blocked if the name does not match exactly.
- Updated a number of underlying libraries to address security vulnerabilities.
- Added foundational functionality to support future AI integrations and products.
