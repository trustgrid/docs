---
title: "Node Agent Deployments"
linkTitle: "Agent Deployments"
Weight: 11
---



## Supported Operating Systems

### Linux 

- Ubuntu 22.04 LTS x86-64/AMD64

{{<alert color="info">}}At this time only x86-64/AMD64 based operating systems are supported.{{</alert>}}

### Container
The Trustgrid agent can be deployed as a container on Linux systems. This provides an isolated environment for the agent without requiring a full operating system instance. 

Trustgrid provides a Docker image that contains the agent pre-configured and optimized to run in Docker or similar container hosting environments like Kubernetes.

## Agent Installation
Installation instructions are also shown in the portal after [adding an agent]({{<ref "/tutorials/agent-deploy/_index.md#adding-agents">}}).
{{< tabpane highlight=true >}}
{{< tab header="Ubuntu 22.04" text=true >}}
Follow the process below to install on Ubuntu 22.04 Jammy Jellyfish operating systems:

{{< readfile file="ubuntu2204-install.md" >}}

{{< /tab >}}
{{< /tabpane >}}


