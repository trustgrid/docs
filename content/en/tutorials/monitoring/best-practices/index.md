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
- **Informational Channel:** For alarms that are useful for awareness but rarely require immediate action, such as DNS or Repository health issues. These can be sent to a general monitoring mailbox or dashboard.


### Recommended Event Types for Critical Systems (e.g., Public Gateways)

For public gateways and other critical nodes, monitor the following event types:

- `Node Connect/Disconnect`
- `Cluster Healthy/Unhealthy`
- `Cluster Failover`
- `Metric Threshold Violation`
- `Network Error`

These events help ensure you are alerted to outages, failovers, and performance issues that could impact connectivity or service availability.

{{<tgimg src="alarm-filter-critical-example.png" width="50%" caption="Example alarm filter for critical systems such as public gateways." >}}

### Recommended Monitoring for All Nodes

For every node, consider the following event types and their recommended alerting priority:

| Event Type                  | Immediate Attention | Monitor Only |
|-----------------------------|:------------------:|:------------:|
| Node Connect/Disconnect     |         ✔️         |              |
| Cluster Healthy/Unhealthy   |         ✔️         |              |
| Cluster Failover            |         ✔️         |              |
| Network Error               |         ✔️         |              |
| Network Route Error         |         ✔️         |              |
| Metric Threshold Violation  |                    |      ✔️      |
| DNS Health                  |                    |      ✔️      |
| Repository Health           |                    |      ✔️      |

This approach ensures that issues impacting traffic or node health are surfaced immediately, while less urgent issues are still tracked for later review.

{{<tgimg src="alarm-filters-example.png" width="50%" caption="Example of organizing alarm filters and channels in the Trustgrid portal." >}}

For configuration details, see the [Alarm Filters]({{<relref "docs/alarms/alarm-filters" >}}) documentation.

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