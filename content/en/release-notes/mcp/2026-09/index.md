---
title: September 2026 MCP Server Release Notes
linkTitle: September 2026
type: docs
date: 2026-09-14
description: "September 2026 MCP Server Release Notes"
---

## September 24, 2026 - Minor Release

### Access Node Debug Logs

The following tools have been added to allow agents to request and retrieve [debug logs]({{<ref "/help-center/ops-logs/debug-logs">}}):

- `publish_debug_logs` asks a node to publish a debug-log archive.
- `list_debug_logs` lists the available archives for a node.
- `get_debug_log_url` returns a signed download URL for an archive.

## September 14, 2026 - Minor Release

- Fixed an issue where long list results would truncate the results and confuse agents.
