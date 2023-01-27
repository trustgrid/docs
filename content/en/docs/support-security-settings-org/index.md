---
Title: "Supportability and Security settings for org"
Date: 1-27-2023
---
{{% pageinfo %}}
The Support page provides options for restricting Portal access and retrieving documents shared by Trustgrid.
{{% /pageinfo %}}

#### Remote Support
Disallowing remote support prevents Trustgrid staff from invoking services (like remote terminal or ping) on nodes. Changing this field is audited (put a link to audits; they're not documented yet though).

![img](support-box.png)

Users will need `support::modify` permissions to change this setting.

#### IP Restriction
Trustgrid Portal access for your organization can be limited to specific IP addresses. Multiple IP addresses can be added. Note that if your IP address changes and is not included on this list, you will be unable to access the Portal.

Users will need `support::modify` permissions to change this setting.