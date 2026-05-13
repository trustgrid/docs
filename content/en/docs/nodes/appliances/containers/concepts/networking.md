---
title: "Container Networking"
weight: 10
---

By default a container is reachable only from the node it runs on. To make it reachable from somewhere else — the LAN, another Trustgrid node over VPN, or the internet — you attach it to the network in one of three ways. This page covers those three ways plus how outbound traffic from the container leaves the node.

## Container addresses and DNS

When a container starts, the node gives it an IP address in `172.18.0.0/16` (the default container network — contact Trustgrid support if you need to change it). The container can reach the node at `172.18.1.2`, and any other container on the same node by **container name** — the name you typed into the Overview screen.

You normally don't need to change any of this. The two relevant overrides on the Overview screen:

- **IP** — pin the container to a specific address within `172.18.0.0/16`. Only needed if another container has to reach this one by address rather than name.
- **DNS** — point the container at a custom DNS server. Doing this means the container can no longer resolve its sibling containers by name, so only set it if you specifically need a different resolver.

## Three ways to expose a container

You can use any of these, or combine them.

### 1. Host port mappings

A host port mapping puts the container on one of the node's network interfaces — so something else on the same LAN as the node (a workstation, another server) can reach it.

| Field | Notes |
| --- | --- |
| **Protocol** | `tcp`, `udp`, or leave blank for both. |
| **Host Interface** | The node's network port to listen on, e.g. `ens192`. The list of interfaces is on the node's **Networking → Interfaces** page. |
| **Host Port** | The port to listen on. |
| **Container Port** | The port inside the container that should receive the traffic. |

Use this when something on the node's LAN needs to reach the container.

{{<alert color="info">}}
Each mapping listens on one specific node interface. To expose a container on more than one interface, add one mapping per interface.
{{</alert>}}

### 2. Virtual networks

Attaching a Trustgrid virtual network to a container lets it talk to other Trustgrid nodes over the VPN overlay, as if it were a peer.

| Field | Notes |
| --- | --- |
| **Virtual Network** | The Trustgrid virtual network to attach. |
| **Virtual IP** | The address the container should use on that network. |
| **Allow Outbound** | When on, the container can also originate connections out onto the virtual network. When off, traffic only flows into the container. |

Use this when a container needs to reach (or be reached by) other Trustgrid nodes. See the [Expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}}) tutorial for the end-to-end walkthrough.

### 3. Virtual interfaces

A virtual interface forwards **all** traffic from a node-level interface (typically a tunnel) directly into the container, where it shows up as one of the container's own network interfaces.

| Field | Notes |
| --- | --- |
| **Name** | The node-side interface to forward. |
| **Destination** | The interface name to use inside the container (e.g. `eth1`). |

Use this when the container itself needs to manage the interface — running its own VPN client, capturing packets, that sort of thing.

## Outbound traffic — where does it go?

When a container makes an outbound connection:

- **To another container on the same node** — goes directly between them.
- **To anything else (the internet, the node's LAN, an internal server)** — leaves the node using the same network path the node itself uses for outbound traffic. Any firewall rules on the node also apply to the container.
- **To a peer over Trustgrid VPN** — only works if the container has a [virtual network attachment](#2-virtual-networks) with Allow Outbound on, or if you've placed it in a VRF that routes there.

### VRFs

The **VRF** field on the Network screen lets you put the container in a specific routing context defined on the node — useful for forcing all of a container's outbound traffic out a particular tunnel, separate from the rest of the node. Leave it blank to use the node's normal routing.

## A common pattern

A web service container that's reachable from the LAN for admin and from other Trustgrid nodes for application traffic:

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
    │      │  bridge IP = 172.18.0.7     │        │
    │      │  vnet IP   = 10.50.0.5      │        │
    │      └─────────────────────────────┘        │
    └─────────────────────────────────────────────┘
                       ▲
                       │ Outbound via the node's normal route
                       ▼
                  Internet
```

The container has two ways in (LAN port 8080, virtual network IP) and one way out (whatever route the node uses).

## Related

- [Expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}})
- [Container troubleshooting — networking section]({{<ref "../troubleshooting#networking">}})
