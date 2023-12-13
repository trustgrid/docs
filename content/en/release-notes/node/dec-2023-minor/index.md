---
title: "December 2023 Appliance Minor Release Notes"
linkTitle: "December 2023 Minor Release"
type: docs
date: 2023-12-13
---
{{< node-release package-version="1.5.20231206-1914" core-version="20231206-183238.65d2bde" release="n-2.17.3" >}}

## Disable UDP Rekeying
 A customer had a misconfiguration that led to the UDP rekeying messages being lost. This compounded the performance impact as the UDP tunnels had to be rebuilt frequently. 

 This release adds the ability for Trustgrid support to disable UDP rekeying on a gateway if needed. Trustgrid will make the UDP rekeying process more resilient and better able to handle lost key exchanges in the future.