---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Domain"
date: 2022-12-19

---

{{% pageinfo %}}
A domain is the logical grouping of [nodes]({{< ref "node" >}}) inside an organization.
{{% /pageinfo %}}

Domains are used to group [nodes]({{< ref "node" >}}) into security groups. Nodes can only connect to other [nodes]({{< ref "node" >}}) within the same domain. Trustgrid supports domains in the typical FQDN format (ie, customer_name.trustgrid.io, group1.customer_name.trustgrid.io). Each Trustgrid deployment will have a default domain.

