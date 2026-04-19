---
title: "Commands"
linkTitle: "Commands"
aliases:
  - /docs/nodes/agents/commands
description: Configure commands for remote execution on a node or cluster
---

{{% pageinfo %}}
Commands let you define an executable and arguments that can be run on a Trustgrid node. Commands can be configured on an individual node or on a cluster when you want both cluster members to share the same command definition.
{{% /pageinfo %}}

Commands are available for [agent]({{<ref "docs/nodes/agents">}}) and appliance-based nodes. On clusters, commands are managed from the cluster view rather than from an individual member node.

Reading and managing commands requires `node-exec::read` and `node-exec::modify` permissions. Running a command requires `node-exec::compute` permission.

## Node-Level and Cluster-Level Commands

Commands can be configured at either scope:

- **Node-level commands** apply only to the selected node.
- **Cluster-level commands** are shared cluster configuration and should be used when both members need the same command definition, variables, network attachments, and limits.

If a node is a member of a cluster, command settings must be managed from the cluster page. The node view will direct you to the cluster-level command instead of allowing separate per-member edits.

Use node-level commands when the command is specific to a single host. Use cluster-level commands when the command should remain consistent across both members of a cluster.

## Management

Navigate to **Compute > Commands** on a node or cluster.

From the commands table you can:

- Add a command
- Enable or disable a command
- Delete a command
- Import an existing command definition
- Open a command to manage its overview, environment variables, network settings, and resource limits

## Command Fields

{{<fields>}}
{{<field "Name">}}Unique name for the command. The name should only include letters, numbers, and `-`.{{</field>}}
{{<field "Description">}}Optional description to explain what the command is used for.{{</field>}}
{{<field "Status">}}Only enabled commands can be started or run on a schedule.{{</field>}}
{{<field "Execution Type">}}Determines how the command runs:

1. **On Demand** - Starts only when a user manually runs it.
1. **Service** - Starts as a managed long-running service.
1. **Recurring** - Runs on a defined schedule using either a rate expression or a cron expression.

Examples:

| Schedule | Description |
| --- | --- |
| `rate(30 minutes)` | Run every 30 minutes |
| `rate(1 hour)` | Run every hour |
| `0 * * * *` | Run at the top of every hour |

{{</field>}}
{{<field "Schedule">}}Required only for **Recurring** commands. Accepts either a rate expression such as `rate(1 hour)` or a cron expression.{{</field>}}
{{<field "Save Output">}}Persist standard output and standard error to the Trustgrid cloud for later review.

**It is the customer's responsibility to ensure no sensitive information is written to command output before enabling this option.**{{</field>}}
{{<field "Executable">}}The executable to run, such as `/bin/sh`, `/usr/bin/curl`, or a locally installed script path. Variable expansion is supported using `${var_name}` syntax.{{</field>}}
{{<field "Arguments">}}Arguments passed to the executable. Arguments can be separated by spaces. Variable expansion is supported using `${var_name}` syntax.{{</field>}}
{{</fields>}}

## Overview

After creating the command you can configure additional settings on the command itself:

- **Environment Variables** to pass runtime configuration values
- **Network** settings to attach virtual networks or interfaces where needed
- **Resource Limits** to constrain how much CPU or memory the command can consume

These settings are especially useful when a command needs access to virtual network resources or when you want to standardize the runtime environment across multiple nodes.

## Running Commands

- **On Demand** commands are started manually from the command view.
- **Service** commands are appropriate for long-running processes that you want Trustgrid to manage.
- **Recurring** commands are appropriate for periodic jobs such as checks, synchronization, or data collection.

If the executable or arguments include variables in `${var_name}` format, the Portal prompts for values when the command is started.

## Example Use Cases

### Node-Level Examples

- Run a diagnostic script on a single agent to verify local application health.
- Launch a one-off support command on a specific node to collect troubleshooting output.
- Run a recurring inventory collection job on one node that writes results to a local system.

### Cluster-Level Examples

- Run the same health-check command on an HA pair so the command definition stays aligned across both members.
- Define a recurring log collection or metrics export job that should be managed as shared cluster compute configuration.
- Configure a service-style command at the cluster level when both members need the same long-running process definition and supporting variables.

## When to Use Node vs Cluster Scope

Use a **node-level** command when:

- The executable exists only on one node
- The command references host-specific paths or local dependencies
- The task is intended for one system only

Use a **cluster-level** command when:

- Both members should share the same definition
- The command needs the same virtual network attachments, variables, or limits
- You want the command managed as part of the cluster's shared compute configuration
