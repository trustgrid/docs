---
Title: "Create Virtual Network Overlay"
Date: 2023-1-13
---
{{% pageinfo %}}
Trustgrid utilizes a virtual network overlay to overcome many of the challenges often faced with traditional ipsec vpn deployments.  This virtual network will be used to carve out smaller networks for each edge or gateway node deployment.
{{% /pageinfo %}}

1. Navigate to the `Domain`. 

2. Click on `Virtual Networks`

![img](add-vn.png)

3. Give the virtual network a name, IP space in CIDR format and a description and then click `Save`.

![img](add-network.png)

4. Ensure the network you create will be large enough for all planned and future connections.

