---
categories: ["node"]
tags: ["layer 4", "networking"]
title: "Connectors"
linkTitle: "Connectors"
---

{{% pageinfo %}}
Connectors are configured in conjunction with [services]({{<ref "docs/nodes/services" >}}) to define a local listening port for layer 4 (L4) connectivity.
{{% /pageinfo %}}

#### Configuration

Connectors are configured under the Networking > Connectors tab of the node or cluster configuration section in the Trustgrid Portal.

{{<imgproc "connector.png" Resize "400x450"/>}}

{{<fields>}}
{{<field "Protocol" >}}
The protocol that the listener for the connector will use. Options are TCP, UDP, FTP, and TFTP.

> FTP must operate in passive mode when using L4 services and connectors.
> {{</field >}}

{{<field "Listen Interface" >}}
The interface that you want the connector to listen on. If set to "All", the connector will listen on 0.0.0.0.
{{</field >}}

{{<field "Listen Port" >}}
The port that you want the connector to listen on.
{{</field >}}

{{<field "Remote Node" >}}
The remote node or cluster that traffic will be proxied to.
{{</field >}}

{{<field "Remote Service" >}}
Either the friendly name of [service]({{<ref "docs/nodes/services" >}}) that is defined on the remote node, or the IP:Port of the remote service to connect to.
{{</field >}}

{{<field "Rate Limit (Mb/s)" >}}
The maximum amount of throughput that will be allowed to traverse the tunnel when connecting to the connector. This can be used to prevent saturating the connection at either the local or remote sites.
{{</field >}}
{{</fields>}}
