---
title: "Container Tools"
description: Operate running containers — start, stop, view logs, and open a shell from the portal.
weight: 50
---

A container is configured at the cluster level (or node level on a standalone node) — that's where you describe what it should do. The buttons described on this page act on the running container on a specific node — Start, Stop, view live logs, open a shell.

## Where to find them

Navigate to a node, then **Compute → Container Management → Containers**, and click the container's name.

{{<alert color="info">}}
If the node is a member of a cluster, the portal shows a banner that the configuration fields can only be changed at the cluster level. The **Start / Stop / Logs / Terminal** buttons described below still work — they act on the live container on this specific node.
{{</alert>}}

{{<tgimg src="container-action-bar.png" caption="Per-container action bar (Start, Stop, Logs, Terminal) on the node-scoped detail view" width="90%">}}

## State

The container list at node scope shows two columns side by side:

- **Status** — `Enabled` or `Disabled`. Whether the node is *supposed* to run it.
- **State** — what's actually happening right now: `Running`, `Stopped`, `Restarting`, or `Pulling` (downloading the image).

A container can show `Enabled` and `Stopped` at the same time — usually because the image couldn't be pulled, or because it's between restart attempts. See [troubleshooting]({{<ref "../troubleshooting">}}) if it stays that way.

{{<tgimg src="node-container-list.png" caption="Container list at node scope, showing runtime State alongside configured Status" width="90%">}}

## Start

Starts the container immediately on the node.

This is the only way to launch an **On Demand** container. For **Service** containers it's useful when you've stopped one manually and want it back up without waiting.

Requires `node-exec::compute` permission.

## Stop

Stops the container. For Service containers the node waits up to **Stop Time** (default 30 seconds) for the container to shut down cleanly, then forces it.

A manual Stop doesn't keep the container down — Service containers will restart on the next config change or node reboot. To keep it down for longer, change its Status to `Disabled` from the cluster-level configuration.

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
This viewer only shows logs from the container's current run — if the container restarts, the history is gone. To keep logs across restarts, enable **Save Output** on the container's Overview. Saved logs are visible in **Observability** afterward. (Note: anything sensitive your container prints will be saved too.)
{{</alert>}}

## Terminal

Opens a new browser tab with an interactive shell running inside the container — useful for debugging, poking around the filesystem, or running ad-hoc commands.

The shell runs as the configured user. If you didn't set one, it runs as whatever user the image was built to run as.

The Terminal button is disabled if the container isn't currently `Running`.

{{<alert color="info">}}
If the container's image doesn't include a shell (some minimal images don't), the Terminal won't be able to open. You can set the **Command** field on the Overview to launch a temporary diagnostic shell instead.
{{</alert>}}

Requires `node-exec::compute` permission.

## History

The **History** section in the left sidebar of a running container shows:

- **Connections** — recent inbound and outbound network connections for this container. Useful to verify that port mappings and virtual networks are wired up right.

## Related

- [Container troubleshooting]({{<ref "../troubleshooting">}}) — what to check when Start does nothing, the logs are empty, or the Terminal refuses to open.
- [Container networking]({{<ref "../concepts/networking">}}) — how port mappings, virtual networks, and DNS resolution work.
