---
Title: "Alarm Filters"
Tags: ["alarms", "alarm filters"]
Date: 2022-12-28
---

{{% pageinfo %}}
Alarm Filters are used to determine which [events]({{< ref "events" >}}) trigger notifications and to define which [channels]({{< ref "channels" >}}) should receive those notifications.
{{% /pageinfo %}}

#### Name and Description

The name needs to be unique. Both the name and description are displayed in the alarm filters table.

#### Enabled

An alarm filter must be enabled for matching [event]({{< ref "events" >}}) to be sent to the selected [channel]({{< ref "channels" >}}). Deselecting the check box can be handy if you wish to suppress a specific type of [alarm]({{< ref "docs/alarms" >}}).

#### Channels

This section determines which [channels]({{< ref "channels" >}}) matching [alarms]({{< ref "docs/alarms" >}}) will be sent to.

#### Alarm Filter Criteria

The criteria determine which [events]({{< ref "events" >}}) will match the filter. These conditions can be set as:

-All (default) - All specified criteria must be true to match. Equivalent to an AND condition.

-Any - Only one criteria must be true to match. Equivalent to an OR condition.

-None - The specified criteria must be false to match. Equivalent to a NOT-condition.

#### Node Name Criteria

The "Node Name" criteria llows you to select one or more specific [node]({{< ref "docs/node" >}}) names. Note, even if the filter is set to `All`, the filter will match any of the selected [node]({{< ref "docs/node" >}}) names is associated with the [event]({{< ref "events" >}}).

#### Event Type Criteria

The ["Event Type"]({{< ref "event-types" >}}) criteria determines which [events]({{< ref "events" >}}) will match the filter. Note, even if the filter is set to `All`, the filter will match any of the selected [event types]({{< ref "event-types" >}}).

#### Tag Matches Criteria

The "Tag Matches" criteria allows you to use [tag]({{< ref "docs/concepts/tag" >}}) name/value pairs to determine if the filter should match [events]({{< ref "events" >}}). For examples, you may what production devices to send to a high priority [channel]({{< ref "channels" >}}) such as PagerDuty or OpsGenie. If your [nodes]({{< ref "docs/node" >}}) have a tag to indicating “prod_status=production”, you can select that name/value pair from the list to properly filter your [alarms]({{< ref "docs/alarms" >}}).

![img](tag-matches.png)

#### Severity Threshold Criteria

Each [event]({{< ref "events" >}}) type has a severity level associated with it. This filter will match any [event]({{< ref "events" >}}) with the selected severity type or higher. This is the only mandatory criteria.

The severity levels are:

1. INFO

2. WARNING

3. ERROR

4. CRITICAL

For example, if you select the severity level of WARNING the filter will match WARNING, ERROR and CRITICAL [events]({{< ref "events" >}}).

{{<alert>}} Some events have a corresponding [event]({{< ref "events" >}}) that will automatically resolve the alert in the portal and in some [channels]({{< ref "channels" >}}) such as PagerDuty. The corresponding event may have a different severity level, so make sure you select the lower severity for the criteria. e.g. [Node]({{< ref "docs/node" >}}) Disconnect is a WARNING but [Node]({{< ref "docs/node" >}}) Connect which resolves it is only INFO. So you’d need to select both [Event Types]({{< ref "event-types" >}}) and set the severity to INFO. {{</alert>}}

#### Contains Text Criteria

This field will accept any single string of text to match to the contents of an [event]({{< ref "events" >}}). For example, if all your gateways include `-gw` in the name you could enter that without quotes in the field and it would match any [event]({{< ref "events" >}}) that includes that text in the event payload. This criteria can also be used if there is another aspect of the node included in the [event]({{< ref "events" >}}) payload that doesn’t match the criteria above. To see the entire payload of an [event]({{< ref "events" >}}) configure a less specific payload and send to an email [channel]({{< ref "channels" >}}) to see the JSON.

> The [event]({{< ref "events" >}}) payload includes the [node’s]({{< ref "docs/node" >}}) unique identifier (UID) which is a string of generated text and numbers. If your "Contains Text" criteria is too short, there is a chance a [node]({{< ref "docs/node" >}}) UID will also match unexpectedly.
