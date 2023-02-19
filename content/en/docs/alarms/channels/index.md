---
Title: "Channels"
Tags: ["alarms", "channels"]
Date: 2022-12-28
---

{{% pageinfo %}}
A channel defines one or more method of delivering alert notifications to external systems.  
{{% /pageinfo %}}

### Notification Delivery Channels

- Email - One or more email address (comma separated) that will receive messages from alerts@trustgrid.io

- PagerDuty - Trustgrid will generate an incident via the PagerDuty API if provided a valid API routing key.

- OpsGenie - Trustgrid will generate an incident via the OpsGenie API if provided a valid API key with read and create and update permissions.

{{<alert>}} For both PagerDuty and OpsGenie the integration will automatically resolve issues if an [event]({{<ref "docs/alarms/events" >}}) occurs that negates the initial triggering event. For example, if an [event]({{<ref "docs/alarms/events" >}}) is triggered by a Node Disconnect and the [node]({{<ref "docs/nodes" >}}) reconnects, the Node Connect [event]({{<ref "docs/alarms/events" >}}) will resolve the incident via the API. {{</alert>}}

- Slack - Trustgrid can post the [event]({{<ref "docs/alarms/events" >}}) data to a configured channel via a webhook.

- Microsoft Teams - Trustgrid can post [event]({{<ref "docs/alarms/events" >}}) data to a configure Teams channel via an incoming webhook

{{<alert>}} Only a single Slack or Teams channel can be targeted by a Trustgrid channel. However, you can create multiple Trustgrid channels if you wish to post the [event]({{<ref "docs/alarms/events" >}}) data to more than one Slack/Teams channel. {{</alert>}}

### Example Event Data

The [event]({{<ref "docs/alarms/events" >}}) data is delivered in JSON, as shown below, which depending on the integration can allow for additional parsing.

{{< highlight json >}}
{
"nodeName": "edge1",
"expires": 1604801325,
"level": "INFO",
"eventType": "Node Disconnect",
"source": "EKG",
"message": "Node disconnected",
"type": "Alert",
"orgId": "8e1c2c05-2c86-4b1b-a0cc-############",
"GS1PK": "Org#8e1c2c05-2c86-4b1b-a0cc-############",
"\_ct": {},
"uid": "1jwV1R2R6itQUjPza9yqTE8a8zu",
"GS1SK": "Alert#1jwV1R2R6itQUjPza9yqTE8a8zu",
"\_md": {},
"domain": "example.trustgrid.io",
"SK": "Alert#Node Disconnect",
"\_tp": "Alert",
"PK": "Node#0895b104-5434-447b-8577-############",
"state": "UNKNOWN",
"nodeId": "0895b104-5434-447b-8577-############",
"timestamp": 1604714923,
"channelID": "bc47ca84-1d04-454b-bedc-a55d1a917c0e",
"notes": [
"Text from Description Field"
],
"alarmIDs": [
"be324011-4bea-4392-a06a-541646decd39"
]
}
{{< /highlight >}}

### Event Data Descriptions

{{<fields>}}
{{<field "Node Name (Line 1)" >}}
This is the name of the [node]({{<ref "docs/nodes" >}}) that the [event]({{<ref "docs/alarms/events" >}}) relates to.
{{</field >}}

{{<field "Expires (Line 2)" >}}
This is the Unix epoch time when this [event]({{<ref "docs/alarms/events" >}}) will expire and automatically resolve.
{{</field >}}

{{<field "Level (Line 4)" >}}
This is the alert severity.
{{</field >}}

{{<field "Event Type (Line 5)" >}}
Matches to the [event types]({{<ref "docs/alarms/event-types" >}}).
{{</field >}}

{{<field "Timestamp (Line 21)" >}}
This is the Unix epoch time when the [event]({{<ref "docs/alarms/events" >}}) was first triggered.
{{</field >}}

{{<field "ChannelID (Line 22)" >}}
This is the unique identifier of the Trustgrid channel that was used to deliver this message. You can match this to the URL of the channel in the portal.
{{</field >}}
{{</fields>}}

![img](random-link1.png)

{{<field "AlarmID (Lines 26-28)" >}}
Since a channel can be used by more than one [alarm filter]({{<ref "alarm-filters" >}}) this will return an array of 1 or more alarm filters that matched the [event]({{<ref "events" >}}) and used this above channel. You can match this to the URL of the [alarm filters]({{<ref "alarm-filters" >}}) in the portal.
{{</field >}}

![img](random-link2.png)

This can be handy for determining which filter or filters' criteria matched the [event]({{<ref "docs/alarms/events" >}}) and sent the notification through this channel.
