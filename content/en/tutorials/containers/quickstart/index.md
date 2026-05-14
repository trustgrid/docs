---
title: "Container Quickstart"
description: Push an image and run it on a Trustgrid node — no prior Trustgrid knowledge assumed.
weight: 5
---

This tutorial walks through the shortest path from a container image on your laptop to a running container on a Trustgrid appliance, with a port mapping you can reach from its LAN. If you already have a Trustgrid appliance enrolled and a Docker image you want to deploy, this should take about 10 minutes.

## What you need

- A Trustgrid organization with at least one enrolled appliance.
- Permissions: `repositories::modify`, `node-exec::modify`, `node-exec::compute` on the target node or cluster.
- Docker on a `linux/amd64` machine. We'll use the public `nginx:alpine` as the example image.

{{<alert color="warning">}}
Images must be pushed as `linux/amd64`. If your workstation is an Apple Silicon Mac or Windows on ARM, push from an amd64 host (a Linux VM, a build server, or CI).
{{</alert>}}

## 1. Push the image to your Trustgrid registry

Each Trustgrid organization gets a private container registry. Your registry path is:

```
docker.<your-domain>/<your-namespace>/<image>:<tag>
```

For example, on the org with domain `acme.trustgrid.io` the path for an image called `myapp` tagged `v1` is `docker.acme.trustgrid.io/acme.trustgrid.io/myapp:v1`. Your exact namespace is shown on the **Repositories** page.

Authenticate. From the Trustgrid portal, navigate to **Repositories** and copy the **Docker Login** command — it includes a short-lived token. Paste it in your terminal:

```bash
docker login -u trustgrid -p <token> https://docker.<your-domain>
```

From your amd64 host, push the image:

```bash
docker pull nginx:alpine
docker tag nginx:alpine docker.acme.trustgrid.io/acme.trustgrid.io/nginx:alpine
docker push docker.acme.trustgrid.io/acme.trustgrid.io/nginx:alpine
```

Verify the tag landed by navigating to **Repositories → nginx** in the portal. You should see the `alpine` tag listed with its digest. If it doesn't appear, the push was not `linux/amd64` — push from an amd64 host.

{{<alert color="info">}}
The token in the `docker login` command expires after about 24 hours. Re-fetch it from the portal when it does.
{{</alert>}}

## 2. Create the container

Navigate to the cluster or node where you want to run the container, then **Compute → Container Management → Containers**, and click **Add Container**.

{{<alert color="warning">}}
If your target node is a member of a cluster, **create the container at the cluster level**, not on the individual node. Cluster-scoped containers are deployed to every member of the cluster, so the container is running on both members regardless of which one is active. The portal will redirect node-level changes back to the cluster for a clustered node.
{{</alert>}}

Fill in:

| Field | Value |
| --- | --- |
| **Name** | `nginx` |
| **Execution Type** | `Service` (will run as a daemon and auto-restart) |
| **Status** | `Enabled` |
| **Image Name** | `<your-namespace>/nginx` |
| **Image Tag** | `alpine` |

Click **Save**. The portal returns to the container list with `nginx` shown as `Enabled`.

## 3. Expose a port

By default the container is reachable only from the appliance itself. To make it reachable from the appliance's local network, add a host port mapping.

1. Click into `nginx`, then **Network** on the left sidebar.
2. Under **Host Port Mappings** click **Add**.
3. Fill in:

   | Field | Value |
   | --- | --- |
   | **Protocol** | `tcp` |
   | **Host Interface** | The LAN-facing NIC on the appliance (e.g. `ens192`). Check **Networking → Interfaces** if you're not sure. |
   | **Host Port** | `8080` |
   | **Container Port** | `80` |

4. Click **Save**.

## 4. Wait for it to start

Go back to the container list at **node scope** (not cluster scope — runtime state lives at the node scope). The container's **State** column transitions through `Initializing` → `Running` within a minute or two for a small image.

## 5. Reach it

From any host on the same LAN as the appliance's host interface (e.g. an admin workstation):

```bash
curl http://<appliance-LAN-IP>:8080/
```

You should see the default nginx welcome page.

## 6. Verify and explore

From the container's detail page at node scope, the action bar at the top has:

- **Start / Stop** — restart the container without changing its configuration.
- **Logs** — open a streaming viewer for the container's `stdout` / `stderr`.
- **Terminal** — open a shell inside the running container.

These are described in [Container Tools]({{<relref "/docs/nodes/appliances/containers/tools">}}).

## Next steps

- [Expose the container over a virtual network]({{<relref "/tutorials/containers/expose-over-vpn">}}) so peer Trustgrid nodes can reach it without going through the LAN.
- Add a [Health Check]({{<relref "/docs/nodes/appliances/containers/#health-check">}}) so the container is flagged in the portal if it stops responding.
- Persist data with a [bind mount or volume]({{<relref "/docs/nodes/appliances/containers/#mounts">}}).
