---
title: Incident.io Webhook Integration
linkTitle: Incident.io Webhook
description: Learn how to integrate Trustgrid with Incident.io to generate Alerts from Trustgrid events using the generic webhook.
---

Trustgrid can be integrated with [Incident.io](https://incident.io/) using the generic webhook channel to send alerts based on Trustgrid events. This allows for streamlined incident management and improved response times. The guide below outlines the steps to set up this integration along with suggestions on gathering additional [attributes](https://help.incident.io/articles/8351115985-attributes-and-priorities#attributes-0) for the alerts.

## Setting Up the Integration

{{<alert color="info">}}
The Incident.io site is subject to change so any screenshots or instructions may become outdated. Please refer to the official [Incident.io documentation](https://help.incident.io) for the most current information.
{{</alert>}}

### Step 1 - Start Creating the Incident.io Alert Source

1. Log in to your [Incident.io](https://app.incident.io/) account.
1. Click on your organization name at the top and then select **Settings**.
1. Under `On Call`, click on `Alerts`.
1. Click the `+` to the right of `Sources` to create a new source. {{<tgimg src="incident-io-webhook-create-source.png" width="40%" caption="Creating a new Alert source in Incident.io">}}
1. Follow these steps to create the source {{<tgimg src="incident-io-webhook-create-source-steps.png" width="800%" caption="Steps to create the source in Incident.io">}}
    1. Search for `http`.
    1. Select the HTTP type from the list and provide a unique name.
    1. For `Type of HTTP Source` select [custom](https://help.incident.io/articles/2353344082-custom-http-alert-sources).  This allows the source to use the native Trustgrid JSON payload which can be transformed into the Incident.io format.
    1. Click **Continue** to proceed to the next step.
1. Capture the generated webhook URL and save it for later use.
    1. Select the `Query Authentication` tab. 
    1. Click the copy icon to the right of the URL.  Save this as you will need it later .
1. Set the `Transform Expression` and `Deduplication key path`
    1. Scroll down until you see the `Transform expression` section.  There will be existing JavaScript there. Select all and delete.
    1. Copy and paste the following code into the `Transform expression` section: {{<codeblock>}}var resolvedFlag = !!$.resolved; // force boolean
var statusValue = resolvedFlag ? "resolved" : "firing";

return {
  title: $.nodeName + " - " + $.eventType,
  description: $.message,
  status: statusValue,
  deduplication_key: $.uid,
  metadata: {
    nodeName: $.nodeName,
    eventType: $.eventType,
    level: $.level,
    domain: $.domain,
    tags: $.tags,
    timestamp: $.timestamp
  }
};
{{</codeblock>}} {{<tgimg src="incident-io-webhook-transform-expression.png" width="80%" caption="Transform expression in Incident.io">}}
    1. Scroll below to the `Deduplication key path` section and set the value to `$.uid`.
    1. Click **Save configuration** to finish creating the webhook source. {{<tgimg src="incident-io-webhook-save-configuration.png" width="80%" caption="Saving the webhook configuration in Incident.io">}}
1. On the right side, you should see a section saying `No alerts received yet`. In the next section, you will create a Webhook channel in Trustgrid and trigger events to generate sample Alerts for Incident.io to use to complete the configuration.

### Step 2 - Create Webhook Channel in Trustgrid
To complete the below step it is idea to identify a node that you can use to generate test events, ideally by performing an action such as restarting the node. In the example below we will create a new Alarm Filter for a specific node and test by restarting it. 

1. Login to the Trustgrid portal.
1. Navigate to Alarms > Channels.
1. Click the `+Create Channel` button. (Or optionally select and choose Edit from Actions to add to an existing Channel)
1. In the channel configuration, set the following:
   1. **Name**: Give your channel a descriptive name.
   1. **Generic Webhook**: Paste the URL you copied from Incident.io.
   1. Click **Save** to create the channel. 
   {{<tgimg src="incident-io-webhook-create-channel.png" width="100%" caption="Creating a new Webhook channel in Trustgrid">}}
1. Setup an Alarm Filter to test the new channel. (If you have an existing Alarm Filter that you can use for testing, just edit and select the channel created above and save. Then proceed to the next step.)
    1. Navigate to Alarms > Alarm Filters.
    1. Click the `+Create Alarm` button.
    1. In the filter configuration, set the following:
       1. **Name**: Give your filter a descriptive name. This example is just for testing.
       1. Make sure **Enabled** is checked.
       1. **Node name**: Select the node you identified for testing. 
       1. **Severity Threshold**: Change to `INFO` so that all relevant events match the filter including `Resolved` notifications.
       1. **Channels**: Select the channel you created above.
       1. Click **Save** to create the filter.
       {{<tgimg src="incident-io-webhook-create-alarm-filter.png" width="100%" caption="Creating a new Alarm Filter in Trustgrid">}}
1. To test the filter, [restart the node]({{<relref "/tutorials/management-tasks/restart-node">}}) you selected in the filter. This will generate a `Node Restart` event, then a `Node Disconnect` event followed by a `Node Connect` event.

{{<alert color="warning">}}The Trustgrid system will only send out matching alerts if there are **no unresolved events of that type**. If the test node has unresolved events, use the Alert Center to clear the:
1. Click the Alert Center icon in the top right corner of the Node's detail page.
1. You can use the check marks to resolve individual Alerts, or
1. Click the **Mark All As Resolved** button to clear all unresolved Alerts of that type.
{{<tgimg src="incident-io-webhook-alert-center.png" width="60%">}}
{{</alert>}}

### Step 3 - Complete Alert Source config in Incident.io
To complete the configuration we will use the generated Alerts to extract additional useful [attributes](https://help.incident.io/articles/8351115985-attributes-and-priorities#attributes-0) from the example events to surface in Incident.io. Feel free to modify the Attribute names to fit your needs.

1. Return to the Alert sources setup page in Incident.io
    1. You should now see example alert in the bottom right.  
    1. Click Continue.
    {{<tgimg src="incident-io-webhook-example-alert.png" width="80%" caption="Example alert in Incident.io">}}
1.  On the `Configure your setup` page, you can map the attributes from the incoming alerts to the fields in Incident.io. Use the example alert to help you identify the relevant fields.
    1. Leave the default `Alert title` and `Description` mappings unless you need to customize them.
    1. Click the **Edit** pencil icon to the right of `Attributes`. {{<tgimg src="incident-io-webhook-edit-attributes.png" width="80%" caption="Editing attributes in Incident.io">}}
        1. In the `Alert Payload` section, click **+"nodeName"**. As the name implies, this will include the name of the Node the event relates to. {{<tgimg src="incident-io-webhook-click-nodename.png" width="70%" caption="Adding a new attribute in Incident.io">}}
            1. Unless you have an existing relevant attribute defined, such as `Server Name`, you can create a new one. Scroll down and select `+ Add new attribute`. {{<tgimg src="incident-io-webhook-add-attribute.png" width="50%" caption="Adding a new attribute in Incident.io">}}
            1. In the new attribute configuration: 
                1. Give the attribute something descriptive, like `Node Name`. Leave all other settings at their defaults.
                1. Click **Add**
                {{<tgimg src="incident-io-webhook-add-attribute-config.png" width="70%" caption="Adding a new attribute configuration in Incident.io">}}
        1. Repeat the above steps to add the following attributes:
            1. **+eventType** - This will include the type of event that triggered the alert. 
            1. **+description** - This will copy the Description, which is the original Event's message, into an attribute.  This will allows this to be displayed when the Alert is forwarded to a tool like Slack without having to open the actual alert.
            1. **+level** - This will include the severity level of the event. 
            1. **+domain** - This is only useful if your company has multiple Trustgrid organizations with different domains. This will allow you to potentially filter and route differently based on this attribute.
            1. **+tags** - This will include all configured [Tags] {{<relref "/docs/nodes/shared/tags">}}. For this attribute, change the `Choose what results should be parsed into` to `Labels`.  {{<tgimg src="incident-io-webhook-tags-labels.png" width="70%" caption="Adding a new attribute configuration in Incident.io">}}
    1. Click **Apply** in the bottom right to complete adding Attributes.
1. Click **Save and finish** to complete the Alert configuration in Incident.io
1. At this point you can follow the Incident.io documentation to [Create a new Alert route](https://help.incident.io/articles/4007103429-creating-escalations-and-incidents-from-alerts) from the new alert source. 

## Addition Information
### Understanding the transformations
{{<highlight javascript "linenos=inline">}}
  var resolvedFlag = !!$.resolved; // force boolean
  var statusValue = resolvedFlag ? "resolved" : "firing";
  
  return {
    title: $.nodeName + " - " + $.eventType,
    description: $.message,
    status: statusValue,
    deduplication_key: $.uid,
    metadata: {
      nodeName: $.nodeName,
      eventType: $.eventType,
      level: $.level,
      domain: $.domain,
      tags: $.tags,
      timestamp: $.timestamp
    }
  };
{{</highlight>}}

* Lines 1-2 are responsible for determining the alert's status based on its resolution state. This is then convert into Incident.io's syntax of `firing` for new alerts and `resolved` for resolved alerts.
* Lines 4-20 define the structure of the alert object that will be sent to Incident.io.
* Line 5 sets the title of the alert using the node name and event type.
* Line 6 sets the description of the alert using the message from the original event.
* Line 7 sets the status of the alert using the resolved flag to determine if it is `firing` or `resolved`.
* Line 8 sets the deduplication key for the alert using the unique ID from the original event. This allows Incident.io to match a `firing` alert with its corresponding `resolved` alert.
* Lines 9-20 set various metadata fields for the alert, including the node name, event type, level, domain, tags, and timestamp.

### Testing via Curl
If you want to test the webhook integration, you can use Curl to send a sample payload to the Incident.io webhook URL. 

This can be useful if you are not seeing the expected alerts in Incident.io after triggering events in the Trustgrid portal as the curl command may return any HTTP error Incident.io is responding with. 

Here's an example command:

``` bash
curl -X POST <WEBHOOK_URL> \
-H "Content-Type: application/json" \
-d '{
	"nodeName": "edge1", 
	"expires": 1604801325, 
	"level": "INFO", 
	"eventType": "Test Event",
	"source": "EKG",
	"message": "This is just a test event. It is not real.",
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
	"timestamp": 1604714923, 
	"channelID": "bc47ca84-1d04-454b-bedc-a55d1a917c0e", 
	"notes": ["Text from Description Field"],
    "tags": { "prod_status":"production","site_name":"Main Datacenter"},
    "resolved": false
}
```
Replace `<WEBHOOK_URL>` with the actual URL of your Incident.io webhook. This command will simulate an alert being sent to Incident.io, allowing you to verify that the integration is working correctly.

If everything is working you should see a message like `{"status":"accepted","message":"Event accepted for processing","deduplication_key":"735cbacc3f07e740d26ff364a19f856aa5af95f929017538214093afb132006e"}` and then the Alert will show in Incident.io.