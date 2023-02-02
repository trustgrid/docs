---
tags: ["domain", "concepts"]
title: "Thresholds"
date: 2023-2-3
---

{{% pageinfo %}}
Thresholds provide a way to trigger events when different measurements exceed a given value. Thresholds configured at the domain level apply to all nodes in the domain, except when overridden. Events will be of type `Metric Threshold Violation`.
{{% /pageinfo %}}

Navigate to your domain, and click `Thresholds` on the left.

![img](list.png)

### Load Thresholds

Load thresholds measure the health of the node itself.

{{< field-def "Name" >}}
The name of the threshold. This will be available in generated events.
{{< /field-def >}}

{{< field-def "Telemetry" >}}
The metric to monitor. Options are CPU usage (%), memory usage (%), disk usage (%), and embryonic flows (absolute count).
{{< /field-def >}}

{{< field-def "Threshold" >}}
The value that must be exceeded for an event to be generated.
{{< /field-def >}}

{{< field-def "Duration" >}}
The time period to measure. If the threshold is exceeded for this duration, an event will be generated.
{{< /field-def >}}

### Network Thresholds

Network thresholds measure the health of the network from the node's perspective.

{{< field-def "Name" >}}
The name of the threshold. This will be available in generated events.
{{< /field-def >}}

{{< field-def "Telemetry" >}}
The metric to monitor. Currently only latency (measured in milliseconds) is available.
{{< /field-def >}}

{{< field-def "Threshold" >}}
The value that must be exceeded for an event to be generated.
{{< /field-def >}}

{{< field-def "Duration" >}}
The time period to measure. If the threshold is exceeded for this duration, an event will be generated.
{{< /field-def >}}

{{< field-def "Target" >}}
The target node to measure the latency to. Each node will measure the latency to the target node.
{{< /field-def >}}
