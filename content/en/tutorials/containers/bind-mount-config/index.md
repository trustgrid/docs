---
title: "Bind-mount a Config File"
description: Make a host config file or directory available inside a container, the most common persistence pattern.
weight: 15
---

This tutorial walks through wiring a host file into a container as a bind mount. We'll use nginx as the example — replacing its default config with one we control from the node — but the same pattern works for any application that reads configuration off disk.

## What you need

- A container running on a Trustgrid node — follow the [Quickstart]({{<ref "/tutorials/containers/quickstart">}}) if you don't have one yet.
- SSH or another way to write files on the node (you'll place the host-side config file there).
- The node's `node-exec::modify` permission.

For the impatient: this entire tutorial works through the portal — no node SSH required if you write the config file via the container itself first, then promote it to a bind mount. We'll show the SSH version first because it's the canonical pattern.

## 1. Stage the config file on the node

SSH to the node and create the directory and file you want to mount. For the nginx example:

```bash
sudo mkdir -p /etc/trustgrid-apps/hello-nginx
sudo tee /etc/trustgrid-apps/hello-nginx/index.html >/dev/null <<'HTML'
<!doctype html>
<html><body>
  <h1>Hello from a bind-mounted file</h1>
  <p>Served by nginx, read from /etc/trustgrid-apps/hello-nginx on the node.</p>
</body></html>
HTML
```

Make sure the path is readable by the user the container runs as. For nginx the default container user is `nginx` (UID 101 in the alpine image), so:

```bash
sudo chown -R 101:101 /etc/trustgrid-apps/hello-nginx
```

If you skip the chown the container will start but nginx will log `Permission denied` and serve `403 Forbidden`. See [troubleshooting]({{<ref "/docs/nodes/appliances/containers/troubleshooting">}}).

{{<alert color="warning">}}
On a **clustered** node, the host path needs to exist on every cluster member. If you stage `/etc/trustgrid-apps/hello-nginx` on edge1 but not edge2, the container will work until failover and then break. For data that needs to be shared across members, use a [volume]({{<ref "/docs/nodes/appliances/containers/concepts/storage#volumes">}}) instead.
{{</alert>}}

## 2. Add the bind mount to the container

Open the container in the portal (**Compute → Container Management → Containers → `hello-nginx`**), navigate to **Mounts**, click **Add**.

| Field | Value |
|---|---|
| **Type** | `BIND` |
| **Source** | `/etc/trustgrid-apps/hello-nginx` |
| **Destination** | `/usr/share/nginx/html` |

Save.

The portal will prompt to apply the change. The container restarts with the new mount in place — for a Service container this happens automatically; for an On Demand container you'll need to Start it manually.

## 3. Verify

From a host on the node's LAN, curl the container:

```bash
curl http://<node-LAN-IP>:8080/
```

You should see the HTML you just wrote. If you still see the default nginx welcome page, the bind mount didn't apply — check the Mounts list shows the entry and the State on the node-scoped container detail shows `Running` (not `Stopped`).

To iterate on the content without restarting:

```bash
sudo vim /etc/trustgrid-apps/hello-nginx/index.html
curl http://<node-LAN-IP>:8080/   # picks up the change immediately
```

The container sees host filesystem changes in real time — nginx serves whatever's currently at `/usr/share/nginx/html/index.html`.

## When to bind-mount vs. volume

| Need | Choice |
|---|---|
| Mount a config file managed by Ansible/Puppet/etc. | **Bind** |
| Mount a directory shared with a logging or monitoring tool on the host | **Bind** |
| Make application state survive container restarts | **Volume** |
| Encrypt the data at rest | **Volume** (encrypted) |
| Move between cluster members | **Volume** |

See [Container storage]({{<ref "/docs/nodes/appliances/containers/concepts/storage">}}) for the full comparison.

## Common pitfalls

- **`Permission denied` in the container logs.** The container's user can't read the bind mount. `chown` the host path to the right UID/GID, or run the container as a user that has access (see [Container security — User]({{<ref "/docs/nodes/appliances/containers/concepts/security#user">}})).
- **Host path doesn't exist.** The mount won't be created automatically. Create the directory on the host first.
- **Symlinks point outside the mount.** The container sees a broken link — Docker's bind mount doesn't follow symlinks out of the bound directory.
- **SELinux on the node.** If you're running on a hardened distribution with SELinux enforcing, you may need to label the host directory with the right context. Trustgrid's default node OS does not enable SELinux.

## Related

- [Container storage concepts]({{<ref "/docs/nodes/appliances/containers/concepts/storage">}})
- [Container security — User]({{<ref "/docs/nodes/appliances/containers/concepts/security#user">}})
- [Container troubleshooting]({{<ref "/docs/nodes/appliances/containers/troubleshooting">}})
