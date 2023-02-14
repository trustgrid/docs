---
categories: ["node"]
tags: ["layer 4", "networking"]
title: "Services"
linkTitle: "Services"
---

{{% pageinfo %}}
Services are configured in conjunction with [connectors]({{< ref "docs/node/connectors" >}}) to define a remote host:port to connect to for layer 4 (L4) connectivity.
{{% /pageinfo %}}

#### Configuration

Services are configured under the Networking > Services tab of the node or cluster configuration section in the Trustgrid Portal.

![Layer 4 - Add a service](service.png)

{{<fields>}}
{{<field "Protocol" >}}
The protocol of the service to connect to. Options are TCP, UDP, FTP, and TFTP, along with pre-defined default ports for RDP, SSH, and VNC.
{{</field >}}

> FTP must operate in passive mode when using L4 services and connectors

{{<field "Service Name" >}}
A friendly name for the service that will be used in the Remote Service field of a [connector]({{< ref "docs/node/connectors" >}}).
{{</field >}}

{{<field "Host" >}}
The IP of the host to connect to.
{{</field >}}

{{<field "Port" >}}
The port to connect to on the host.
{{</field >}}

{{<field "Health Check" >}}
A health check can be configured for TCP services. The health check will attempt a TCP connectivity check once per minute. After 5 consecutive failures, an Outbound L4 Service Health Check alert will be triggered and clustered nodes will be marked as unhealthy. Clustered nodes will return to a Healthy status after 5 consecitive connection attempts are successful.
{{</field >}}
{{</fields>}}
