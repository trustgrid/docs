---
categories: ["virtual machines"]
title: "Volumes"
linkTitle: "Volumes"
---

Volumes are used to allocate disks or to attach CD-ROMs to virtual machines.

![img](vm_volumes.png)

![img](vm_add_volume.png)

{{<fields>}}
{{<field "Name">}}
Name for the volume.
{{</field>}}
{{<field "Device Type">}}
Virtual machine type.
1. Disk
1. CD-ROM
{{</field>}}
{{<field "Device Bus">}}
Bus for the device.
1. IDE
1. VirtIO
1. SCSI
1. SATA
{{</field>}}
{{<field "Size (Only for Disk)">}}
Size in GB, MB, KB, or Bytes for the disk.
{{</field>}}
{{<field "Disk Provisioning (Only for Disk)">}}
Allocation method used for provisioning a disk.
1. Thin Provision - Disk starts small and expands as more disk space is required up to the size specified. Use  this format to save storage space.
2. Thick Provision Lazy Zeroed - Space required for the  virtual disk is allocated when the virtual disk  is created. Data that remains on the physical device is not erased during creation, but is zeroed out on demand at a later time on first write from the virtual  machine.
3. Thick Provision Eager Zeroed - Space required for the virtual disk is allocated at creation time. The data remaining on the  physical device is zeroed out when the virtual disk  is created. It  might take much longer to create disks in this format than to create other types of disks.
{{</field>}}
{{<field "Encrypted (Only for Disk)">}}
True if disk should be encrypted, false otherwise.
{{</field>}}
{{<field "ISO File (Only for CD-ROM)">}}
ISO file selected from the node data store.
{{</field>}}
{{</fields>}}

{{<alert color="info">}}
Windows virtual machines do not support thin provisioned disks.
{{</alert>}}