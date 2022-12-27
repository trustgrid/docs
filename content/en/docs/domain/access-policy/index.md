---
categories: ["concepts"]
tags: ["domain", "concepts", "rewrite"]
title: "Access Policy"
date: 2022-12-22
---

{{% pageinfo %}}
An access policy is a list of rules that each allow, reject, or drop traffic at the VPN level.
{{% /pageinfo %}}
 
If no rules are defined, traffic will not pass through the [virtual network]({{< ref "docs/domain/virtual-networks" >}}).


 **A rule can match traffic based on the following fields:**


{{< field-def "Action" >}}
Whether to ALLOW, DROP, or REJECT the traffic.
{{< /field-def >}}
  
{{< field-def "Protocol" >}}
You can choose ANY, UDP, ICMP, or TCP.
{{< /field-def >}}
  
{{< field-def "Source" >}}
The source of the traffic. You can select a network object (link), a network group (link), or provide a CIDR.
{{< /field-def >}}
  
{{< field-def "Destination" >}}
The destination of the traffic. You can select a network object (link), a network group (link), or provide a CIDR.
{{< /field-def >}} 

{{< field-def "Line Number" >}}
Rules are evaluated starting with the lowest numbered - rule. Once a rule matches, later rules are ignored even if they might also match the traffic (confirm w/ Steven)
{{< /field-def >}}

{{< field-def "Port Range">}}
For TCP and UDP traffic, you can specify a port (e.g., 80) or a range (e.g., 8000-9000).
{{< /field-def >}}


![img](access-policy.png)

**Note that _all specified_ fields must be matched for the rule to match. (confirm w steven).**