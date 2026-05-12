---
title: "May 2026 Cloud Release Notes"
linkTitle: "May 2026"
date: 2026-05-05
description: "May 2026 Cloud Release Notes"
type: docs
---

## May 12, 2026 - Minor Release

## Restart Cluster Server
The [Restart Cluster Server]({{<relref "docs/clusters#restart-cluster-server">}}) tool allows you to restart the cluster server service on a cluster member without restarting the entire node. This can be useful in situations where both members are reporting as active or standby, and a restart is needed to restore normal operation. The control is now launched from the **Actions** menu in the cluster **Nodes** table and requires confirmation before the restart request is sent, which helps prevent accidental restarts during failover or failback.

## May 8, 2026 - Minor Release
### Time Range and Search Bar Preservation
The selected time range and search bar content in the [Events]({{<relref "/docs/alarms/events">}}), [Flow Logs]({{<relref "/help-center/flow-logs">}}), and [Changes]({{<relref "/docs/operations/changes">}}) pages are now preserved when changing time ranges. Previously, changing the time range would reset any search parameters to their defaults.

## May 5, 2026 - Minor Release
### Other Improvements and Fixes
- Resolves an issue that prevented saving changes to JVM settings on appliance-based nodes.
