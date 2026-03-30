---
title: March 2026 Minor Appliance Release Notes
linkTitle: March 2026 Minor
type: docs
date: 2026-03-30
description: "Release notes for the March 2026 Minor Trustgrid Appliance release"
---

{{< node-release package-version="1.5.20260116-2361" core-version="20260116-153629.64ac9d4" release="n-2.23.6" >}}

<br />

 - **Improved UDP tunnel recovery after AWS Network Firewall maintenance** - During AWS Network Firewall maintenance events, established tunnel state can be lost, causing the firewall to
  unintentionally block the existing UDP tunnel's source port. Trustgrid now better detects this type of disruption and automatically reconnects using a new source port, eliminating
  orphaned tunnels that previously required manual intervention. Additionally, when nodes are deployed behind an [AWS Network Firewall and UDP tunnels are used, it is required to add explicit bidirectional rules]({{<relref "/tutorials/deployments/deploy-aws-ami/#aws-network-firewall-and-udp-tunnels">}}) to allow UDP traffic on the expected IPs and ports. 
 - **Fixed UDP rekey interval calculation** — The gateway now correctly calculates rekey timing based on the last rekey event rather than session creation time, preventing unnecessary rekey
  churn.
