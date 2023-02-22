---
tags: ["applications", "ztna"]
title: "RDP App"
date: 2023-01-11
---

{{% pageinfo %}}
A RDP app is a ZTNA application that allows remote access to an internal RDP server. RDP servers can be hosted internally but exposed to authorized users.
{{% /pageinfo %}}

#### General

{{<fields>}}
{{<field "Name" >}}
app name
{{</field >}}
{{<field "Description" >}}
app description
{{</field >}}
{{<field "Icon" >}}
the application's icon (optional) to show in the application dashboard
{{</field >}}
{{</fields>}}

#### Connectivity

- Connectivity type:
  - local to gateway - the application is hosted on the same network as the gateway
  - remote node - the application is hosted on an edge node's network
  - virtual network - the application is accessible over the Trustgrid virtual network from the ZTNA gateway
- ZTNA Gateway - the ZTNA gateway node that will be used to connect to the application
- Destination Node - only available if connectivity type is Remote Node. The edge node with access to the application
- VRF - only available if connectivity type is Remote Node. The VRF used to connect to the application.

- Internal server hostname or IP - the internal hostname or IP address of the RDP server and port number
- Virtual Network - only available if connectivity type is Virtual Network. The virtual network that will be used to connect to the application
- Client Virtual IP - only available if connectivity type is Virtual Network -defaults to the virtual management IP of the ZTNA gateway if it exists. The source virtual IP address that will be used to connect to the application
- Virtual server hostname or IP - only available if connectifity type is Virtual Network. The destination virtual IP address and port number that will be used to connect to the application

#### Security

- Identity Provider - the [Identity Provider]({{<ref "docs/idps" >}}) to authenticate users

1. Click `Add Application`

![img](add-app.png)

2. Click the `Remote Desktop - RDP` option

![img](rdp.png)

3. Fill out all details in the sheet provided, then click `Save` when satisfied with the information entered

![img](rdp-app1.png)
