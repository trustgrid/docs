---
categories: ["concepts"]
tags: ["VPN", "concepts"]
title: "VPN"
date: 2022-12-19
description: >
  Trustgrid VPN
---

## Summary
Trustgrid utilizes Network Address Translation to alter the source and/or destination IP address of Layer 3 traffic as it enters or leaves the Virtual Private Network (VPN). These NATs are defined under either an individual [node]({{< ref "node" >}}) or at the [cluster]({{< ref "cluster" >}}) level. If a [node]({{< ref "node" >}}) is part of a [cluster]({{< ref "cluster" >}}) only the NATs defined at the [cluster]({{< ref "cluster" >}}) level will be effective.

## Outside NATs
An Outside NAT will convert:

- The **Source IP** address of traffic leaving the Virtual Network. This determines how traffic appears inside the local network 
- The **Destination IP** of traffic entering the Virtual Network.

**An Outside NAT is required for all [nodes]({{< ref "node" >}})/[clusters]({{< ref "cluster" >}}) using the Virtual Network.**

## Inside NATs
An Inside NAT will convert:

- The **Source IP** address of traffic entering the Virtual Network. 
- The **Destination IP** address of traffic leaving the Virtual Network. 


