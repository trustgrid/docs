---
title: "Container Lifecycle"
description: How Trustgrid containers start, stop, restart, and run on a schedule.
weight: 20
---

A container's **Execution Type** controls when and how often the node runs it. Pick the one that matches the workload.

## Execution types

### Service

The container runs continuously and the node keeps it running.

- Starts automatically with the node.
- If the container stops for any reason, the node starts it again right away.
- A failing [Health Check]({{<ref "../#health-check">}}) counts as a stop and triggers a restart.
- **Use for:** web services, background agents, anything that should always be up.

### Recurring

The container runs on a schedule, like a cron job.

- The **Schedule** field accepts either a simple rate or a cron expression.
- Each run is independent — the node starts a fresh container, lets it finish, then waits for the next time.
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

- Does not start with the node.
- Does not restart after it stops.
- **Use for:** diagnostic tools, one-time tasks, testing a new container before switching it to Service.

## Status vs. State

The container list shows two columns that look similar but mean different things:

- **Status** — whether the container is *configured* to run. `Enabled` means the node will run it; `Disabled` means the node will leave it alone.
- **State** — what's actually happening on the node right now: `Running`, `Stopped`, or `Pulling` (downloading the image).

Common combinations:

| Status | State | What's going on |
|---|---|---|
| `Enabled` | `Running` | Healthy — a Service container that's up. |
| `Enabled` | `Pulling` | Node is downloading the image. Will become `Running` if successful. |
| `Enabled` | `Stopped` | The node is trying to run it but it isn't running — either between restart attempts or unable to start. |
| `Disabled` | `Stopped` | Configured but the node won't try to run it. Useful when you're staging changes. |

## Manual Start and Stop

Each container has **Start** and **Stop** buttons on its detail page at node scope:

- **Stop** halts the container. The node waits up to **Stop Time** (default 30 seconds) for a clean shutdown before forcing it. Note: a manual Stop on a Service container is just an exit, so the node will start it again — set Status to `Disabled` if you want it to stay down.
- **Start** launches the container. This is the only way to launch an On Demand container. For Service or Recurring containers, it's useful when you want to test a new image immediately.

See [Container Tools]({{<ref "../tools">}}) for the full set of per-container actions.

## When a Service container restarts in a tight loop

Service containers restart immediately when they stop. If a container has a config problem and exits as soon as it starts, the node will keep restarting it as fast as it can. To break the loop:

- Add a [Health Check]({{<ref "../#health-check">}}) with a **Start Period** long enough for your application to come up, so the node gives it time before considering it failed.
- Change Execution Type to **On Demand** while you debug, then change it back when fixed.
- Set Status to **Disabled** while you fix the underlying problem.

Recurring and On Demand containers do not restart on exit, so they don't have this problem.

## Related

- [Container Tools]({{<ref "../tools">}}) — Start, Stop, Logs, Terminal
- [Health Check]({{<ref "../#health-check">}}) — field reference
