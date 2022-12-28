---
Title: "Edge Deployment Scenarios"
Tags: ["edge nodes", "Networking", "deployment"]
Date: 2022-12-28
---

When deciding on an Edge deployment architecture, there are several key variables to consider:

- Behind the Firewall or Beside the Firewall

- HA or Single [Node]({{< ref "docs/concepts/node" >}})

- Multi-interface or Single-interface

- Layer 3 or Layer 4 Edge [networking]({{< ref "docs/overview/networking" >}})

Trustgrid supports unique configurations per site to minimize the changes required at the Edge network. The easiest scenario to deploy is a Layer 4, behind the firewall, single [node]({{< ref "docs/concepts/node" >}}) deployment. This is the standard recommendation. When deployed behind a firewall, the firewall must allow egress traffic to Trustgrid's Cloud environment (see Network Requirements) and the [network(s)]({{< ref "docs/overview/networking" >}}) where the gateway nodes are installed.

### Beside a Firewall
In some scenarios, you may want to deploy your Trustgrid [node(s)]({{< ref "docs/concepts/node" >}}) beside the firewall. In this configuration, the management interface of each [node]({{< ref "docs/concepts/node" >}}) will be connected directly to the internet and the data interface will get connected to a LAN such as the inside/LAN network.

##### Edge Network Topology:

![img](/docs/domain/edge-topology.png)

### Considerations in this Configuration:
- A unique, static, public IP address is required for each [node]({{< ref "docs/concepts/node" >}}).
- Typical deployments beside a firewall involve a switch connected to the ISP’s ethernet handoff, the firewall, and the [nodes]({{< ref "docs/concepts/node" >}}).
- [Clustered nodes]({{< ref"docs/concepts/cluster" >}}) in this configuration may require using the data interface for [cluster]({{< ref"docs/concepts/cluster" >}}) communication.
- It is possible to configure L4 proxy connectors that are bound to the management interface. Care should be taken to ensure that L4 proxy connectors aren’t inadvertently configured to allow access to a service from the public internet.  If using L4 proxy connectors on the management interface, the source CIDR should be used to limit access to the service as desired.

#### Requirements to Configure:
- Standard
    - 1 Public IP Address
    - 1 Private IP Address
    - 2 DNS Servers
- HA
    - 2 Public IP Addresses
    - 2 Private IP Addresses
    - 1 Private IP Address for Cluster IP
    - 2 DNS Servers

### Behind a Firewall
The most typical configuration for edge nodes is to place the [nodes]({{< ref "docs/concepts/node" >}}) behind a firewall with the appropriate access controls in place to allow communication on the management interface.  In this configuration, the management configuration can be placed in a DMZ network and the data interface would get connected to a LAN such as the inside/LAN network.

##### Edge Network Topology:

![img](/docs/domain/edge-topology2.png)

#### Considerations in this Configuration:
- Multiple [nodes]({{< ref "docs/concepts/node" >}}) can NAT to a single public IP address.
- Firewalls in front of gateway nodes must be configured to allow access to the IP:Port of the gateway [nodes]({{< ref "docs/concepts/node" >}}) from the public IP address(es) that the [node(s)]({{< ref "docs/concepts/node" >}}) NAT to.
- The firewall must be configured to allow outbound traffic from the management interface of the [node(s)]({{< ref "docs/concepts/node" >}}) based on Trustgrid's network requirements.

#### Requirements to Configure:
- Standard
    - 1 Public IP Address
    - 1 Private IP Address
    - 2 DNS Servers
- HA
    - 2 Public IP Addresses
    - 2 Private IP Addresses
    - 1 Private IP Address for Cluster IP
    - 2 DNS Servers

### Firewalled Nodes
Some organizations may choose to put the management interface behind a firewall and also put a firewall between the data interface and the local [network]({{< ref "docs/overview/networking" >}}).  

##### Edge Network Topology:

![img](/docs/domain/edge-topology3.png)

#### Considerations in this Configuration:
- Multiple [nodes]({{< ref "docs/concepts/node" >}}) can NAT to a single public IP address.
- Firewalls in front of gateway [nodes]({{< ref "docs/concepts/node" >}}) must be configured to allow access to the IP:Port of the gateway nodes from the public IP address(es) that the [node(s)]({{< ref "docs/concepts/node" >}}) NAT to.
- The firewall between the [node(s)]({{< ref "docs/concepts/node" >}}) and the internet must be configured to allow outbound traffic from the management interface of the [node(s)]({{< ref "docs/concepts/node" >}}) based on Trustgrid's network requirements.
- The firewall between the [node(s)]({{< ref "docs/concepts/node" >}}) and the LAN must allow connectivity from all IP addresses associated with the [node(s)]({{< ref "docs/concepts/node" >}}) data interface to resources on the necessary services on the LAN. This includes any IP addresses used in outside NAT configurations as well as the cluster IP.

