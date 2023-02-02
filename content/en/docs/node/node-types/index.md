---
categories: ["concepts"]
tags: ["node", "concepts"]
title: "Node Types"
date: 2022-12-19
---

### Gateway Nodes

Gateway Nodes accept incoming TLS tunnels from Edge Nodes. Traffic on these tunnels is bidirectional, as permitted by ACLs or policy. Gateway Nodes usually require a firewall change to permit the incoming traffic. Gateway Nodes are identical to Edge Nodes but with a gateway configuration applied. Gateway nodes only connect to other gateways if one is of type `hub`. Read more about [gateway configuration]({{< ref "docs/node/gateway" >}}).

### Edge Nodes

Edge Nodes build outgoing TLS tunnels to Gateway Nodes. Bidirectional traffic is supported through this tunnel, subject to ACLs and [security]({{< ref "docs/overview/security">}}) policy restrictions. Edge Nodes will only require a firewall rule change if outbound internet restrictions are in place. Edge nodes can be a target for software deployment. Edge Nodes can be deployed with a single or multiple ethernet connections to support deployments behind a firewall or adjacent to a firewall (public WAN / private LAN).

### Management Nodes

Management Nodes are not like other [nodes]({{< ref "docs/node" >}}) because they do not connect to the data plane, only the control plane. Management Nodes are deployed by Trustgrid for each customer and are multi-tenant like other control plane components. Customers may elect to deploy their own Management Nodes in place of multi-tenant Management Nodes. Management Nodes facilitate the monitoring, management, and support of Edge and Gateway Nodes.
