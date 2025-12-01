---
title: December 2025 Minor Appliance Release Notes
linkTitle: December 2025 Minor
type: docs
date: 2025-12-01
description: "Release notes for the December 2025 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20251124-2340" core-version="20251124-234159.0f29e39" release="n-2.23.3" >}}

This is a minor release that fixes the following issues:
- There was a situation where a bad certificate for the control plane could cause a device to stop attempting to connect.
- Simultaneous disconnections of the cluster heartbeat connections could cause a device to enter a deadlock state where it would not attempt to reconnect to its cluster peers.
- Resolves an issue where non-cluster appliance nodes would log an error about looking for a cluster IP address.
