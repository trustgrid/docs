---
title: Edge Node Behavior When UDP Tunnels are Blocked by Application Inspection
linkTitle: UDP Blocked by Layer 7
---

{{% pageinfo %}}
Application-aware (Layer 7) firewalls can pass the UDP tunnel handshake and then drop the traffic that follows. The node sees a healthy tunnel that carries nothing.
{{% /pageinfo %}}

## Symptoms

- The [Gateway UDP Tunnel Error]({{<relref "docs/alarms/event-types#gateway-udp-tunnel-error" >}}) event repeats. Each `UDP Tunnel has timed out` is followed by a re-established message shortly after.
- The UDP tunnel rebuilds roughly every three minutes, each time from a new source port.
- Traffic to the peer only moves while the node is on TCP fallback. Once the UDP tunnel comes back up, traffic stalls again.
- A port-based rule permitting UDP to the gateway is already in place and shows as allowed in the firewall.

## Cause

The tunnel is encrypted, so an application-aware firewall cannot identify it. Palo Alto and similar firewalls pass the first packets of a flow while classifying it. That window is long enough for the tunnel handshake to complete and the tunnel to report healthy. The firewall then classifies the flow as unidentified UDP and drops everything after.

A port-based allow does not prevent this. The drop comes from application policy, not the port rule. More than one firewall vendor behaves this way.

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
- About two minutes later the tunnel times out. The edge restarts from a new source port and the pattern repeats.

The firewall's traffic log confirms it. Look for the session to the gateway UDP port with an application of `unknown-udp` or similar and a deny action from the application rule rather than the port rule.

## Resolution

- Allow UDP to the gateway addresses on the configured port with a rule that does not depend on application identification. Check your firewall vendor's documentation for how to write that rule. See [Application-Aware (Layer 7) Firewalls]({{<relref "help-center/kb/site-requirements#application-aware-layer-7-firewalls" >}}).
- Until the firewall change is in place, turn off **Enable UDP** in the affected edge node's [Gateway settings]({{<relref "docs/nodes/appliances/gateway#global-settings" >}}). The node stays on TCP tunnels. Turn it back on once the rule is confirmed.
