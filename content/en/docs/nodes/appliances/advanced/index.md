---
title: Advanced Options
linkTitle: Advanced
aliases: 
  - /docs/nodes/advanced
description: Configure advanced settings on appliance-based nodes
---
The Advance section provides additional configuration options for customizing the node behavior from the [default settings]({{<relref "/help-center/kb/default-settings">}}). 
{{<alert color="warn">}} Default settings work in most environments and should only be changed if needed based on your specific requirements. {{</alert>}}

## Config Options
{{<tgimg src="config-options.png" width="80%" caption="Advanced Config Options">}}

The Config Options panel allows customizing various advanced settings for the node such as [network flow defaults]({{<relref "/help-center/kb/default-settings/network-flows">}}). Each setting provides a description of what it controls and the value field.  Additionally, the button on the far right allows resetting a setting back to the default value. 

After changing a setting click the Save button. **Some changes may require a node restart to take effect.**

## JVM Memory
{{<tgimg src="jvm-memory.png" width="40%" caption="Java Virtual Machine (JVM) memory settings">}}

### Execute Garbage Collection
This button forces the Java process to clean up memory and execute garbage collection. This is useful if you are seeing memory issues and want to force the JVM to clean up memory.

{{<tgimg src="garbage-collection.png" width="40%" caption="Garbage Collection button">}}

### JVM Memory Settings

The JVM Memory panel allows changing the [default JVM settings]({{<relref "/help-center/kb/default-settings/jvm">}}) for the node process. 

**The node must be restarted for any change to be effective.**