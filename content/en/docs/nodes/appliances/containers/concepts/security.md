---
title: "Container Security"
description: Capabilities, privileges, user identity, and connectivity gating for containers on Trustgrid nodes.
weight: 40
---

Trustgrid nodes run containers with least-privilege defaults. This page covers the levers you have to tighten or relax those defaults — and the cases where you actually need to.

## The default sandbox

Out of the box, a container on a Trustgrid node runs with:

- A reduced Linux capability set (no `CAP_SYS_ADMIN`, no `CAP_NET_ADMIN`, etc.).
- Its own PID, network, IPC, and mount namespaces.
- No access to the host filesystem unless explicit [bind mounts]({{<ref "storage#bind-mounts">}}) say otherwise.
- As the image's default user (typically `root` *inside* the container, but isolated from the node's `root` by namespacing).

Most containers don't need anything beyond that. Reach for the controls below only when something doesn't work without them.

## User

The **User** field on the Overview screen sets the user/group the container's main process runs as. Accepted formats:

- `user` — username defined inside the image
- `user:group` — username and group both defined inside the image
- `uid` — numeric user ID (does not need to exist inside the image)
- `uid:gid` — numeric user and group
- `user:gid` or `uid:group` — mixed

If left blank, the container runs as the image's default user (specified by `USER` in the Dockerfile, defaulting to `root`).

**Common reasons to set this:**

- Image runs as root by default but you'd rather a specific UID for ownership of mounted volumes.
- Multiple containers share a volume and you need consistent UIDs across them.
- Compliance requirement that nothing run as UID 0.

## Linux Capabilities

The **Linux Capabilities** screen lets you add or drop specific kernel capabilities the container has access to.

- **Add Caps** — grant a capability not in the default set. Pick narrowly; granting `CAP_SYS_ADMIN` is effectively granting root.
- **Drop Caps** — remove a capability from the default set. Use for paranoid hardening.

Capability names follow the Linux convention (`CAP_NET_BIND_SERVICE`, `CAP_SYS_PTRACE`, etc.). See [capabilities(7)](https://man7.org/linux/man-pages/man7/capabilities.7.html) for the full list and what each grants.

**Examples of when to add:**

| Capability | Why you'd add it |
|---|---|
| `CAP_NET_BIND_SERVICE` | Container needs to bind to a privileged port (< 1024) without running as root. |
| `CAP_NET_ADMIN` | Container manages network interfaces or routing inside its namespace — VPN clients, packet sniffers. |
| `CAP_SYS_PTRACE` | Container needs to attach a debugger or strace to a process inside it. |
| `CAP_NET_RAW` | Container uses raw sockets (e.g., ping, traceroute, custom protocols). |

**Always prefer adding specific capabilities over enabling `Privileged`.** Each cap you add is documented; `Privileged` removes the entire sandbox.

## Privileged

The **Privileged** toggle on the Overview screen relaxes the container sandbox so workloads that need elevated access can run. Common cases: nested container runtimes, low-level kernel-state managers, anything that needs more than the named [Linux Capabilities](#linux-capabilities) list can express. If your image needs it, turn it on; if you don't know whether you need it, you probably don't.

When you do know which specific capabilities your workload needs, prefer adding them via [Linux Capabilities](#linux-capabilities) rather than enabling Privileged — narrower grants are easier to audit.

## Use Init

The **Use Init** toggle enables a tiny init process (PID 1) inside the container.

By default, the container's main process is PID 1. PID 1 in Linux has special responsibilities — reaping zombie children, handling signals — that most application processes don't implement correctly. Symptoms of needing Use Init:

- Defunct/zombie processes accumulating inside the container over time.
- Container doesn't shut down cleanly on Stop; instead waits the full Stop Time grace period before being killed.
- Container is using shell scripts as the entrypoint that don't forward signals to child processes.

When enabled, a small init binary runs as PID 1 and your application runs as PID 2. **Cheap to enable; enable it for any service that spawns child processes.**

## Require Connectivity

The **Require Connectivity** toggle on the Overview gates container startup on the node having control-plane connectivity to Trustgrid.

- **Disabled (default):** the container starts whenever the node tries to start it, regardless of whether the node can reach the Trustgrid control plane.
- **Enabled:** the container will not start if the node has lost control-plane connectivity. Used in combination with encrypted volumes to ensure sensitive data never decrypts on an offline node.

This only affects **startup**. A running container that loses control-plane connectivity mid-run is not stopped automatically. To enforce continuous connectivity, also configure a Health Check that probes a control-plane-only endpoint.

See [storage — encrypted volumes]({{<ref "storage#encrypted-volumes">}}) for the typical use case.

## Save Output

The **Save Output** toggle on the Overview persists the container's `stdout` and `stderr` to the Trustgrid cloud, where it's visible in **Observability** and survives container restarts.

This is a **security knob** as well as an observability one:

- **It is the customer's responsibility to ensure no privileged information is included in the output.** If your container logs API keys, customer PII, or session tokens, those end up in the Trustgrid log store.
- The log retention follows your org's log retention policy.
- If you're unsure whether the container's output contains sensitive data, leave Save Output disabled and use the [portal Logs viewer]({{<ref "../tools#logs">}}) for live debugging instead — that streams logs without persisting.

## Putting it together

A common hardening profile for an internet-facing service:

1. Set **User** to a non-root UID (or use an image that sets it).
2. Leave **Privileged** off.
3. Under **Linux Capabilities**, drop everything you don't need (`CAP_CHOWN`, `CAP_DAC_OVERRIDE`, etc.) and add only the specific capabilities the service requires (e.g., `CAP_NET_BIND_SERVICE` if it listens on port 80).
4. Enable **Use Init** so signal handling and zombie reaping work correctly.
5. Leave **Save Output** off unless you've confirmed nothing sensitive ends up in stdout/stderr.
6. Use [bind mounts]({{<ref "storage#bind-mounts">}}) for config (read-only paths) and [encrypted volumes]({{<ref "storage#encrypted-volumes">}}) for application state.
7. Add a [Health Check]({{<ref "../#health-check">}}) so the node restarts the container on application-level failure.

## Related

- [Container Tools]({{<ref "../tools">}}) — Logs / Terminal access to a running container
- [Container troubleshooting]({{<ref "../troubleshooting">}}) — `Permission denied` and capability-related symptoms
- [Container storage]({{<ref "storage">}}) — encrypted volumes interaction with Require Connectivity
