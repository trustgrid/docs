---
title: "Container Quickstart"
description: Push an image and run it on a Trustgrid node — no prior Trustgrid knowledge assumed.
weight: 5
---

This tutorial walks through the shortest path from a Docker image on your laptop to a running container on a Trustgrid node, with a port mapping you can reach from the node's LAN. If you already have a Trustgrid node enrolled and a Docker image you want to deploy, this should take about 10 minutes.

## What you need

- A Trustgrid organization with at least one enrolled appliance node.
- Permissions: `repositories::modify`, `node-exec::modify`, `node-exec::compute` on the target node or cluster.
- Docker (or another OCI registry client like [`crane`](https://github.com/google/go-containerregistry/tree/main/cmd/crane)) on your local machine.
- The image you want to run, available locally. We'll use the public `nginx:alpine` as the example.

## 1. Push the image to your Trustgrid registry

Each Trustgrid organization gets a private container registry. Your registry path is:

```
docker.<your-domain>/<your-namespace>/<image>:<tag>
```

For example, on the org with domain `acme.trustgrid.io` the path for an image called `myapp` tagged `v1` is `docker.acme.trustgrid.io/acme.trustgrid.io/myapp:v1`. Your exact namespace is shown on the **Repositories** page.

{{<alert color="warning">}}
**Pushing from Apple Silicon?** Docker Desktop and Colima on M-series Macs push `linux/arm64` OCI manifests by default. The Trustgrid registry only indexes `linux/amd64` Docker schema 2 manifests — an arm64 push completes silently but the tag will be invisible to the portal and unreachable by nodes. Use `docker buildx build --platform linux/amd64 --push ...` or run the push from an amd64 Linux host. See [Repositories — Supported image platforms]({{<ref "/docs/repositories#supported-image-platforms">}}).
{{</alert>}}

Authenticate. From the Trustgrid portal, navigate to **Repositories** and copy the **Docker Login** command — it includes a short-lived token. Paste it in your terminal:

```bash
docker login -u trustgrid -p <token> https://docker.<your-domain>
```

Tag and push:

```bash
docker pull nginx:alpine
docker tag nginx:alpine docker.acme.trustgrid.io/acme.trustgrid.io/nginx:alpine
docker push docker.acme.trustgrid.io/acme.trustgrid.io/nginx:alpine
```

Verify the tag landed by navigating to **Repositories → nginx** in the portal. You should see the `alpine` tag listed with its digest.

{{<alert color="info">}}
The token in the `docker login` command expires after about 24 hours. Re-fetch it from the portal when it does.
{{</alert>}}

## 2. Create the container

Navigate to the cluster or node where you want to run the container, then **Compute → Container Management → Containers**, and click **Add Container**.

{{<alert color="warning">}}
If your target node is a member of a cluster, **create the container at the cluster level**, not on the individual node. Cluster-scoped containers automatically follow the active cluster member and survive failover. The portal will redirect node-level changes back to the cluster for a clustered node.
{{</alert>}}

Fill in:

| Field | Value |
| --- | --- |
| **Name** | `hello-nginx` |
| **Execution Type** | `Service` (will run as a daemon and auto-restart) |
| **Status** | `Enabled` |
| **Image Name** | `<your-namespace>/nginx` |
| **Image Tag** | `alpine` |

Click **Save**. The portal returns to the container list with `hello-nginx` shown as `Enabled`.

## 3. Expose a port

By default the container is reachable only from the node itself. To make it reachable from the node's local network, add a host port mapping.

1. Click into `hello-nginx`, then **Network** on the left sidebar.
2. Under **Host Port Mappings** click **Add**.
3. Fill in:

   | Field | Value |
   | --- | --- |
   | **Protocol** | `tcp` |
   | **Host Interface** | The LAN-facing NIC on the node (e.g. `ens192`). Check **Networking → Interfaces** on the node if you're not sure. |
   | **Host Port** | `8080` |
   | **Container Port** | `80` |

4. Click **Save**.

## 4. Wait for it to start

Go back to the container list at **node scope** (not cluster scope — the runtime state lives on the node). The container's **State** column transitions through `Pulling` → `Stopped` → `Running` within a minute or two for a small image.

If it sticks at `Stopped`, see [Troubleshooting]({{<ref "/docs/nodes/appliances/containers/troubleshooting">}}).

## 5. Reach it

From any host on the same LAN as the node's host interface (e.g. an admin workstation):

```bash
curl http://<node-LAN-IP>:8080/
```

You should see the default nginx welcome page.

## 6. Verify and explore

From the container's detail page at node scope, the action bar at the top has:

- **Start / Stop** — restart the container without changing its configuration.
- **Logs** — open a streaming viewer for the container's `stdout` / `stderr`.
- **Terminal** — open a shell inside the running container.

These are described in [Container Tools]({{<ref "/docs/nodes/appliances/containers/tools">}}).

## Next steps

- [Expose the container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}}) so peer Trustgrid nodes can reach it without going through the LAN.
- Add a [Health Check]({{<ref "/docs/nodes/appliances/containers/#health-check">}}) so the node restarts the container if it stops responding.
- Persist data with a [bind mount or volume]({{<ref "/docs/nodes/appliances/containers/#mounts">}}).
