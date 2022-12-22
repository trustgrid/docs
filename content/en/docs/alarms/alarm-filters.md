---
Title: "Alarm Filters"
Tags: ["alarms", "alarm filters"]
Date: 2022-12-22
---

{{% pageinfo %}}
Alarm Filters are used to determine which events trigger notifications and to define which channels should receive those notifications. 
{{% /pageinfo %}}

#### Name and Description
The name needs to be unique. Both the Name and Description are displayed in the Alarm Filters table.

#### Enabled
An alarm filter must be enabled for matching events to be sent to the selected channel.  Deselecting the check box can be handy if you wish to suppress a specific type of [alarm]({{< ref "docs/alarms" >}}).

#### Channels
This section determines which Channels matching [alarms]({{< ref "docs/alarms" >}}) will be sent to.

#### Alarm Filter Criteria
The criteria determine which events will match the filter. These conditions can be set as:

-All (default) - All specified criteria must be true to match. Equivalent to an AND condition.

-Any - Only one criteria must be true to match. Equivalent to an OR condition.

-None - The specified criteria must be false to match. Equivalent to a NOT condition.

#### Node Name Criteria
Allows you to select one or more specific [node]({{< ref "docs/concepts/node" >}}) names. Note, even if the filter is set to "All", the filter will match any of the selected [node]({{< ref "docs/concepts/node" >}}) names is associated with the event.

#### Event Type Criteria
Determines which [Events]({{< ref "docs/alarms/events" >}}) will match the filter. Note, even if the filter is set to "All", the filter will match any of the selected Event Types.

#### Tag Matches Criteria
The Tag Matches criteria allows you to use [tag]({{< ref "docs/concepts/tag" >}}) name/value pairs to determine if the filter should match [events]({{< ref "docs/alarms/events" >}}). For examples, you may what production devices to send to a high priority channel such as PagerDuty or OpsGenie. If your [nodes]({{< ref "docs/concepts/node" >}}) have a tag to indicating “prod_status=production”, you can select that name/value pair from the list to properly filter your [alarms]({{< ref "docs/alarms" >}}). 

![img](/docs/alarms/tag-matches.png)

#### Severity Threshold Criteria 
Each [event]({{< ref "docs/alarms/events" >}}) type has a severity level associated with it. This filter will match any [event]({{< ref "docs/alarms/events" >}}) with the selected severity type or higher. This is the only mandatory criteria.

The severity levels are:

1. INFO

2. WARNING

3. ERROR

4. CRITICAL

For example, if you select the severity level of WARNING the filter will match WARNING, ERROR and  CRITICAL events.

>- Some events have a corresponding [event]({{< ref "docs/alarms/events" >}}) that will automatically resolve the alert in the portal and in some channels such as PagerDuty. The corresponding event may have a different severity level, so make sure you select the lower severity for the criteria. e.g. [Node]({{< ref "docs/concepts/node" >}}) Disconnect is a WARNING but [Node]({{< ref "docs/concepts/node" >}}) Connect which resolves it is only INFO. So you’d need to select both Event Types and set the Severity to INFO.

#### Contains Text Criteria
This field will accept any single string of text to match to the contents of an [event]({{< ref "docs/alarms/events" >}}). For example, if all your gateways include “-gw” in the name you could enter that without quotes in the field and it would match any event that includes that text in the event payload. This criteria can also be used if there is another aspect of the node included in the event payload that doesn’t match the criteria above. To see the entire payload of an event configure a less specific payload and send to an email channel to see the JSON.

>- The event payload includes the [node’s]({{< ref "docs/concepts/node" >}}) unique identifier (UID) which is a string of generated text and numbers. If your Contains Text criteria is too short there is a chance a [node]({{< ref "docs/concepts/node" >}}) UID will also match unexpectedly.

