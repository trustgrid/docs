---
Title: "Alarm Filters"
Tags: ["alarms", "alarm filters"]
Date: 2022-12-28
---

{{% pageinfo %}}
Alarm Filters are used to determine which [events]({{< ref "docs/alarms/events" >}}) trigger notifications and to define which [channels]({{< ref "docs/alarms/channels" >}}) should receive those notifications. 
{{% /pageinfo %}}

#### Name and Description
The name needs to be unique. Both the name and description are displayed in the alarm filters table.

#### Enabled
An alarm filter must be enabled for matching [event]({{< ref "docs/alarms/events" >}}) to be sent to the selected [channel]({{< ref "docs/alarms/channels" >}}).  Deselecting the check box can be handy if you wish to suppress a specific type of [alarm]({{< ref "docs/alarms" >}}).

#### Channels
This section determines which [channels]({{< ref "docs/alarms/channels" >}}) matching [alarms]({{< ref "docs/alarms" >}}) will be sent to.

#### Alarm Filter Criteria
The criteria determine which [events]({{< ref "docs/alarms/events" >}}) will match the filter. These conditions can be set as:

-All (default) - All specified criteria must be true to match. Equivalent to an AND condition.

-Any - Only one criteria must be true to match. Equivalent to an OR condition.

-None - The specified criteria must be false to match. Equivalent to a NOT-condition.

#### Node Name Criteria
The "Node Name" criteria llows you to select one or more specific [node]({{< ref "docs/concepts/node" >}}) names. Note, even if the filter is set to `All`, the filter will match any of the selected [node]({{< ref "docs/concepts/node" >}}) names is associated with the [event]({{< ref "docs/alarms/events" >}}).

#### Event Type Criteria
The ["Event Type"]({{< ref "docs/alarms/event-types" >}}) criteria determines which [events]({{< ref "docs/alarms/events" >}}) will match the filter. Note, even if the filter is set to `All`, the filter will match any of the selected [event types]({{< ref "docs/alarms/event-types" >}}).

#### Tag Matches Criteria
The "Tag Matches" criteria allows you to use [tag]({{< ref "docs/concepts/tag" >}}) name/value pairs to determine if the filter should match [events]({{< ref "docs/alarms/events" >}}). For examples, you may what production devices to send to a high priority [channel]({{< ref "docs/alarms/channels" >}}) such as PagerDuty or OpsGenie. If your [nodes]({{< ref "docs/concepts/node" >}}) have a tag to indicating “prod_status=production”, you can select that name/value pair from the list to properly filter your [alarms]({{< ref "docs/alarms" >}}). 

![img](/docs/alarms/tag-matches.png)

#### Severity Threshold Criteria 
Each [event]({{< ref "docs/alarms/events" >}}) type has a severity level associated with it. This filter will match any [event]({{< ref "docs/alarms/events" >}}) with the selected severity type or higher. This is the only mandatory criteria.

The severity levels are:

1. INFO

2. WARNING

3. ERROR

4. CRITICAL

For example, if you select the severity level of WARNING the filter will match WARNING, ERROR and  CRITICAL [events]({{< ref "docs/alarms/events" >}}).

{{<alert>}} Some events have a corresponding [event]({{< ref "docs/alarms/events" >}}) that will automatically resolve the alert in the portal and in some [channels]({{< ref "docs/alarms/channels" >}}) such as PagerDuty. The corresponding event may have a different severity level, so make sure you select the lower severity for the criteria. e.g. [Node]({{< ref "docs/concepts/node" >}}) Disconnect is a WARNING but [Node]({{< ref "docs/concepts/node" >}}) Connect which resolves it is only INFO. So you’d need to select both [Event Types]({{< ref "docs/alarms/event-types" >}}) and set the severity to INFO. {{</alert>}}

#### Contains Text Criteria
This field will accept any single string of text to match to the contents of an [event]({{< ref "docs/alarms/events" >}}). For example, if all your gateways include `-gw` in the name you could enter that without quotes in the field and it would match any [event]({{< ref "docs/alarms/events" >}}) that includes that text in the event payload. This criteria can also be used if there is another aspect of the node included in the [event]({{< ref "docs/alarms/events" >}}) payload that doesn’t match the criteria above. To see the entire payload of an [event]({{< ref "docs/alarms/events" >}}) configure a less specific payload and send to an email [channel]({{< ref "docs/alarms/channels" >}}) to see the JSON.

> The [event]({{< ref "docs/alarms/events" >}}) payload includes the [node’s]({{< ref "docs/concepts/node" >}}) unique identifier (UID) which is a string of generated text and numbers. If your "Contains Text" criteria is too short, there is a chance a [node]({{< ref "docs/concepts/node" >}}) UID will also match unexpectedly.

