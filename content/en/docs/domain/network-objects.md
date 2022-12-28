---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Network Objects"
date: 2022-12-28
---

{{% pageinfo %}}
A network object allows naming a CIDR inside a [VPN]({{< ref "docs/concepts/VPN" >}}) for easier rule management.
{{% /pageinfo %}}

A network object requires the following fields:

{{< field-def "Name" >}}
The label to use when referencing this object.
{{< /field-def >}}

{{< field-def "CIDR" >}}
Specifically an IPv4 CIDR.
{{< /field-def >}}

![img](/docs/domain/network-objects.png)


Multiple network objects may be clustered into a [network group]({{< ref "docs/domain/network-groups" >}}).