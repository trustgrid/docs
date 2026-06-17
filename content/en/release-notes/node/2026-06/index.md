---
title: June 2026 Major Appliance Release Notes
linkTitle: June 2026 Major
type: docs
date: 2026-06-08
description: "Release notes for the June 2026 Major Trustgrid Appliance release"
---
{{< node-release package-version="1.5.20260602-2472" core-version="20260606-003838.226.3724894" release="n-2.24.0" >}}

## Hop Monitoring
This release reworks how [Hop Monitoring]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers">}}) probes the network path. On some networks the previous behavior was interpreted as abnormal traffic and triggered `Peer Timed Out` events that briefly disrupted the data plane.

The node now behaves more like MTR. It no longer sends resets to intermediate hops, completes the TCP handshake when a hop replies with a SYN/ACK, slows the probe rate as it works further down the path, and stops probing once it reaches the gateway. Together these changes make hop monitoring far less likely to be flagged by firewalls in the path.

A new **Monitor Hops Protocol** setting in the [Client]({{<relref "/docs/nodes/appliances/gateway/gateway-client#hop-monitoring-settings">}}) panel lets you choose how the node probes the path. **SYN** uses TCP SYN packets as before, **ICMP** requires both ends to allow the traffic and makes a low path MTU easier to detect, and **SACK** builds a full TCP session and sends 1-byte packets within it until it gets a response, for paths where network devices mishandle bare SYN traces. Because the node no longer resets intermediate hops, the **Support Monitor Hops Resets** setting is deprecated and has no effect.

The `gateway latency exceeded` event now includes the destination gateway name, so you can tell which path is slow without cross-referencing other data.

## Clustering and Failover
- AWS clusters now support a floating [Cluster IP]({{<relref "/tutorials/deployments/deploy-aws-ami/ip-failover">}}), assigned as a secondary private IP on the ENI, bringing them to parity with Azure. The previous Elastic-IP-based cluster IP has been removed.
- Azure clusters could get permanently stuck attaching the floating cluster IP after a master change when Azure returned `PrivateIPAddressIsBeingCleanedUp`. The attach logic now backs off, retries the correct interface, and recovers on its own.
- A race between a preferred-active change and a cluster server restart could leave both members active after a failback, causing split-brain and dropped connections. This is fixed, and cluster server restarts now write an [audit entry]({{<relref "/docs/nodes/shared/audits">}}) recording the time, node, and reason.
- Cluster logic now handles two simultaneous unhealthy conditions (for example an interface down at the same time as a control and data plane outage) without ending in an unresolved dual-active state.
- A node now throws an event if it cannot reach its cluster peer on the gossip port.
- A node refreshes its node info when transitioning from unhealthy back to healthy, avoiding stale data that could confuse cluster decisions.

## Gateway Latency Monitors
[Gateway latency monitors]({{<relref "/docs/nodes/appliances/gateway/gateway-client#gateway-latency-monitors">}}) are now available in the portal under **Gateway > Client**. A monitor watches the round trip time of a node's tunnel to a gateway and reacts when it crosses a threshold you set. For a cluster member, a triggered monitor can mark the node unhealthy and fail traffic over to a healthier member. On a standalone node it raises a [Gateway Latency Exceeded]({{<relref "/docs/alarms/event-types#gateway-latency-exceeded">}}) event while traffic stays in place.

## Node Services for Automated Operations
A number of node services were reworked to return structured, actionable results rather than a simple pass or fail. When these services are driven through the Trustgrid API, monitoring tooling and AI agents can read the specific reason a check failed and respond to it directly, instead of a person reading through logs.

- The DNS health service reports whether DNS is misconfigured, a configured server is unreachable, or a server fails to resolve a Trustgrid hostname, and reports the status of the primary and secondary servers individually.
- The repository connectivity service returns the specific reason it is unhealthy, such as `repo.trustgrid.io` not resolving, the repository being unreachable on port 8443, or an unrecognized client certificate.
- A node can capture filtered `tcpdump` output and return it as trimmed text sized for automated analysis.
- The outbound TCP port check can return its result as a successful event with the connection details, rather than erroring on a failed connection.
- Additional node tools, including the remote and local TCP port tests, can now be invoked as events.
- New services retrieve node startup error logs and report the list of installed packages to the portal.
- The gateway routes service reports per-route bytes sent and received, and node stats report whether gateway client or server traffic shaping is currently active.

