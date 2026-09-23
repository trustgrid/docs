---
title: Installed Packages
linkTitle: Installed Packages
description: View the software packages installed on an appliance node.
---
The **Installed Packages** page lists packages installed on an appliance node. In the node's **Advanced** section, select **Installed Packages** to view the list.

The table includes the package name, description, version, and architecture. Use the search field to find packages, or select a column heading to sort the list.

This page requires the `nodes::service:installed-packages` permission. This permission is included in every built-in role that grants `node::read`.

The node must run [release n-2.24.0]({{<relref "release-notes/node/2026-06" >}}) or later, which introduced the installed-packages node service. The corresponding package version for n-2.24.0 is `1.5.20260608-2475`. The page is shown only for nodes that support the installed-packages service.
