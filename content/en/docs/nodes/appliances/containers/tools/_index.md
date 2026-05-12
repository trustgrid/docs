---
title: "Container Tools"
description: Operate running containers — start, stop, view logs, and open a shell from the portal.
weight: 50
---

Once a container is configured it can be controlled and inspected from the portal. The configuration screens (Overview, Network, Mounts, etc.) live at **cluster scope** — they describe what a container *should* be. The runtime controls described here live at **node scope** — they act on the container *as it's running on a specific node*.

## Where to find them

Navigate to a node, then **Compute → Container Management → Containers**, and click the container's name.

{{<alert color="info">}}
If the node is a cluster member, the portal shows a banner reading *"This node is a member of a cluster. These settings must be changed at the cluster level here."* The configuration fields are read-only at the node — but the **Start / Stop / Logs / Terminal** buttons described below remain available, because they operate on the running container on this specific node.
{{</alert>}}

{{<tgimg src="container-action-bar.png" caption="Per-container action bar (Start, Stop, Logs, Terminal) on the node-scoped detail view" width="90%">}}

## State

Above the action bar, the container list at node scope shows two columns:

- **Status** — `Enabled` / `Disabled`. The configured state. A disabled container will not be started on the node.
- **State** — the runtime state on this node. Common values: `Running`, `Stopped`, `Restarting`, `Pulling`.

A container can be `Enabled` and `Stopped` — for example, between a restart and the next pull, or if the image cannot be pulled.

{{<tgimg src="node-container-list.png" caption="Container list at node scope, showing runtime State alongside configured Status" width="90%">}}

## Start

Starts the container immediately on the node. For `Service` containers, this is the same path as the automatic startup that runs on node boot — useful after manually stopping the container or for `On Demand` containers that don't auto-start.

Requires `node-exec::compute` permission.

## Stop

Stops the container immediately. For `Service` containers, the node will respect the configured `Stop Time` grace period (default 30 seconds) before sending SIGKILL.

The container remains `Enabled` after a manual stop. To prevent it from being restarted automatically on the next node reboot or config change, set its **Status** to `Disabled` from the Overview screen at the cluster level.

Requires `node-exec::compute` permission.

## Logs

Opens a new browser tab streaming `stdout` and `stderr` from the running container.

{{<tgimg src="logs-modal.png" caption="Container Logs options" width="60%">}}

{{<fields>}}
{{<field "Follow Log Output">}}When **True**, the stream stays open and prints new log lines as they arrive. When **False**, the modal prints the most recent N lines and disconnects.{{</field>}}
{{<field "Number of Lines">}}How many lines of recent history to print before tailing. Default is 100.{{</field>}}
{{</fields>}}

The viewer is a streaming session opened through the Trustgrid control plane to the node — your browser never connects directly to the node's WAN address. A **CONNECTED** badge in the top-right indicates the WebSocket is live; clicking **Terminate** closes the stream.

{{<tgimg src="logs-stream.png" caption="Streaming log viewer with the CONNECTED indicator and Terminate button" width="90%">}}

{{<alert color="warning">}}
The portal-side log viewer only shows output for the **currently running** container instance. To collect logs across restarts, enable **Save Output** on the container's Overview (logs are then persisted to the Trustgrid cloud and visible in **Observability**) — but be aware that any sensitive output will be persisted as well.
{{</alert>}}

## Terminal

Opens a new browser tab with an interactive shell inside the running container. Equivalent to `docker exec -it <container> /bin/sh`.

The terminal uses the same WebSocket session protocol as the log viewer, tunneled through the Trustgrid control plane. The container's shell is launched as the configured user — if no `User` is set on the Overview, the shell runs as the image's default user (typically `root`).

The Terminal button is disabled when the container is not in the `Running` state.

{{<alert color="info">}}
If the container's image does not include a shell (e.g. a distroless image), the Terminal session will fail to start. Use **Command** on the Overview screen to set an alternative entrypoint for diagnostic builds, or override the shell binary with the **Container Shell Override** node feature.
{{</alert>}}

Requires `node-exec::compute` permission.

## History

The **History** section in the left sidebar of a running container exposes:

- **Connections** — recent inbound and outbound network connections observed for this container, useful for verifying that port mappings or virtual-network attachments are doing what you expect.

## Related

- [Container troubleshooting]({{<ref "../troubleshooting">}}) — what to check when Start does nothing, logs are empty, or the Terminal refuses to open.
- [Container networking]({{<ref "../concepts/networking">}}) — explains the bridge, the resolver, and how port mappings and virtual networks attach to a container.
