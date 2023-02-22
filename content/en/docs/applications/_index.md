---
tags: ["applications", "ztna"]
title: "Applications"
linkTitle: "Applications"
date: 2023-1-11
weight: 1
---

ZTNA Applications are managed by clicking the `Applications` link on the left nav bar.

![img](apps1.png)
![img](apps2.png)

Trustgrid provides several ZTNA application types that allow remote access to internal resources through a continuously validated connection.

Each application has an [Access Policy]({{<ref "docs/applications/access-policy" >}}) and a [Visibility]({{<ref "docs/applications/visibility" >}}) list. For a user to connect to a ZTNA application, they must (1) be permitted access by the application's Access Policy and (2) be able to see the application.

Access to an application is authenticated through an [Identity Provider]({{<ref "docs/idps" >}}) (IdP).
