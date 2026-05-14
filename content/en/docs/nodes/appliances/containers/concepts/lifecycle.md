---
title: "Container Lifecycle"
description: How Trustgrid containers start, stop, restart, and run on a schedule.
---

A container's **Execution Type** controls when and how often it runs. Pick the one that matches the workload.

## Execution types

### Service

The container runs continuously and is kept running.

- Starts automatically when the appliance comes up.
- If the container exits — clean or not — it's restarted automatically. A container with a config problem that exits immediately will keep being restarted; to break the cycle, set **Status** to `Disabled` while you fix it, or switch **Execution Type** to `On Demand` for debugging.
- A [Health Check]({{<relref "../#health-check">}}) reports the container as healthy or unhealthy in the portal. Restarts are driven by the container exiting; the health-check result is reporting only.
- **Use for:** web services, background agents, anything that should always be up.

### Recurring

The container runs on a schedule, like a cron job.

- The **Schedule** field accepts either a simple rate or a cron expression.
- Each run is independent — a fresh container starts, runs to completion, then waits for the next time.
- If a previous run is still going when the next one is due, the new run is skipped.
- **Use for:** scheduled batch jobs, periodic data pulls, cleanup scripts, reports.

| Rate | Description |
|---|---|
| `rate(30 minutes)` | Run every 30 minutes |
| `rate(1 hour)` | Run every hour |
| `rate(1 day)` | Run once a day |

For cron expressions, [crontab.guru](https://crontab.guru/examples.html) is a useful reference.

### On Demand

The container only runs when you start it manually from the portal.

- Does not start with the appliance.
- Does not restart after it stops.
- **Use for:** diagnostic tools, one-time tasks, testing a new container before switching it to Service.

## Status vs. State

The container list shows two columns that look similar but mean different things:

- **Status** — whether the container is *configured* to run. `Enabled` means it will be run; `Disabled` means it will be left alone.
- **State** — what's actually happening right now, such as `Running` or `Stopped`.

Common combinations:

| Status | State | What's going on |
|---|---|---|
| `Enabled` | `Running` | Healthy — a Service container that's up. |
| `Enabled` | `Stopped` | Configured to run but isn't running right now — between restart attempts, or unable to start. |
| `Disabled` | `Stopped` | Configured but won't be run. Useful when you're staging changes. |

## Manual Start and Stop

Each container has **Start** and **Stop** buttons on its detail page at node scope:

- **Stop** halts the container. There's a wait of up to **Stop Time** (default 30 seconds) for a clean shutdown before it's forced.
- **Start** launches the container. For an On Demand container this is the only way to run it. For Service or Recurring containers, Start is useful when you want to run it immediately rather than waiting for the next scheduled run.

See [Container Tools]({{<relref "../tools">}}) for the full set of per-container actions.

## Related

- [Container Tools]({{<relref "../tools">}}) — Start, Stop, Logs, Terminal
- [Health Check]({{<relref "../#health-check">}}) — field reference
