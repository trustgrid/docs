---
Title: "Alarms"
Tags: ["alarms"]
date: 2022-12-22
---

{{% pageinfo %}}
Trustgrid’s Alarms system notifies users of [events]({{< ref "docs/alarms/events" >}}) both within our portal and via notification integrations such as email, PagerDuty, OpsGenie, Slack, and Teams.
{{% /pageinfo %}}

1. Changes in status (such as a [node]({{< ref "docs/concepts/node" >}}) disconnecting or reconnecting) create [events]({{< ref "docs/alarms/events" >}}) within the Trustgrid portal. 

2. These [events]({{< ref "docs/alarms/events" >}}) trigger **Alerts** which are the method of notification within the Trustgrid portal.  

3. Trustgrid customers can define [alarm filters]({{< ref "docs/alarms/alarm-filters" >}}) with a set a conditions that are matched against alerts.  Example of the criteria include the [node]({{< ref "docs/concepts/node" >}}) name, event type, event severity, [tags]({{< ref "docs/concepts/tag" >}}), or even a basic text match.  

4. If an Alert matches the criteria of an Alarm Filter then notifications will be sent to configured **Channels**.