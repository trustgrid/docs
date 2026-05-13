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

**When to use a bind mount:**

- You need to give the container a config file you maintain on the node.
- You want logs or output written somewhere you can grab from the node side.
- The data is managed by something outside the container — config rolled out by Ansible, certs rotated by another process, etc.

**Limitations:**

- The path has to exist on every node the container might run on. For a clustered container that means every cluster member; if `/etc/myapp` exists on one cluster member but not the other, the container will break after failover.
- The container needs permission to read or write the path. If it can't, you'll see a `Permission denied` error in the container's logs.
- Bind-mounted data is not part of node backups. If the node is replaced, the data goes with it unless you've backed it up separately.

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

When creating a volume you can mark it **Encrypted**. Trustgrid encrypts the data on disk using a key it fetches from the Trustgrid control plane when the container starts. Two things to know:

- A volume can only be unlocked while the node is connected to Trustgrid. If the node is offline, the data stays encrypted.
- This protects the data if someone gets the node's disk — they can't read the volume without also having a working node and a connection to the control plane.

There's one related setting, **Require Connectivity** on the container's Overview, which controls what happens to a stopped container when the node is offline:

- **Require Connectivity: On** + encrypted volume → if the node loses its connection while the container is stopped, it won't start until the connection comes back. Use this when running offline is worse than not running.
- **Require Connectivity: Off** + encrypted volume → a container that's already running stays running, but one trying to start while disconnected will fail because the volume can't unlock.

Require Connectivity has no effect on unencrypted volumes.

## Bind or volume — quick comparison

| | Bind | Volume |
|---|---|---|
| Storage location | A path you pick on the node | Trustgrid manages it |
| Survives container restart / image change | Yes | Yes |
| Survives node replacement | No | No |
| Encryption available | No | Yes |
| Works automatically on all cluster members | Only if you keep the path in sync yourself | Yes — volumes are defined at the cluster level |
| Good fit for config files | ✔ | Overkill |
| Good fit for application data (databases, app state) | Not ideal — no encryption, manual cluster setup | ✔ |

A common pattern: **bind-mount config files, use a volume for application data.** A database container might bind-mount `/etc/myapp/config.yaml` from the host and use a volume for `/var/lib/myapp/data`.

## Importing volumes between nodes

Under **Container Management → Volumes** there's an **Import** action that copies a volume's name and settings from another node or cluster — handy when you're rolling out the same container to a new edge. It copies the **definition only**; the new volume starts empty.

## Related

- [Container Tools]({{<ref "../tools">}}) — opening a shell inside a running container
- [Volumes reference]({{<ref "../volumes">}}) — every field on the Volumes screen
- [Tutorial: bind-mount a config file]({{<ref "/tutorials/containers/bind-mount-config">}}) — end-to-end walkthrough
