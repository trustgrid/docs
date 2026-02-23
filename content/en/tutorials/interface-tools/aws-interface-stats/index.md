---
title: "AWS Interface ENA Stats"
linkTitle: "AWS Stats"
description: "View AWS Elastic Network Adapter (ENA) statistics to troubleshoot network performance issues on AWS-hosted nodes"
tags: ["aws", "networking", "troubleshooting", "interface"]
categories: ["interface-tools"]
---

{{% pageinfo %}}
The AWS Stats tool displays Elastic Network Adapter (ENA) statistics for AWS-hosted nodes. These statistics help identify when AWS is throttling network traffic due to instance limits being exceeded. Access requires the `service:network-status` permission.
{{% /pageinfo %}}

## Overview

When running Trustgrid nodes on AWS EC2 instances, network performance is subject to instance-type-specific limits for bandwidth, packets per second, and connection tracking. The AWS Interface ENA Stats tool allows you to view counters that indicate when these limits have been exceeded, helping diagnose network performance issues.

{{<alert title="Note" color="info">}}
The **AWS Stats** button is only visible on nodes running on AWS EC2 instances with ENA-enabled network interfaces. If you do not see this button on an AWS node, upgrade the device to a version that includes AWS Interface ENA Stats support.
{{</alert>}}

## Accessing AWS Stats

1. Log in to the Trustgrid portal and navigate to the node you want to inspect.
1. Select **Interfaces** under the **Network** section in the left navigation.
1. In the **Interface Tools** section on the right side, click the **AWS Stats** button.
{{<tgimg src="aws-stats-button.png" width="95%" caption="Interfaces page showing the AWS Stats button in the Interface Tools section">}}
1. The **AWS Interface ENA Stats** dialog will open, displaying the statistics for the selected network interface.
{{<tgimg src="aws-stats-modal.png" width="75%" caption="AWS Interface ENA Stats dialog showing ENA statistics">}}

### Interface Selection

Use the **Network Interface** dropdown at the top of the dialog to select which network interface you want to view statistics for. The dropdown lists all ENA interfaces available on the node (e.g., `ens5`, `eth0`).

Click the refresh button to update the statistics with the latest values from the instance.

## ENA Statistics

The statistics table displays five metrics that track when AWS has throttled traffic due to instance limits being exceeded:

{{<fields>}}
{{<field "Bandwidth In Allowance Exceeded">}}The number of packets queued or dropped because the inbound aggregate bandwidth exceeded the maximum for the instance. This counter increments when incoming traffic surpasses the instance's bandwidth allocation.{{</field>}}
{{<field "Bandwidth Out Allowance Exceeded">}}The number of packets queued or dropped because the outbound aggregate bandwidth exceeded the maximum for the instance. This counter increments when outgoing traffic surpasses the instance's bandwidth allocation.{{</field>}}
{{<field "Connection Tracking Allowance Exceeded">}}The number of packets dropped because connection tracking exceeded the maximum for the instance and new connections could not be established. This occurs when the number of tracked connections exceeds the instance's limit.{{</field>}}
{{<field "Link Local Allowance Exceeded">}}The number of packets dropped because the PPS of the traffic to local proxy services (such as the DNS service or Instance Metadata Service) exceeded the maximum for the network interface.{{</field>}}
{{<field "PPS Allowance Exceeded">}}The number of packets queued or dropped because the bidirectional PPS (packets per second) exceeded the maximum for the instance. This counter increments when packet rate limits are exceeded regardless of bandwidth utilization.{{</field>}}
{{</fields>}}

{{<alert title="Important" color="warning">}}
Non-zero values for any of these statistics indicate that AWS is actively throttling network traffic on this instance. This can cause packet loss, increased latency, and degraded application performance.
{{</alert>}}

## Troubleshooting Use Cases

### When to Use This Tool

The AWS Interface ENA Stats tool is useful when investigating:

- **Unexplained packet loss** - If applications report intermittent connectivity issues or packet drops that cannot be explained by network configuration
- **Bandwidth limitations** - When throughput appears capped below expected levels despite adequate network path capacity
- **High connection rates** - Applications that create many short-lived connections may hit connection tracking limits
- **Metadata service issues** - Problems accessing EC2 instance metadata or DNS resolution failures
- **VPN or tunnel performance** - When VPN throughput is lower than expected on AWS-hosted gateway nodes

### Interpreting the Statistics

| Statistic | Non-Zero Value Indicates | Potential Remediation |
|-----------|-------------------------|----------------------|
| Bandwidth In/Out Allowance Exceeded | Instance bandwidth limit reached | Upgrade to a larger instance type with higher bandwidth allocation |
| Connection Tracking Allowance Exceeded | Too many concurrent connections | Upgrade instance or optimize application connection handling |
| Link Local Allowance Exceeded | High rate of metadata/DNS queries | Implement local caching for metadata and DNS |
| PPS Allowance Exceeded | Packet rate limit exceeded | Upgrade instance or reduce packet rate (e.g., larger packets) |

### AWS Instance Network Specifications

Each EC2 instance type has specific network performance characteristics. Refer to the [AWS EC2 Instance Types documentation](https://aws.amazon.com/ec2/instance-types/) for details on:

- Baseline and burst bandwidth limits
- Packets per second limits  
- Connection tracking limits

When ENA statistics show consistent throttling, consider upgrading to an instance type with higher network allocations that match your workload requirements.

## Additional Resources

- [AWS ENA Driver Documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/enhanced-networking-ena.html)
- [Monitoring Network Performance for EC2 Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-network-performance-ena.html)
- [EC2 Network Performance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-network-bandwidth.html)
