---
title: November 2025 Second Minor Appliance Release Notes
linkTitle: November 2025 Minor
type: docs
date: 2025-11-21
description: "Release notes for the November 2025 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20251118-2333" core-version="20251118-232513.ebc829d" release="n-2.23.3" >}}

This is a minor release that fixes two issues:
- There was a situation where a bad certificate for the control plane could cause a device to stop attempting to connect.
- Simultaneous disconnections of the cluster heartbeat connections could cause a device to enter a deadlock state where it would not attempt to reconnect to it's cluster peers.
