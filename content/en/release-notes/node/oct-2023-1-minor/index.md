---
title: "October 2023 Appliance Second Release Notes"
linkTitle: "October 2023 Second Release"
type: docs
date: 2023-10-19
---
{{< node-release package-version="1.5.20231017-1855" core-version="20231017-165331.5772aac" release="n-2.17.1" >}}

## Resolved Issues
- Fixed an issue where the appliance would stop responding to SNMP traffic on the virtual management IP if [virtual network port forwards]({{< relref "/docs/domain/virtual-networks/remote-port-forward" >}}) were configured.