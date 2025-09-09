---
title: Additional Metrics
linkTitle: Metrics
description: The Metrics section under History shows stats in addition to those displayed on the overview page
aliases: 
    - /docs/nodes/metrics
description: View additional metrics and statistics for the node
---

The Metrics page will be used to display additional available statistics that are not displayed on the [overview]({{<relref "/docs/nodes/appliances/overview">}}) page. Similar to the overview page these stats are viewable in 1, 2, 6, 12, and 24-hour views, and 1 week and 1 month views. 

## JVM Heap 
The Trustgrid service on appliance-based nodes runs via the Java Virtual Machine (JVM). When launched the java process claims a minimum amount of memory from the operating system. As the process runs additional memory can be claimed up to a configured maximum. Not all of this memory is actively in use by Java. Once the maximum amount of memory has been claimed or a period of time has elapsed Java will run a garbage collection process to free up pages of memory that are no longer being actively used.  

The JVM Heap statistic shows how much of the claimed memory is actively being used. The sharp reductions in the graph below indicate times where the garbage collection process has run freeing up available memory by reducing the current heap size.
{{<tgimg src="metrics-jvm-heap.png" caption="Example JVM Heap chart" width="75%">}}