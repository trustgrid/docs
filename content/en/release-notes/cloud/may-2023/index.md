---
title: May 2023 Release Notes
linkTitle: 'May 2023 Release'
type: docs
date: 2023-05-07
description: May Release focusing on Audit and UI improvements, plus bug fixes
---

## Accessibility Improvements
Prior to this release we used only the color red to indicate if the control or data plane was disconnected.  We now use different icons to indicate there is an issue with connectivity.
### Control Plane Disconnected 
{{<tgimg src="control-plane-disconnected.png" caption="Control Plane Disconnect icon" alt="Red circle with an exclamation mark(!) inside to indicate control plane disconnected">}}

### Data Plane Disconnected
{{<tgimg src="data-plane-disconnected.png" caption="Data Plane Disconnected icon" alt="Red triangle with an exclamation mark(!) to indicate data plane is disconnected ">}}

## Change Audit Improvements
Several improvements where made around our change auditing system including:
* Tag changes are now audited.
* Change records can now be replicated into a customer's AWS S3 bucket. This works much like the existing [S3 Flow Log Export]({{<ref "/docs/operations/flow-logs#flow-log-export">}}) and requires the same bucket policy and versioning settings.  To have this setup contact Trustgrid support. 
* Changes to Container settings now include the name of the container instead of its unique ID.
* Changes to a policy will now show what resources are covered by the policy.
* Group deletion is now audited with name.
* Authentication records using a configured [Identity Provider(IdP)]({{<ref src="/docs/idps">}}) now include the IdP used.
* The advanced search now allows you to select more object types.

## L4 Proxy UI Improvements
This release also makes several improvements to the UI for configuring and managing our Layer 4 (L4) proxy [Services]({{<ref "/docs/nodes/services">}}) and [Connectors]({{<ref "/docs/nodes/connectors">}}). 

These improvement include:
* More consistency in the names of fields between the add/edit prompt and the table listings.
* The addition of links to [automatically test connectivity of a TCP service]({{<ref "/docs/nodes/services#test-connectivity">}}) and [sniff traffic for a connector listening port]({{<ref "/docs/nodes/connectors#sniff-traffic">}})

## Flow Log Advanced Search
With this release you can now set the source or destination node to `local` for an Advanced Flow Log Search. This will cover flows that do not traverse the data plane, such as traffic that is forwarded between interfaces.

## Formatted Slack Alarm Events
Previously when an [event]({{<ref "/docs/alarms/events">}}) was forwarded to a [Slack Channel]({{<ref "/docs/alarms/channels#slack-channel">}}) it was sent as raw JSON that was difficult to read. There is now an option to format messages sent to Slack to be human readable. 
{{<tgimg src="/docs/alarms/channels/slack-format-option.png" width="50%" caption="Checkbox to enable Slack formatting">}}

With this selected message will appear in Slack like the below example.
{{<tgimg src="/docs/alarms/channels/formatted-slack-example.png" width="80%" caption="Example formatted slack event">}}

## WireGuard Tunnel Client Config
This release adds an [example WireGuard client configuration]({{<ref "/docs/nodes/tunnels/WireGuard#example-client-config">}}) to both the add and update dialogues for WireGuard tunnel interfaces. 

{{<tgimg src="/docs/nodes/tunnels/WireGuard/WireGuard-tunnel-example-config.png" width="90%" caption="Example auto-generated WireGuard client configuration">}}