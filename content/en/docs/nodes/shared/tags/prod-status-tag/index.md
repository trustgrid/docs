---
title: "Production Status Tags"
date: 2022-12-28
description: Production status tags recommended for identifying the current lifecycle state of Trustgrid nodes.
---

It is recommended that customers use a [tag]({{<ref "/docs/nodes/shared/tags" >}}) to indicate if [nodes]({{<ref "docs/nodes" >}}) are currently in production or not. For example, you may wish to have a [tag]({{<ref "/docs/nodes/shared/tags" >}}) such as prod_status with possible values like

- `deploying` for devices still being deployed
- `production` for devices actively in user
- `decommission` for devices that are being removed

{{<alert>}}
It’s important to have these [tag]({{<ref "/docs/nodes/shared/tags" >}}) names and values be consistent within the organization including case. `Production` and `production` would be viewed as two different values, as an example.
{{</alert>}}
