---
title: November 2026 Minor Appliance Release Notes
linkTitle: November 2026 Minor
type: docs
date: 2026-11-06
description: "Release notes for the November 2026 Minor Trustgrid Appliance release"
---
{{< node-release package-version="TBD" core-version="TBD" release="n-2.26.1" >}}

## Multipath TCP Disabled
Anthropic published a report, [Discovering vulnerabilities with Claude](https://www.anthropic.com/research/discovering-vulnerabilities-with-claude), that describes a remote root exploit against Linux servers with a Multipath TCP (MPTCP) listener. The report tracks it as CVE-2026-98432. Trustgrid does not use MPTCP. This release disables it on the node host and in every network namespace the node manages. The setting is applied at every node start and persists across reboots. Node traffic is not affected. Kernels without MPTCP support, including Ubuntu 18.04 appliances, are unchanged.

The report does not give enough detail to confirm that disabling MPTCP prevents the exploit. Disabling it removes the listener type the report names. It does not patch the kernel bug. Kernel updates remain necessary and continue to arrive through operating system security updates.

The setting is scoped to each network namespace. A workload with enough privilege can re-enable it in its own namespace.

## Drop IP-Option SYNs
The exploit in the report also depends on a TCP SYN packet that carries IPv4 options reaching a kernel TCP listener. The node now watches for these packets on its own listeners. By default it counts them, logs them in the node's kernel log with rate limiting, and lets them through.

A new **Drop IP-option SYNs** checkbox under **Advanced** > **Config Options** drops these packets instead. It is off by default. A change takes effect within a few minutes and does not need a restart. Clearing the checkbox returns the node to counting and logging. The rule only matches SYN packets addressed to a listener on the node itself. Traffic through outside NATs and traffic that transits the node are not affected. The checkbox is available only for nodes on this release or later.

Trustgrid has tested known Trustgrid control and data plane traffic and example application traffic across the Trustgrid connection with this option enabled. We cannot guarantee how other application traffic might be affected. Monitor for issues after enabling it.

See [Advanced Options]({{<relref "docs/nodes/appliances/advanced#config-options" >}}).

## L4 Connector Stalls
An [L4 connector]({{<relref "docs/nodes/shared/connectors" >}}) session to an unresponsive service could stall every other connector on the same node.

On an edge node, each session open was handled on the thread that carries the gateway connection. If the service did not answer, that thread blocked for the full [Connect Timeout]({{<relref "docs/nodes/shared/services" >}}), 5 seconds by default, and every tunnel on that connection stopped moving. A site with a connector to a dead host that retried every 20 seconds was frozen for about half of every cycle. Session opens are now asynchronous, so a slow service no longer blocks the connection.

On a gateway, session opens were queued and were still delivered after the client had disconnected. A burst of connections to a dead service could keep the healthy connectors on the same gateway down for hours after the clients were gone, and only a restart of the active gateway cleared it. The gateway now discards opens whose client has disconnected or timed out before dispatching them. Opens to different destinations are isolated, so a healthy connector does not wait behind a dead one, and each destination has an admission limit. The node log now reports rejected and timed-out opens with counts.

## Fixes
- The latest L4 connector configuration is preserved across a connector stop and restart, and a connector that times out waiting for an ephemeral port no longer leaks the listener.
