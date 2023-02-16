---
categories: ["virtual machines"]
title: "Virtual Machines"
linkTitle: "Virtual Machines"
---

Trustgrid nodes can deploy virtual machines (if supported) using the open source virtualization technology `Kernel-base Virtual Machine (KVM)`which allows the node to act as a hypervisor to run multiple, isolated virtual guests.

The virtual machine can be attached to both the local and virtual network space which allows both local and remote resources to communicate with the virtual machine. For example an API could be deployed on a Trustgrid Gateway which sends API calls via the virtual network space to a virtual machine running on a Trustgrid Edge Node. The API call could then be translated to make a call to a database running on the local network and passed back up to the gateway host.

{{<alert color="info">}}
Before being able to deploy a virtual machine to a node, an [image]({{<ref "images">}}) needs to be created.
{{</alert>}}

## Management

Navigate to `VM Management` > `Virtual Machines` on a node.

![img](vm_manage.png)

![img](vm_add.png)

{{<fields>}}
{{<field "Name">}}
Name of the virtual machine.
{{</field>}}
{{<field "Description">}}
Description for the virtual machine.
{{</field>}}
{{<field "Image">}}
Image selected from the ones available.
{{</field>}}
{{<field "Firmware">}}
Boot firmware for the virtual machine.
1. BIOS
1. UEFI
{{</field>}}
{{<field "Device Bus">}}
Device Bus used for deployment of the virtual machine.
1. VirtIO
1. IDE
1. SATA
{{</field>}}
{{</fields>}}

## Overview

Navigating into a virtual machine, the overview section allows editing basic information about the virtual machine's execution environment.

![img](vm_overview.png)

{{<fields>}}
{{<field "Name">}}
Name of the virtual machine.
{{</field>}}
{{<field "Description">}}
Description for the virtual machine.
{{</field>}}
{{<field "Image">}}
Image selected from the ones available.
{{</field>}}
{{<field "Guest OS (read-only)">}}
Operating System of the Image.
{{</field>}}
{{<field "Hostname (Optional)">}}
Hostname given to the virtual machine on deploy.
{{</field>}}
{{<field "Stop Timeout (Optional)">}}
Denotes how long to wait for the virtual machine to gracefully stop before killing the process.  Default is 2 minutes.
{{</field>}}
{{<field "Firmware (read-only)">}}
Boot firmware of the virtual machine.
{{</field>}}
{{<field "Device Bus (read-only)">}}
Device Bus used for deployment of the virtual machine.
{{</field>}}
{{<field "Auto Start (Optional)">}}
Indicates the virtual machine should automatically start after deploy.
{{</field>}}
{{</fields>}}

## CPU

Virtual machine CPU allocation.

![img](vm_cpu.png)

## Memory

Virtual machine Memory allocation in GB, MB, KB, or Bytes.

![img](vm_mem.png)

## Network

The networking section allows you to configure the virtual machine's VRF, its port mappings, and its virtual networks and interfaces.

![img](vm_net.png)

### Host Port Mappings

Host port mappings allow you to expose a port on the host to the virtual machine. This is useful for exposing a service running in the container to the local network.

{{<fields>}}
{{<field "Protocol">}}The protocol to listen for. If not specified, all traffic is forwarded to the virtual machine.{{</field>}}
{{<field "Host Interface">}}The host interface to listen on.{{</field>}}
{{<field "Host Port">}}The host port to listen on.{{</field>}}
{{<field "Virtual Machine Port">}}The virtual machine port that will receive the mapped traffic.{{</field>}}
{{</fields>}}

### Virtual Networks

Attaching a virtual network to a virtual machine allows virtual network traffic to reach it.

{{<fields>}}
{{<field "Virtual Network">}}The virtual network to attach.{{</field>}}
{{<field "Virtual IP">}}The virtual IP to assign to the virtual machine.{{</field>}}
{{<field "Allow Outbound">}}Whether the virtual machine should be allowed to make outbound connections into the virtual network.{{</field>}}
{{</fields>}}

### Virtual Interfaces

A virtual interface can be mapped to a virtual machine to forward all traffic.

{{<fields>}}
{{<field "Name">}}The virtual interface name.{{</field>}}
{{<field "Destination">}}The interface destination inside the virtual machine.{{</field>}}
{{</fields>}}

## Volumes

Allows mapping disks or cd-roms to a virtual machine.

![img](vm_vols.png)

## Snapshots

Allows the creation, restoration, and deletion of virtual machine snapshots.

![img](vm_snapshots.png)

{{<fields>}}
{{<field "Name">}}The snapshot name.{{</field>}}
{{<field "Creation Time">}}The timestamp when the snapshot was created.{{</field>}}
{{<field "VM State">}}The state of the virtual machine when the snapshot was created.{{</field>}}
{{<field "Snapshot Mode">}}
1. Extenal Disk Only - Mode applied to the snapshot when the firmware of the virtual machine is UEFI.
1. Internal - Mode applied to the snapshot when the firmware of the virtual machine is BIOS.  This mode includes the guest virtual machine state.
{{</field>}}
{{<field "Parent Snapshot">}}If this is a child snapshot, this references the parent snapshot.{{</field>}}
{{<field "Description">}}The snapshot description.{{</field>}}
{{</fields>}}

Every snapshot operation is tracked asynchronously.  You can keep track of the status of every snapshot creation, restoration, and deletion via the `Recent Tasks` table.

![img](vm_snapshots_tasks.png)

## Cloud Init

Cloud-init is a simple and powerful way to configure virtual machines during deployment.  You can manage user creation, install updates, configure network, etc.

{{<alert color="info">}}
You can find the full cloud-init documentation **[here](https://cloudinit.readthedocs.io/en/latest)**
{{</alert>}}

![img](vm_cloud_init.png)


***

## History

Shows active and completed traffic flows for the virtual machine.

{{<alert color="info">}}
Active flows can be manually terminated.
{{</alert>}}

![img](vm_history_flows.png)
