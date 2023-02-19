---
categories: ["virtual machines"]
title: "Images"
linkTitle: "Images"
---

The images section allows a user to add/delete virtual machine images.

![img](vm_images.png)

{{<alert color="info">}}
The `Downloaded` column shows whether or not an image is present on the node.
{{</alert>}}

![img](vm_add_image.png)

{{<fields>}}
{{<field "Display Name">}}
Friendly name associated to the image.
{{</field>}}
{{<field "Description">}}
Description for the image.
{{</field>}}
{{<field "Image File">}}
Image file selected from the node data store.
{{</field>}}
{{<field "Operating System">}}
Operating System for the image.  This is rather important to optimize the virtual machine configuration.
{{</field>}}
{{</fields>}}

{{<alert color="info">}}
If an incorrect `Operating System` is selected for the image, the virtual machine might not be able to detect install media or perform critical device controller operations
{{</alert>}}

#### Delete Image

If the image is present on the node, executing a delete of the image will prompt the deletion of the image from the node.

![img](vm_image_delete_confirm.png)

1. Yes - The image will be attempted to be deleted from the node first before the image entry is removed.
1. No - The image will not be deleted from the node and the image entry will be removed.