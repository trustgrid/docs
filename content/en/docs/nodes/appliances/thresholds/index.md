---
tags: ["node", "concepts"]
title: "Thresholds"
linkTitle: "Thresholds"
aliases: 
    - /docs/nodes/thresholds
description: Define threshold limits for node resource utilization and trigger events
---

{{% pageinfo %}}
Thresholds provide a way to trigger events when different measurements exceed a given value. Thresholds configured at the node level can override [Domain Thresholds]({{<ref "/docs/domain/thresholds" >}}).
{{% /pageinfo %}}

{{<tgimg src="node-thresholds.png" width="80%" caption="Example Thresholds page on a node">}}

## Node Level Threshold Status
{{<tgimg src="node-threshold-status.png" width="40%">}}
{{<fields>}}
{{<field Status>}}
- Disabled (default) - Only thresholds defined at the [domain level]({{<relref "/docs/domain/thresholds">}}) will be used for generating events.
- Enabled - Thresholds defined at the node level will be used for generating events, including overriding any domain-level values.
{{</field>}}
{{</fields>}}

## Configuring Node Level Thresholds

{{<alert>}}To view and configure thresholds, a user will need `nodes::configure::thresholds` permissions.{{</alert>}}

To configure thresholds at the node level:
1. Navigate to the node in the portal.
1. In the left-side navigation bar find Thresholds in the System section and click it. {{<tgimg src="system-threshold.png" width= "40%" alt="System > Thresholds">}}
1. Change the Status to "Enabled". 
1. Click the + button on the far right to add either a Load or Network threshold. 
1. Provide the required fields. 
1. Optionally, repeat with any other thresholds you wish to set.
1. Click save.



### Load Thresholds

{{<readfile "/docs/domain/thresholds/load-threshold.md">}}
### Network Thresholds

{{<readfile "/docs/domain/thresholds/network-threshold.md">}}

