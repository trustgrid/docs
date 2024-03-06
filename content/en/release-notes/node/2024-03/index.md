---
title: March 2024 Appliance Release Notes
linkTitle: March 2024
type: docs
date: 2024-03-05
description: "Release notes for the March 2024 Trustgrid Appliance release focused"
---
{{< node-release package-version="1.5.20240306-1999" core-version="20240306-161006.62ba5a5" release="n-2.19.0" >}}

## Dynamic Access Policy Evaluation
Previously [Access Policies]({{<relref "/docs/domain/virtual-networks/access-policy">}}) were only evaluated at the time a flow was created. With this release, access policies will be re-evaluated in realtime for active flows. This allows dynamically changing policies to immediately take effect without requiring existing flows to timeout or reset.

## System Resources Overview
Appliances running this release will now share with the control plane the number of CPU cores, total system memory (RAM), and total disk size. This information will be displayed on the [Overview graph]({{<relref "/docs/nodes/appliances/overview">}}). This will provide additional context to the percentage utilization statistics the node provides in the graph.
{{<tgimg src="overview-resources.png" width="60%">}}


## BGP Status Tool Improvement
With this release nodes running BGP will now report [rejected routes]({{<relref "/docs/nodes/appliances/bgp#rejected-routes">}}) when a peer advertises a route that doesn't match the import policy. This makes it easier to determine if a route is not being advertised by the peer or if the import policy is misconfigured. 

## Gateway Path Improvements
This release allows for disabling additional [gateway paths]({{<relref "/docs/nodes/appliances/gateway/gateway-client#gateway-paths">}}) rather than requiring them to be completely removed.   It also resolves an issue where the Ping and Test Performance tools on the [Data Plane panel]({{<relref "/docs/nodes/appliances/data-plane">}}) would not actually utilize the path selected for the test.

## Other Fixes and Changes
- Resolves an issue with Wireguard ZTNA access not releasing ephemeral ports properly. 
- Resolves an issue with the Virtual Sniff tool not showing Wireguard ZTNA traffic.
- Resolves an issue where VPN attachments to VLANs were not completely removed without a restart.
- The Trustgrid service will no longer bring up interfaces that are not enabled and configured.
- The [data plane status SNMP OID .2.4.0]({{<relref "docs/nodes/appliances/snmp/oids">}}) now reflects a degraded status if some but not all of the data plane peers are connected. It also now factors in additional gateway paths.