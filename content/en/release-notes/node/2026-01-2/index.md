---
title: January 2026 Second Minor Appliance Release Notes
linkTitle: January 2026 Second Minor
type: docs
date: 2026-01-08
description: "Release notes for the January 2026 Second Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20260108-2356" core-version="20260108-162524.eeaec95" release="n-2.23.5" >}}

<br />

This release resolves an issue seen when upgrading appliances using the NCA-1210 and Dell R250 platforms running Ubuntu 22.04 to [version n-2.23.4]({{<relref "release-notes/node/2026-01">}}), which caused the upgrade of a particular OS package to fail due to an artifact left over from the imaging process. Appliances using these platforms can now successfully upgrade to this release.

Like version n-2.23.4, this release updates the APT repository to a mirror from **December 15, 2025**, which includes packages available on that date. 
