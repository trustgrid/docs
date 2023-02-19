---
title: "Containers"
---

Trustgrid nodes support running Docker containers which allows for ease of deployment across an organization. Any Docker container is supported with the exception of containers that require root level access.

The container can be attached to both the local and virtual network space which allows both local and remote resources to communicate with the container. For example an API could be deployed on a Trustgrid Gateway which sends API Calls via the virtual network space to a container running on a Trustgrid Edge Node. The API call could then be translated to make a call to a database running on the local network and passed back up to the gateway host.

Before adding a container to a node, push an image to your [repository]({{<ref "repositories">}}).

Reading and managing containers requires `node-exec::read` and `node-exec::modify` permissions, respectively. Executing a container requires `node-exec::compute` permission.

## Management

Navigate to `Container Management` under `Compute` on a node or cluster.

![Containers List View](containers-list.png)

Here you can add, enable, disable, delete, and import a container.

![Add Container Modal](add-container.png)

{{<fields>}}
{{<field "Name" >}}The name of the container.{{</field>}}
{{<field "Execution Type" >}}A container can be deployed as one of three types.

1. **Service** - the container is created and will run as a daemon. If the node reboots, the container will start up automatically.
1. **Recurring** - the container will run on a defined schedule. A schedule parameter is defined for the frequency, either as a rate or a [cron expression](https://crontab.guru/examples.html).

   | Rate               | Description          |
   | ------------------ | -------------------- |
   | `rate(30 minutes)` | Run every 30 minutes |
   | `rate(1 hour)`     | Run every hour       |
   | `rate(1 day)`      | Run every day        |

1. **On Demand** - typically used for testing. The container is executed for a single session, but will not restart on its own.
   {{</field>}}
   {{<field "Status" >}}Only enabled containers will run.{{</field>}}
   {{<field "Image Name" >}}The name of the image to execute.{{</field>}}
   {{<field "Image Tag" >}}The image tag to execute.{{</field>}}
   {{</fields>}}

## Overview

Navigating into a container, the overview section allows editing basic information about the container's execution environment.

![Container Overview](overview.png)

In addition to the fields above, you can modify:

{{<fields>}}
{{<field "Command">}}The command to execute inside the container.{{</field>}}
{{<field "Hostname">}}The hostname set inside the container.{{</field>}}
{{<field "Stop Time">}}The grace period (in seconds) to allow a container to stop before killing it. Defaults to 30 seconds.{{</field>}}
{{<field "User">}}Sets the username or group or uid or gid in the container.{{</field>}}
{{<field "Save Output">}}Persist standard output/standard error.{{</field>}}
{{<field "Privileged">}}Grant the container extended privileges.{{</field>}}
{{<field "Use Init">}}Use an init process inside the container as PID 1. This ensures responsibilities of an init system are performed inside the container (e.g., handling exit signals).{{</field>}}
{{<field "Require Connectivity">}}Ensures that the container will not start if it has encrypted volumes and is unable to reach the control plane.{{</field>}}
{{</fields>}}

## Environment Variables

Environment variables can be added to a container to provide configuration information to the container at runtime.

![Environment Variables](envvars.png)

## Network

The networking section allows you to configure the container's VRF, its port mappings, and its virtual networks and interfaces.

![Container Network](network.png)

### Host Port Mappings

Host port mappings allow you to expose a port on the host to the container. This is useful for exposing a service running in the container to the local network.

{{<fields>}}
{{<field "Protocol">}}The protocol to listen for. If not specified, all traffic is forwarded to the container.{{</field>}}
{{<field "Host Interface">}}The host interface to listen on.{{</field>}}
{{<field "Host Port">}}The host port to listen on.{{</field>}}
{{<field "Container Port">}}The container port that will receive the mapped traffic.{{</field>}}
{{</fields>}}

### Virtual Networks

Attaching a virtual network to a container allows virtual network traffic to reach it.

{{<fields>}}
{{<field "Virtual Network">}}The virtual network to attach.{{</field>}}
{{<field "Virtual IP">}}The virtual IP to assign to the container.{{</field>}}
{{<field "Allow Outbound">}}Whether the container should be allowed to make outbound connections into the virtual network.{{</field>}}
{{</fields>}}

### Virtual Interfaces

A virtual interface can be mapped to a container to forward all traffic.

{{<fields>}}
{{<field "Name">}}The virtual interface name.{{</field>}}
{{<field "Destination">}}The interface destination inside the container.{{</field>}}
{{</fields>}}

## Mounts

Mounts allow a container to persist data either as an externally defined [volume]({{<ref "volumes">}}), or a bind mount of the node's filesystem.

![Container Mounts](mounts.png)

{{<fields>}}
{{<field "Type">}}Either `BIND` or `VOLUME`. For type `VOLUME`, the mount must reference an existing [volume]({{<ref "volumes">}}).{{</field>}}
{{<field "Source">}}For volumes, the name of the volume. For bind mounts, the path on the node's filesystem.{{</field>}}
{{<field "Destination">}}The mount location inside the container.{{</field>}}
{{</fields>}}

## Resource Limits

Containers can be restricted to limit the amount of resources they can consume from the host.

![Container Resource Limits](limits.png)

{{<fields>}}
{{<field "CPU Max %">}}Maximum CPU allocation. Default is 50%.{{</field>}}
{{<field "Memory Max (MB)">}}Hard limit for RAM allocation. Default is 50% of the host's memory.{{</field>}}
{{<field "Memory High (MB)">}}Soft limit for RAM allocation. Cannot exceed hard limit. Default is 45% of the host's memory.{{</field>}}
{{<field "IO Max Read (B/s)">}}Max allowed bytes per second of IO reads. Disabled by default.{{</field>}}
{{<field "IO Max Write (B/s)">}}Max allowed bytes per second of IO writes. Disabled by default.{{</field>}}
{{<field "IO Max Read Operations (ops/s)">}}Max allowed IO read operations per second. Disabled by default.{{</field>}}
{{<field "IO Max Write Operations (ops/s)">}}Max allowed IO write operations per second. Disabled by default.{{</field>}}
{{</fields>}}

Linux ulimits can be set for each container. Supported ulimits are:

- CORE
- DATA
- FSIZE
- LOCKS
- MEMLOCK
- MSGQUE
- NICE
- NOFILE
- NPROC
- RSS
- RTPRIO
- RTTIME
- SIGPENDING
- STACK

## Health Check

A health check can be configured to monitor the container's health. If the health check fails, the container will be restarted.

![Container Health Check](health-check.png)

{{<fields>}}
{{<field "Command">}}The command to run. A non-zero return code indicates a health check failure.{{</field>}}
{{<field "Interval">}}The frequency (in seconds) to run the health check.{{</field>}}
{{<field "Timeout">}}How long (in seconds) to wait for the health check to complete. A timeout is considered a failure.{{</field>}}
{{<field "Start Period">}}Grace period (in seconds) during container startup before health checks should start.{{</field>}}
{{<field "Retries">}}Number of allowed health check failures before marking the container unhealthy.{{</field>}}
{{</fields>}}

## Linux Capabilities

Linux capabilities can be added to or removed from a container, allowing fine-grained control over kernel-level features and device access.

![Container Linux Capabilities](capabilities.png)

## Logging Configuration

Log files (when persisting container output) are rotated based on a size threshold.

![Container Logging Configuration](logging.png)

{{<fields>}}
{{<field "Max File Size (MB)">}}The maximum size (in MB) of a log file before it is rotated.{{</field>}}
{{<field "Max Files">}}The maximum number of log files to keep.{{</field>}}
{{</fields>}}
