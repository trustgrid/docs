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
The Java Virtual Machine (JVM) is responsible for managing the memory used by the Trustgrid service on appliances. Two key parameters that control this memory management are the minimum (min) and maximum (max) memory values.

1. **Minimum Memory:** This is the initial memory allocation for the JVM. When your Java application starts, the JVM allocates this amount of memory right away. It's like setting a base size for the memory footprint of your application.

2. **Maximum Memory:** This is the maximum amount of memory the JVM is allowed to use. If your application needs more memory than the initial amount (set by min), it will continue to use more up to this maximum limit.


|Parameter| Default | Recommendations|
|---|---|---|
|Minimum | 512MB | 512MB is the recommended smallest value | 
|Maximum | 1GB | - Start with the total available system memory <br>- Subtract ~ 1G for OS and other processes <br> - If running containers or virtual machines on the appliance subtract their expected memory usage <br> - Set the value so that it does not exceed about 75% of remaining memory|

### Why Increase the Max Memory?

- **Handling Larger Workloads:** If your Trustgrid appliance is managing a large number of flows or processing large amounts of data, increasing the maximum memory can help accommodate spikes in memory usage without triggering garbage collection as frequently.

- **Performance Optimization:** Increasing the max memory can reduce the frequency of garbage collection (a process that frees up memory by removing unused objects), which can improve performance.

Remember, while increasing max memory can help, it's important to balance it with the available system resources to avoid starving other processes or causing system-wide issues.

## Java Garbage Collection
The Java Garbage Collection system:
- Automatically frees up memory by removing unused objects.
- Essential for managing memory in the JVM.

There are multiple different garbage collectors, each with its own strategy and performance characteristics. The default garbage collector on Trustgrid appliances is the Parallel GC.  Optionally, we also support using the G1 Garbage Collector.

### Switching to G1 Garbage Collector
- Better Performance: G1 offers more predictable garbage collection pauses, reducing latency.
- Handling Large Memory: Efficient for applications with large heaps.

Despite these advantages, Trustgrid has not yet been able to see significant performance improvements by switching to G1 over Parallel GC in most customer environments. Unless you are experiencing specific performance issues, we recommend leaving the default Parallel GC.


{{<alert color="info">}}Contact Trustgrid support if you need to adjust these settings{{</alert>}}