---
Title: "Alarms"
Tags: ["alarms"]
date: 2022-12-28
weight: 2
---

{{% pageinfo %}}
Trustgrid’s Alarms system notifies users of [events]({{< ref "event-types" >}}) both within our portal and via notification integrations such as email, PagerDuty, OpsGenie, Slack, and Teams.
{{% /pageinfo %}}

1. Changes in status (such as a [node]({{< ref "docs/node" >}}) disconnecting or reconnecting) create [events]({{< ref "event-types" >}}) within the Trustgrid portal.

2. These [events]({{< ref "event-types" >}}) trigger **alerts** which are the method of notification within the Trustgrid portal.

3. Trustgrid customers can define [alarm filters]({{< ref "alarm-filters" >}}) with a set a conditions that are matched against alerts. Example of the criteria include the [node]({{< ref "docs/node" >}}) name, event type, event severity, [tags]({{< ref "docs/concepts/tag" >}}), or even a basic text match.

4. If an alert matches the criteria of an [alarm filter]({{< ref "alarm-filters" >}}) then notifications will be sent to configured [channels]({{< ref "docs/alarms/channels" >}}).
