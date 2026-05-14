---
title: "Container Networking"
weight: 10
---

By default a container is reachable only from the appliance it runs on. To make it reachable from somewhere else — the LAN, another Trustgrid node over VPN, or the internet — attach it to the network in one of three ways. This page covers those three ways plus how outbound traffic from the container leaves the appliance.

## Container addresses and DNS

When a container starts, an IP address in `172.18.0.0/16` (the default container network — contact Trustgrid support if you need to change it) is assigned. The container can communicate with its parent appliance and with any other container running on the same appliance by **container name** — the name you typed into the Overview screen.

You normally don't need to change any of this. The two relevant overrides on the Overview screen:

- **IP** — pin the container to a specific address within `172.18.0.0/16`. Only needed if another container has to reach this one by address rather than name.
- **DNS** — point the container at a custom DNS server. Doing this means the container can no longer resolve its sibling containers by name, so only set it if you specifically need a different resolver.

## Three ways to expose a container

You can use any of these, or combine them.

### 1. Host port mappings

A host port mapping puts the container on one of the appliance's network interfaces — so something else on the same LAN (a workstation, another server) can reach it.

| Field | Notes |
| --- | --- |
| **Protocol** | `tcp`, `udp`, or leave blank for both. |
| **Host Interface** | The appliance's network port to listen on, e.g. `ens192`. The list of interfaces is on the **Networking → Interfaces** page. |
| **Host Port** | The port to listen on. |
| **Container Port** | The port inside the container that should receive the traffic. |

Use this when something on the appliance's LAN needs to reach the container.

{{<alert color="info">}}
Each mapping listens on one specific appliance interface. To expose a container on more than one interface, add one mapping per interface.
{{</alert>}}

### 2. Virtual networks

Attaching a Trustgrid virtual network to a container lets it talk to other Trustgrid nodes over the VPN overlay, as if it were a peer.

| Field | Notes |
| --- | --- |
| **Virtual Network** | The Trustgrid virtual network to attach. |
| **Virtual IP** | The address the container should use on that network. |
| **Allow Outbound** | When on, the container can also originate connections out onto the virtual network. When off, traffic only flows into the container. |

Use this when a container needs to reach (or be reached by) other Trustgrid nodes. See the [Expose a container over a virtual network]({{<relref "/tutorials/containers/expose-over-vpn">}}) tutorial for the end-to-end walkthrough.

### 3. Virtual interfaces

A virtual interface forwards **all** traffic from an appliance-level interface (typically a tunnel) directly into the container, where it shows up as one of the container's own network interfaces.

| Field | Notes |
| --- | --- |
| **Name** | The appliance-side interface to forward. |
| **Destination** | The interface name to use inside the container (e.g. `eth1`). |

Use this when the container itself needs to manage the interface — running its own VPN client, capturing packets, that sort of thing.

## Outbound traffic — where does it go?

When a container makes an outbound connection:

- **To another container on the same appliance** — goes directly between them.
- **To anything else (the internet, the appliance's LAN, an internal server)** — leaves the appliance using the same network path the appliance itself uses for outbound traffic.
- **To a peer over Trustgrid VPN** — only works if the container has a [virtual network attachment](#2-virtual-networks) with Allow Outbound on, or if you've placed it in a VRF that routes there.

### VRFs

The **VRF** field on the Network screen lets you put the container in a specific routing context — useful for forcing all of a container's outbound traffic out a particular tunnel, separate from the rest of the appliance. Leave it blank to use the appliance's normal routing.

## Deployment example

A web service container with three traffic paths:

- **LAN clients** reach it through a port mapping on the appliance's LAN interface.
- **Other Trustgrid nodes** reach it on its virtual network address.
- **The container itself** reaches the internet through the appliance's WAN — for package updates, calling external services, etc.

```
      LAN client          Trustgrid peer node                       Internet
           │                       │                                   ▲
           │ port 8080             │ port 443 to 10.50.0.10            │ container
           │                       │ over the virtual network          │ outbound
           ▼                       ▼                                   │
   ┌─────────────────────┬───────────────────────────────────────────────┐
   │   LAN interface     │              WAN interface                    │
   │  192.168.100.209    │              172.16.0.10                      │
   ├─────────────────────┴───────────────────────────────────────────────┤
   │                            Trustgrid node                           │
   │                                                                     │
   │            │ NAT in          │ NAT in            ▲ NAT out          │
   │            │ (port map)      │ (virtual network) │ (to internet)    │
   │            ▼                 ▼                   │                  │
   │       ┌───────────────────────────────────────────────────┐         │
   │       │  Container nginx                                  │         │
   │       │      Container address        172.18.0.7          │         │
   │       │      Virtual-network address  10.50.0.10          │         │
   │       └───────────────────────────────────────────────────┘         │
   │                                                                     │
   └─────────────────────────────────────────────────────────────────────┘
```

## Related

- [Expose a container over a virtual network]({{<relref "/tutorials/containers/expose-over-vpn">}})
