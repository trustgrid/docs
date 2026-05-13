---
title: "Container Storage"
description: Bind mounts, volumes, and the interaction with encryption and connectivity requirements.
weight: 30
---

A container has only an ephemeral filesystem by default — anything written inside the container is lost when it stops and is recreated. To persist data, or to make host files available to the container, configure a **mount** on the container's Mounts screen.

There are two mount types: **bind** and **volume**.

## Bind mounts

A bind mount maps a path on the node's filesystem directly into the container.

| Field | Notes |
|---|---|
| **Type** | `BIND` |
| **Source** | Absolute path on the node, e.g. `/etc/myapp` or `/var/log/myapp`. The path must already exist on the node before the container starts. |
| **Destination** | Mount location inside the container. |

**When to use a bind mount:**

- You need to share an existing host file or directory with the container — a config file from `/etc`, a log directory, a TLS cert bundle.
- You need the path to be inspectable from the node side (e.g., to copy a log file off the node manually).
- The data is managed by something outside the container's lifecycle (config managed by Ansible, certs rotated by another process, etc.).

**Limitations:**

- The host path must exist on every node the container can run on. For a clustered container, that means **every cluster member** — if a path exists on edge1 but not edge2, the container will fail on edge2 (or after failover).
- Permissions follow the host filesystem. If the container's user can't read or write the path, it will fail at runtime — typically with a `Permission denied` in the container logs.
- Bind mounts are **not** included in node backups. If the node is replaced, you lose them unless you've backed them up out-of-band.

## Volumes

A volume is a managed storage location with a lifecycle independent of any container.

| Field | Notes |
|---|---|
| **Type** | `VOLUME` |
| **Source** | The name of a pre-existing volume on the node or cluster (configured under **Container Management → Volumes**). |
| **Destination** | Mount location inside the container. |

Volumes are created under **Compute → Container Management → Volumes**. Each volume is named and can be reused across containers — typical pattern is one volume per application that needs persistent state.

**When to use a volume:**

- You want Trustgrid to manage the storage rather than wiring up a host directory yourself.
- You want the data to survive container replacement, node config changes, or image updates.
- You want to share storage between containers on the same node.
- You want the option to **encrypt** the data at rest (see below).

### Encrypted volumes

When creating a volume you can mark it **Encrypted**. Encrypted volume data is encrypted on the node disk using a key the node fetches from the Trustgrid control plane at runtime. This means:

- An attacker with raw disk access to the node cannot read volume data without also obtaining a working node identity and a path to the control plane.
- The node must have **control plane connectivity** to unlock the volume. If the node is offline from Trustgrid's control plane, the volume is unreadable.

There's an interaction with container behavior on disconnect, controlled by the **Require Connectivity** field on the container's Overview:

- **Require Connectivity: enabled** + encrypted volume → if the node loses control-plane connectivity while the container is stopped, the container will not start. Once connectivity returns, the volume unlocks and the container starts normally. **Use this when the data in the volume is sensitive enough that running offline is worse than running not at all.**
- **Require Connectivity: disabled** + encrypted volume → behavior on disconnect depends on whether the volume was already unlocked. Containers with unlocked encrypted volumes keep running; containers attempting to start while disconnected will fail to mount the volume.

For unencrypted volumes Require Connectivity has no effect on storage.

## Choosing between bind and volume

| | Bind | Volume |
|---|---|---|
| Storage managed by | You (on the node filesystem) | Trustgrid |
| Survives container replacement | Yes (host path persists) | Yes |
| Survives node replacement | No (data on the old node) | No (data on the old node) |
| Can be encrypted | No | Yes, with control-plane key |
| Available on all cluster members | Only if you ensure it | Yes (volume defined at cluster level) |
| Inspect from host shell | Yes (it's a host path) | Yes, but path is under Trustgrid-managed location |
| Right for config files | ✔ | ✘ (overkill) |
| Right for application state / databases | ✘ (no encryption, manual cluster sync) | ✔ |

A common pattern: **bind mount for static config, volume for application state.** For example, a database container would bind-mount `/etc/myapp/config.yaml` from the host and use a volume for `/var/lib/myapp/data`.

## Importing volumes between nodes

Under **Container Management → Volumes**, the **Import** action copies a volume's definition (name, encryption flag) from another node or cluster. This is useful for replicating a container configuration to a new edge — import the volumes first, then the container that references them.

Importing copies the **definition**, not the data. The new volume starts empty.

## Related

- [Container Tools]({{<ref "../tools">}}) — how to access the container filesystem from the portal Terminal
- [Volumes reference]({{<ref "../volumes">}}) — field-by-field reference for the Volumes screen
- [Tutorial: bind-mount a config file]({{<ref "/tutorials/containers/bind-mount-config">}}) — end-to-end walkthrough
