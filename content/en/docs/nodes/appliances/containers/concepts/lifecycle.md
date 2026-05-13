---
title: "Container Lifecycle"
description: How Trustgrid containers start, stop, restart, and run on a schedule.
weight: 20
---

A container's **Execution Type** controls when and how often the node runs it. Pick the one that matches the workload — they have different restart and persistence semantics.

## Execution types

### Service

The node treats the container as a long-running daemon.

- Starts automatically when the node software starts.
- If the container exits — for any reason, including a clean exit code — the node restarts it. The restart loop has no backoff; expect immediate respawn.
- A configured [Health Check]({{<ref "../#health-check">}}) that fails counts as an exit and triggers a restart.
- **Use for:** web services, agents, anything that should always be up.

### Recurring

The node starts the container on a schedule.

- The **Schedule** field accepts either a rate or a cron expression.
- A rate looks like `rate(30 minutes)`, `rate(1 hour)`, `rate(1 day)`.
- A cron expression follows the standard 5-field format. [crontab.guru](https://crontab.guru/examples.html) is a useful reference.
- Each invocation is independent — the node starts a fresh container, runs to exit, then waits for the next scheduled time.
- If the previous invocation is still running when the next schedule fires, the new invocation is skipped.
- **Use for:** batch jobs, periodic data pulls, cleanup scripts, scheduled reports.

| Rate | Description |
|---|---|
| `rate(30 minutes)` | Run every 30 minutes |
| `rate(1 hour)` | Run every hour |
| `rate(1 day)` | Run every day |

### On Demand

The container only runs when explicitly started from the portal or API.

- Does not auto-start with the node.
- Does not restart on exit.
- Each Start is a single invocation; the container runs until it exits or you stop it.
- **Use for:** diagnostic workloads, one-shot tasks, testing a container's configuration before promoting it to a Service.

## Status vs. State

These two columns on the container list mean different things and they can disagree:

- **Status** — the configured state of the record. `Enabled` means the node should run it (per its Execution Type); `Disabled` means the node should leave it alone.
- **State** — the live runtime state on a node. Values include `Running`, `Stopped`, `Pulling`, `Restarting`.

A few combinations and what they mean:

| Status | State | Meaning |
|---|---|---|
| `Enabled` | `Running` | Healthy steady state for a Service. |
| `Enabled` | `Stopped` | Either between restarts, mid-pull, or unable to start. Check the node log and [troubleshooting]({{<ref "../troubleshooting">}}). |
| `Enabled` | `Pulling` | Node is downloading the image. Transitions to `Running` on success or `Stopped` on failure. |
| `Disabled` | `Stopped` | Container is configured but the node won't try to run it. Typical for staging changes you don't want to take effect yet. |
| `Disabled` | `Running` | Transient; the next config evaluation should stop it. If it persists, the disable didn't propagate. |

## Manual control

The **Start** and **Stop** buttons on the per-container detail view at node scope override the Execution Type momentarily:

- **Stop** halts the container. For a Service, the node respects the configured **Stop Time** (default 30 seconds) before sending SIGKILL. Because Service containers are configured to auto-restart, a manual Stop will only hold until the next config evaluation — to keep it down permanently, set Status to `Disabled`.
- **Start** starts the container. For On Demand containers this is the only way to launch one. For Service or Recurring containers it's useful when you want to verify a new image immediately rather than wait for the next scheduled run or for the node to detect the config change.

See [Container Tools]({{<ref "../tools">}}) for the full UI surface.

## Restart semantics in detail

Service containers restart unconditionally on exit. There is no `max_restarts` or exponential backoff configuration today — if a container exits in a tight loop (e.g., a config error causes immediate exit), the node will restart it as fast as the runtime can spawn a new one. To avoid hammering, either:

- Add a Health Check with a **Start Period** long enough to let your application stabilize, so the node only considers the container failed after a real health check window has elapsed.
- Set Execution Type to On Demand while you debug — the container will only run when you Start it manually.
- Set Status to Disabled while you fix the underlying config.

Recurring containers do **not** restart on exit. They run, exit, and wait for the next scheduled invocation. A failure during one invocation is independent of the next.

On Demand containers do not restart at all — each Start is a single invocation.

## Related

- [Container Tools]({{<ref "../tools">}}) — Start, Stop, Logs, Terminal
- [Container troubleshooting]({{<ref "../troubleshooting">}}) — when containers won't start, won't pull, or restart-loop
- [Health Check]({{<ref "../#health-check">}}) — field reference
