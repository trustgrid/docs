---
title: May 2023 Release Notes
linkTitle: 'May 2023 Release'
type: docs
date: 2023-05-07
description: May Release focusing on Audit and UI improvements, plus bug fixes
---

## Change Audit Improvements
Several improvements where made around our change auditing system including:
* Tag changes are now Audited.
* Change records can now be replicated into a customer's AWS S3 bucket. This works much like the existing [S3 Flow Log Export]({{<ref "/docs/operations/flow-logs#flow-log-export">}}) and requires the same bucket policy and versioning settings.  To have this setup contact Trustgrid support. 
* Changes to Container settings now include the name of the container instead of its unique ID.
* Changes to a policy will now show what resources are covered by the policy.

## Flow Log Advanced Search
With this release you can now set the source or destination node to `local` for an Advanced Flow Log Search. This will cover flows that do not traverse the data plane, such as traffic that is forwarded between interfaces.

## Formatted Slack Alarm Events
Previously when an [event]({{<ref "/docs/alarms/events">}}) was forwarded to a [Slack Channel]({{<ref "/docs/alarms/channels#slack-channel">}}) it was sent as raw JSON that was difficult to read. There is now an option to format messages sent to Slack to be human readable. 
{{<tgimg src="/docs/alarms/channels/slack-format-option.png" width="50%" caption="Checkbox to enable Slack formatting">}}

With this selected message will appear in Slack like the below example.
{{<tgimg src="/docs/alarms/channels/formatted-slack-example.png" width="80%" caption="Example formatted slack event">}}
