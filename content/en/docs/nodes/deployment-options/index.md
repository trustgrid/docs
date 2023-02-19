---
categories: ["concepts"]
tags: ["node", "concepts"]
title: "Node Deployment Options"
linkTitle: "Node Deployment Options"
---

### Virtual Appliance

Trustgrid supports deploying [nodes]({{<ref "docs/nodes" >}}) as a virtual appliance in vSphere 5.5+. The virtual appliance is pre-configured by Trustgrid and delivered as a zip file containing the OVF template and supporting files for deployment. For deployment information, see "Deploy a Trustgrid [Node]({{<ref "/docs/nodes" >}})" in vSphere in the configuration section.

### Hardware Appliance

Trustgrid [nodes]({{<ref "docs/nodes" >}}) can be deployed on various hardware provided that the hardware is amd64-based. Hardware [nodes]({{<ref "docs/nodes" >}}) are typically deployed on small devices with 1 or 2 [network]({{<ref "docs/overview/networking" >}}) interfaces, an intel processor with at least 2 cores, at least 2 GB of RAM, and 32 GB of storage. Hardware [nodes]({{<ref "docs/nodes" >}}) are delivered pre-configured by Trustgrid for easy "plug-n-play" deployment.

### Amazon AMI

Trustgrid provides a CloudFormation template that automates the provisioning of a Trustgrid [node]({{<ref "docs/nodes" >}}) using an AWS AMI. For deployment information, see "Deploy a Trustgrid [Node]({{<ref "docs/nodes" >}})" in Amazon Web Services using an AMI in the configuration section.
