---
title: "Expose a Container over a Virtual Network"
description: Reach a container running on one Trustgrid node from a peer node over the Trustgrid overlay.
---

A host port mapping makes a container reachable on the appliance's **local network**. To reach the container from a peer Trustgrid node — for example, a gateway in another datacenter — give the container a **virtual IP** on a virtual network instead.

This tutorial assumes you have completed the [Container Quickstart]({{<relref "/tutorials/containers/quickstart">}}) and have a running `nginx` container on an edge node, and a gateway node connected to the same organization with a working VPN tunnel to the edge.

## Topology

```
   ┌──────────────┐
   │  test host   │
   │ 10.60.0.42   │ (routes 10.50.0.0/24 via the gateway)
   └──────┬───────┘
          │  curl 10.50.0.10:80
          ▼
   ┌─────────────────────┐                            ┌─────────────────────────────────┐
   │       gateway       │   Trustgrid VPN tunnel     │              edge               │
   │                     │ ─────────────────────────► │                                 │
   │                     │   my-vnet (10.50.0.0/24)   │   reaches container at          │
   │                     │                            │   10.50.0.10:80                 │
   │                     │                            │                                 │
   │                     │                            │   ┌──────────────────────────┐  │
   │                     │                            │   │  Container nginx         │  │
   │                     │                            │   │    listening on :80      │  │
   │                     │                            │   └──────────────────────────┘  │
   └─────────────────────┘                            └─────────────────────────────────┘
```

The container is reachable on the virtual network at `10.50.0.10` — the IP you assign on the container's **Network** screen. From the gateway, traffic to `10.50.0.10:80` arrives at the edge and is delivered to the container.

## 1. Confirm the virtual network is attached to both nodes

The container can only be exposed on a virtual network that is already attached to the appliance running it.

1. On the edge node, navigate to **Networking → VPN**. Verify `my-vnet` is listed.
2. On the gateway, do the same. Verify the same `my-vnet` is listed.

## 2. Attach the container to the virtual network

Open the container's **Network** screen at cluster scope (or node scope for a standalone node). Under **Virtual Networks**, click **Add**.

| Field | Value |
| --- | --- |
| **Virtual Network** | `my-vnet` |
| **Virtual IP** | `10.50.0.10` |
| **Allow Outbound** | Enable if the container also needs to make connections out onto the virtual network (e.g. to fetch from another peer). Leave disabled if the container is purely inbound-serving. |

Save. The container is reachable at `10.50.0.10` as soon as the config update lands on the appliance.

{{<alert color="warning">}}
The container does not need a host port mapping for this to work. The virtual network attachment is independent — you can leave the port mapping for LAN access, or remove it if the container should only be reachable over the overlay.
{{</alert>}}

## 3. Reach the container from the gateway

From any host that can route to the virtual network — typically a test host behind the gateway — make a request to the container's virtual IP on the port the container is actually listening on inside the container (`80` for the default nginx image):

```bash
curl http://10.50.0.10:80/
```

You should see the nginx welcome page. The traffic path:

1. `curl` sends to `10.50.0.10`.
2. The gateway's routing table sends `10.50.0.0/24` out the `my-vnet` tunnel toward the edge.
3. The edge appliance delivers the traffic to the container.
4. nginx replies on the same path.

## What's different from a host port mapping?

| | Host Port Mapping | Virtual Network |
| --- | --- | --- |
| Reachable from | Same node, same LAN | Any node attached to the same virtual network |
| In a cluster | Each member's NIC needs the mapping; reach the active member's IP | Container runs on every member, all attached to the virtual network |
| Address you reach | `<node-IP>:<host-port>` | `<container-VIP>:<container-port>` |
| Common use | Local admin tools, sidecar agents, on-prem services | Microservices, APIs reached by other Trustgrid sites |

You can do both — a typical pattern is a port mapping for local admin access plus a virtual network attachment for application traffic.

## Outbound

If the container needs to reach a service on a peer node (rather than serve traffic from one), enable **Allow Outbound** on the attachment. With it enabled, the container can originate connections to anything reachable on `my-vnet`. With it disabled, the container can only respond to incoming traffic on the virtual IP.

For more granular outbound control, set a **VRF** on the container's **Network** screen. See [Container networking — Outbound traffic]({{<relref "/docs/nodes/appliances/containers/concepts/networking#outbound-traffic--where-does-it-go">}}).
