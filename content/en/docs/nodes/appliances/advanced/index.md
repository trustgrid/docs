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

{{<fields>}}
{{<field "TCP TTL" >}} Idle timeout, in seconds, for established TCP flows. The connection must be idle for the whole period before the flow is removed. Default `900`. {{</field>}}
{{<field "TCP Handshake Timeout" >}} Timeout, in seconds, for TCP flows that have not completed the 3-way handshake (SYN, SYN/ACK, ACK). Default `15`. {{</field>}}
{{<field "UDP TTL" >}} Idle timeout, in seconds, for UDP flows. Default `60`. {{</field>}}
{{<field "ICMP TTL" >}} Idle timeout, in seconds, for ICMP flows. Default `5`. {{</field>}}
{{<field "ARP Interval" >}} How often, in seconds, the node sends a gratuitous ARP request for each interface that has a cluster IP set. Default `60`. {{</field>}}
{{<field "Drop IP-option SYNs" >}} Drops incoming TCP SYN packets that carry IP options and are addressed to a service on the node itself. When **Disabled**, the node counts and logs these packets and lets them through. Default **Disabled**. See [Drop IP-option SYNs](#drop-ip-option-syns). {{</field>}}
{{</fields>}}

After changing a setting click the Save button. **Some changes may require a node restart to take effect.**

### Drop IP-option SYNs
Set **Drop IP-option SYNs** to **Enabled** to drop incoming TCP SYN packets that carry IP options and are addressed to a service on the node itself. A change takes effect within a few minutes and does not need a restart. Only packets addressed to the node itself are checked. Traffic through outside NATs and traffic that passes through the node are not affected.

Trustgrid has tested known Trustgrid control and data plane traffic and example application traffic across the Trustgrid connection with this option enabled. We cannot guarantee how other application traffic might be affected. Monitor for issues after enabling it.

{{<alert color="info">}}The node must run the [November 2026 Minor Appliance release]({{<relref "release-notes/node/2026-11" >}}) (n-2.26.1) or later. For older nodes the option cannot be enabled and the panel shows an upgrade notice.{{</alert>}}

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

{{<alert color="info">}}The node must run [release n-2.24.0]({{<relref "release-notes/node/2026-06" >}}) or later, which introduced the installed-packages node service. The corresponding package version for n-2.24.0 is `1.5.20260608-2475`. The page is shown only for nodes that support the installed-packages service.{{</alert>}}

The Installed Packages panel lists the software packages installed on the node. The table includes the following columns:

{{<fields>}}
{{<field "Name" >}} The package name. {{</field>}}
{{<field "Description" >}} A short summary of the package. {{</field>}}
{{<field "Version" >}} The installed version of the package. {{</field>}}
{{<field "Architecture" >}} The CPU architecture the package was built for (e.g., `amd64`). {{</field>}}
{{</fields>}}
