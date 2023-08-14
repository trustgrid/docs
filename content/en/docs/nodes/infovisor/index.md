---
title: Infovisor
linkTitle: Infovisor
description: The infovisor displays details about the Trustgrid node
weight: 8
---

The infovisor displays details about the Trustgrid node and can be displayed by clicking the Info button in the top right or hitting the backtick `` ` `` key.
{{<tgimg src="info-button.png" caption="Button to open Infovisor" width="25%">}} 

It contains information such as:

* Name
* Control and Data Plane Status
* Public IP - This is the IP observed by the Trustgrid Control plane as the source IP by traffic from the node
* Version and Package Version - The Trustgrid software version running on the node
* [Enabled / Disabled Status]({{<ref "/tutorials/management-tasks/changing-node-status" >}})
* Device Type 
* Operating System (OS) Version
* [TGRN or Trustgrid Resource Name]({{<ref "/docs/user-management/policies#trustgrid-resource-names-or-tgrn">}})
* Upgrade Status
* [Tags]({{<ref "/docs/nodes/tags">}})
* On AWS Nodes additional information is reported based on the metadata service
    * Availability Zone
    * Instance ID
    * Instance type

{{<tgimg src="infovisor.png" caption="Example infovisor" width="80%">}}