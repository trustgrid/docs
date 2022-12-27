---
categories: ["overview"]
tags: ["architecture"]
title: "Networking"
date: 2022-12-27
description: >
  Trustgrid networking overview
---

{{% pageinfo %}}
Trustgrid delivers a software defined wide area network (SD-WAN) that enables connectivity between cloud environments, data centers, and edge or on-premise hosts. Our robust networking feature set provides layer 3 and layer 4 network overlays with cloud-like management tools.
{{% /pageinfo %}}

### Types of Connectivity

Trustgrid establishes TLS tunnels between [nodes]({{< ref "/docs/concepts/node" >}}) using Mutual Authentication. These tunnels between Edge Nodes and Gateway Nodes constitute the data plane of a Trustgrid network. Over this tunnel two distinct types of connectivity are supported, Layer 3 and Layer 4. Layer 3 supports [routed]({{< ref "docs/domain/routes" >}}) traffic and Network Address Translation (NAT) to replicate traditional WAN connections. Trustgrid's Layer 4 networking uses reverse proxies to map local IP-port number combinations to remote hosts. Each connectivity type has distinct advantages and may be used for either Gateway Nodes or Edge Nodes.

### Layer 3 Overview

The Trustgrid firewall feature provides the capability to securely route IP packets between networks. In this configuration, Trustgrid [nodes]({{< ref "docs/concepts/node" >}}) can operate as a distributed mesh virtual private network ([VPN]({{< ref "docs/concepts/VPN" >}})) that can allow applications to access remote data and services at layer 3 (L3) of the network OSI model. This is done by defining a virtual L3 network (similar to an Amazon VPC) and then selecting how Edge node networks are exposed and translated into the the virtual address space.

Layer 3 configuration is required for all Trustgrid networking deployments on the Gateway (data center / cloud) side. This allows the use of L3 (Firewall) or Layer 4 (L4) configurations for Edge Node deployments. Layer 3 connectivity enables easy support of many hosts at each side of the connection. This allows the replication of existing [routes]({{< ref "docs/domain/routes" >}}) and subnets to minimize reconfiguration of any remote or local networks.

Layer 3 Edge configuration is ideal in situations where there is an incomplete understanding of all traffic (source and destination hosts/ports) as it supports more permissive ACLs. It is required when using stateful protocols (except FTP) and other protocols such as SIP.

All Gateway Nodes must be deployed using L3 networking in a two interface (management/data NIC) configuration.

### Layer 4 Overview

Trustgrid Layer 4 (L4) is only supported in Edge configurations. Layer 4 networking enables a reverse proxy connection between the Gateway Node IP address:port number that maps directly to an on-premise host:port. Specific network plug-ins support stateful features like FTP using L4 connectivity. With Trustgrid's L4 networks there is no concern of overlapping local IP subnets on either end of the tunnel because Trustgrid proxies the traffic. The source of all TCP/UDP traffic is the Edge Node interface, not the originating host. This is transparent to the hosts on either side.

Layer 4 connectivity offers slightly better performance than L3 in specific types of traffic and is ideal when there is a single or few remote/edge hosts that are being accessed with deterministic traffic. Layer 4 Edge Nodes are usually easier to deploy at an Edge site as they can avoid the configuration of a [Virtual Network]({{< ref "docs/domain/virtual-networks" >}}) Overlay (NATs, ACLs, [Routes]({{< ref "docs/domain/routes" >}}), etc). Remote sites that desire to view the actual source of traffic should use Layer 3 for Edge configurations as the source IP for all traffic in a Layer 4 connection is the Edge Node's data IP address.

### Layer 3 + 4 Overview

Trustgrid supports the use of Layer 3 and 4 configurations simultaneously (Gateway side is Layer 3, Edge side Layer 4), and is the recommended default configuration for networking. This enables use of Layer 3 in the data center or cloud and either Layer 3 or Layer 4 at the Edge.
