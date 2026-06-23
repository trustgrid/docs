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
- As of the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}), the [Source Interface](#configuration) setting lets you choose the egress interface and source connections from the cluster IP instead of the interface IP.

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

{{<field "Source Interface" >}}
Determines the interface and source IP used for the connection to the host. Requires the [June 2026 release]({{<ref "/release-notes/node/2026-06/index.md">}}) or later, and is not shown on earlier versions.

Select the interface to source the connection from, then choose how its IP is determined:
- **Use Interface IP** - Source connections from the selected interface's IP. This is the default and matches the behavior of earlier releases.
- **Use Cluster IP** - Source connections from the [cluster IP]({{<relref "/docs/clusters/cluster-only-config#cluster-ip">}}) instead of the interface IP. Available for clustered nodes with a configured cluster IP.
{{<tgimg src="add-service-source-interface.png" width="40%" caption="The Source Interface options: Use Interface IP or Use Cluster IP." alt="Add Service dialog showing the Source Interface dropdown with Use Interface IP and Use Cluster IP options">}}
{{</field >}}

{{<field "Connect Timeout" >}}
(Optional) The maximum amount of time, in seconds, a TCP connection will wait for a response before timing out. Defaults to 5 seconds.
{{</field >}}

{{<field "Description">}}
(Optional) User friendly description of the service.
{{</field>}}

{{<field "Test Connectivity">}}
Show only on the Services table for TCP Protocol services. When clicked, the client node will attempt a TCP port connection to the configured host and port and display success or failure.
{{<tgimg src="service-test-connectivity.png" width="80%">}}
{{</field>}}

{{</fields>}}
