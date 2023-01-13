---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Applying Changes"
date: 2023-1-9
---

{{% pageinfo %}}
Because changes to a domain [virtual network]({{< ref "docs/domain/virtual-networks" >}}) affect _all_ [nodes]({{< ref "docs/node" >}}), [domain]({{< ref "docs/domain" >}}) configuration changes must be reviewed before applied, and are applied all at once.
{{% /pageinfo %}}

After making changes to a [virtual network]({{< ref "docs/domain/virtual-networks" >}}), you can review outstanding changes. A list of changes will show who made each change, with a before/after table if appropriate.

![img](outstanding-changes.png)

Individual changes can be rejected by clicking the delete icon at the top.

Once changes are applied, outstanding changes will disappear from the review page, and [nodes]({{< ref "docs/node" >}}) will receive the configuration updates.

