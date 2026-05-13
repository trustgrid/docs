---
title: "Containers"
aliases: 
  - /docs/nodes/containers
description: Configure containers to run on appliance-based nodes
---

Trustgrid nodes can run container images built to the Docker/OCI image spec, which allows for ease of deployment across an organization. Containers run with least-privilege defaults; workloads that need elevated access can opt into specific [Linux Capabilities]({{<ref "concepts/security#linux-capabilities">}}) or full [Privileged]({{<ref "concepts/security#privileged">}}) mode — see [Container security]({{<ref "concepts/security">}}) for the implications of each.

The container can be attached to both the local and virtual network space which allows both local and remote resources to communicate with the container. For example an API could be deployed on a Trustgrid Gateway which sends API Calls via the virtual network space to a container running on a Trustgrid Edge Node. The API call could then be translated to make a call to a database running on the local network and passed back up to the gateway host.

Before adding a container to a node, push an image to your [repository]({{<ref "repositories">}}). Pushes must be `linux/amd64` — see [Repositories — Supported image platforms]({{<ref "repositories#supported-image-platforms">}}) for the Apple-Silicon caveat.

Reading and managing containers requires `node-exec::read` and `node-exec::modify` permissions, respectively. Executing a container requires `node-exec::compute` permission.

## Where to start

- **New to running containers on Trustgrid?** Start with the [Container Quickstart]({{<ref "/tutorials/containers/quickstart">}}) — push an image, run it, reach it from the LAN in about 10 minutes.
- **Need to reach the container from a peer node over VPN?** Follow [Expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}}).
- **Persist data or mount a config file?** See [Bind-mount a config file]({{<ref "/tutorials/containers/bind-mount-config">}}) and [Container storage]({{<ref "concepts/storage">}}).
- **Operating a running container — logs, terminal, restart?** See [Container Tools]({{<ref "tools">}}).
- **Running into an issue?** See [Container Troubleshooting]({{<ref "troubleshooting">}}).

## Concepts

The container model has four moving parts; each gets its own page:

- [**Container networking**]({{<ref "concepts/networking">}}) — bridge, resolver, VRF, port mappings, virtual networks, egress.
- [**Container lifecycle**]({{<ref "concepts/lifecycle">}}) — Service vs Recurring vs On Demand, restart behavior, Status vs State.
- [**Container storage**]({{<ref "concepts/storage">}}) — bind mounts vs volumes, encrypted volumes, Require Connectivity.
- [**Container security**]({{<ref "concepts/security">}}) — capabilities, privileged, user identity, init, save output.

## Cluster scope vs node scope

Container configuration is set at the **cluster** level on a cluster, or at the **node** level on a standalone node. The runtime controls — Start, Stop, Logs, Terminal — are always on the **node** that's running the container. The portal redirects you between these views as appropriate. See [Container Tools]({{<ref "tools">}}) for the full breakdown.

---

The rest of this page is a **field reference** for the Add Container modal and the per-container detail screens. For the *why* behind each section, follow the concept-page links.

## Management

Navigate to **Container Management** under **Compute** on a node or cluster.

![Containers List View](containers-list.png)

Here you can add, enable, disable, delete, and import a container.

![Add Container Modal](add-container.png)

{{<fields>}}
{{<field "Name" >}}The name of the container.{{</field>}}
{{<field "Execution Type" >}}`Service`, `Recurring`, or `On Demand`. Determines when and how often the container runs and whether it restarts on exit. See [Container lifecycle]({{<ref "concepts/lifecycle">}}).{{</field>}}
{{<field "Status" >}}Only `Enabled` containers will run.{{</field>}}
{{<field "Image Name" >}}The name of the image to execute, in the form `<your-namespace>/<image>`.{{</field>}}
{{<field "Image Tag" >}}The image tag to execute.{{</field>}}
{{</fields>}}

## Overview

The overview section allows editing basic information about the container's execution environment.

{{<tgimg src="overview.png" caption="Container Overview" alt="Container overview section" width="90%">}}

{{<fields>}}
{{<field "Save Output">}}Persist standard output/standard error to the Trustgrid cloud for analysis. **It is the customer's responsibility to ensure no privileged information is included in the output.** See [Container security — Save Output]({{<ref "concepts/security#save-output">}}).{{</field>}}
{{<field "Command">}}The command to execute inside the container. Overrides the image's entrypoint. Useful for troubleshooting.{{</field>}}
{{<field "Hostname">}}The hostname set inside the container. Defaults to the node name.{{</field>}}
{{<field "Stop Time">}}Grace period (in seconds) to allow a container to stop before killing it. Defaults to 30 seconds.{{</field>}}
{{<field "User">}}Sets the username/group/UID/GID the container's main process runs as. See [Container security — User]({{<ref "concepts/security#user">}}).{{</field>}}
{{<field "DNS">}}DNS server for resolution inside the container. By default the container uses the node-side resolver at `172.18.1.2` which resolves other containers by name and forwards external lookups. See [Container networking — DNS resolver]({{<ref "concepts/networking#dns-resolver">}}).{{</field>}}
{{<field "IP">}}Pins the container to a specific IP in `172.18.0.0/16` (cannot be `172.18.1.2`). By default the address is assigned dynamically. See [Container networking — The container bridge]({{<ref "concepts/networking#the-container-bridge">}}).{{</field>}}
{{<field "Privileged">}}Grant the container extended privileges — disables most of the sandbox. **Almost no workload should need this.** Prefer [Linux Capabilities]({{<ref "concepts/security#linux-capabilities">}}).{{</field>}}
{{<field "Use Init">}}Run an init process as PID 1 inside the container. Recommended for any service that spawns child processes. See [Container security — Use Init]({{<ref "concepts/security#use-init">}}).{{</field>}}
{{<field "Require Connectivity">}}Gates container startup on the node having control-plane connectivity. Used with encrypted volumes. See [Container storage — Encrypted volumes]({{<ref "concepts/storage#encrypted-volumes">}}).{{</field>}}
{{</fields>}}

