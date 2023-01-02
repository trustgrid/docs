---
categories: ["concepts"]
tags: ["node", "concepts"]
title: "Node"
linkTitle: "Node"
date: 2022-12-07
weight: 14
---

{{% pageinfo %}}
A node is an instance of software used for building connections, managing compute resources, and deploying software.
{{% /pageinfo %}}

A node can be installed on a hardware appliance, as a virtual appliance (Vsphere and Hyper-V), and deployed in a public cloud (Amazon AMI, Google, Microsoft). There are three main types of nodes - Edge, Gateway, and Management.

Edge nodes create TLS tunnels to all gateway nodes provisioned in a domain. Edge nodes also create TLS tunnels to Management nodes operated by Trustgrid.

Gateway nodes accept incoming tunnels from Edge nodes.

All Trustgrid nodes run on a hardened Ubuntu operating system. Nodes are managed through the Trustgrid portal or via the API (requires authentication token). A local UI exists to support the reconfiguration of an ethernet interface and diagnostic/network tests prior to connection to the Portal.
