---
title: "Expose a Container over a Virtual Network"
description: Connect a container to a Trustgrid virtual network so it's reachable from another site — the overlay is the only inbound path, while the container can still call out to its local LAN and the public internet.
---

This tutorial walks through making a container reachable across sites by attaching it to a Trustgrid virtual network.

- **Inbound:** the container is reachable across the virtual network via the configured virtual IP.
- **Outbound to the LAN:** the container can call services on the appliance's local network.
- **Outbound to the internet:** the container can call public services through the appliance's WAN.
- **No other inbound path:** there is no inbound access to the container from outside the appliance unless you explicitly configure a host port mapping. The overlay is the only way in.

The walk-through uses a small example: `acme-api`, a container running on Acme Bank's edge cluster, queried by a fintech provider via the overlay. The container fetches data from Acme's on-prem core (LAN) and is able to make outbound calls to a SaaS endpoint (WAN).

**Starting point**

Both clusters already exist — nodes are deployed and paired into a cluster, but no virtual-network configuration is applied yet. See [Clusters]({{<relref "/docs/clusters">}}) for cluster setup.

- An edge cluster running the container, set up via the [Container Quickstart]({{<relref "/tutorials/containers/quickstart">}}).
- A gateway cluster with nothing configured on it yet beyond cluster membership.

{{<tgimg src="topology.svg" caption="The fintech app reaches the container at its overlay IP over the Trustgrid Data Plane. From there, the container can reach the FI Core on the institution LAN and SaaS apps on the public internet, both through the edge appliance. The container has no other inbound path." width="95%">}}

## 1. Enable each gateway appliance as a public gateway

**Nodes → Node Name → Gateway → Server**

| Field | Value |
| --- | --- |
| Enable | `Enabled` |
| Public IP or DNS | The public endpoint that terminates the TLS / UDP Data Plane tunnels from connecting nodes |
| Port | `8443` |
| UDP Port | `8443` |
| Gateway Type | `Public` |

See [Gateway Server]({{<relref "/docs/nodes/appliances/gateway/gateway-server">}}) for the full options reference.

{{<tgimg src="tutorial-gw-server.png" caption="Gateway Server settings on a gateway node" width="90%">}}

## 2. Create the virtual network

**Domain → Virtual Networks → Add**: name it `acme-vnet-0515-1538`, CIDR `172.20.0.0/16`. A subnet is allocated to each cluster.

{{<tgimg src="tutorial-vnet-list.png" caption="Virtual networks" width="90%">}}

## 3. Add virtual network routes

Routes tell the overlay where to send traffic for each destination subnet — packets are forwarded to the active cluster master:

| CIDR | Destination |
| --- | --- |
| `172.20.0.0/24` | gateway cluster |
| `172.21.0.0/24` | edge cluster |

{{<tgimg src="tutorial-vnet-routes.png" caption="Virtual network routes — one per side" width="90%">}}

## 4. Add access policies

Trustgrid virtual networks are zero trust — all traffic is denied by default. Add a single rule that permits exactly the traffic the integration needs — the fintech side reaching the container on its API port:

| Protocol | Source | Destination | Ports | Action |
| --- | --- | --- | --- | --- |
| `tcp` | `172.20.0.0/24` (fintech slice) | `172.21.0.10/32` (container VIP) | `80` | Allow |

{{<tgimg src="tutorial-vnet-policies.png" caption="A single scoped allow rule — fintech site to the acme-api container on port 80" width="90%">}}

## 5. Attach virtual network to gateway cluster interface

**Clusters → Gateway Cluster → Network → VPN → Actions → Attach**

- Virtual Network: `acme-vnet-0515-1538`
- Validation CIDR: `172.20.0.0/24`
- Interface: gateway-side LAN NIC (e.g. `ens192`)
- Inside NAT: `172.20.0.0/24 ← 192.168.200.0/24`

{{<tgimg src="tutorial-gw-vpn-nats.png" caption="Gateway cluster VPN attachment" width="90%">}}

## 6. Attach the virtual network to the edge cluster

**Clusters → Edge Cluster → Network → VPN → Actions → Attach**

- Virtual Network: `acme-vnet-0515-1538`
- Validation CIDR: `172.21.0.0/24`

{{<tgimg src="tutorial-edge-vpn-nats.png" caption="Edge cluster VPN attachment" width="90%">}}

## 7. Attach the container to the virtual network

Open the container on the edge cluster → **Network → Virtual Networks → Add Entry**:

| Field | Value |
| --- | --- |
| Virtual Network | `acme-vnet-0515-1538` |
| Virtual IP | `172.21.0.10` |
| Allow Outbound | off |

{{<tgimg src="tutorial-container-network-vnet.png" caption="Container attached at 172.21.0.10" width="90%">}}

## Verify — the fintech queries the FI Core through the container API

Run the test from the fintech side. The fintech provider's backend — a host on the gateway-side LAN — calls the `acme-api` container at its overlay IP:

```bash
$ curl http://172.21.0.10/v1/accounts
{
  "service": "acme-fintech-core",
  "host": "edge-host",
  "timestamp": "2026-05-15T17:26:19.329926Z",
  "accounts": [
    { "id": "ACC-1001", "balance": 4218.55 },
    { "id": "ACC-1002", "balance": 19234.10 }
  ]
}
```

That response is what Acme Bank's on-prem core returned. The request path:

1. The fintech provider's backend calls `http://172.21.0.10/v1/accounts`. The call is routed to the Trustgrid gateway cluster IP, which is always served by the cluster's active node.
2. The Trustgrid virtual network overlay carries the request across the secure, encrypted Data Plane to Acme's edge appliance.
3. The `acme-api` container receives the request at its overlay VIP.
4. The container's nginx forwards the request to Acme's FI Core at `192.168.100.72:8443` over the appliance's LAN.
5. The FI Core returns the account data, which retraces the path back to the fintech.

One request, end to end. The fintech never touches Acme's LAN; Acme never exposes the core publicly. The container is the only thing either side talks to, and the integration's policy is enforced by the single zero-trust ACL.

Traffic flows one way. The FI cannot initiate connections back to the fintech — the container's **Allow Outbound** is off and the ACL only permits fintech → container.
