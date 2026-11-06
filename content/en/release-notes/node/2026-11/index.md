---
title: November 2026 Minor Appliance Release Notes
linkTitle: November 2026 Minor
type: docs
date: 2026-11-06
description: "Release notes for the November 2026 Minor Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20261106-2578" core-version="20261106-145459.7b32679" release="n-2.26.1" >}}

## L4 Connector Stalls
An [L4 connector]({{<relref "docs/nodes/shared/connectors" >}}) pointed at a service that does not respond could stall every other connector on the same node. This release fixes two causes.

On an edge node, each new connection waited on the shared gateway link for the full [Connect Timeout]({{<relref "docs/nodes/shared/services" >}}), 5 seconds by default. While it waited, all traffic over that link stopped. A connector that retried a dead host every 20 seconds froze the node for about half of every cycle. New connections no longer hold up the link.

On a gateway, connection requests stayed queued after the client gave up. A burst of connections to a dead service could keep the healthy connectors on that gateway down for hours after the clients were gone, and only a gateway restart cleared it. The gateway now discards requests whose client has disconnected or timed out, and keeps each destination in its own queue so a dead service does not delay a healthy one. The node log reports how many requests were rejected or timed out.

Other L4 fixes:

- L4 connector settings are kept across a connector stop and restart, and a connector that times out waiting for a port no longer leaves the port in use.

## CVE-2026-98432 Mitigations
Anthropic published a report, [Discovering vulnerabilities with Claude](https://www.anthropic.com/research/discovering-vulnerabilities-with-claude), that describes a remote root exploit against Linux servers. It is tracked as CVE-2026-98432. The report names two conditions the exploit relies on: a Multipath TCP (MPTCP) listener, and a TCP SYN packet that carries IP options. This release addresses both. The report does not give enough detail to confirm that either change prevents the exploit. The real fix is a kernel update. As soon as Ubuntu publishes the kernel security packages, Trustgrid will fast track a release to make them available.

### Multipath TCP Disabled
Trustgrid does not use MPTCP. This release turns it off on the node and in every network namespace the node manages. The setting is applied at every node start and survives reboots. Node traffic is not affected. Ubuntu 18.04 appliances have no MPTCP support and are unchanged.

### Drop IP-Option SYNs
The node now watches for incoming TCP SYN packets that carry IP options and are addressed to a service on the node itself. By default it counts and logs them and lets them through.

A new **Drop IP-option SYNs** option under **Advanced** > **Config Options** drops them instead. It is off by default. A change takes effect within a few minutes and does not need a restart. Only packets addressed to the node itself are checked. Traffic through outside NATs and traffic that passes through the node are not affected.

Trustgrid has tested known Trustgrid control and data plane traffic and example application traffic across the Trustgrid connection with this option enabled. We cannot guarantee how other application traffic might be affected. Monitor for issues after enabling it.

See [Advanced Options]({{<relref "docs/nodes/appliances/advanced#config-options" >}}).

## Security Updates
This release updates the APT security repository to a mirror from **October 26, 2026**, giving appliances the latest security patches for underlying system packages.

- The mirror includes the updated `ca-certificates` package that removes the root certificates cross-signing the compromised eMudhra certificate authority. See [Ubuntu's eMudhra CA compromise notice](https://ubuntu.com/blog/emudhra-ca-compromise-fixes-available).
