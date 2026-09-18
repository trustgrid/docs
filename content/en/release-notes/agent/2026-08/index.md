---
title: August 2026 Agent Release Notes
linkTitle: August 2026
type: docs
date: 2026-08-31
description: "August release of the Trustgrid Agent"
---
{{<agent-release package-version="0.2.20260825-2536" release="a-0.5.0">}}

## Connector Source Blocks

Agent-based nodes now support restricting connector access by source CIDR. Set the **Source Block** field on a connector to a comma separated list of network CIDRs, and only those source IPs will be allowed to connect. This matches the behavior already available on appliance-based nodes. See [Connectors]({{<relref "docs/nodes/shared/connectors" >}}) for details.

## Local Service Auto-Discovery

The agent can now scan its host for running services and report the results to the Trustgrid Portal. This makes it easier to find candidates when setting up connectors, without needing console access to the machine.

## New Data Plane Tools

Three tools were added to the agent [Data Plane]({{<relref "docs/nodes/agents/data-plane" >}}) panel:

- `mtr` for hop-by-hop path analysis from the agent, matching the tool already available on appliance-based nodes.
- A listening ports service that lists every port the host is listening on along with its protocol. Useful when troubleshooting connectors.
- A `gateway-perf` service for measuring throughput between a gateway and an agent.

## Dependency Updates

This release updates third-party dependencies to address published vulnerabilities. None of these vulnerabilities are exploitable in the Trustgrid Agent. The Rust toolchain used to build the agent was also updated.

## Other Fixes and Improvements

- Resolved an issue where an agent interface IP address was not updating, even after a restart.
- Corrected handling of gateway routes on the agent.
- Agents now support gateway DNS host names for UDP tunnels.
- Resolved an issue where the `trustgrid/agent` container did not shut down in response to `CTRL+C`.
- Resolved an issue where the systemd unit did not invoke the agent executable correctly.
- Resolved an issue where the Sniff Traffic tool failed on filters containing single quotes.
- Improved the VPN error message shown when a multi-tenant gateway is unreachable.
- Agents running under Docker no longer show the View Virtual Routes tool or the Static Routes navigation, neither of which applies to a container deployment.
- Agents with no domain information are now handled correctly.
- The `tg-agent diag` output now includes the agent's public IP. See [Agent Diagnostic]({{<relref "tutorials/management-tasks/agent-mgmt/agent-diag" >}}).
