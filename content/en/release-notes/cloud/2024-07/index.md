---
title: July 2024 Release Notes
linkTitle: July 2024 Release
date: 2024-07-11
description: "July 2024 cloud release introducing Alert Center"
type: docs
---

## Portal Backend Library Update
With this release, we are updating the React library used in the Portal to a more recent version. This change should be mostly invisible to users there are a few notable improvements including:
- Resizable columns: In most tables including the Nodes and Flow Logs tables. This will make it easier to read the data in the tables.
- Notifications are improved: Node actions show in-progress status and node alerts appear as stackable bottom-right alerts expanding on hover.
- Faster page rendering.

## Bulk Upgrade Improvement
Previously if you selected multiple nodes from the Nodes table and used the Actions>Upgrade option the confirmation dialog didn't include the actual names of the selected nodes. Now each selected node is listed prior to confirmation. 
{{<tgimg src="bulk-upgrade-confirm.png" width="50%" caption="List of nodes to be upgraded">}}

## Threshold Alert Improvements
This release will change alert resolution messages to include the **new** value of the threshold. Previously the message would only include the **old** value. E.g. if CPU utilization was over 90% but dropped back to 55% that new value will be included in the event message.

## JVM Garbage Collection Action
There is a new button in the [Advanced > JVM Memory]({{<relref "/docs/nodes/appliances/advanced#execute-garbage-collection">}}) panel that will force the JVM to clean up memory and execute garbage collection. This is useful if you are seeing memory issues and want to force the JVM to clean up memory.

## Node Deletion Event
This release introduces a [new event type]({{<relref "/docs/alarms/event-types">}}) generated whenever a node is deleted. Deletions were always tracked as part of [Changes]({{<relref "/docs/operations/changes">}}) but this new event type allows for an [alarm filter]({{<relref "docs/alarms/alarm-filters">}}) to be configured to alert a channel when a deletion occurs.

## Improved NAT Validation
NAT validation previously allowed invalid CIDR such as 10.0.0.8/8. This release will now prevent invalid CIDRs from being saved.

## Status Page Notification Banner
The status notification banner for maintenance and incidents has been moved from a pop-up in the bottom right to a banner at the bottom of the screen.  This will make it a bit less intrusive when working within the portal.
{{<tgimg src="status-notification.png" width="80%">}}

## Node Health Check Visibility
Node information about SSH connectivity, DNS access, and repo connectivity are now surfaced in the Portal, both in the nodes table and on each node's overview page.

* SSH Lockdown - When `disabled`, the node is listening for SSH on an external interface.
* DNS Resolution - When `unhealthy`, the node cannot resolve DNS names. This can prevent things like control plane failovers from working correctly.
* Repo Connectivity - When `disconnected`, the node is unable to fetch software updates from Trustgrid.

## Other Fixes and Improvements
- Flow log exports include TCP Flags
- Removing a VLAN sub-interface no longer requires saving after confirmation
- VRF Traffic Rules now allow for selecting Public or Private covering the respective address spaces
