---
categories: ["overview"]
tags: ["architecture"]
title: "Basic Architecture"
date: 2022-12-28
description: >
  Trustgrid basic architecture overview
---

### Cloud Components

{{< field-def "Portal" >}}
The cloud management UI.
{{< /field-def >}}

{{< field-def "Gatekeeper" >}}
Provides configuration updates to edge nodes.
{{< /field-def >}}

{{< field-def "Zuul" >}}
Maintains persistent connection with [nodes]({{< ref "docs/node" >}}) for reporting, [events]({{< ref "docs/alarms/events" >}}), and updates.
{{< /field-def >}}

{{< field-def "API" >}}
The management API that exposes 100% of UI elements to automation.
{{< /field-def >}}

{{< field-def "Repo" >}}
The APT repository that stores all firmware, OS, and [node]({{< ref "docs/node" >}}) updates.
{{< /field-def >}}

### Edge Components

{{< field-def "Node" >}}
The software that provides core functionality in the edge including [networking]({{< ref "docs/overview/networking" >}}), [security]({{< ref "docs/overview/security" >}}), compute, and management features. Edge nodes build outbound connections to gateway nodes.
{{< /field-def >}}

### Use Cases

{{< field-def "Software Defined Networking" >}}
Create a mesh [network]({{< ref "docs/overview/networking" >}}) that connects cloud applications to edge data with load balancing, [clustering]({{< ref "/docs/cluster" >}}), and failover managed through a portal or API.
{{< /field-def >}}

{{< field-def "Edge Compute" >}}
Deploy applications to the edge to access datasets not appropriate for replication to the cloud due to [security]({{< ref "docs/overview/security" >}}) or compliance concerns, latency, or cost.
{{< /field-def >}}

{{< field-def "Device Management">}}
Manage thousands of [nodes]({{< ref "docs/node" >}}) with advanced tools to reduce the burden of operations at enterprise scale.
{{< /field-def >}}

{{< field-def "Edge API" >}}
Integrate thousands of edge datasets with a single API interface and ETL functions executing at the edge.
{{< /field-def >}}

{{< field-def "Security" >}}
Leverage Trustgrid's advanced [security]({{< ref "docs/overview/security" >}}) to protect against a wide range of threats.
{{< /field-def >}}
