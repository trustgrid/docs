---
title: "Container Security"
weight: 40
---

By default, a container on a Trustgrid node runs with safe, restrictive settings — it can't see the host filesystem, can't touch host devices, and can't interfere with the node or other containers. The toggles on this page let you grant additional access when a specific workload needs it.

## User

The **User** field on the Overview screen sets which user inside the container runs the application. By default it uses the user the image was built to run as.

**When to set it:**

- The image runs as `root` by default and you want it to run as a specific UID instead — usually for ownership of files in a mounted volume.
- Several containers share a volume and need to agree on the same UID so they can all read/write the files.
- You have a compliance requirement that nothing runs as UID 0.

Accepted formats: `user`, `user:group`, `uid`, `uid:gid`, or a mix.

## Linux Capabilities

The **Linux Capabilities** screen lets you grant the container specific abilities it doesn't have by default, or remove abilities it does have.

You only need to touch this if your image's documentation says it requires a specific capability — for example "requires `NET_ADMIN`" or "needs `SYS_PTRACE`."

**Examples of when an image might ask for one:**

| Capability | Typical use |
|---|---|
| `CAP_NET_BIND_SERVICE` | Listen on a privileged port (under 1024) without running as `root`. |
| `CAP_NET_ADMIN` | Manage network interfaces or routing inside the container — VPN clients, packet capture tools. |
| `CAP_SYS_PTRACE` | Attach a debugger or `strace` to a process inside the container. |
| `CAP_NET_RAW` | Use raw sockets — `ping`, `traceroute`, custom network tools. |

Adding a specific capability is preferred over enabling [Privileged](#privileged) — it follows the principle of least access.

## Privileged

The **Privileged** toggle on the Overview screen grants the container elevated access at the container level. It's commonly needed for workloads that manage kernel-level state, run a nested container runtime, or otherwise need more access than a specific Linux Capability can grant. Most container deployments do not require it. Where possible, use [Linux Capabilities](#linux-capabilities) instead — granting only the specific capability a workload needs follows the principle of least access.

## Use Init

The **Use Init** toggle on the Overview helps containers shut down cleanly.

A typical case: an application launched through a shell script wrapper (e.g. `start.sh` that runs `java -jar app.jar`). When you click **Stop**, the shutdown signal goes to the shell instead of the application — so the application keeps running until the node force-kills it after the Stop Time grace period. Use Init forwards the signal correctly so the application shuts down as expected.

It's cheap to turn on and a safe default for any container whose entrypoint isn't the application itself.

## Require Connectivity

The **Require Connectivity** toggle on the Overview gates container startup on the node being online — meaning the node has a live connection to the Trustgrid management plane and shows as online in the portal.

- **Off (default):** the container starts whenever the node tries to start it, whether or not the node is online.
- **On:** the container won't start until the node is online.

Use this with encrypted volumes — see [Container storage — Encrypted volumes]({{<ref "storage#encrypted-volumes">}}). It only affects startup; a container that's already running keeps running if the node goes offline.

## Save Output

The **Save Output** toggle on the Overview saves the container's log output (everything it prints to the terminal) to Trustgrid, where you can view it later in **Observability** even after the container restarts.

**Be careful what you save.** If your container prints API keys, customer information, or other sensitive data, that ends up in Trustgrid's log store. It is the customer's responsibility to ensure no sensitive information appears in the output.

If you're unsure about a container's output, leave Save Output off and use the [Logs viewer]({{<ref "../tools#logs">}}) for live debugging instead — that shows logs as they happen without keeping a copy.

## A secure setup for an internet-facing service

If you want a starting point for a service that's exposed to the public internet:

1. Set **User** to a non-root user (or use an image that already does).
2. Leave **Privileged** off.
3. Leave **Linux Capabilities** at defaults unless the image asks for one.
4. Enable **Use Init**.
5. Leave **Save Output** off unless you've checked the container's output for sensitive data.
6. Use [bind mounts]({{<ref "storage#bind-mounts">}}) for read-only config files and [encrypted volumes]({{<ref "storage#encrypted-volumes">}}) for application data.
7. Add a [Health Check]({{<ref "../#health-check">}}) so the node restarts the container if it stops responding.

## Related

- [Container Tools]({{<ref "../tools">}}) — viewing logs and opening a shell in a running container
- [Container troubleshooting]({{<ref "../troubleshooting">}}) — `Permission denied` and related errors
- [Container storage]({{<ref "storage">}}) — encrypted volumes and Require Connectivity
