---
title: "Packet Simulator Interface Tool"
linkTitle: "Packet Simulator"
description: "Simulate a packet entering an interface and trace its path through the node"
---

The **Packet Simulator** shows what a node would do with a packet that enters one of its interfaces. You supply the interface, protocol, and addresses. The node walks the packet through its NAT, ACL, and routing rules without sending any traffic, then reports each step and whether the packet was forwarded or blocked.

Use it to confirm a virtual network configuration before sending real traffic, or to find which rule is dropping a flow.

{{<alert color="info">}}The Packet Simulator requires a node running the [June 2026 Major Appliance release]({{<relref "release-notes/node/2026-06" >}}) or newer. It covers VPN traffic only, not the L4 Proxy.{{</alert>}}

{{<alert color="info">}}Using the Packet Simulator requires the `nodes::service:packet-sim` permission. The built-in policies **tg-monitor**, **tg-operator**, and **tg-node-admin** grant this permission by default. The **Packet Simulator** button is hidden for users without it.{{</alert>}}

## Using the Packet Simulator

1. Log in to the Trustgrid portal and navigate to the node you want to test.
1. Select **Interfaces** under the **Network** section. From the dropdown, select the interface the packet should enter.
1. Click the **Packet Simulator** button in the **Interface Tools** section. {{<tgimg src="packet-sim-button.png" width="50%" caption="Packet Simulator button in Interface Tools.">}}
1. Fill in the form. {{<tgimg src="packet-sim-form.png" width="50%" caption="Packet Simulator form.">}}
    {{<fields>}}
    {{<field "Interface">}}Required. The interface the simulated packet enters. Defaults to the interface selected on the page.{{</field>}}
    {{<field "Protocol">}}`TCP`, `UDP`, or `ICMP`. Defaults to `TCP`.{{</field>}}
    {{<field "Source IP">}}Required. The IPv4 source address of the packet.{{</field>}}
    {{<field "Source Port">}}Optional. The source port, from `1` to `65535`.{{</field>}}
    {{<field "Destination IP">}}Required. The IPv4 destination address of the packet.{{</field>}}
    {{<field "Destination Port">}}Required. The destination port, from `1` to `65535`.{{</field>}}
    {{</fields>}}
1. Click **Execute** to run the simulation.
1. Review the result. Click **Back** to adjust the inputs and run again, or **Close** to finish.

## Reading the Result

{{<tgimg src="packet-sim-result.png" width="50%" caption="Example Packet Simulator result.">}}

The top of the result shows the source and destination with a protocol badge, followed by the interface and the virtual network that handled the packet.

A status line summarizes the outcome:

- `Packet traversed the <network> network successfully` means the packet would be forwarded.
- `Packet blocked at the <stage> stage` names the step that stopped the packet.
- `Packet blocked: <error>` shows an error returned by the node, such as `Invalid sourceIp specified`.

Below the status line, the path lists each step the packet took in order. Steps that block the packet are shown in red.

{{<fields>}}
{{<field "Ingress">}}The packet entering the node, with the interface, source, and destination.{{</field>}}
{{<field "NAT">}}A matching [NAT]({{<relref "docs/nodes/appliances/vpn/nats" >}}) rule, with its inside and outside addresses, networks, and descriptions.{{</field>}}
{{<field "Translate">}}The source and destination after NAT is applied.{{</field>}}
{{<field "ACL">}}The [access policy]({{<relref "docs/domain/virtual-networks/access-policy" >}}) evaluation. Shows **Allowed** or **Denied** and the matching line. When a specific rule matched, its action, protocol, source, destination, and ports are also shown.{{</field>}}
{{<field "Route">}}The matching [route]({{<relref "docs/domain/virtual-networks/routes" >}}) with its destination, next hop, metric, and description. Shows `No route to destination` when no route matches.{{</field>}}
{{<field "Forward">}}Where the packet is forwarded, including the destination type and protocol.{{</field>}}
{{<field "Egress">}}The packet leaving the node, with the interface, source, destination, and protocol.{{</field>}}
{{<field "No Rule">}}No rule matched the packet. Includes a message from the node explaining why.{{</field>}}
{{</fields>}}

A denied ACL, a missing route, and **No Rule** all block the packet.

## Related Tools

The [Trustgrid MCP server]({{<relref "docs/mcp/tools" >}}) can run the same simulation from an AI agent.