## Platform Support
- Added the gen3 AWS platform, which supports most current-generation x86_64 EC2 instance types, including `c7i`, `c7a`, `m8i`, and `m8a`. Nodes on this platform appear in the portal with the `AWS Gen3` device type. See the [supported instance types]({{<relref "/tutorials/deployments/deploy-aws-ami#instance-type">}}).
- Added a KVM node adapter for deployments that pin NICs to consistent names.
- On GCP, interfaces with no configured MTU now default to 1460 rather than 1500.
- Removed legacy KVM functionality from the node, excluding the Datastore.

## Security Updates
This release updates the APT security repository to a mirror from **April 20, 2026**, giving appliances the latest security patches for underlying system packages.

- Security patches from the Ubuntu ESM repositories were sometimes skipped by unattended upgrades. The node now runs an additional `apt update` after adding the ESM repositories and matches repositories by label rather than origin, so those patches install correctly.
- Unattended upgrade timing was tuned to reduce failed or interrupted upgrades: a longer kill timeout, a short delay between retries, and fewer retries.
- Upgrading a node from the local console no longer hangs at an interactive apt prompt.

## Improvements
- Added a Packet Simulator tool. Enter a source IP, destination IP, and port, and the node returns the NAT rules and routing that would be applied. It currently covers VPN traffic only, not the L4 Proxy. The tool is available through the [MCP servers]({{<relref "/docs/mcp">}}) today, with access from the portal coming soon.
- L4 [services]({{<relref "/docs/nodes/shared/services">}}) can now source traffic as the cluster IP instead of the interface IP. The interface IP remains the default.
- L4 services and [Port Forwards]({{<relref "/docs/nodes/appliances/vpn/port-forwarding">}}) now send an event when the node runs out of ephemeral source ports, matching the existing warning for VPN traffic.
- The node supports additional sub-interface IP configurations, including sub-interfaces defined at the node level on a clustered node.
- Added the ability to manually set the local name resolver's host cache, for cases where DNS keeps failing back to bad entries.
- Added `hping3` to the console network tools shell for control plane troubleshooting.
- The `ping` tools now always return output, even when there is no reply.
- After a disk resize, the node re-checks its disk size so the Overview graphs reflect the new value.

## Fixes
- Updating the metric of a dynamic route now updates the metric in the receiving node's route table. Previously the change required removing and re-adding the route.
- VLAN sub-interfaces again appear in Interface Traffic Graphs and SNMP. This had regressed in an earlier release.
- VRF forwarding was dropping packets even when forwarding was enabled on the VRF. This is fixed.
- Fixed a null pointer exception when a tunnel interface used a masquerade NAT forwarding out a virtual network tunnel interface.
- An ICMP [route monitor]({{<relref "/docs/nodes/appliances/vpn/static-routing#route-monitors">}}) no longer falsely reports its latency threshold exceeded when the actual latency is well below the limit.
- A container's previous image is no longer reported as in use after the container is updated to a new image tag, so the old image can be removed without restarting the node.
- A timeout on a configured second [gateway path]({{<relref "/docs/nodes/appliances/gateway/gateway-client#gateway-paths">}}) no longer disconnects both tunnels to that gateway.
- Setting an interface to ignore health checks now takes effect without a restart.
- Fixed a memory leak in the VPN websocket handler.
- Opening the Resize Root Partition tab no longer creates an [audit]({{<relref "/docs/nodes/shared/audits">}}) entry on its own.
- Reduced log and event noise across several areas, including `All Gateways Disconnected` events during WAN IP changes, repeated max-connection events, BGP attribute warnings, and DNS cache refresh logging.
