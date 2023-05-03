---
categories: ["node"]
tags: ["layer 4", "networking"]
title: "Services"
linkTitle: "Services"
weight: 19
---

{{% pageinfo %}}
Services are configured in conjunction with [connectors]({{<ref "docs/nodes/connectors" >}}) to define a remote host:port to connect to for layer 4 (L4) connectivity.
{{% /pageinfo %}}

#### Configuration

Services are configured under the Networking > Services tab of the node or cluster configuration section in the Trustgrid Portal.

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
A friendly name for the service that will be used in the Remote Service field of a [connector]({{<ref "docs/nodes/connectors" >}}). Can only contain letters, numbers, and the `-` character.
{{</field >}}

{{<field "Host" >}}
The IP of the host to connect to.
{{</field >}}

{{<field "Port" >}}
The port to connect to on the host.
{{</field >}}

{{<field "Description">}}
(Optional) User friendly description of the service.
{{</field>}}

{{<field "Test Connectivity">}}
Show only on the Services table for TCP Protocol services. When client the node will attempt a TCP port connection test to the configured host and port, and then display success or failure.
{{<tgimg src="service-test-connectivity.png" width="80%">}}
{{</field>}}

{{</fields>}}
