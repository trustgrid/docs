---
Title: "Overview Section"
Date: 1-27-2023
---

### Stats

The [Node overview page](https://portal.dev.trustgrid.io/#/node/a9f1fedd-a474-455b-a123-37a035409229) shows performance and network traffic data.

Supported time windows are selectable at the top. VPN and network statistics can be targeted to specific virtual networks and interfaces.

![img](node-overview.png)

![img](node-overview2.png)

{{< field-def "Node Performance" >}}
shows CPU, disk, and memory usage percentages
{{< /field-def >}}
{{< field-def "VPN Traffic Volume" >}}
shows data usage sent and received, across all VPNs and for the selected virtual network
{{< /field-def >}}
{{< field-def "Traffic Volume" >}}
shows data sent and receives, across all interfaces and for the selected interface
{{< /field-def >}}
{{< field-def "Connected Peers" >}}
shows the number of other nodes connected. This will change based on the node type - gateways connect to all edge nodes, while edge nodes only connect to gateways.(LINK HUB GATEWAY PAGE WHEN DOCUMENTED)
{{< /field-def >}}
{{< field-def "VPN Flows" >}}
shows new and active flows, across all VPNs and for the selected virtual network
{{< /field-def >}}
{{< field-def "TCP Errors" >}}
shows TCP errors across all interfaces
{{< /field-def >}}