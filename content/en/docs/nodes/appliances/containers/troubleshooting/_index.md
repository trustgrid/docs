---
title: "Container Troubleshooting"
description: Diagnose containers that won't start, won't pull, can't reach the network, or can't be reached.
weight: 90
---

This page is organized by symptom. Each section starts with a one-line summary and lists what to check first.

## `RegistryMediator: Retries exhausted downloading https://repo.*/v2/...`

This appears in `/var/log/trustgrid/tg-default.log` on the node and is the most common cause of a container staying in `Stopped` after a fresh deploy. The exception text is generic — the underlying cause is *not* the same across reports — but in most cases it means **the node could not retrieve the image manifest from the configured pull host** (the value of `repo.uri` in the node's profile, typically `repo.<env>.trustgrid.io`).

What it does **not** tell you:
- The HTTP status code returned by the registry.
- Whether the failure was authentication (401), missing manifest (404), TLS handshake, or timeout.
- Which `Accept` header the node sent.

Known causes, in order of likelihood:

1. **Image was not pushed as `linux/amd64`.** Pushes from Apple Silicon Macs or Windows on ARM default to `arm64` and the tag will not appear in the portal nor be pullable by nodes. Re-push with `docker buildx build --platform linux/amd64 --push ...` or push from an amd64 host. See [Repositories — Supported image platforms]({{<ref "/docs/repositories#supported-image-platforms">}}).
2. **Multi-segment image name** (e.g. `mynamespace/sub/dir/image`). Historical bug in node URL construction. Fixed in node release `n-2.23.0` (Aug 2025). If you're seeing this on an older node, upgrade.
3. **Tag does not exist.** Confirm via **Repositories** in the portal that the tag is actually listed for your image. The portal list lags pushes by a few seconds.
4. **Registry unreachable from the node.** Run `Actions → Test Repo Connectivity` on the node — it should return `connected`. If not, fix control-plane connectivity first.

To collect more detail before raising:

```bash
# From the node shell:
sudo grep -h "RegistryMediator\|manifests" /var/log/trustgrid/tg-default.log | tail -20
sudo cat /var/lib/trustgrid/config/node-profile.json | grep repo.uri
```

The manifest URL the node tries follows the pattern `<repo.uri>/v2/<namespace>/<image>/manifests/<tag>` — verify it matches the image you pushed.

## Container stays in "Stopped" state

The configuration is present and the container is `Enabled`, but **State** in the node-scoped container list never moves off `Stopped`.

**Check, in order:**

1. **Image and tag exist in the registry.** Open **Repositories**, click into the image, and verify the tag is listed. The list of tags can lag the registry by a few minutes; if you just pushed, wait or verify directly with `crane ls docker.<your-domain>/<namespace>/<image>`.
2. **Node can reach the registry.** Run **Actions → Test Repo Connectivity** on the node. The Trustgrid registry is reached through the control plane, so internet connectivity isn't enough — control-plane connectivity is what matters. If repo connectivity fails, fix that before continuing.
3. **Node has the latest config.** Cluster-scoped containers are propagated to nodes asynchronously. If you just created the container at cluster scope, wait 30–60 seconds and refresh. The node-scoped container list will show the container once it has the config.
4. **Node startup error.** Check the node's overview for the `startup.error` flag. If set, the node's container runtime may not have started cleanly — restart the node service from the node's action bar (**Restart**) and re-check.
5. **Image is incompatible.** A pulled image that targets a different CPU architecture (e.g. an ARM-only image on an x86 node) will pull successfully but fail to start. Re-tag the image as multi-arch or push a build for the node's architecture.

## Image pull fails

State alternates between `Pulling` and `Stopped`, or stays at `Pulling` indefinitely.

**Check:**

1. **Image name format.** The full path is `<your-namespace>/<image>` (the registry hostname is added by the node automatically). For example, `acme.trustgrid.io/nginx`, not `docker.acme.trustgrid.io/acme.trustgrid.io/nginx`.
2. **Registry control-plane connectivity.** As above — **Test Repo Connectivity**.
3. **Disk space.** Run **Actions → Node Info** and check the storage gauge. A node with full disk cannot pull. Old container layers are cleaned automatically but can take time; if needed, manually remove unused volumes from **Compute → Container Management → Volumes**.

## Container starts, then crashes immediately

State reads `Running` briefly, then `Stopped`. Repeats every few seconds.

**Check:**

1. **Logs.** Open the **Logs** viewer for the container. Most application-level startup failures (missing env var, missing file, permission denied) print to stdout/stderr before exiting.
2. **Save Output.** If the container crashes before you can attach the live log viewer, enable **Save Output** on the Overview screen. Logs will be persisted and visible in **Observability** even across restarts.
3. **Required environment variables.** Compare the container's **Environment Variables** screen against the image's documented requirements.
4. **Required mounts.** A container that expects a config file at `/etc/myapp/config.yaml` will exit if no mount provides it. Add a **Bind** or **Volume** mount on the Mounts screen.
5. **User / UID mismatch.** If the image expects to run as a specific UID and you've set **User** on the Overview to something else, the entrypoint may fail. Either match the image's expected user or use **Linux Capabilities** to grant the needed permissions.

## Container is running but unreachable from the LAN

State is `Running`. `nc -zv <node-LAN-IP> <host-port>` returns connection refused or times out.

**Check:**

1. **Port mapping bound to the right interface.** Open **Network → Host Port Mappings**. The **Host Interface** field must be the NIC that's actually on the LAN you're connecting from. A mapping on `ens160` (WAN) will not be reachable from a host plugged into `ens192` (LAN).
2. **Container is actually listening.** Open the **Terminal** for the container and run `ss -tlnp` (or `netstat -tlnp`). The container must be listening on `0.0.0.0:<container-port>`, not just on `127.0.0.1`. nginx, Postgres, and Redis all default to localhost-only on some distributions.
3. **Node firewall.** Node-level firewall rules apply to inbound traffic before it reaches the port-mapping rules. Check **Networking → Interfaces → Firewall** on the node.
4. **Cluster active member.** If the node is a clustered edge, only the active member is forwarding traffic on the cluster's shared IP. Verify the active member from **Cluster → Overview** and try the active member's individual IP if the shared IP is unreachable.

## Container is running but unreachable from a peer node over a virtual network

State is `Running`. The host port mapping works locally, but a peer node can't reach the container's virtual IP.

**Check:**

1. **Container has the virtual network attachment.** Open **Network → Virtual Networks** on the container. Verify the network name and **Virtual IP**.
2. **Virtual network is attached to the node.** **Networking → VPN** on the node should list the same network. The container can only join a virtual network already attached to the node.
3. **VPN tunnel is up.** Use **Actions → Trustgrid Ping** from the peer node, targeting the *node's* virtual IP (not the container's). If the node-to-node ping fails, fix VPN connectivity first.
4. **Address inside the container.** Open the **Terminal** and run `ip -br addr`. You should see an additional interface with the configured virtual IP. If it's missing, the attachment didn't apply — try toggling **Status** to `Disabled` and back to `Enabled` to force a restart.
5. **Allow Outbound (for replies in some configurations).** Inbound-only attachments should respond to incoming connections fine, but if you're seeing one-way traffic, verify the attachment isn't NAT-translating in an unexpected direction. Use **Actions → VPN NATs** on the node to see the active translation table.

