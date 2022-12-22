---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Applying Changes"
date: 2022-12-22
---

{{% pageinfo %}}
Because changes to a domain virtual network affect *all* [nodes]({{< ref "docs/concepts/node" >}}), [domain]({{< ref "docs/domain" >}}) configuration changes must be reviewed before applied, and are applied all at once.
{{% /pageinfo %}}


After making changes to a virtual network, you can review outstanding changes. A list of changes will show who made each change, with a before/after table if appropriate.


![img](/docs/domain/outstanding-changes.png)


Individual changes can be rejected by clicking the delete icon at the top. 


Once changes are applied, outstanding changes will disappear from the review page, and nodes will receive the configuration updates.