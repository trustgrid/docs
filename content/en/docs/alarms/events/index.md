---
Title: "Events"
Tags: ["events", "alarms"]
Date: 2022-12-22
---

{{% pageinfo %}}
Events are emitted from nodes and from the Trustgrid control plane when actionable things happen. Events are the basis for alarms and notifications.
{{% /pageinfo %}}

Events can be viewed for individual nodes by navigating to `Events` under the `History` section.

![img](node-events-list.png)

Events for the entire organization can be viewed by navigating to `Events` under the `Alarms` section.

Clicking the `Test` button will send the event through your configured alarms to help verify channels are configured as expected.

#### Event Times

{{<fields>}}
{{<field "Generated Time" >}}
The time the event was created.
{{</field >}}

{{<field "Received Time" >}}
The time the event was received by the control plane. This can be later than the generated time in the event of a network disruption, for eample.
{{</field >}}
{{</fields>}}
