---
title: "Container Tools"
description: Operate running containers — start, stop, view logs, and open a shell from the portal.
weight: 50
---

A container is configured at the cluster level (or node level on a standalone appliance) — that's where you describe what it should do. The buttons described on this page act on the running container on a specific node — Start, Stop, view live logs, open a shell.

## Where to find them

Navigate to a node, then **Compute → Container Management → Containers**, and click the container's name.

{{<alert color="info">}}
If the node is a member of a cluster, the portal shows a banner that the configuration fields can only be changed at the cluster level. The **Start / Stop / Logs / Terminal** buttons described below still work — they act on the live container at node scope.
{{</alert>}}

{{<tgimg src="container-action-bar.png" caption="Per-container action bar (Start, Stop, Logs, Terminal) on the node-scoped detail view" width="90%">}}

## Status, State, and Health

The container list at node scope shows three columns side by side:

- **Status** — `Enabled` or `Disabled`. Whether the container is *configured* to run.
- **State** — what's actually happening right now: `Running`, `Initializing` (image is downloading or the container is starting up), or `Stopped`.
- **Health** — `Healthy` or `Unhealthy`, populated by the configured [Health Check]({{<ref "../#health-check">}}). Blank if no health check is set.

A container can show `Enabled` and `Stopped` at the same time — usually because the image couldn't be pulled, or because it's between restart attempts.

{{<tgimg src="node-container-list.png" caption="Container list at node scope, showing runtime State alongside configured Status" width="90%">}}

## Start

Starts the container immediately.

This is the only way to launch an **On Demand** container. For **Service** containers it's useful when you've stopped one manually and want it back up without waiting.

Requires `node-exec::compute` permission.

## Stop

Stops the container. For Service containers there's a wait of up to **Stop Time** (default 30 seconds) for the container to shut down cleanly, then it's forced.

A manual Stop on a Service container is temporary — the container is brought back up on the next configuration sync. To keep it down, change its **Status** to `Disabled` from the cluster-level configuration.

Requires `node-exec::compute` permission.

## Logs

Opens a new browser tab that streams the container's log output as it happens.

{{<tgimg src="logs-modal.png" caption="Container Logs options" width="60%">}}

{{<fields>}}
{{<field "Follow Log Output">}}When **True**, the viewer stays open and prints new lines as they appear. When **False**, the viewer prints recent lines and disconnects.{{</field>}}
{{<field "Number of Lines">}}How many recent lines to show before starting to follow. Default is 100.{{</field>}}
{{</fields>}}

A **CONNECTED** badge in the corner means the stream is live. Click **Terminate** to close it.

{{<tgimg src="logs-stream.png" caption="Streaming log viewer with the CONNECTED indicator and Terminate button" width="65%">}}

{{<alert color="warning">}}
This viewer only shows logs from the container's current run — if the container restarts, the history is gone. To keep logs across restarts, enable **Save Output** on the container's Overview. Saved logs are downloadable from **Compute → Output Artifacts** at node scope. (Note: anything sensitive your container prints will be saved too.)
{{</alert>}}

## Terminal

Opens a new browser tab with an interactive shell running inside the container — useful for debugging, poking around the filesystem, or running ad-hoc commands.

{{<tgimg src="terminal.png" caption="Terminal session inside an Alpine container" width="65%">}}

The shell runs as the configured user. If you didn't set one, it runs as whatever user the image was built to run as. The Terminal button is disabled if the container isn't currently `Running`.

{{<alert color="info">}}
The Terminal needs a shell to exist inside the image. Most images include one, but some intentionally don't. If the Terminal won't open, check that the image has `/bin/sh` or similar.
{{</alert>}}

Requires `node-exec::compute` permission.

## History

The **History** section in the left sidebar of a running container shows:

- **Connections** — recent inbound and outbound network connections for this container.

{{<tgimg src="connections.png" caption="Active and Completed connections for a container" width="90%">}}

## Related

- [Container networking]({{<ref "../concepts/networking">}}) — how port mappings, virtual networks, and DNS resolution work.
