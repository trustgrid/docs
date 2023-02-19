---
Title: "Cluster Site Health State"
Tags: ["concepts", "clusters", "cluster health", "cluster management"]
Date: 2022-12-28
---

{{% pageinfo %}}
A healthy [cluster]({{<ref "docs/clusters">}}) should have two online Trustgrid [nodes]({{<ref "docs/nodes" >}}) in a healthy state. The Trustgrid control plane can monitor the member [nodes]({{<ref "docs/nodes" >}}) and generate [events]({{<ref "docs/alarms/events" >}}) when their state changes such that the health of the [cluster]({{<ref "docs/clusters" >}}) is impacted.
{{% /pageinfo %}}

### Basics

{{<alert>}}
These events are based on the information reported by each member node to the control plane. Because of this:

- There can be a delay between a change and the [Cluster]({{<ref "docs/clusters" >}}) Site Health events.
- If member [nodes]({{<ref "docs/nodes" >}}) have issues preventing control plane connections but not impacting data plane or other services the [events]({{<ref "docs/alarms/events" >}}) may not accurately reflect the state of the [cluster]({{<ref "docs/clusters" >}}).
  {{</alert>}}

### Cluster Health State

The Trustgrid will now display one of three states in the portal:

| State         | Meaning                                                                                                                                                        |
| ------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Healthy**   | Both members are online with a single member active.                                                                                                           |
| **Unhealthy** | There is at least one online and healthy member of the [cluster]({{<ref "docs/clusters" >}}) with the active role AND one or more of these conditions is true: |
|               | - One of the [cluster]({{<ref "docs/clusters" >}}) members is offline, unhealthy or disabled                                                                   |
|               | - There are **two** [cluster]({{<ref "docs/clusters" >}}) members with the active role                                                                         |
| **Offline**   | Both [nodes]({{<ref "docs/nodes">}}) are offline, unhealthy or disabled.                                                                                       |
