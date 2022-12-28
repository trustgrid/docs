---
title: "Production Status Tags"
date: 2022-12-28
---

{{% pageinfo %}}
Production status tags are used for for organizing large Trustgrid deployments.
{{% /pageinfo %}}

It is recommended that customers use a [tag]({{< ref "docs/concepts/tag" >}}) to indicate if [nodes]({{< ref "docs/concepts/node" >}}) are currently in production or not. For example, you may wish to have a [tag]({{< ref "docs/concepts/tag" >}}) such as prod_status with possible values like

- `deploying` for devices still being deployed
- `production` for devices actively in user
- `decommission` for devices that are being removed

{{<alert>}}
It’s important to have these [tag]({{< ref "docs/concepts/tag" >}}) names and values be consistent within the organization including case. `Production` and `production` would be viewed as two different values, as an example.
{{</alert>}}
