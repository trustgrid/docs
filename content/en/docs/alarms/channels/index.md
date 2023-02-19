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

- Microsoft Teams - Trustgrid can post [event]({{<ref "docs/alarms/events" >}}) data to a configured Teams channel via an incoming webhook

{{<alert>}} Only a single Slack or Teams channel can be targeted by a Trustgrid channel. However, you can create multiple Trustgrid channels if you wish to post the [event]({{<ref "docs/alarms/events" >}}) data to more than one Slack/Teams channel. {{</alert>}}

### Example Event Data

The [event]({{<ref "docs/alarms/events" >}}) data is delivered in JSON, as shown below, which depending on the integration can allow for additional parsing.

{{<highlight json>}}
{
	"nodeName": "edge1", /* Name of the node that the event relates to */
	"expires": 1604801325, /* Unix epoch timestamp when this event will expire and automatically resolve */
	"level": "INFO", /* Alert severity */
	"eventType": "Node Disconnect", /* Matches to the event types */
	"source": "EKG",
	"message": "Node disconnected",
	"type": "Alert",
	"orgId": "8e1c2c05-2c86-4b1b-a0cc-############",
	"GS1PK": "Org#8e1c2c05-2c86-4b1b-a0cc-############",
	"_ct": {},
	"uid": "1jwV1R2R6itQUjPza9yqTE8a8zu",
	"GS1SK": "Alert#1jwV1R2R6itQUjPza9yqTE8a8zu",
	"_md": {},
	"domain": "example.trustgrid.io",
	"SK": "Alert#Node Disconnect",
	"_tp": "Alert",
	"PK": "Node#0895b104-5434-447b-8577-############",
	"state": "UNKNOWN",
	"nodeId": "0895b104-5434-447b-8577-############",
	"timestamp": 1604714923, /* Unix epoch timestamp when the event was first triggered */
	"channelID": "bc47ca84-1d04-454b-bedc-a55d1a917c0e", /* The unique id of the channel used to deliver this message. */
	"notes": ["Text from Description Field"],
	"alarmIDs": [ 
		/* A list of alarm filters that matched the event */
		"be324011-4bea-4392-a06a-541646decd39"
	]
}
{{</highlight>}}

### Event Data Descriptions
