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

### Execute Garbage Collection
This button forces the Java process to clean up memory and execute garbage collection. This is useful if you are seeing memory issues and want to force the JVM to clean up memory.

{{<tgimg src="garbage-collection.png" width="40%" caption="Garbage Collection button">}}

### Memory Settings

The JVM Memory panel allows changing the default JVM settings for the node process. See the [JVM Knowledge Base article]({{<relref "/help-center/kb/default-settings/jvm" >}}) for more information on the default settings and recommendations for changing them.
{{<fields>}}
{{<field "Minimum Memory" >}} The starting amount of memory the JVM process will consume for heap usage. {{</field>}}
{{<field "Maximum Memory" >}} The maximum amount of memory the JVM process can use for heap usage. {{</field>}}
{{<field "Garbage Collector" >}} 
<ul>
  <li>Default (Parallel GC)</li>
  <li>G1 GC</li>
</ul>
 {{</field>}}
{{</fields>}}
**The node must be restarted for any change to take effect.**
{{<tgimg src="jvm-memory.png" width="40%" caption="Java Virtual Machine (JVM) memory settings">}}

## Installed Packages

{{<alert color="info">}}Viewing installed packages requires the `nodes::service:installed-packages` permission. The built-in policies **tg-monitor**, **tg-operator**, **tg-node-admin**, and **tg-admin** grant this permission by default.{{</alert>}}

The node must run [release n-2.24.0]({{<relref "release-notes/node/2026-06" >}}) or later, which introduced the installed-packages node service. The corresponding package version for n-2.24.0 is `1.5.20260608-2475`. The page is shown only for nodes that support the installed-packages service.

The Installed Packages panel lists the software packages installed on the node. The table includes the following columns:

{{<fields>}}
{{<field "Name" >}} The package name. {{</field>}}
{{<field "Description" >}} A short summary of the package. {{</field>}}
{{<field "Version" >}} The installed version of the package. {{</field>}}
{{<field "Architecture" >}} The CPU architecture the package was built for (e.g., `amd64`). {{</field>}}
{{</fields>}}

{{<tgimg src="installed-packages.png" width="80%" caption="Installed packages list under Node Advanced." >}}
