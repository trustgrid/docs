---
categories: ["node"]
tags: ["layer 4", "networking"]
title: "Services"
linkTitle: "Services"
aliases: 
    - /docs/nodes/services
description: Configure Layer 4 (L4) services to make them accessible via peer nodes
---

{{% pageinfo %}}
Services are configured in conjunction with [connectors]({{<ref "docs/nodes/shared/connectors" >}}) or [port forwards]({{<ref "docs/nodes/appliances/vpn/port-forwarding">}}) to define a host (IP or DNS) and port to connect to for layer 4 (L4) connectivity.
{{% /pageinfo %}}

## Description
Services define the host name or IP and port of a server that can be accessed by the Trustgrid node (or members of the cluster) where it is defined.  Remote clients can then access the service via peer nodes using [connectors]({{<ref "docs/nodes/shared/connectors" >}}) or [VPN port forwards]({{<ref "docs/nodes/appliances/vpn/port-forwarding">}}). 

## Functionality
On the local network, a connection's **source IP** will be the node's interface IP.
- If the service is defined on a cluster, the active cluster members IP will be used.
- If the node has multiple interfaces, the source will be the interface that has a matching [interface route]({{<ref "docs/nodes/appliances/interfaces#interface-routes">}}) for the target IP.

If the service's [host field]({{<ref "#host">}}) is set to a DNS name, the node will use its WAN interface DNS servers to resolve the address.


## Configuration

Services are configured under the Networking > Services panel of the node or cluster configuration section in the Trustgrid Portal.

{{<tgimg src="service.png" width="40%" caption="Add Service dialogue" alt="Dialogue to add a service with fields for enabled" >}}

{{<fields>}}
{{<field "Enabled">}}
Values: Yes or No. Allows each service to be individually disabled if desired"
{{</field>}}

{{<field "Protocol" >}}
The protocol of the service to connect to. Options are TCP, UDP, FTP, and TFTP, along with pre-defined default ports for RDP, SSH, and VNC.

> FTP must operate in passive mode when using L4 services and connectors.
{{</field >}}



{{<field "Service Name" >}}
A friendly name for the service that will be used in the Remote Service field of a [connector]({{<ref "docs/nodes/shared/connectors" >}}). Can only contain letters, numbers, and the `-` character.
{{</field >}}

{{<field "Host" >}}
The IP or DNS address of the host to connect to.
{{</field >}}

{{<field "Port" >}}
The port to connect to on the host.
{{</field >}}

{{<field "Description">}}
(Optional) User friendly description of the service.
{{</field>}}

{{<field "Test Connectivity">}}
Show only on the Services table for TCP Protocol services. When clicked, the client node will attempt a TCP port connection to the configured host and port and display success or failure.
{{<tgimg src="service-test-connectivity.png" width="80%">}}
{{</field>}}

{{</fields>}}
