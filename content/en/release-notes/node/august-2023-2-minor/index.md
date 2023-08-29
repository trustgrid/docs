---
title: "August 2023 Second Appliance Release Notes"
linkTitle: "August 2023 Second Release"
type: docs
date: 2023-08-29
---
{{< node-release package-version="1.5.20230824-1797" core-version="20230824-150008.3c2d422" release="n-2.16.2" >}}

## Update Fix
This minor release resolves several issues:
- Fixed an issue where client devices could experience increased CPU utilization if their peer gateways were offline or unreachable.
- Resolved issues with ARP on VLAN interface cluster IPs that could cause both active and standby members to respond to ARP requests. This prevented proper failover behavior.
- Resolved SNMP issues for VMware VM8 nodes that didn't have all 8 virtual network interfaces attached to the virtual machine.