---
linkTitle: Best Practices
title: "Best Practices for Monitoring and Alerting"
description: "Learn best practices for monitoring and alerting to ensure system reliability and performance."

type: docs
---


## Monitoring with Trustgrid Alarm Filters

Alarm filters in Trustgrid help you ensure the right people are notified about the right events. Instead of documenting configuration steps, this section focuses on what you should monitor and how to organize your alerting for maximum effectiveness.



### Recommended Alarm Channels

- **Critical Channel:** For alarms that require immediate action, such as issues with public gateways or core infrastructure. This channel should notify on-call or primary support staff.
- **Standard Channel:** For important but less urgent alarms, such as edge node disconnects or non-critical service interruptions. This channel can notify a broader group or secondary contacts.
- **Informational/Follow up Channel:** For alarms that are useful for awareness but rarely require immediate action, such as DNS or Repository health issues. These can be sent to a general monitoring mailbox or dashboard.

### Criteria Recommendations

- **Use Tag Filters for Production Systems:**
	- Limit alarm filters to nodes with a `prod_status` or `ProdStatus` tag. This ensures alarms are only sent for production systems, reducing noise from test or development nodes. See the [Production Status Tags documentation]({{<relref "docs/nodes/shared/tags/prod-status-tag" >}}) for more on the recommended `prod_status` tag.
- **Set Severity Threshold to INFO:**
	- Set the Severity Threshold to `INFO` for alarm filters. Many resolve messages are sent with this threshold, so using a higher threshold would prevent those messages from being sent to the channel and could result in unresolved alarms.


### Recommended Event Types for Critical Systems (e.g., Public Gateways)

For public gateways and other critical nodes, monitor the following event types:

- `Node Connect/Disconnect`
- `Cluster Healthy/Unhealthy`
- `Cluster Failover`
- `Metric Threshold Violation`
- `Network Error`
- `Connection Flapping`
- `Network Route Error`
- `All Peers Disconnected` (for gateways) or `All Gateways Disconnected` (for edge nodes)

These events help ensure you are alerted to outages, failovers, and performance issues that could impact connectivity or service availability.

Either select each node individually or use tag filters to include all critical systems in this alarm filter.

{{<tgimg src="alarm-filter-critical-example.png" width="80%" caption="Example alarm filter for critical systems such as public gateways." >}}

### Recommended Monitoring for All Nodes

For every node, consider the following event types and their recommended alerting priority:

| Event Type                  | Immediate Attention | Monitor for Follow Up |
|-----------------------------|:------------------:|:------------:|
| Node Connect/Disconnect     |         ✔️         |              |
| Cluster Healthy/Unhealthy   |         ✔️         |              |
| Cluster Failover            |         ✔️         |              |
| Network Error               |         ✔️         |              |
| Network Route Error         |         ✔️         |              |
| All Peers/Gateways Disconnected |     ✔️         |              |
| Connection Flapping         |          ✔️        |              |
| Metric Threshold Violation¹  |                    |      ✔️      |
| DNS Resolution              |                    |      ✔️      |
| Repo Connectivity           |                    |      ✔️      |

The Immediate Attention alerts should be sent to a channel that will get immediate visibility. The monitoring events should be sent somewhere like a ticketing system (via email or webhook) or Slack/Teams channel. 

This approach ensures that issues impacting traffic or node health are surfaced immediately, while less urgent issues are still tracked for later review.

{{<tgimg src="alarm-filters-immediate-example.png" width="80%" caption="Example of Immediate Attention filter" >}}

{{<tgimg src="alarm-filters-follow-up-example.png" width="80%" caption="Example of Follow Up filter" >}}

> ¹ Metric Threshold Violations can be noisy depending on your thresholds. Consider monitoring these for follow up rather than immediate attention unless you have specific thresholds that indicate critical issues.

For configuration details, see the [Alarm Filters]({{<relref "docs/alarms/alarm-filters" >}}) documentation.

---

## Network Threshold Alerts for Gateways

For any node acting as a [public or hub gateway]({{<relref "docs/nodes/appliances/gateway/gateway-server" >}}), it is recommended to configure a [Domain-level Network Threshold]({{<relref "docs/domain/thresholds#network-thresholds" >}}) alert monitoring latency. This provides early warning of network degradation that could impact all connected edge nodes.

### Recommended Configuration

- **Telemetry:** Latency
- **Threshold:** 500 ms
- **Duration:** 5 minutes or greater
- **Target:** Each public or hub gateway node

A 500 ms latency threshold over a 5-minute duration helps identify sustained network performance issues without triggering on brief, transient spikes. Because these gateways serve as connection points for many edge nodes, elevated latency on a gateway can have a widespread impact across your environment.

Configuring this at the [domain level]({{<relref "docs/domain/thresholds" >}}) ensures the threshold applies to all nodes, but you can also [override it per node]({{<relref "docs/nodes/appliances/thresholds" >}}) if specific gateways require different values. When the threshold is exceeded, a `Metric Threshold Violation` [event]({{<relref "docs/alarms/event-types" >}}) is generated and can be routed through your [alarm filters]({{<relref "docs/alarms/alarm-filters" >}}).

---

## Monitoring Individual Nodes with SNMP


SNMP (Simple Network Management Protocol) allows you to pull statistics about individual Trustgrid nodes into third-party network monitoring tools. This enables integration with your existing Network Management System (NMS) or observability platforms for centralized monitoring.

Using SNMP is recommended if your team already uses an NMS for troubleshooting other systems, as it allows you to view Trustgrid node health and metrics alongside the rest of your infrastructure in a single pane of glass.

SNMP monitoring should be used in addition to—not as a replacement for—Trustgrid alarms. If you use both, be aware that some event types (such as Metric Threshold Violations) may be duplicated in both systems. You can choose to limit these notifications to one system or the other to avoid unnecessary noise.

Polling can be performed against either a node's interface IP address or its virtual management IP. When your SNMP collector is in a different data center or location than the edge node, use the virtual management IP to monitor across the VPN data plane. This approach is also useful for monitoring the health of the data plane connection itself.

For supported OIDs, setup, and more details, see the [SNMP documentation]({{<relref "docs/nodes/appliances/snmp" >}}).

---

## Using OpenTelemetry to Export Data to External Systems

Trustgrid supports exporting metrics, events, configuration changes, and audit entries to an OpenTelemetry (OTel) collector. This enables you to parse, analyze, or display Trustgrid telemetry in your external observability platforms.

Similar to SNMP, OTel export allows you to achieve a single pane of glass view by combining Trustgrid data with telemetry from the applications that rely on Trustgrid connectivity. This unified approach helps correlate network events with application performance and security.

{{<tgimg src="otel-integration-example.png" width="50%" caption="Configuring OpenTelemetry export in Trustgrid." >}}

For more on OTel integration, see the [Observability guide]({{<relref "docs/observability" >}}).