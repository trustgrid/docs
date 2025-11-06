---
title: November 2025 Minor Appliance Release Notes
linkTitle: November 2025 Minor
type: docs
date: 2025-11-03
description: "Release notes for the November 2025 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20251105-2319" core-version="20251105-135239.a7de5ae" release="n-2.23.2" >}}


## Cluster Fixes and Improvements

Fixed an issue where, after the configured active cluster member was changed while the original active node was offline, the returning node could incorrectly claim the active role. This occurred because the returning node was unaware of the change and briefly assumed the active role before releasing it, but the newly configured active node did not reclaim the role.

The fix introduces a version counter that increments whenever the configured active member changes. Nodes now compare this version and will not attempt to claim the active role if running an older configuration version. This change requires both members to be updated to this release to be effective.

Additionally, cluster members will now log more clearly when the active role is being claimed or released, providing better visibility into the cluster's state.

Finally, added better handling for the cluster server connections to prevent getting stuck in a close-wait state after network interruptions. This change also allows for modifying the cluster timeout without requiring a restart. 

## VPN Route Sorting Fix
This release changes the method used to sort VPN routes. The previous method could fail when multiple routes had the same destination CIDR and metric leading to a delay in processing updated routes. 

## BGP Fix
Resolves an issue that was causing BGP to fail to export routes on standby cluster members even if the export policy was configured with the `Cluster` setting to `No`.
