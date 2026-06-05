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
The Java Virtual Machine (JVM) manages the memory used by the Trustgrid service on appliances. The two customer-tunable settings are the minimum heap size (Xms) and maximum heap size (Xmx).

|Parameter|Default|Recommendations|
|---|---|---|
|Minimum heap (Xms)|512 MiB|512 MiB is the recommended smallest value.|
|Maximum heap (Xmx)|1 GiB|Use the formula and sizing table below. Do not set Xmx close to total RAM.|

### Maximum heap (Xmx)

To raise Xmx, first reserve memory for the operating system, JVM off-heap usage, and any container workload, then cap by appliance size:

```text
Xmx ≤ Total RAM − 1 GiB (OS reserve) − 512 MiB (JVM off-heap) − Container memory budget
```

|Total RAM|Suggested Xmx ceiling|% of RAM|
|---|---|---|
|2 GiB|0.7 GiB|35%|
|4 GiB|2.0 GiB|50%|
|8 GiB|5.0 GiB|62%|
|16 GiB|11 GiB|69%|
|32 GiB|22 GiB|69%|

The percentage rises with appliance size because the OS and off-heap reserves (~1.5 GiB) are roughly fixed in absolute terms. On a 2 GiB appliance this leaves little room for the Java heap, so use a larger appliance for sustained throughput workloads.

**Do not set Xmx close to total RAM.** Xmx at 75% of a 4 GiB appliance has caused production memory pressure incidents.

### Container memory budgeting

Container workloads vary widely in memory use. Subtract the container's expected memory budget (from its developer's documentation) before computing Xmx. If the resulting Xmx falls below 1 GiB, use a larger appliance.

{{<alert color="info" title="Note:">}}
On OpenJDK 8 (Trustgrid's current JDK), the JVM does not return committed heap memory to the operating system during normal operation. The portal's **JVM Heap Usage** chart shows the application's view of the heap and oscillates with GC activity, but the operating system continues to see the JVM holding the expanded heap until the service restarts. When sizing Xmx, treat it as the memory the appliance will eventually carry, not a soft target.
{{</alert>}}

### Why this matters / how to recognize misconfiguration

Oversized Xmx is the most common Trustgrid memory misconfiguration. Symptoms:

- Process RSS climbs to near Xmx and stays there even when traffic drops.
- The portal's heap usage chart shows usage well below Xmx.
- Other processes get squeezed and system swap usage rises.
- Forcing a full GC briefly drops RSS, which then re-climbs under load.

If you see this pattern, lower Xmx using the formula above and restart the node.

## Java Garbage Collection

Trustgrid nodes use **G1 GC** by default and it should remain the default for any node processing live network traffic. Do not switch to ParallelGC. ParallelGC keeps memory bounded via frequent stop-the-world full collections, which add latency and drop packets on a node processing live traffic.


{{<alert color="info">}}Contact Trustgrid support if you need to adjust these settings{{</alert>}}
