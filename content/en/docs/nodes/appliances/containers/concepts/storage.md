---
title: "Container Storage"
description: How to give a container persistent storage or access to host files.
weight: 30
---

By default, anything a container writes is lost when it stops — the filesystem is reset every time the container restarts. To keep data around, or to share a file from the node with the container, configure a **mount** on the container's Mounts screen.

There are two mount types: **bind** and **volume**. Pick a volume for application data you want to keep; pick a bind mount when you want to share a specific file or folder that already exists on the node.

## Bind mounts

A bind mount makes a file or folder from the node visible inside the container.

| Field | Notes |
|---|---|
| **Type** | `BIND` |
| **Source** | The path on the node, e.g. `/etc/myapp` or `/var/log/myapp`. The path must already exist on the node before the container starts. |
| **Destination** | Where the container should see it. |

**Limitations:**

- The path has to already exist on the node — there's no portal mechanism to populate an arbitrary node path. For shipping configuration in, use environment variables or bake the config into the image.
- On a cluster, the path has to exist on every member or the container will fail on the member that's missing it.
- The container needs permission to read or write the path. If it can't, you'll see a `Permission denied` error in the container's logs.

## Volumes

A volume is a piece of managed storage that lives alongside the container, set up under **Compute → Container Management → Volumes**. Volumes have a name and can be reused across containers.

| Field | Notes |
|---|---|
| **Type** | `VOLUME` |
| **Source** | The name of a volume you already created. |
| **Destination** | Where the container should see it. |

**When to use a volume:**

- You want Trustgrid to handle the storage location instead of picking a host path yourself.
- You want the data to survive container restarts, image updates, or container reconfiguration.
- You want to share storage between two containers running on the same node.
- You want to encrypt the data at rest (see below).

### Encrypted volumes

Mark a volume **Encrypted** to have Trustgrid encrypt its contents at rest. The key is fetched from the Trustgrid control plane when the container starts, so the volume can only be unlocked while the node is connected to the control plane.

The **Require Connectivity** toggle on the container's Overview controls what happens when the node is disconnected from the control plane:

- **On** — the container won't start until the node reconnects to the control plane. Use this when running offline is worse than not running.
- **Off** — a container that's already running keeps running; one trying to start while disconnected will fail because the volume can't unlock.

Require Connectivity has no effect on unencrypted volumes.

## Bind or volume — quick comparison

| | Bind | Volume |
|---|---|---|
| Storage location | A path you pick on the node | Trustgrid manages it |
| Survives container restart / image change | Yes | Yes |
| Encryption available | No | Yes |
| Works automatically on all cluster members | No | Yes — volumes are defined at the cluster level |
| Good fit for application data (databases, app state) | Not ideal — no encryption, manual cluster setup | ✔ |

For most workloads, **use volumes**. Bind mounts are a fit only when there's already a node-side path the container needs to see (e.g. a path created by another Trustgrid feature).

## Importing volumes between nodes

Under **Container Management → Volumes** there's an **Import** action that copies a volume's name and settings from another node or cluster — handy when you're rolling out the same container to a new edge. It copies the **definition only**; the new volume starts empty.

## Related

- [Container Tools]({{<ref "../tools">}}) — opening a shell inside a running container
- [Volumes reference]({{<ref "../volumes">}}) — every field on the Volumes screen