## Container can't resolve DNS

`getaddrinfo` failures inside the container; logs show `Could not resolve host`.

**Check:**

1. **Default resolver.** Inside the container, `/etc/resolv.conf` should point to `172.18.1.2` (the node-side resolver). If you've set the **DNS** field on the Overview, the container uses that resolver instead — verify the address is correct and reachable from the container's network space.
2. **Node DNS is healthy.** **Actions → Test DNS** on the node. The container's resolver forwards external lookups to the node's resolvers — if those are broken, container DNS breaks too.
3. **DNS over a virtual network.** If you set **DNS** to an address only reachable via a virtual network, the container must also be attached to that virtual network with at least the resolver as a reachable route.

## Container can't reach the internet

State is `Running` but the container can't `curl https://example.com`.

**Check:**

1. **Node can reach the internet.** **Actions → Speed Test**. Containers egress through the node's default route — if the node is offline, the container is offline.
2. **VRF.** If you set a **VRF** on the container's Network screen, outbound traffic is restricted to that VRF's routing table. Confirm the VRF has a default route, or unset it to use the node's default routing.
3. **Allow Outbound on virtual networks.** If the container has a virtual network attachment with **Allow Outbound** enabled, some traffic may be routed onto the virtual network instead of egress. Check the virtual network's routes.

## Health check keeps restarting the container

The container is healthy from a user's perspective but the configured health check fails and the node restarts it on the configured interval.

**Check:**

1. **Health check command.** Run it manually from the **Terminal** to see what it returns. A non-zero exit code is a failure. Common gotchas: `wget`/`curl` not present in the image, the command needs an absolute path, the check probes `localhost` but the service binds to `0.0.0.0` only.
2. **Start Period.** The check starts running this many seconds after the container starts. If your application takes 30s to come up but **Start Period** is `0`, the first checks will fail and accumulate against **Retries**.
3. **Interval / Timeout.** A `curl` health check against a slow endpoint with **Timeout: 1** will fail intermittently.

## Container hits a resource limit

OOM kills, throttling, or slow IO.

**Check:**

1. **Resource Limits screen.** Default CPU max is 50% and default Memory max is 50% of the host. For a container that genuinely needs more, raise the limits explicitly.
2. **IO limits.** Disabled by default. If you set `IO Max Read (B/s)` or similar and the container is doing more disk activity than expected, it may be throttled.
3. **Health-check restart loops.** A restarting container repeatedly burns the same memory budget — the OOM you see may be a symptom of the health-check failure above, not a real over-budget condition.

## Diagnostic checklist for any container issue

When you don't know what's wrong, collect these in order:

1. **Container state** — from the node-scoped container list.
2. **Container logs** — Logs viewer, with **Follow** off and a high line count to grab history.
3. **Node startup.error** — from the node's Overview.
4. **Node service status** — **Actions → Test Runtime Status**.
5. **Repo connectivity** — **Actions → Test Repo Connectivity**.
6. **DNS health** — **Actions → Test DNS**.
7. **Active flows** — **Actions → Active Flows**, filtered by `containerId`.

A container issue almost always falls out from one of these seven views.
