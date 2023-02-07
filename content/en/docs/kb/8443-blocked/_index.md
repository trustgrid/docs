---
title: Edge Node Behavior When Port 8443 is blocked
---

## Symptoms

- Node will not connect to the portal
- Packet capture shows the following behavior:
- - DNS request and response for gatekeeper.trustgrid.io and zuul.trustgrid.io
- - TCP connection attempt to port 8443 on the returned IP address with no response
- - A 2 minute delay before the above cycle repeats

## Cause

- The Trustgrid edge node cannot connect to the [Trustgrid control plane]({{<ref "site-requirements">}}) over 8443

## Verification

- Place a device in the same network (or assume the IP address of) the edge node. Confirm you can make a successful connection to gatekeeper.trustgrid.io:8443 using telnet, netcat or similar tool

## Resolution

- Adjust firewall settings to ensure TCP traffic over port 8443 is allowed to all of the [Trustgrid control plane]({{<ref "site-requirements">}})
