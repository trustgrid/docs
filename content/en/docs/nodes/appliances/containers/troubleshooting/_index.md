---
title: "Container Troubleshooting"
description: Diagnose containers that won't start, won't pull, can't reach the network, or can't be reached.
weight: 90
---

This page is organized by symptom — find the section that matches what you're seeing.

## Container stays in "Stopped" state

The container is `Enabled` but its **State** never moves off `Stopped`.

**Check, in order:**

1. **The image and tag exist in the registry.** Open **Repositories**, click into the image, and verify the tag is listed. The list can take a few minutes to update after a push.
2. **The node can reach the registry.** Run **Actions → Test Repo Connectivity** on the node. The Trustgrid registry is reached through the Trustgrid control plane — internet access alone isn't enough.
3. **The node has the latest config.** Cluster-level container configs take 30–60 seconds to propagate to nodes. If you just created the container, wait and refresh.
4. **The image works on this node's architecture.** A container built only for ARM won't run on an x86 node. Re-push as `linux/amd64` (see [Repositories — Supported image platforms]({{<ref "/docs/repositories#supported-image-platforms">}})).

## Image pull fails

State alternates between `Pulling` and `Stopped`, or stays at `Pulling` indefinitely. Or the node log shows `RegistryMediator: Retries exhausted` (the generic "couldn't get the image" error).

**Check:**

1. **Image name format.** The full path is `<your-namespace>/<image>` — the registry hostname is added by the node. Example: `acme.trustgrid.io/nginx`, not `docker.acme.trustgrid.io/acme.trustgrid.io/nginx`.
2. **Image was pushed as `linux/amd64`.** Pushes from Apple Silicon Macs or Windows on ARM are `arm64` by default and won't appear in the portal or be pullable. Re-push from an amd64 host. See [Repositories — Supported image platforms]({{<ref "/docs/repositories#supported-image-platforms">}}).
3. **Repo connectivity.** Run **Actions → Test Repo Connectivity** on the node.
4. **Disk space.** Check **Actions → Node Info** on the node. A full disk can't pull. Unused container layers are cleaned up automatically but it takes time; if needed, remove unused volumes under **Compute → Container Management → Volumes**.

## Container starts, then crashes immediately

State shows `Running` briefly, then `Stopped`, repeatedly.

**Check:**

1. **Logs.** Open the **Logs** viewer. Most application startup errors (missing environment variable, missing file, permission denied) print before the container exits.
2. **Save Output.** If the container crashes too fast to see in the live viewer, turn on **Save Output** on the Overview. Restart attempts will then be visible in **Observability**.
3. **Environment variables.** Check the container's **Environment Variables** screen against what the image's documentation says it needs.
4. **Required files.** A container that expects a config file at, say, `/etc/myapp/config.yaml` will crash if no mount provides it. Add a bind mount or volume.
5. **User mismatch.** If the image expects a specific UID and you set **User** to something else, the entrypoint may fail.

## Container is running but unreachable from the LAN

State is `Running`, but trying to reach the container from another machine on the LAN times out.

**Check:**

1. **Port mapping is on the right interface.** Under **Network → Host Port Mappings**, the **Host Interface** has to be the one that's on the LAN you're connecting from. A mapping on the WAN interface won't be reachable from the LAN.
2. **The container is actually listening.** Open the **Terminal** for the container and run `ss -tlnp` (or `netstat -tlnp` if `ss` isn't available). The container needs to be listening on `0.0.0.0:<container-port>`, not just `127.0.0.1`. nginx, Postgres, and Redis all default to localhost-only in some configurations.
3. **Node firewall.** Node-level firewall rules apply before the port mapping — check **Networking → Interfaces → Firewall**.
4. **Cluster member.** On a cluster, only the active member forwards traffic on the shared cluster IP. Try the active member's individual IP if the cluster IP isn't reachable.

## Container is running but unreachable from another Trustgrid node

The container's host port mapping works locally, but a peer node can't reach the container's virtual IP.

**Check:**

1. **The virtual network is attached to the container.** Under **Network → Virtual Networks** on the container, verify the network name and **Virtual IP**.
2. **The virtual network is also attached to the node.** Look at **Networking → VPN** on the node — the same network needs to be there. The container can only join networks the node has.
3. **VPN is up.** Try **Actions → Trustgrid Ping** from the peer node targeting the **node's** virtual IP (not the container's). If that fails, fix VPN connectivity first.
4. **The container actually got the address.** Open the **Terminal** and run `ip -br addr`. You should see an interface with the virtual IP. If not, toggle the container's Status to `Disabled` and back to force a restart.

## Container can't resolve DNS

Inside the container, lookups fail with "Could not resolve host."

**Check:**

1. **The container's resolver.** Inside the container, `/etc/resolv.conf` should show `172.18.1.2` — the node's resolver. If you set a custom **DNS** address on the container's Overview, that resolver is used instead — make sure it's reachable.
2. **The node's own DNS works.** Run **Actions → Test DNS** on the node. The container's resolver forwards external lookups to the node — so if the node's DNS isn't resolving, the container's won't either.

## Container can't reach the internet

State is `Running` but the container can't reach external sites.

**Check:**

1. **The node can reach the internet.** Run **Actions → Speed Test** on the node. Containers leave the node through the node's normal route.
2. **VRF.** If you set a VRF on the container's Network screen, the container is limited to that VRF's routing. Either make sure the VRF has a default route or unset it to use the node's normal routing.

## Health check keeps restarting the container

The container looks fine in the logs but the health check fails and the node restarts it.

**Check:**

1. **Run the health check command manually.** Open the **Terminal** and run it yourself. Non-zero exit = failure. Common issues: `wget` or `curl` isn't installed in the image, the path needs to be absolute, the check probes `localhost` but the service only listens on `0.0.0.0`.
2. **Start Period.** The check kicks in this many seconds after the container starts. If the app takes 30 seconds to come up but Start Period is `0`, the first checks fail and count toward Retries.
3. **Timeout.** A 1-second timeout on a slow endpoint will fail intermittently.

## Container is being killed for using too much memory or CPU

OOM kills, slow performance, or throttling.

**Check:**

1. **Resource Limits screen.** Defaults are CPU 50% and Memory 50% of the node. Raise them if your workload legitimately needs more.
2. **IO limits.** Off by default. If you set IO limits explicitly and the container is doing heavier disk activity than expected, it'll get throttled.
3. **Restart loop.** A container that's crashing and restarting will keep hitting the same memory limit. Fix the underlying crash first.

## Diagnostic checklist for any container issue

When you're not sure what's wrong, gather these in order:

1. **Container State** — from the node-scoped container list.
2. **Container logs** — Logs viewer, with a high line count to capture history.
3. **Node Repo connectivity** — **Actions → Test Repo Connectivity**.
4. **Node DNS** — **Actions → Test DNS**.
5. **Active flows** — **Actions → Active Flows** filtered by `containerId`.

A container issue almost always shows up in one of these.