## Environment Variables

Environment variables can be added to a container to provide configuration at runtime.

![Environment Variables](envvars.png)

## Network

Configure the container's VRF, port mappings, virtual networks, and virtual interfaces. **Conceptual background:** [Container networking]({{<ref "concepts/networking">}}).

![Container Network](network.png)

### Host Port Mappings

Expose a port on the node to the container.

{{<fields>}}
{{<field "Protocol">}}`tcp` or `udp`. If unspecified, all traffic is forwarded.{{</field>}}
{{<field "Host Interface">}}The host NIC to listen on (e.g. `ens192`).{{</field>}}
{{<field "Host Port">}}The host port to listen on.{{</field>}}
{{<field "Container Port">}}The container port that receives the mapped traffic.{{</field>}}
{{</fields>}}

### Virtual Networks

Attach a Trustgrid virtual network so peer nodes can reach the container.

{{<fields>}}
{{<field "Virtual Network">}}The virtual network to attach.{{</field>}}
{{<field "Virtual IP">}}The virtual IP to assign to the container.{{</field>}}
{{<field "Allow Outbound">}}Whether the container may originate connections onto the virtual network.{{</field>}}
{{</fields>}}

See [Tutorial: expose a container over a virtual network]({{<ref "/tutorials/containers/expose-over-vpn">}}).

### Virtual Interfaces

Forward all traffic from a node-side virtual interface into the container as a dedicated interface.

{{<fields>}}
{{<field "Name">}}The virtual interface name on the node.{{</field>}}
{{<field "Destination">}}The interface name presented inside the container.{{</field>}}
{{</fields>}}

## Mounts

Persist data either as an externally defined [volume]({{<ref "volumes">}}), or a bind mount of the node's filesystem. **Conceptual background:** [Container storage]({{<ref "concepts/storage">}}).

![Container Mounts](mounts.png)

{{<fields>}}
{{<field "Type">}}`BIND` or `VOLUME`. For `VOLUME`, the source must reference an existing [volume]({{<ref "volumes">}}).{{</field>}}
{{<field "Source">}}For volumes, the volume name. For bind mounts, the absolute path on the node's filesystem.{{</field>}}
{{<field "Destination">}}The mount location inside the container.{{</field>}}
{{</fields>}}

See [Tutorial: bind-mount a config file]({{<ref "/tutorials/containers/bind-mount-config">}}).

## Resource Limits

Restrict the resources a container can consume from the host.

![Container Resource Limits](limits.png)

{{<fields>}}
{{<field "CPU Max %">}}Maximum CPU allocation. Default is 50%.{{</field>}}
{{<field "Memory Max (MB)">}}Hard RAM limit. Default is 50% of host memory.{{</field>}}
{{<field "Memory High (MB)">}}Soft RAM limit. Cannot exceed hard limit. Default is 45% of host memory.{{</field>}}
{{<field "IO Max Read (B/s)">}}Max IO read bytes/sec. Disabled by default.{{</field>}}
{{<field "IO Max Write (B/s)">}}Max IO write bytes/sec. Disabled by default.{{</field>}}
{{<field "IO Max Read Operations (ops/s)">}}Max IO read ops/sec. Disabled by default.{{</field>}}
{{<field "IO Max Write Operations (ops/s)">}}Max IO write ops/sec. Disabled by default.{{</field>}}
{{</fields>}}

Linux `ulimit`s can also be set per container. Supported ulimits: `CORE`, `DATA`, `FSIZE`, `LOCKS`, `MEMLOCK`, `MSGQUE`, `NICE`, `NOFILE`, `NPROC`, `RSS`, `RTPRIO`, `RTTIME`, `SIGPENDING`, `STACK`.

## Health Check

Configure a probe that monitors container health. A failing health check restarts the container. **Conceptual background:** [Container lifecycle — Restart semantics]({{<ref "concepts/lifecycle#restart-semantics-in-detail">}}).

![Container Health Check](health-check.png)

{{<fields>}}
{{<field "Command">}}Command to run. Non-zero exit means failure.{{</field>}}
{{<field "Interval">}}Seconds between checks.{{</field>}}
{{<field "Timeout">}}Seconds to wait for a check; timeout counts as failure.{{</field>}}
{{<field "Start Period">}}Grace seconds during container startup before checks begin.{{</field>}}
{{<field "Retries">}}Failures allowed before marking the container unhealthy.{{</field>}}
{{</fields>}}

## Linux Capabilities

Add or drop specific Linux capabilities. Always prefer this over enabling `Privileged`. See [Container security — Linux Capabilities]({{<ref "concepts/security#linux-capabilities">}}).

![Container Linux Capabilities](capabilities.png)

## Logging Configuration

Rotate persisted container logs (when **Save Output** is enabled).

![Container Logging Configuration](logging.png)

{{<fields>}}
{{<field "Max File Size (MB)">}}Maximum log file size before rotation.{{</field>}}
{{<field "Max Files">}}Maximum number of rotated log files to keep.{{</field>}}
{{</fields>}}
