---
tags: ["applications", "ztna"]
title: "Access Policy"
date: 2023-01-06
---

{{% pageinfo %}}
An Access Policy is comprised of several rules that determine whether a user is permitted to access a ZTNA application. If no rule matches a user's access attempt, the attempt will be denied.
{{% /pageinfo %}}

Each rule has three types of criteria:

- Include - criteria in this section are logically OR'd together. For a rule to apply to an access attempt, at least one of the include criteria must match.
- Exception - criteria in this section are logically NOT'd. If any exception criterion matches, the rule will not apply to the access attempt.
- Require - criteria in this section are logically AND'd. ALL of the require criteria must match for the rule to apply to the access attempt.

Rules also have an action:

- Allow - if the rule matches, allow access
- Block - if the rule matches, block access

Rules are evaluated in the order they are listed in the policy. Rules may be moved higher or lower in the list by clicking the up or down arrows next to the rule.

INSERT SCREENSHOT

### Criteria

{{< field-def "Emails" >}}
a list of email addresses, comma separated. Eg, "user1@company.com, user2@company.com"
{{< /field-def >}}
{{< field-def "Country" >}}
the user's country, determined by their IP address
{{< /field-def >}}
{{< field-def "Emails ending in" >}}
a required email suffix, like "@company.com"
{{< /field-def >}}
{{< field-def "Everyone" >}}
this rule always matches
{{< /field-def >}}
{{< field-def "IDP Groups" >}}
a list of [IDP Groups](put-link-here) that the user must be a member of
{{< /field-def >}}
{{< field-def "IP Ranges" >}}
a list of IP ranges either in CIDR notation or a single IP address, comma separated. Eg, "10.10.44, 10.10.8.0/24"
{{< /field-def >}}

INSERT SCREENSHOT OF CRITERIA FILLED OUT