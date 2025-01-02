---
title: January 2025 Minor Appliance Release Notes
linkTitle: January 2025 Minor
type: docs
date: 20242025-01-02
description: "Release notes for the January 2025 Minor Trustgrid Appliance release"
---
{{< node-release package-version="ReplaceMe" core-version="ReplaceMe" release="n-2.21.1" >}}
## Other Improvements and Fixes
- Fixes an issue where connections to the Trustgrid control plane could get stuck in a CLOSE_WAIT state. This could cause the node to show as disconnect despite a working data plane connection. 
- Adds an early release feature to allow clustered Azure-based appliances to manage a floating, local cluster IP address similar to on-premises appliances.