---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Network Objects"
date: 2023-1-9
weight: 10
---

{{% pageinfo %}}
A network object allows naming a CIDR inside a VPN for easier rule management.
{{% /pageinfo %}}

A network object requires the following fields:

{{<fields>}}
{{<field "Name" >}}
The label to use when referencing this object.
{{</field >}}

{{<field "CIDR" >}}
Specifically an IPv4 CIDR.
{{</field >}}
{{</fields>}}

![img](network-objects.png)

Multiple network objects may be clustered into a [network group]({{<ref "docs/domain/virtual-networks/network-groups" >}}).
