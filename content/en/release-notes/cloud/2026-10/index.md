---
title: "October 2026 Cloud Release Notes"
linkTitle: "October 2026"
date: 2026-10-05
description: "October 2026 Cloud Release Notes"
type: docs
---

## October 28, 2026 - Minor Release

### Packet Simulator in Interface Tools

A new **Packet Simulator** button in the node **Interface Tools** section simulates a TCP, UDP, or ICMP packet entering the selected interface. The node traces the packet through its NAT, ACL, and routing rules without sending traffic, and the portal shows each step along with whether the packet was forwarded or blocked. The built-in **tg-monitor**, **tg-operator**, and **tg-node-admin** policies grant the required `nodes::service:packet-sim` permission by default. The tool requires a node running the [June 2026 Major Appliance release]({{<relref "release-notes/node/2026-06" >}}) or newer. [Learn more]({{<relref "tutorials/interface-tools/packet-simulator" >}}).

## October 21, 2026 - Minor Release

### Negotiated TLS details in Data Plane

The Data Plane panel now offers **TLS Version** and **TLS Cipher** as optional per-peer columns for nodes running the [September 2026 Major Appliance release]({{<relref "release-notes/node/2026-09" >}}) or newer. [Learn more]({{<relref "docs/nodes/appliances/data-plane" >}}).

## October 5, 2026 - Minor Release

### Custom Slack alarm messages

The alarm-channel editor now supports custom Slack messages built from Section, Context, and Message blocks. Section blocks lay out selected fields in two columns with editable labels. Context blocks show supporting fields as compact metadata. Message blocks let you write free-form text and optionally insert alert values. Choose and arrange the blocks and fields to fit your alert; the examples in the [Slack channel documentation]({{<relref "docs/alarms/channels#custom-slack-message-blocks" >}}) are not required defaults.
