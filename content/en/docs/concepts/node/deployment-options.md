---
categories: ["concepts"]
tags: ["node", "concepts"]
title: "Node Deployment Options"
date: 2022-12-20

---

## Virtual Appliance

Trustgrid supports deploying nodes as a virtual appliance in vSphere 5.5+. The virtual appliance is pre-configured by Trustgrid and delivered as a zip file containing the OVF template and supporting files for deployment. For deployment information, see Deploy a Trustgrid [Node]({{< ref "/docs/concepts/node" >}}) in vSphere in the configuration section.

## Hardware Appliance

Trustgrid [nodes]({{< ref "/docs/concepts/node" >}}) can be deployed on various hardware provided that the hardware is amd64-based. Hardware nodes are typically deployed on small devices with 1 or 2 network interfaces, an Intel processor with at least 2 cores, at least 2 GB of RAM, and 32 GB of storage. Hardware nodes are delivered pre-configured by Trustgrid for easy "plug-n-play" deployment.

## Amazon AMI

Trustgrid provides a CloudFormation template that automates the provisioning of a Trustgrid node using an AWS AMI. For deployment information, see Deploy a Trustgrid [Node]({{< ref "/docs/concepts/node" >}}) in Amazon Web Services using an AMI in the configuration section.
