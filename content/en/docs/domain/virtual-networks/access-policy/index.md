---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Access Policy"
date: 2022-12-28
weight: 40
---

{{% pageinfo %}}
An access policy is a list of rules that each allow, reject, or drop traffic at the VPN level.
{{% /pageinfo %}}

If no rules are defined, traffic will not pass through the [virtual network]({{<ref "docs/domain/virtual-networks" >}}).

**A rule can match traffic based on the following fields:**

{{<fields>}}
{{<field "Action" >}}
Whether to ALLOW, DROP, or REJECT the traffic.
{{</field >}}

{{<field "Protocol" >}}
You can choose ANY, UDP, ICMP, or TCP.
{{</field >}}

{{<field "Source" >}}
The source of the traffic. You can select a [network object]({{<ref "docs/domain/virtual-networks/network-objects" >}}), a [network group]({{<ref "docs/domain/virtual-networks/network-groups" >}}), or provide a CIDR.
{{</field >}}

{{<field "Destination" >}}
The destination of the traffic. You can select a [network object]({{<ref "docs/domain/virtual-networks/network-objects" >}}), a [network group]({{<ref "docs/domain/virtual-networks/network-groups" >}}), or provide a CIDR.
{{</field >}}

{{<field "Line Number" >}}
Rules are evaluated starting with the lowest numbered rule. Once a rule matches, later rules are ignored even if they might also match the traffic. 
{{</field >}}

{{<field "Port Range">}}
For TCP and UDP traffic, you can specify a port (e.g., 80) or a range (e.g., 8000-9000).
{{</field >}}
{{</fields>}}

![img](access-policy.png)

**Note that _all specified_ fields must be matched for the rule to match.**
