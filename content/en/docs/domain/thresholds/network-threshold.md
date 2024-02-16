Network thresholds measure the health of the network from the node's perspective.

#### Network Metrics

|Metric Type| Description|
|-|-|-|
|**Latency (ms)**| Monitors the round trip tunnel latency between this node and the target |
|**Bandwidth IN Usage (Mbps)**| Monitor the amount of received bandwidth on the specified interface |
|**Bandwidth Out Usage (Mbps)**| Monitors the amount of sent bandwidth on the specified interfaces |

#### Network Fields
{{<fields>}}
{{<field "Name" >}}
The name of the threshold. This will be available in generated events.
{{</field >}}

{{<field "Telemetry" >}}
The metric to monitor. Currently only latency (measured in milliseconds) is available.
{{</field >}}

{{<field "Threshold" >}}
The value that must be exceeded for an event to be generated.
{{</field >}}

{{<field "Duration" >}}
The time period to measure. If the threshold is exceeded for this duration, an event will be generated.
{{</field >}}

{{<field "Target" >}}
- For Latency - The target node to measure the latency to. Each node will measure the latency to the target node.
- For Bandwidth In/Out - The network interface to monitor usage on. 
{{</field >}}
{{</fields>}}