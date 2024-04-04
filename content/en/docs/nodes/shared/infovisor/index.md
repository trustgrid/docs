---
title: Infovisor
linkTitle: Infovisor
description: The infovisor displays details about the Trustgrid node
weight: 8
aliases:
    - /docs/nodes/infovisor
---

The infovisor displays details about the Trustgrid node and can be displayed by clicking the Info button in the top right or hitting the backtick `` ` `` key.
{{<tgimg src="info-button.png" caption="Button to open Infovisor" width="25%">}} 

It displays the following information in a card layout:
{{<tgimg src="infovisor.png" caption="Example infovisor" width="80%">}}

It contains the below information but can vary between appliance and agent-based nodes. The table below lists the fields and if they are displayed depending on node type.

|Field | Agent | Appliance |
|---|---|---|
|Name|✅|✅|
|Control and Data Plane Status| ✅ |✅|
|Public IP - Source IP observed by the Trustgrid control plane| ✅ |✅|
|Version| ✅ |✅|
|Package Version| ❌ |✅|
|[Enabled / Disabled Status]({{<ref "/tutorials/management-tasks/changing-node-status" >}})| ✅ |✅|
|Device Type | Show "Agent" as type|✅|
|Operating System (OS) Version| ✅ |✅|
|[TGRN or Trustgrid Resource Name]({{<ref "/docs/user-management/policies#trustgrid-resource-names-or-tgrn">}})| ✅ |✅|
| Upgrade Status| ❌ |✅|
|[Tags]({{<ref "/docs/nodes/shared/tags">}})| ✅ |✅|
| <ul><li>AWS metadata (AWS Nodes only)</li><ul><li>Availability Zone</li><li>Instance ID</li><li>Instance type</li></ul></ul>| ❌ |✅|

