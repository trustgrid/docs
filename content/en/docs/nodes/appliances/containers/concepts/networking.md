---
title: "Container Networking"
description: How containers reach the network and how the network reaches containers on a Trustgrid node.
weight: 10
---

A container running on a Trustgrid node sits behind a node-managed bridge. This page explains the bridge, the DNS resolver, the three ways traffic can enter the container, and how outbound traffic leaves.

## The container bridge

Every container is attached to a node-managed bridge network in the **`172.18.0.0/16`** range. The node assigns each container an IP in this network automatically unless you override it via the **IP** field on the Overview screen.

- The node side of the bridge holds **`172.18.1.2`**. This is also the address of the [DNS resolver](#dns-resolver) — containers should send queries here, not to a public resolver, in order to benefit from container-to-container name resolution.
- Containers can address each other by **container name** within the same node. The node-side resolver maintains the mapping.
- The node itself reaches the container at the container's bridge IP. The container reaches the node at `172.18.1.2`.

You can override a container's IP by setting **IP** on the Overview to an address within `172.18.0.0/16` (other than `172.18.1.2`). This is rarely needed — pin an IP only when another container needs to reach this one by address rather than by name.

## DNS resolver

Inside the container, `resolv.conf` points at `172.18.1.2`. The node-side resolver:

1. Resolves any container running on the same node by its `name` (the value set on the Overview screen).
2. Forwards anything else to the node's configured DNS servers (see **Networking → Interfaces** for the node's resolvers).

You can override this by setting **DNS** on the container's Overview — for example, to point a container directly at an internal DNS server reachable over a virtual network. This bypasses the node's resolver and disables container-to-container name resolution.

## Inbound traffic — three attach modes

There are three independent ways to expose a container to traffic. They can be combined.

### 1. Host port mappings

A host port mapping listens on one of the **node's** physical interfaces and forwards matching traffic to a port inside the container.

| Field | Notes |
| --- | --- |
| **Protocol** | `tcp`, `udp`, or unspecified (forwards all protocols). |
| **Host Interface** | The node NIC to listen on, e.g. `ens192`. The node's Networking → Interfaces page lists available NICs. |
| **Host Port** | Port on the host interface. |
| **Container Port** | Port inside the container that receives the traffic. |

Use this when something on the node's local network needs to reach the container — for example, a workstation on the same LAN curling an HTTP API.

{{<alert color="info">}}
Host port mappings are bound to a specific node NIC, not to `0.0.0.0`. To expose a container on more than one node interface, add one mapping per interface.
{{</alert>}}

### 2. Virtual networks

Attaching a virtual network to a container grafts a second interface into the container. The container is assigned an IP within the virtual network's CIDR, and traffic arriving on that virtual network is delivered to the container.

| Field | Notes |
| --- | --- |
| **Virtual Network** | The virtual network to attach. |
| **Virtual IP** | The address to assign on that network. Must be within the network's configured CIDR. |
| **Allow Outbound** | When enabled, the container can also originate connections out onto the virtual network. When disabled, the attachment is inbound-only. |

This is how you expose a container to other Trustgrid nodes over a VPN — see [Expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}}).

### 3. Virtual interfaces

A virtual interface forwards **all** traffic from a node-level interface (typically a tunnel adapter) directly into the container, where it appears as a dedicated interface.

| Field | Notes |
| --- | --- |
| **Name** | The node-side virtual interface to attach. |
| **Destination** | The interface name to present inside the container (e.g. `eth1`). |

Use this when the container itself needs to manage the interface — for example, a container running its own VPN client or a packet sniffer.

## Outbound traffic — where does it go?

When a container originates a connection:

1. **To another container on the same node** — routed across the bridge directly, no NAT.
2. **To anything else** — the connection is source-NAT'd onto whichever node interface routes to the destination. By default this is the node's default gateway, i.e. the same path the node itself uses to reach the internet.
3. **To a peer over a virtual network** — only available if the container has a [virtual network attachment](#2-virtual-networks) with **Allow Outbound** enabled, or has the destination route inside its **VRF**.

### VRFs

The **VRF** field on the Network screen scopes the container's routing table to a specific VRF defined on the node. Use this to keep a container's outbound traffic isolated from the node's default routing — for example, forcing all traffic out a specific tunnel.

When no VRF is selected, the container uses the node's default routing table.

### Internet egress and firewalling

There is no separate "container firewall" — outbound traffic from a container that exits the node goes through the node's network stack, so any node-level firewall rules apply. To block specific destinations from a container, configure node-level rules or detach **Allow Outbound** from any virtual network attachments the container doesn't actually need.

## Putting it together

A common pattern: a service container with port mappings on the node's LAN interface (for local administrative access) plus a virtual network attachment (for application traffic over the Trustgrid overlay):

```
       LAN clients                    Trustgrid peers
            │                                │
       :8080│                                │ :443 (vnet IP)
            ▼                                ▼
    ┌─────────────────────────────────────────────┐
    │  Node nate-edge1                            │
    │  ens192=192.168.100.209  vnet0=10.50.0.5    │
    │                                             │
    │      ┌─────────────────────────────┐        │
    │      │  Container docs-nginx       │        │
    │      │  br0  = 172.18.0.7          │        │
    │      │  vnet0 = 10.50.0.5          │        │
    │      │  resolver: 172.18.1.2       │        │
    │      └─────────────────────────────┘        │
    └─────────────────────────────────────────────┘
                       ▲
                       │ NAT'd through ens160 (WAN)
                       │
                  Internet egress
```

The container has three ways in (LAN port 8080, virtual network 10.50.0.5, no virtual interface in this example) and one way out (NAT'd through the node's default route).

## Related

- [Expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}})
- [Container troubleshooting — networking section]({{<ref "../troubleshooting#networking">}})
