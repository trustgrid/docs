Load thresholds measure the health of the node itself.

#### Load Metrics

|Metric Type| Description|Default Value|
|-|-|-|
|**CPU Usage**| Monitors percent of CPU usage across all cores | 95% for 10 minutes|
|**Memory Usage**| Monitors percent of total memory (RAM) used | 90% for 30 minutes|
|**Disk Usage**| Monitors percent of total disk usage for the root partition | 80% for 1 minute|
|**Embrionic Flows**| Monitors the number of TCP flows (connections) that are in the embryonic state (waiting for ACK) | none |
|**JVM Heap**|Monitors the percent of allocated [JVM memory]({{<relref "/help-center/kb/default-settings/jvm">}}) used | none |


#### Load Fields
{{<fields>}}
{{<field "Name" >}}
The name of the threshold. This will be available in generated events.
{{</field >}}

{{<field "Telemetry" >}}
The metric to monitor. Options are CPU usage (%), memory usage (%), disk usage (%), and embryonic flows (absolute count).
{{</field >}}

{{<field "Threshold" >}}
The value that must be exceeded for an event to be generated.
{{</field >}}

{{<field "Duration" >}}
The time period to measure. If the threshold is exceeded for this duration, an event will be generated.
{{</field >}}
{{</fields>}}
