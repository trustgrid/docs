---
title: "Speed Test"
linkTitle: "Speed Test"
description: "Measure upload and download throughput between a node and the Trustgrid cloud over a specific interface"
---

{{% pageinfo %}}
The Speed Test tool measures upload and download throughput between the node and the Trustgrid cloud over a selected local interface. Because it runs directly from the node against the same control-plane gatekeeper endpoint the node normally uses for management traffic, it's useful for confirming real-world achievable bandwidth on a given interface, independent of any VPN or virtual network overhead.
{{% /pageinfo %}}

## Usage

1. Login to the Trustgrid portal and navigate to the Node you want to test.
1. Select **Interfaces** under the **Network** section, then use the interface dropdown to select the interface you want to test.
1. In the **Interface Tools** section, click the **Speed Test** button. {{<tgimg src="/tutorials/interface-tools/tcp-port-test/network-tools.png" width="90%" caption="Selecting Speed Test">}}
1. Click **GO** to start the test.
1. The test first measures download throughput, then upload throughput, animating a gauge for each phase as the measured speed comes in.
1. When the test completes, the results panel shows the peak download and upload speeds achieved. Click **GO** again to re-run the test.

## How It Works

Speed Test transfers data between the node and Trustgrid's cloud gatekeeper - the node's control-plane endpoint - rather than an arbitrary internet target. The download and upload legs run one after the other, each over the interface selected on the Interfaces page, so the result reflects that interface's actual path to the Trustgrid cloud rather than aggregate throughput across all interfaces.
