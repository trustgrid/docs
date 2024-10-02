---
title: September 2024 Agent Release Notes
linkTitle: September 2024
type: docs
date: 2024-09-16
description: "September release of the Trustgrid Agent"
---
{{<agent-release package-version="0.2.20240905-2146" release="a-0.3.0">}}

## Disable UDP Tunnels
This release introduces an option to disable UDP tunnels for agent-based nodes. This is useful for nodes that are behind a NAT or firewall that does not support UDP. Agent-based nodes will continue to have UDP enabled by default.

## Other Fixes and Improvements
- Resolved an issue where agents running as containers would attempt to remove IP tables rules for layer 3 VPN forwarding that did not exist.
- Changed a the logging of closing a gateway connection from INFO to DEBUG to reduce noise in the logs.

