---
title: January 2025 Minor Appliance Release Notes
linkTitle: January 2025 Minor
type: docs
date: 2025-01-02
description: "Release notes for the January 2025 Minor Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20250109-2213" core-version="20250109-221438.4c724ae" release="n-2.21.1" >}}
## Other Improvements and Fixes
- Fixes an issue where connections to the Trustgrid control plane could get stuck in a CLOSE_WAIT state. This could cause the node to show as disconnect despite a working data plane connection. 
- Adds an early release feature to allow clustered Azure-based appliances to manage a floating, local cluster IP address similar to on-premises appliances. Currently there is no UI support for this feature.  Contact Trustgrid support for more information.
- Update the APT security packages to those available on 2024-12-02.