#### Requirements to Configure:
- Standard
    - 1 Public IP Address
    - 1 Private IP Address
    - 2 DNS Servers
- HA
    - 2 Public IP Addresses
    - 2 Private IP Addresses
    - 1 Private IP Address for Cluster IP
    - 2 DNS Servers

### On-a-stick - Single Node Interface
Our best practice recommendation is to separate the management and data traffic when possible, but in some cases the use of both network interfaces on edge nodes isn’t desirable. The most common scenario for this configuration is when the edge network is a small, basic, network where utilizing both interfaces on the [node(s)]({{< ref "docs/concepts/node" >}}) would mean connecting both interfaces to the same network. In this scenario, an on-a-stick configuration can be used.

##### Edge Network Topology:

![img](/docs/domain/edge-topology4.png)

#### Considerations in this Configuration:
- Only the management interface of the [node]({{< ref "docs/concepts/node" >}}) is used, requiring a minimum of three IP addresses to support both management/data traffic.
- The management interface must be connected to a network that allows outbound access to the internet based on our network documentation, as well as access to a LAN/inside network.
- Clustered [nodes]({{< ref "docs/concepts/node" >}}) should be configured with the [cluster]({{< ref "docs/concepts/cluster" >}}) communication IP set to the IP of the management interface.
  
#### Requirements to Configure:
- Standard
    - 1 Private IP Address
    - 2 DNS Servers
- HA
    - 2 Private IP Addresses
    - 1 Private IP Address for Cluster IP
    - 2 DNS Servers

### Hosted FI Core - Access to Remote Network via Edge
A common deployment scenario involves the need to access a network that is routable from the edge location, but outside of the [network]({{< ref "docs/overview/networking" >}}) that the Trustgrid edge node is configured on. An example of this is a financial institution with a hosted core that may be available from the financial institution over a connection such as a VPN or MPLS.

##### Edge Network Topology:

![img](/docs/domain/edge-topology5.png)

#### Considerations in this Configuration:
- Inside NAT(s) for the remote address should be added on the data interface to map the hosted core IP address(es) to virtual IPs on the [Trustgrid virtual network]({{< ref "docs/domain/virtual-networks" >}}).
- In this scenario setting a [route]({{< ref "docs/domain/routes" >}}) on the data interface from the Trustgrid portal may be necessary.
- In most cases the hosted core provider will need to allow access to the hosted core from any IP addresses associated with a [node's]({{< ref "docs/concepts/node" >}}) data interface. This includes any IP addresses used in outside NAT configurations as well as the cluster IP.


### Outside NAT - Avoiding Routing at the Edge
In some scenarios, it may not be desirable to have to add [routes]({{< ref "docs/domain/routes" >}}) on the edge network in order to [route]({{< ref "docs/domain/routes" >}}) traffic through the [Trustgrid network]({{< ref "docs/overview/networking" >}}). We recommend using L4 proxy in these scenarios when possible, however outside NAT can be utilized to allow traffic to and from the edge network without adding [routes]({{< ref "docs/domain/routes" >}}) on the edge network. 

##### Edge Network Topology:

![img](/docs/domain/edge-topology6.png)

#### Considerations in this Configuration:
- In this configuration you can configure PAT, which allows you to source all traffic from the [node]({{< ref "docs/concepts/node" >}}) on a single IP address that exists on the local network.
- The IP address that you select must be a free IP address on the network that is also excluded from any DHCP scope on the subnet that it's in.
- In some cases it may be desirable to use multiple outside NATs in order to identify traffic from a gateway-side host using a unique source IP at the edge. Using more than a handful of outside NATs is not recommended as it may present a problem with connectivity during a HA failover situation.

### L3 to L4 Proxying - Function with 1 IP
In deployments where only a few services need to be accessed via an edge node, it may make sense to deploy the edge node in a mode in which the edge node intercepts layer 3 TCP traffic and proxies the traffic to a TCP service defined in the L4 Proxy configuration. The advantage of configuring the edge node in this fashion is that the traffic effectively does a source pat out of the [node]({{< ref "docs/concepts/node" >}}) interface IP address, thus simplifying the network configuration at the edge node location. The edge node appears to the server as a single client machine accessing the server, rather than a range of different IP clients for which the network in that location must be configured to forward traffic to. One of the biggest differences between this and the previous outside nat configuration is that all traffic is coming from the single interface IP address, rather than a separate IP address.

#### Considerations in this Configuration:
- In this configuration all traffic for a configured service will come from a single IP address. If you need to differentiate different clients by IP address, this configuration will not suffice.
- Only TCP services can be used in this configuration. UDP and ICMP will not be usable with this configuration.















