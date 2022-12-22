---
title: "Trustgrid Access to Nodes"
weight: 2
date: 2022-12-22
---
{{% pageinfo %}}
By default Trustgrid support has access to all [nodes]({{< ref "docs/concepts/node" >}}) via the Control Plane and subject to Trustgrid policy and procedures. Trustgrid customers can choose if they want to prohibit Trustgrid support to access their [node]({{< ref "docs/concepts/node" >}}) for support purposes. This feature does not prevent automated management features such as patching/updating, logging, or authentication.
{{% /pageinfo %}}

Trustgrid customers can specify access to individual organizations and sub-organizations. If support is disabled at the organization, support for all sub-organizations and any native-parent-organization [nodes]({{< ref "docs/concepts/node" >}}) will be disallowed for users who belong to Trustgrid. If enabled on the parent-organization level, but disabled at the sub-organization level, [nodes]({{< ref "docs/concepts/node" >}}) that belong to that sub-organization will not allow Trustgrid staff remote access. This enables a Trustgrid customer to customize support access for each of their customers.

Changes to the supportable flag are audited.
