---
title: June 2024 Agent Release Notes
linkTitle: June 2024
type: docs
date: 2024-06-18
description: " release of the Trustgrid Agent"
---
{{<agent-release package-version="0.2.20240618-2083" release="a-0.2.0">}}

## Display Available Resource Statistics
The agent has previously reported the percent utilization of CPU, memory, and disk space. With this release, we also report the total available resources to provide more context to those statistics. This information includes:
- The number of CPU cores
- The total amount of memory
- The total amount of disk space

This information is displayed in the Host Performance graph on the Overview page.

{{<tgimg src="resource-stats.png" width="70%" caption="Header to the Host Performance graph showing the total available resources">}}

## Other Resolved Issues:
- Layer 4 [Connector]({{<relref "/docs/nodes/shared/connectors">}}) rate limiting now works as expected
- The Data Plane status indicator now reports the correct count and state of peers
- Prevents the debug log service from being called multiple times which causes issues
- The agent will now adjust the `iptables` rules it creates if the host IP address changes