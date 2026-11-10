---
title: Edge Node Behavior When UDP Tunnels are Blocked by Application Inspection
linkTitle: UDP Blocked by Layer 7
---

{{% pageinfo %}}
Application-aware (Layer 7) firewalls can pass the UDP tunnel handshake and then drop the traffic that follows. The node sees a healthy tunnel that carries nothing.
{{% /pageinfo %}}

{{% alert title="Behavior changed in the October 2026 release" color="info" %}}
Nodes running [n-2.26.0]({{<relref "release-notes/node/2026-10" >}}) or later test a UDP tunnel before using it, so a tunnel blocked this way never carries traffic in the first place. Both the node and the gateway must be on that release. The symptoms below differ depending on which release the nodes are running.
{{% /alert %}}

## Symptoms

A port-based rule permitting UDP to the gateway is already in place and shows as allowed in the firewall.

### On n-2.26.0 and Later

- Traffic to the peer keeps flowing over the TLS tunnel and never moves to UDP.
- The UDP tunnel is retried on a widening interval rather than every few minutes, out to once every 10 minutes.
- The [Gateway UDP Tunnel Error]({{<relref "docs/alarms/event-types#gateway-udp-tunnel-error" >}}) event reports the tunnel as flapping and then reports once a day instead of on every attempt.
- Node logs show the UDP health check failing for the gateway.

Connectivity is not affected, because the node never moves traffic onto the blocked tunnel. The firewall still needs fixing to get the performance benefit of UDP, but the alert noise and the stalls are gone.

### Before n-2.26.0

- The [Gateway UDP Tunnel Error]({{<relref "docs/alarms/event-types#gateway-udp-tunnel-error" >}}) event repeats. Each `UDP Tunnel has timed out` is followed by a re-established message shortly after.
- The UDP tunnel rebuilds roughly every three minutes, each time from a new source port.
- Traffic to the peer only moves while the node is on the TLS tunnel. Once the UDP tunnel comes back up, traffic stalls again.

## Cause

The tunnel is encrypted, so an application-aware firewall cannot identify it. Palo Alto and similar firewalls pass the first packets of a flow while classifying it. That window is long enough for the tunnel handshake to complete and the tunnel to report healthy. The firewall then classifies the flow as unidentified UDP and drops everything after.

A port-based allow does not prevent this. The drop comes from application policy, not the port rule. More than one firewall vendor behaves this way.

What changed in n-2.26.0 is how long the node believes the tunnel. Earlier releases activated a tunnel as soon as the first keep alive came back, which is inside the window the firewall passes. The node now sends probe packets at the tunnel's full size for several seconds and requires three consecutive replies before it moves traffic, which is long enough for the firewall to classify the flow and start dropping. The tunnel fails the check and traffic stays on the TLS tunnel. See [UDP Tunnels]({{<relref "docs/nodes/appliances/gateway#udp-tunnels" >}}).

## Verification

Capture at both ends at the same time using the [Sniff Traffic]({{<relref "tutorials/interface-tools/sniff-interface-traffic" >}}) tool on each node's WAN interface. Substitute the gateway's public IP and configured [UDP port]({{<relref "docs/nodes/appliances/gateway/gateway-server#udp-port" >}}) and the edge node's public IP.

Edge node filter:

```
udp port <gateway-udp-port> and host <gateway-public-ip>
```

Gateway node filter:

```
udp port <gateway-udp-port> and host <edge-public-ip>
```

When application inspection is the problem, the captures show:

- Both sides see the first packets in each direction, and the tunnel comes up.
- Within seconds, the edge keeps sending on the same source port but the gateway stops seeing those packets, or the gateway's replies stop reaching the edge.

On n-2.26.0 and later, the edge then stops sending on that source port within a few seconds, and the next attempt comes after the backoff interval rather than immediately. Before n-2.26.0, the tunnel times out about two minutes later and the edge restarts from a new source port right away, repeating the pattern.

The firewall's traffic log confirms it. Look for the session to the gateway UDP port with an application of `unknown-udp` or similar and a deny action from the application rule rather than the port rule.

## Resolution

- Allow UDP to the gateway addresses on the configured port with a rule that does not depend on application identification. Check your firewall vendor's documentation for how to write that rule. See [Application-Aware (Layer 7) Firewalls]({{<relref "help-center/kb/site-requirements#application-aware-layer-7-firewalls" >}}).
- On releases before n-2.26.0, turn off **Enable UDP** in the affected edge node's [Gateway settings]({{<relref "docs/nodes/appliances/gateway#global-settings" >}}) until the firewall change is in place. The node stays on TLS tunnels. Turn it back on once the rule is confirmed. On n-2.26.0 and later the node already keeps traffic on the TLS tunnel, so disabling UDP only stops the retries.

{{% alert title="Re-enabling UDP needs a reconnect" color="warning" %}}
On n-2.26.0, turning **Enable UDP** back on does not rebuild the tunnels on its own. The node logs that UDP is enabled and then stays on the TLS tunnel indefinitely. Restart the node, or wait for its control plane connection to reconnect, to get UDP back.
{{% /alert %}}
