---
title: January 2026 Minor Appliance Release Notes
linkTitle: January 2026 Minor
type: docs
date: 2026-01-05
description: "Release notes for the January 2026 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20251216-2350" core-version="20251216-211950.61a1833" release="n-2.23.4" >}}

<br />

{{<alert color="warning">}}Shortly after publishing we discovered that appliances using the NCA-1210 and Dell R250 platforms running Ubuntu 22.04 hit a significant package conflict that caused the upgrade to fail. Because of the failure mode, this release has been effectively revoked until the conflict can be resolved; please do not upgrade devices at this time.{{</alert>}}

This release updates the APT repository to a mirror from **December 15, 2025**, which includes packages available on that date. There are no changes to Trustgrid software in this release.

This update ensures appliances have access to the latest security patches for underlying system packages without waiting for a major release.