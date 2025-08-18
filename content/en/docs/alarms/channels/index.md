---
Title: "Channels"
Tags: ["alarms", "channels"]
Date: 2022-12-28
---

{{% pageinfo %}}
A channel defines one or more method of delivering alert notifications to external systems.  
{{% /pageinfo %}}

## Notification Delivery Channels

### Email Channel
One or more email address (comma separated) that will receive messages from alerts@trustgrid.io

### PagerDuty Channel
Trustgrid will generate an incident via the PagerDuty API if provided a valid API routing key. To procure a routing key, create a [service in PagerDuty](https://support.pagerduty.com/main/docs/services-and-integrations) and add an `Events API V2` integration. After adding the integration, the `Integration Key` is your API routing key.

Copy the routing key to the Trustgrid channel definition.

### OpsGenie Channel
Trustgrid will generate an incident via the OpsGenie API if provided a valid API key with read and create and update permissions.

{{<alert>}} For both PagerDuty and OpsGenie the integration will automatically resolve issues if an [event]({{<ref "docs/alarms/events" >}}) occurs that negates the initial triggering event. For example, if an [event]({{<ref "docs/alarms/events" >}}) is triggered by a Node Disconnect and the [node]({{<ref "docs/nodes" >}}) reconnects, the Node Connect [event]({{<ref "docs/alarms/events" >}}) will resolve the incident via the API. {{</alert>}}

### Slack Channel
Trustgrid can post the [event]({{<ref "docs/alarms/events" >}}) data to a configured channel via an incoming webhook. First, [create the webhook](https://api.slack.com/messaging/webhooks), and then copy the webhook URL into the Trustgrid channel definition.

Optionally, you can configure the slack event to be posted with formatting to make it [easier to read]({{<ref "#example-formatted-slack-event">}}) as opposed to [raw JSON]({{<ref "#example-event-json">}}). {{<tgimg src="slack-format-option.png" width="50%" caption="Slack format option checkbox">}}

### Microsoft Teams Channel
Trustgrid can post [event]({{<ref "docs/alarms/events" >}}) data to a configured Teams channel via an incoming webhook. First, [create the webhook](https://learn.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook), and then copy the webhook URL into the Trustgrid channel definition.


### Google Chat Channel 
Trustgrid can post [event]({{<ref "docs/alarms/events" >}}) data to a configured Google Chat (gChat) space via an incoming webhook.  First [create the webhook](https://developers.google.com/workspace/chat/quickstart/webhooks#create-webhook) and then copy the webhook URL into the Trustgrid channel definition.

{{<alert>}} Only a single Slack or Teams channel, or Google Chat Space can be targeted by a Trustgrid channel definition. However, you can create multiple Trustgrid channels if you wish to post the [event]({{<ref "docs/alarms/events" >}}) data to more than one Slack/Teams channel. {{</alert>}}

### Generic Webhook 

Trustgrid can send event data to any HTTP endpoint using a generic webhook channel. This allows for integration with a wide variety of systems and services. The Event Data payload below is sent in `application/json` format to the URL specified. Any authentication credentials should be specified in the URL. 

## Example Event Data

The [event]({{<ref "docs/alarms/events" >}}) data is delivered in JSON, as shown below, which depending on the integration can allow for additional parsing.
### Example Event JSON
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

### Example Formatted Slack Event
Below is an example of a formatted Slack event.
{{<tgimg src="formatted-slack-example.png" width="80%" caption="Example Slack Event">}}