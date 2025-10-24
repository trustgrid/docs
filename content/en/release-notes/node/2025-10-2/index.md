---
title: October 2025 Second Minor Appliance Release Notes
linkTitle: October 2025 Minor
type: docs
date: 2025-10-24
description: "Release notes for the October 2025 Second Minor Trustgrid Appliance release"
---

{{< node-release package-version="FIX ME" core-version="FIX ME" release="n-2.23.2" >}}


## Cluster Fixes and Improvements

Fixed an issue where, after the configured active cluster member was changed while the original active node was offline, the returning node could incorrectly claim the active role. This occurred because the returning node was unaware of the change and briefly assumed the active role before releasing it, but the newly configured active node did not reclaim the role.

The fix introduces a version counter that increments whenever the configured active member changes. Nodes now compare this version and will not attempt to claim the active role if running an older configuration version. This change requires both members to be updated to this release to be effective.

Additionally, cluster members will now log more clearly when the active role is being claimed or released, providing better visibility into the cluster's state.

## VPN Route Sorting Fix
This release changes the method used to sort VPN routes. The previous method could fail when multiple routes had the same destination CIDR and metric leading to a delay in processing updated routes. 

## BGP Fix
Resolves an issue that was causing BGP to fail to export routes on standby cluster members even if the export policy was configured with the `Cluster` setting to `No`.
