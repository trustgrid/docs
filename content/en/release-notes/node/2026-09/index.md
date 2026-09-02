---
title: September 2026 Major Appliance Release Notes
linkTitle: September 2026 Major
type: docs
date: 2026-09-02
description: "Release notes for the September 2026 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20260827-2540" core-version="20260827-205421.cfa4c19" release="n-2.25.0" >}}

## TLS 1.3 and Post-Quantum Cryptography
Nodes now negotiate TLS 1.3 for the TCP data plane, falling back to TLS 1.2 when the peer does not support it. TLS 1.3 connections use the `TLS_AES_128_GCM_SHA256` cipher. TLS 1.2 connections continue to use `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`.

TLS 1.3 connections use `X25519MLKEM768` hybrid key exchange, which protects traffic captured today from being decrypted later by a quantum computer. Peers that do not support it fall back to a classical key exchange.

The [UDP data plane]({{<relref "getting-started/security#udp-data-plane-encryption" >}}) is unchanged.

A node can be pinned back to TLS 1.2 if something in the path cannot handle TLS 1.3. Contact support if you need this.

See [Security]({{<relref "getting-started/security#tls-encryption" >}}).

## Clustering and Failover
- On Azure, an unresponsive instance metadata endpoint could leave Azure API calls retrying indefinitely on the same threads that carry cluster traffic. Members lost sight of each other and could both go active, which effectively blocks traffic. A **Restart Cluster Server** request would then hang with the server stopped, and only a node reboot recovered it. Azure API calls now run off the cluster threads under enforced timeouts.
- Restarting one member of an Azure cluster no longer causes the other member to flap in and out of the master role.
- A brief loss of connectivity between cluster peers could leave a member with two connections to that peer, tracking the wrong one. Its view of which member was active then diverged from reality, leaving the cluster dual-active. A node now refuses a duplicate connection to a peer it already has an open session with, and closes a connection attempt that times out.
- A gateway or cluster server holding a dead client connection would reject that client's reconnect with `Already have max number of client connections`, leaving the node unable to reconnect until the old connection timed out. If it never closed, the node stayed locked out. The server now closes the old session and accepts the reconnect.

## Gateway Latency Monitors
[Gateway latency monitors]({{<relref "docs/nodes/appliances/gateway/gateway-client#gateway-latency-monitors" >}}) now separate reporting from health policy. A [Gateway Latency Exceeded]({{<relref "docs/alarms/event-types#gateway-latency-exceeded" >}}) event fires on every monitor trigger and recovery, whatever **Failure Mode** and **Critical** are set to. **Failure Mode** now decides only whether the node marks itself unhealthy, and counts only monitors marked **Critical**.

If you use **Any** or **All** with no monitor marked **Critical**, the node will no longer mark itself unhealthy on a latency trigger. Enable **Critical** on the monitors that should drive a failover.

Also fixed: in **All** mode the node recovered as soon as a single monitor cleared, and a standby cluster member could send a cleared event with no matching exceeded event.

## Security Updates
This release updates the APT security repository to a mirror from **July 20, 2026**, giving appliances the latest security patches for underlying system packages.

- Bundled third-party libraries were updated to versions with no open security findings. The earlier findings were not exploitable in the way the node uses those libraries, but the updates keep the appliance clean for vulnerability scanners.
- Unattended upgrades no longer treat a "reboot required" result from the package manager as a failure. The node now reboots to finish the upgrade instead of logging an error.

## Improvements
- The **Node Startup Errors** alert now carries the errors themselves in its detail, so you can see what failed without opening the node.

## Fixes
- Nodes can now pull container images published in the OCI image index format. Previously a registry that returned an OCI manifest could fail the download.
- Removed a spurious error and stack trace logged when both ends of an L4 proxy session closed it at the same moment.
