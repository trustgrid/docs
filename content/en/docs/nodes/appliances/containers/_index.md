---
title: "Containers"
aliases: 
  - /docs/nodes/containers
description: Configure containers to run on appliance-based nodes
---

Trustgrid nodes can run container images built to the Docker/OCI image spec, which allows for ease of deployment across an organization. Containers run with least-privilege defaults; workloads that need elevated access can opt into specific Linux Capabilities or full Privileged mode. [Container security]({{<relref "concepts/security">}}) covers the implications of each.

The container can be attached to both the local and virtual network space which allows both local and remote resources to communicate with the container. For example an API could be deployed on a Trustgrid Gateway which sends API Calls via the virtual network space to a container running on a Trustgrid Edge Node. The API call could then be translated to make a call to a database running on the local network and passed back up to the gateway host.

Before adding a container to a node, push an image to your [repository]({{<relref "repositories">}}). Pushes must be `linux/amd64`; [supported image platforms]({{<relref "repositories#supported-image-platforms">}}) covers the Apple-Silicon caveat.

Reading and managing containers requires `node-exec::read` and `node-exec::modify` permissions, respectively. Executing a container requires `node-exec::compute` permission.

## Where to start

- **New to running containers on Trustgrid?** Start with the [Container Quickstart]({{<relref "/tutorials/containers/quickstart">}}) — push an image, run it, reach it from the LAN in about 10 minutes.
- **Need to reach the container from a peer node over VPN?** Follow [Expose a container over a virtual network]({{<relref "/tutorials/containers/expose-over-vpn">}}).
- **Persist data?** See [Container storage]({{<relref "concepts/storage">}}).
- **Operating a running container — logs, terminal, restart?** See [Container Tools]({{<relref "tools">}}).

## Concepts

- [**Container networking**]({{<relref "concepts/networking">}})
- [**Container lifecycle**]({{<relref "concepts/lifecycle">}})
- [**Container storage**]({{<relref "concepts/storage">}})
- [**Container security**]({{<relref "concepts/security">}})

## Cluster scope vs node scope

Container configuration is set at the **cluster** level on a cluster, or at **node scope** on a standalone appliance. The runtime controls — Start, Stop, Logs, Terminal — are always at the node scope of the appliance running the container. The portal redirects you between these views as appropriate. See [Container Tools]({{<relref "tools">}}) for the full breakdown.

## Management

Navigate to **Container Management** under **Compute** on a node or cluster.

![Containers List View](containers-list.png)

Available actions from the **Actions** menu:

- **Add Container** — opens the Add Container modal.
- **Delete** — removes the selected container after a confirmation.
- **Enable** — enables the selected container so it will run.
- **Disable** — disables the selected container so it will not run.
- **Import** — copies a container definition from another node or cluster.
- **Export** — exports the selected container definitions to a file.

![Add Container Modal](add-container.png)

{{<fields>}}
{{<field "Name" >}}The name of the container.{{</field>}}
{{<field "Description" >}}Free-text description for the container.{{</field>}}
{{<field "Execution Type" >}}`Service`, `Recurring`, or `On Demand`. Determines when and how often the container runs and whether it restarts on exit. See [Container lifecycle]({{<relref "concepts/lifecycle">}}).{{</field>}}
{{<field "Status" >}}Only `Enabled` containers will run.{{</field>}}
{{<field "Image Name" >}}The name of the image to execute, in the form `<your-namespace>/<image>`.{{</field>}}
{{<field "Image Tag" >}}The image tag to execute.{{</field>}}
{{</fields>}}

## Overview

The overview section allows editing basic information about the container's execution environment.

{{<tgimg src="overview.png" caption="Container Overview" alt="Container overview section" width="90%">}}

