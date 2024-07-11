---
title: June 2024 Minor Appliance Release Notes
linkTitle: June 2024 Minor
type: docs
date: 2024-06-01
description: "Release notes for the June 2024 Minor Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20240514-2064" core-version="20240514-213500.a967cd2" release="n-2.19.2" >}}
## Issues Resolved
- Flow logs related to [port forwards]({{<relref "/docs/nodes/appliances/vpn/port-forwarding">}}) were not showing on the receiving device
- An issue was preventing some security packages from being updated
- Updated the AWS library being utilized in the AWS dependency package