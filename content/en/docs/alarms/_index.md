---
Title: "Alarms"
Tags: ["alarms"]
---

{{% pageinfo %}}
Trustgrid’s Alarms system notifies users of [events]({{<ref "event-types" >}}) both within our portal and via notification integrations such as email, PagerDuty, OpsGenie, Slack, and Teams.
{{% /pageinfo %}}

1. Changes in status (such as a [node]({{<ref "docs/nodes" >}}) disconnecting or reconnecting) create [events]({{<ref "event-types" >}}) within the Trustgrid portal.

1. These [events]({{<ref "event-types" >}}) are evaluated against [alarm filters]({{<ref "alarm-filters">}}), which define criteria to match events, like the node name or event severity.

1. If an event matches the criteria of an alarm filter, then an **alert** is generated and notifications are broadcast to the configured [channels]({{<ref "docs/alarms/channels" >}}).
