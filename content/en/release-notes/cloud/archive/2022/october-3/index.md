---
title: "October Third Release"
linktitle: "October - Third Release"
date: 2022-10-28
type: docs
weight: 10
---

## Repository Performance Improvements

This release includes a significant overhaul of the the repository control plane service. This service is responsible for providing upgrade packages and pushing/pulling container images. 

## Data Plane Stats for Private and Hub Gateways

Prior to this release, viewing the Data Plane stats such as tunnel latency and [network hops]({{<ref "/tutorials/gateway-tools/monitoring-network-hops-to-peers">}} would not display. These stats will now be displayed the just like edge and public gateway stats.

## Changing Configured Master (Active) Logging

With this release, if a user changes which node is configured as the master (active) member of the cluster it will be logged on `Changes` for the cluster.

<div style="display: flex; flex-direction: row;">
<div style="width: 25%;">
{{<card header="Changing the configured master (active) node">}}
![set active](set-active.png)
{{</card>}}
</div>
<div style="width: 75%;">
{{<card header="Change entry for active member change">}}
![change list](changes.png)
{{</card>}}
</div>
</div>
