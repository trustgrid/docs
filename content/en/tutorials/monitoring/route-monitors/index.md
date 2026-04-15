---
title: "Route Monitor Best Practices"
linkTitle: Route Monitors
description: "Recommendations, limitations, and expected behavior for Trustgrid route monitors"
type: docs
---

{{% pageinfo %}}
Route monitors help determine whether a configured route should be considered available. This page focuses on monitor behavior, limitations, and how to avoid false failures.
{{% /pageinfo %}}

## What Route Monitors Actually Test

Route monitors run from the node or cluster member that owns the route and send traffic over the virtual network using that route.

- **ICMP monitors** send echo requests to the configured destination IP.
- **TCP monitors** send a TCP SYN to the configured destination IP and port.

This means a route monitor validates more than simple route existence. It validates that the destination can be reached over the Trustgrid path using the configured protocol and that the reply arrives within the allowed latency.

## Prerequisites

- The route must already exist.
- The node should have a virtual management IP on the virtual network.
- The selected destination must be reachable through the route.
- Firewalls and host policy must allow the selected monitor traffic and its replies.
- For TCP monitors, the destination port must represent a service that can safely tolerate repeated connection attempts.

## Failure and Recovery Behavior

Route monitors do not fail or recover on a single result unless the failure count is set to `1`.

- A monitor fails after it accumulates the configured number of failed checks.
- A successful check reduces the accumulated failure count by one.
- A monitor recovers only after enough successful checks reduce the failure count back to zero.

This behavior helps smooth out brief packet loss or transient latency spikes, but a very aggressive interval and low failure count can still create route flapping.

## Choosing Between ICMP and TCP

### Prefer ICMP when

- the destination host reliably responds to ping
- you want a lightweight reachability check
- you do not want to interact with an application socket or server listener

### Use TCP when

- ICMP is blocked
- you need to verify that a specific TCP service is accepting connections
- the destination service and firewall policy are designed to tolerate frequent connection attempts

## Known Limitations and Caveats

### TCP reset responses count as failures

If a TCP monitor receives a reset from the destination, the monitor treats that as a failure. Do not point a TCP monitor at a port that intentionally resets connections unless failure is what you want to detect.

### TCP monitors can bother fragile services

TCP monitors repeatedly initiate new connections. Some services or security devices may log those attempts, throttle them, or behave poorly when many short-lived handshakes occur. Be careful using TCP monitors on applications that are sensitive to connection churn.

### Latency thresholds can create false failures

If **Max Latency** is set too low, monitors can fail even though the route is otherwise usable. This is especially common on long-distance or internet-routed paths with variable latency.

### Route monitor success is not an application health check

A passing monitor only proves the selected IP and protocol are reachable within the configured threshold. It does not prove the entire application behind that route is healthy.

## Recommendations

- Start with **ICMP** unless there is a clear reason to validate a TCP port.
- Use a destination IP that is stable and expected to answer consistently.
- For TCP, choose a port on a service that is expected to remain up and can tolerate monitor traffic.
- Do not set the interval lower than necessary.
- Avoid a failure count of `1` unless you explicitly want very fast failover and accept the risk of false positives.
- Be conservative with max latency thresholds on links that can have jitter.

## Route Failover Considerations

If multiple routes exist for the same destination CIDR, route monitors influence which routes are considered available. A lower metric route can be bypassed when its monitor fails, and it can become preferred again once its monitor recovers.

That automatic return is useful for healthy failback, but if the destination is unstable it can create repeated route changes. In those situations:

- increase the failure count
- increase the interval
- relax the latency threshold
- or temporarily force traffic to a backup route by adjusting route metrics

## Domain Routes vs Node or Cluster Routes

Route monitor behavior is the same after configuration reaches the node, but the workflow is different:

- **Node or cluster routes** are updated directly on that route.
- **Domain routes** must be saved and then applied through the domain change workflow before the monitor configuration is published.

For configuration steps, see:

- [Static Routing]({{<relref "docs/nodes/appliances/vpn/static-routing">}})
- [Virtual Network Routes]({{<relref "docs/domain/virtual-networks/routes">}})
