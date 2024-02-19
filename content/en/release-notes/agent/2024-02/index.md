---
title: February 2024 Agent Release Notes
linkTitle: February 2024
type: docs
date: 2024-02-19
description: "Initial production release of the Trustgrid Agent"
---
{{<agent-release package-version="0.2.20240218-1985" release="a-0.1.0">}}

This is the initial release of the agent-based Trustgrid node software.  

## Production Features include:
- Support for Ubuntu 22.04
- Basic [agent VPN]({{<relref "/docs/nodes/agents/vpn">}}) functionality
- Layer-4(L4) Proxy support with [services]({{<relref "/docs/nodes/shared/services">}}) and [connectors]({{<relref "/docs/nodes/shared/connectors">}})

## Beta Support:
- Running the agent in Red Hat Enterprise Linux 9 and in containers like Docker/k8s 