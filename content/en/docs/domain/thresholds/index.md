---
tags: ["domain", "concepts"]
title: "Thresholds"
---

{{% pageinfo %}}
Thresholds provide a way to trigger events when different measurements exceed a given value. Thresholds configured at the domain level apply to all nodes in the domain, except when overridden. Events will be of type `Metric Threshold Violation`.
{{% /pageinfo %}}

To view thresholds, a user will need `domains::read` permissions. To configure them, they will need `domains::configure:threshold` permissions.

Navigate to your domain, and click `Thresholds` on the left.

{{<tgimg src="domain-thresholds.png" width="85%">}}

### Load Thresholds

{{<readfile "./load-threshold.md">}}

### Network Thresholds

{{<readfile "./network-threshold.md">}}
