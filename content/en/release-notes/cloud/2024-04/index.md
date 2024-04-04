---
title: April 2024 Release Notes
linkTitle: April 2024 Release
date: 2024-04-04
description: "April 2024 cloud release introducing Alert Center"
type: docs
---
## Introducing Alert Center
A new [Alert Center]({{<relref "/docs/nodes/shared/alert-center">}}) has been added to the node page to centralize alert messages and simplify resolving multiple alerts. This removes the alerts from the [Infovisor]({{<relref "/docs/nodes/shared/infovisor">}}) menu and makes resolving multiple alerts easier.
{{<tgimg src="/docs/nodes/shared/alert-center/alert-center.png" width="60%" caption="Example alert center with 3 open alerts">}}

## Other Improvements and Fixes
- Further improvements in [flow log]({{<relref "/help-center/flow-logs">}}) query performance for faster results. 
- Detaching Policies now prompts for confirmation.
- Renaming nodes now forces all characters to lower-case. Lower-case has been enforced on creation but this wasn't previously enforced on rename.
