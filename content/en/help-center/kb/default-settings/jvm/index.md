---
title: "Java Virtual Machine (JVM)"
linkTitle: JVM 
description: >
  Default settings for the Java Virtual Machine (JVM) running on Trustgrid appliances
---

{{<alert color="info">}}
These settings only apply to appliance-based Trustgrid nodes
{{</alert>}}

## Memory Settings
The Java Virtual Machine (JVM) is responsible for managing the memory used by the Trustgrid service on appliances. Two key parameters that control this memory management are the minimum (min) and maximum (max) memory values. These settings can be adjusted under [Advanced > JVM Memory]({{<relref "/docs/nodes/appliances/advanced#memory-settings" >}}) in the Trustgrid Management Portal.

1. **Minimum Memory:** This is the initial memory allocation for the JVM. When the Trustgrid service starts, the JVM allocates this amount of memory right away. It's like setting a base size for the service's memory footprint.

2. **Maximum Memory:** This is the maximum amount of memory the JVM is allowed to use. If the Trustgrid service needs more memory than the initial amount (set by Minimum), it will continue to use more up to this maximum limit.


|Parameter| Default | Recommendations|
|---|---|---|
|Minimum | 512MB | 512MB is the recommended smallest value | 
|Maximum | 1GB | - The 1GB default runs reliably on appliances with as little as 2GB of total memory <br>- To raise it, see [Suggested Maximum Memory by Total System Memory](#suggested-maximum-memory-by-total-system-memory) below|

### Suggested Maximum Memory by Total System Memory

These are starting points assuming the appliance isn't running any containers.

| Total System Memory | Suggested Maximum Memory | % of Total |
|---|---|---|
| 2GB | 1GB (default) | 50% |
| 4GB | 2GB | 50% |
| 8GB | 5GB | 62% |
| 16GB | 11GB | 69% |
| 32GB | 22GB | 69% |

If the appliance also runs containers, subtract the sum of **all configured containers'** [Memory Max]({{<relref "docs/nodes/appliances/containers#resource-limits" >}}) limits from the suggested value above. If the result falls below 1GB, use a larger appliance instead. For sizes not listed, cap Maximum Memory at 50-70% of total system memory, trending toward the higher end on larger appliances.

### A Note on G1 and the Memory Usage Chart

{{<alert color="info">}}
Trustgrid appliances run the G1 garbage collector by default, which doesn't return memory to the operating system once it's allocated during normal operation. Because of this, the memory usage shown on the appliance [Overview]({{<relref "docs/nodes/appliances/overview#stats" >}}) can stay high even when the appliance is idle. A high, flat reading on its own isn't a sign of a problem. Size **Maximum Memory** for what you're comfortable committing long-term, not as a soft ceiling.
{{</alert>}}

### Diagnosing a Suspected Memory Issue

If the node is processing traffic normally and you have no other signs of trouble, leave the defaults alone. If you want to check further, look at the [JVM Heap metric]({{<relref "docs/nodes/appliances/metrics#jvm-heap" >}}) for one of these patterns:

- **Heap stays pinned high** (above 80% of Maximum Memory) and doesn't drop after a GC run. This can indicate the appliance genuinely needs more memory.
- **Heap rapidly fills and empties** in a repeating sawtooth pattern. This can indicate garbage collection is running very frequently, which adds overhead.

If you see either pattern:

1. [Force a GC run]({{<relref "docs/nodes/appliances/advanced#execute-garbage-collection" >}}) and watch the heap for a few minutes afterward.
2. If it stays pinned high, or keeps sawtoothing under normal traffic, increase **Maximum Memory** using the [sizing guidance above](#suggested-maximum-memory-by-total-system-memory).
3. Restart the service only as a last resort. The Java process can also request memory outside the heap for certain activities, and a restart is the only way to release that.

## Java Garbage Collection
The Java Garbage Collection system:
- Automatically frees up memory by removing unused objects.
- Essential for managing memory in the JVM.

Trustgrid appliances run **G1 GC** by default. G1 offers more predictable, shorter garbage collection pauses than the alternative Parallel GC. That matters for nodes handling live network traffic: Parallel GC's collections are stop-the-world and can introduce latency or drop packets while they run. Unless you're troubleshooting a specific issue, leave the Garbage Collector setting on G1.

{{<alert color="info">}}
The Garbage Collector dropdown in the portal labels the Parallel GC option "Default." That label is left over from an earlier release. G1 GC is what Trustgrid appliances actually run out of the box.
{{</alert>}}
