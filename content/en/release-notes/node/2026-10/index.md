---
title: October 2026 Major Appliance Release Notes
linkTitle: October 2026 Major
type: docs
date: 2026-10-19
description: "Release notes for the October 2026 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20261015-2565" core-version="20261015-232909.d05f29b" release="n-2.26.0" >}}

## UDP Tunnel Health and Backoff
A UDP tunnel is now tested before any traffic uses it. Some firewalls allow a UDP flow to establish and then start dropping it once application inspection classifies the traffic, and a low path MTU on the route is not always obvious. Previously the node activated a tunnel as soon as the first keep alive came back, which both of those conditions pass.

When a new tunnel comes up, the node sends probe packets at the tunnel's full packet size, 1420 bytes by default, one second apart. The peer echoes each one. The tunnel has to return three consecutive probes to activate, within a budget of six attempts, so a single dropped packet is tolerated but a path that only passes the first few packets is not. Traffic stays on the [TLS tunnel]({{<relref "docs/nodes/appliances/gateway#tls-tunnels" >}}) for the few seconds this takes, and stays there permanently if the tunnel fails. The check runs only when a tunnel is built on a new source port. Routine rekeys skip it.

A failed tunnel now backs off instead of retrying every few seconds. Retries go 30 seconds, 1 minute, 2 minutes, 5 minutes, then every 10 minutes, each with a little jitter so a site full of nodes does not retry in lockstep. The interval resets only after the tunnel has been healthy for 30 minutes.

A tunnel that keeps coming and going is marked as flapping, after three healthy-to-lost cycles or five straight failures within an hour. While flapping, the node keeps retrying every 10 minutes but suppresses repeat alerts to one per day, and the event says when it will report again. Flapping clears after an hour of health. A firewall that silently drops UDP no longer produces a continuous stream of alerts.

A recovered UDP tunnel now clears its [Gateway UDP Tunnel Error]({{<relref "docs/alarms/event-types#gateway-udp-tunnel-error" >}}). Earlier releases raised the error but never sent the resolution, so the alert stayed open until it was cleared by hand. Establishing a tunnel is also logged, so you can tell whether UDP was working before an incident.

Both ends need this release for the tunnel test to run. Against an older peer the node activates the tunnel immediately as it did before.

See [UDP Tunnels]({{<relref "docs/nodes/appliances/gateway#udp-tunnels" >}}).

## Control Plane Stability
A node could get stuck in a reconnect loop with the control plane, dropping and reconnecting roughly every three minutes and raising an urgent node disconnect alert each time. Traffic kept flowing, but the node was effectively unusable for management and the alerts did not stop until it was restarted. Any brief pause on the node could start it, including a VM snapshot or a live migration. This is fixed, and a pause now costs one reconnect.

## Audit Records
An audit entry is now written when a node service is requested rather than when it finishes. A request that hung, such as a cluster server restart that never completed, previously left no record in the portal that anyone had asked for it.

See [Audits]({{<relref "docs/nodes/shared/audits" >}}).

## L4 Connectors
Editing [L4 connectors]({{<relref "docs/nodes/shared/connectors" >}}) could leave ports unbound. A connector is restarted when its config changes, and the restart could fail to reclaim its listening port. The connector then stayed down until a later config change touched it or the node was restarted, and nothing alerted. Sites that renumber many connectors in one change were most exposed, since every connector in the change rolled the same odds.

The node now waits for the port to be released before rebinding, and retries if it is still in use.

A connector that still cannot bind raises a new [L4 Connector Bind]({{<relref "docs/alarms/event-types#l4-connector-bind" >}}) event, with the port, listen address, and protocol in the event details. It clears when the connector binds. A connector that is down is now visible instead of silent.

## Security Updates
This release updates the APT security repository to a mirror from **October 12, 2026**, giving appliances the latest security patches for underlying system packages.

## Fixes
- Fixed a memory leak on gateways serving WireGuard tunnels. A gateway with heavy roaming agent activity could exhaust host memory over several months.
- Fixes to the portal's **Try** option for WAN changes. A WAN change that restarts the node now raises a Node Restart alert before it restarts, and the node reconnects to the control plane immediately after a Try instead of waiting. A custom WAN MTU now survives a Try followed by Save, an MTU set locally through `tg-console` and marked permanent is no longer overwritten by the portal, and an unrelated portal change such as a tag edit no longer saves network settings you have not confirmed.
