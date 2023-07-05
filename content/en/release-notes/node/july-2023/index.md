---
title: "July 2023 Appliance Release Notes"
linkTitle: "July 2023"
type: docs
date: 2023-07-05
---

## New Event Types
Several new [event types]({{<ref "docs/alarms/event-types">}}) are generated with this version of the node appliance software.

### DNS Failure
If the node is unable to resolve DNS via its configured DNS servers it will generate a DNS Resolution error event. 
{{<tgimg src="dns-event.png" caption="Example DNS resolution events" width="70%">}}

### Interface Down Notification
This event is triggered when any configured interface's link state goes from UP to DOWN. 
{{<tgimg src="interface-event.png" caption="Example Interface health events" width="70%">}}

### Trustgrid Unhealthy Stop Notification
This event is triggered if the Trustgrid service fails to stop completely prior to the most recent start. This could indicate a hard power cycle event or something caused the Trustgrid service to halt unexpectedly.

{{<tgimg src="trustgrid-crash.png" caption="Example Node Stop Error" width="70%">}}

### Cluster Health Events Self-Resolve
The following cluster health related events will now resolve themselves once the triggering condition clears.
- Network Health Check
- Data and Control Plane Disruption
