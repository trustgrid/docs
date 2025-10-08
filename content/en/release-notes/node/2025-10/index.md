---
title: October 2025 Minor Appliance Release Notes
linkTitle: October 2025 Minor
type: docs
date: 2025-10-08    
description: "Release notes for the October 2025 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20251006-2299" core-version="20251006-145121.01eb7c5" release="n-2.23.1" >}}


## Resolve Inaccurate UDP Events and Debug Logs
This release will fix an issue that caused some nodes with UDP enabled to generate a large number of events indicating:
- `UDP Tunnel has timed out for node <nodename> and endpoint <endpoint>`
- `UDP Tunnel connection has been re-established for node <nodename> and endpoint <endpoint>`


These events were inaccurate and could be safely ignored. **UDP tunnel traffic was NOT disrupted.** However, they could clutter the event log and make it harder to identify real issues. This release will also reduce the verbosity of debug logs related to UDP tunnels, making it easier to troubleshoot other issues.

## Disables Ubuntu Services Causing Unexpected Network Traffic
Several Ubuntu services that are not required for Trustgrid appliances have been found to generate unexpected network traffic. Specifically traffic to the IP addresses related to these DNS names:
- `cdn.fwupd.org`
- `archive.ubuntu.com`
- `motd.ubuntu.com`

This release will disable the services in question and prevent this traffic from occurring. 

As a minor release, this update focuses on addressing these specific issues without introducing new features or significant changes. This includes no changes to the Ubuntu security packages. 