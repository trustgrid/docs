---
title: "Expose a Container over a Virtual Network"
description: Reach a container running on one Trustgrid node from a peer node over the Trustgrid overlay.
weight: 10
---

A host port mapping makes a container reachable on the **node's local network**. To reach the container from a peer Trustgrid node — for example, a gateway in another datacenter — attach the container to a **virtual network** instead.

This tutorial assumes you have completed the [Container Quickstart]({{<ref "/tutorials/containers/quickstart">}}) and have a running `hello-nginx` container on an edge node, and that you have a gateway node connected to the same organization with a working VPN tunnel to the edge.

## Topology

```
   ┌──────────────────┐       Trustgrid          ┌──────────────────┐
   │  gateway-node    │       virtual            │  edge-node       │
   │                  │       network            │                  │
   │ vIP=10.50.0.1    │◄────────────────────────►│ vIP=10.50.0.2    │
   │                  │       my-vnet            │                  │
   │                  │                          │   ┌────────────┐ │
   │                  │                          │   │ container  │ │
   │                  │                          │   │ vIP=10.50  │ │
   │                  │                          │   │     .0.10  │ │
   │                  │                          │   └────────────┘ │
   └──────────────────┘                          └──────────────────┘
```

The container will get its own address inside the virtual network — `10.50.0.10` in this example. From the gateway, traffic to `10.50.0.10:80` arrives at the edge node, is delivered directly to the container, and is answered without ever touching the edge node's LAN interface.

## 1. Confirm the virtual network is attached to both nodes

The container can only join a virtual network that is already attached to the node it runs on.

1. On the edge node, navigate to **Networking → VPN**. Verify `my-vnet` is listed.
2. On the gateway, do the same. Verify the same `my-vnet` is listed.
3. From the gateway, ping the edge's virtual IP (`10.50.0.2` here) using **Actions → Trustgrid Ping**. If this fails, fix VPN connectivity first — see the [gateway diagnostic tools]({{<ref "/docs/nodes/appliances/gateway/gateway-diag">}}).

## 2. Attach the container to the virtual network

Open the container's **Network** screen at cluster scope (or node scope for a standalone node). Under **Virtual Networks**, click **Add**.

| Field | Value |
| --- | --- |
| **Virtual Network** | `my-vnet` |
| **Virtual IP** | `10.50.0.10` |
| **Allow Outbound** | Enable if the container also needs to make connections out onto the virtual network (e.g. to fetch from another peer). Leave disabled if the container is purely inbound-serving. |

Save. The container will restart to pick up the new interface.

{{<alert color="warning">}}
The container does not need a host port mapping for this to work. The virtual network attachment is independent — you can leave the port mapping for LAN access, or remove it if the container should only be reachable over the overlay.
{{</alert>}}

## 3. Reach the container from the gateway

From any host that can route to the virtual network — typically a test host behind the gateway, or the gateway itself via SSH — make a request to the container's virtual IP on the port it's actually listening on inside the container (port `80` for nginx, not the host-mapped `8080`):

```bash
curl http://10.50.0.10:80/
```

You should see the nginx welcome page. The traffic path:

1. `curl` sends to `10.50.0.10`.
2. The gateway's routing table sends `10.50.0.0/24` out the `my-vnet` tunnel toward the edge.
3. The edge node delivers the packet to the container's `my-vnet` interface.
4. nginx replies; the response is NAT'd back through the tunnel.

## What's different from a host port mapping?

| | Host Port Mapping | Virtual Network |
| --- | --- | --- |
| Listens on | Node's physical NIC | Container's own vnet interface |
| Reachable from | Same node, same LAN | Any node attached to the same virtual network |
| In a cluster | Each member's NIC needs the mapping; reach the active member's IP | Container runs on every member, all attached to the virtual network |
| Address you reach | `<node-IP>:<host-port>` | `<container-vIP>:<container-port>` |
| Common use | Local admin tools, sidecar agents, on-prem services | Microservices, APIs reached by other Trustgrid sites |

You can do both — a typical pattern is a port mapping for local admin access plus a virtual network attachment for application traffic.

## Outbound

If the container needs to reach a service on a peer node (rather than serve traffic from one), enable **Allow Outbound** on the attachment. With it enabled, the container can originate connections to anything reachable on `my-vnet`. With it disabled, the container can only respond to incoming traffic on the virtual IP.

For more granular outbound control, set a **VRF** on the container's **Network** screen. See [Container networking — Outbound traffic]({{<ref "/docs/nodes/appliances/containers/concepts/networking#outbound-traffic--where-does-it-go">}}).
