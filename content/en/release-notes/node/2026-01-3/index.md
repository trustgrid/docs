---
title: January 2026 Third Minor Appliance Release Notes
linkTitle: January 2026 Third Minor
type: docs
date: 2026-01-16
description: "Release notes for the January 2026 Third Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20260116-2361" core-version="20260116-153629.64ac9d4" release="n-2.23.6" >}}

<br />

It was discovered that upgrading appliances via the [Local Console Utility]({{<relref "tutorials/local-console-utility">}}) caused some systems to fail to boot. This issue resolves that problem and allows appliances to be upgraded successfully using the Local Console Utility.

Additionally, this release changes how OS packages are upgraded so that future upgrades should apply faster. 

Like version n-2.23.4, this release updates the APT repository to a mirror from **December 15, 2025**, which includes packages available on that date. 
