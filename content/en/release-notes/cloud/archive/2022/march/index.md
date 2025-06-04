---
title: "March"
linktitle: "March"
date: 2022-10-03
type: docs
weight: 3
---

## Data Plane Panel Improvements

In the February 2022 release we introduced the [Data Plane status panel]({{<ref "february">}}) for investigating connections between edge and gateway nodes.  With this release we’ve added some tools to this panel. The new Ping and Test Performance tools are linked in the far right column of each gateway’s row.

<img src="ping-test-performance.png" width="400px" />

### Gateway Ping

<img src="gateway-ping-modal.png" width="400px" />

The gateway ping tool will launch a ping inside the data plane tunnel between the node and the gateway, similar to ICMP ping.  This can provide a more realtime view of latency between node and gateway compared to the RTT column (only loaded on refresh) and the historical data which provides a minute by minute view.

![Ping Output](ping-output.png)

### Gateway Performance Tool

The gateway performance tool attempts to estimate the available upload bandwidth between a node and gateway. 

<img src="gateway-performance-modal.png" width="400px" />

By default this tool will attempt to send 10MB of data to the destination device. This then used to calculate the available bandwidth.  Sending more data provides a more accurate estimate but be aware that other devices using the same internet circuit may experience degraded performance while the test is running. 

![Gateway Performance Output](performance-output.png)

The periods are posted for about every 1MB of data sent. If you see exclamation marks this can indicate lost packets.

## Tool Improvements

### MTR Improvements

MTR, which was [recently added to the portal]({{<ref "february">}}), can be powerful troubleshooting tool, but it has many options which can be tricky for newer users.  To simplify the use we’ve added a new user interface that provides graphical options for selecting common parameters.

<img src="mtr-modal.png" width="400px" />

## AWS Instances Metadata

When the Trustgrid software is deployed in AWS we can gather additional metadata about the underlying instance. With this release we have added the availability zone, instances type and instance ID to the Info panel available on every node. 

![AWS Metadata](node-visor.png)

## Timestamps for Virtual Network Changes

Before changes are committed to a virtual network setting, such as the virtual network route table, they are listed in the Changes panel for approval. This made it easier to see if other users had outstanding changes that might be in conflict. However, it wasn’t clear if the other user just made that change or if it had been left accidentally. With this release we’ve added a timestamp that shows when the user saved the change.

![Change Example](modified-route.png)