{{<fields>}}
{{<field "Save Output">}}Persist standard output/standard error to the Trustgrid cloud for analysis. **It is the customer's responsibility to ensure no privileged information is included in the output.** See [Container security — Save Output]({{<relref "concepts/security#save-output">}}).{{</field>}}
{{<field "Schedule">}}Cron expression or rate (e.g. `rate(1 hour)`) that governs when the container runs. Shown only when **Execution Type** is `Recurring`. See [Container lifecycle — Recurring]({{<relref "concepts/lifecycle#recurring">}}) for the accepted formats.{{</field>}}
{{<field "Command">}}The command to execute inside the container. Overrides the image's entrypoint. Useful for troubleshooting.{{</field>}}
{{<field "Hostname">}}The hostname set inside the container. Defaults to the appliance's name.{{</field>}}
{{<field "Stop Time">}}Grace period (in seconds) to allow a container to stop before killing it. Defaults to 30 seconds.{{</field>}}
{{<field "User">}}Sets the username/group/UID/GID the container's main process runs as. See [Container security — User]({{<relref "concepts/security#user">}}).{{</field>}}
{{<field "DNS">}}DNS server for resolution inside the container. By default the container uses the appliance-side resolver at `172.18.1.2` which resolves other containers by name and forwards external lookups. See [Container networking — DNS resolver]({{<relref "concepts/networking#dns-resolver">}}).{{</field>}}
{{<field "IP">}}Pins the container to a specific IP in `172.18.0.0/16`. By default the address is assigned dynamically. See [Container networking — The container bridge]({{<relref "concepts/networking#the-container-bridge">}}).{{</field>}}
{{<field "Privileged">}}Grant the container extended privileges — disables most of the sandbox. **Almost no workload should need this.** Prefer [Linux Capabilities]({{<relref "concepts/security#linux-capabilities">}}).{{</field>}}
{{<field "Use Init">}}Run an init process as PID 1 inside the container. Recommended for any service that spawns child processes. See [Container security — Use Init]({{<relref "concepts/security#use-init">}}).{{</field>}}
{{<field "Require Connectivity">}}Gates container startup on the appliance having control-plane connectivity. Used with encrypted volumes. See [Container storage — Encrypted volumes]({{<relref "concepts/storage#encrypted-volumes">}}).{{</field>}}
{{</fields>}}

## Environment Variables

Environment variables can be added to a container to provide configuration at runtime.

![Environment Variables](envvars.png)

## Network

Configure the container's VRF, port mappings, virtual networks, and virtual interfaces. **Conceptual background:** [Container networking]({{<relref "concepts/networking">}}).

![Container Network](network.png)

### Host Port Mappings

Expose a port on the appliance to the container.

{{<fields>}}
{{<field "Protocol">}}`tcp` or `udp`. If unspecified, all traffic is forwarded.{{</field>}}
{{<field "Host Interface">}}The appliance NIC to listen on (e.g. `ens192`).{{</field>}}
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

See [Tutorial: expose a container over a virtual network]({{<relref "/tutorials/containers/expose-over-vpn">}}).

### Virtual Interfaces

Forward all traffic from an appliance-side virtual interface into the container as a dedicated interface.

{{<fields>}}
{{<field "Name">}}The virtual interface name on the appliance.{{</field>}}
{{<field "Destination">}}The interface name presented inside the container.{{</field>}}
{{</fields>}}

## Mounts

Persist data either as an externally defined [volume]({{<relref "volumes">}}), or a bind mount of the appliance's filesystem. **Conceptual background:** [Container storage]({{<relref "concepts/storage">}}).

![Container Mounts](mounts.png)

{{<fields>}}
{{<field "Type">}}`BIND` or `VOLUME`. For `VOLUME`, the source must reference an existing [volume]({{<relref "volumes">}}).{{</field>}}
{{<field "Source">}}For volumes, the volume name. For bind mounts, the absolute path on the appliance's filesystem.{{</field>}}
{{<field "Destination">}}The mount location inside the container.{{</field>}}
{{</fields>}}

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

Linux `ulimit`s can also be set per container. Supported ulimits: `CORE`, `DATA`, `FSIZE`, `LOCKS`, `MEMLOCK`, `MSGQUEUE`, `NICE`, `NOFILE`, `NPROC`, `RSS`, `RTPRIO`, `RTTIME`, `SIGPENDING`, `STACK`.

## Health Check

Configure a probe that monitors container health. A health check periodically runs a command inside the container and uses its exit code to mark the container `Healthy` or `Unhealthy` in the portal. It's a reporting mechanism only — the container is not automatically restarted on failure.

![Container Health Check](health-check.png)

{{<fields>}}
{{<field "Command">}}Command to run. Non-zero exit means failure.{{</field>}}
{{<field "Interval">}}Seconds between checks.{{</field>}}
{{<field "Timeout">}}Seconds to wait for a check; timeout counts as failure.{{</field>}}
{{<field "Start Period">}}Grace seconds during container startup before checks begin.{{</field>}}
{{<field "Retries">}}Failures allowed before marking the container unhealthy.{{</field>}}
{{</fields>}}

## Linux Capabilities

Add or drop specific Linux capabilities. Always prefer this over enabling `Privileged`. See [Container security — Linux Capabilities]({{<relref "concepts/security#linux-capabilities">}}).

![Container Linux Capabilities](capabilities.png)

## Logging Configuration

Rotate persisted container logs (when **Save Output** is enabled).

![Container Logging Configuration](logging.png)

{{<fields>}}
{{<field "Max File Size (MB)">}}Maximum log file size before rotation.{{</field>}}
{{<field "Max Files">}}Maximum number of rotated log files to keep.{{</field>}}
{{</fields>}}
