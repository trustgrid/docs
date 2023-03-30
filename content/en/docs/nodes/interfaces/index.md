---
tags: ["node"]
title: "Interfaces"
linkTitle: "Interfaces"
description: Configure and manage the node's local network configuration
weight: 9
---

## Network Interfaces

Navigate to a node, then select `Interfaces` under the `Network` section.

![Interfaces Page](interfaces.png)

The `Interface` dropdown at the top allows you to select which interface to manage.

![Interface Dropdown](select-interface.png)

### ETH0 - WAN Interface

This interface is always used to build the outbound TLS tunnels for control and data plane connectivity. If the node is deployed in a single interface configuration, this interface will also be used for local network connectivity.

### ETH# - LAN Interface

Used in configurations with more than one interface for local network connectivity.

#### VLAN Subinterface 
 Virtual under a physical LAN interface that will apply a VLAN tag to traffic


## Configuration

{{<fields>}}
{{<field "Hardware Address">}}MAC address of the interface.{{</field>}}
{{<field "Interface VRF">}}Selects to which [VRF]({{<ref "/docs/nodes/vrfs">}}) the interface is attached.{{</field>}}
{{<field "VLAN ID (Subinterfaces Only)">}} Sets the VLAN ID/tag for the subinterface.{{</field>}}
{{<field "IP Assignment (ETH0 only)">}}
* DHCP - requires a DHCP server or relay in the same broadcast domain as the interface connection
* Static - configured using the fields below
{{</field>}}
{{<field "IPv4 Address CIDR">}}IP address using [classless inter-domain routing (CIDR) notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).{{</field>}}
{{<field "Gateway IP">}}
* ETH0 - The IP address of the default gateway for this device. All traffic not in the same network as an interface IP or with a specified route will use this gateway.
* ETH1 - This is only used in a 2-interface setup to access local resources outside the interface's configured network. Routes must be defined.{{</field>}}
{{<field "DNS Server 1/2 IP (ETH0 only)">}}Used to resolve DNS names to connect to the Trustgrid control plane.{{</field>}}
{{<field "Cluster Virtual IP">}}Used in [clusters]({{<ref "/docs/clusters">}}), this IP address is moved between member nodes based on the active cluster member. It can be used to route traffic between the virtual and local networks. This IP address can only be changed within the cluster interfaces page.{{</field>}}
{{<field "Auto-Negotiation">}}Sets if the interface will use auto-negotiation to determine the speed and duplex settings. Values can be AUTO _(default)_ or MANUAL.{{</field>}}
{{<field "Duplex">}}Locked unless Auto-Negotiation set to MANUAL. Sets the interface to either FULL _(default)_ or HALF duplex{{</field>}}
{{<field "Speed">}}Locked unless Auto-Negotiation set to MANUAL. Sets the interface to 10, 100, 1000 Mbps. {{</field>}}
{{<field "Ignore Health Check">}} Value can be Disable _(default)_ or Enable. If set to Enable, the interface's link status will no longer be used by the node to determine its [cluster member health state]({{<ref "/docs/clusters/#cluster-member-health-conditions">}}). Useful if a single node has a network connection it's peer does not such as an alternate WAN path.  {{</field>}}
{{</fields>}}

## Interface Routes

> If the node is part of a cluster, these must be configured at the cluster level

These routes are used to access networks that are not in the same network as a configured interface **and that are not accessed using the default gateway IP configured for ETH0 - WAN Interface**. It is rare to need to define these on the ETH0 interface.

1. Select the interface from the Network Interface dropdown.
1. Click the `+ Add Entry` button.
1. Enter the route information There are three fields:

	{{<fields>}}
	{{<field "Destination CIDR">}}This is the CIDR notation of the network or host that needs to be routed. For a single host, use /32 as the CIDR suffix.{{</field>}}
	{{<field "Next Hop">}}The IP address to which traffic destined for the above defined network will be sent. No CIDR suffix is needed.{{</field>}}
	{{<field "Description">}}Optional field to provide additional information about the route.{{</field>}}
	{{</fields>}}

1. Click the green check <i class="fa-solid fa-circle-check" style="color: green;"></i> to confirm the route.
1. Repeat steps 2-4 with any additional routes.
1. Click the `Save` button.

## Additional IPs
IP address entries added to the list below will be bound to the current interface.
{{<alert type="note">}}Only visible on LAN interfaces and Subinterfaces{{</alert>}}

## Advanced Network Settings
{{<fields>}}
{{<field "Dark Mode">}}Value can be Disable _(default)_ or Enable. When set to Enable the node will restrict ICMP responses so that it will not respond to ICMP (Ping) echo requests or respond to failed TCP/UDP connection attempts with reset or ICMP Destination Port Unreachable {{</field>}}
{{</fields>}}
