---
Title: "Limit Node Functionality to Current Public IP"
Date: 2023-1-9
---

{{% pageinfo %}}
This security feature allows restricting node functionality to the current public IP address. If the public IP changes the data plane connectivity will cease to function and no data plane traffic will pass. It is the equivalent to disabling a node in the trustgrid portal. This feature should not be used on networks where the Public IP is controlled via DHCP. It should only be used where the public address is statically assigned to the node and is not expected to change. 
{{% /pageinfo %}}

###### Process to Restrict
- Click `Info` in the top right corner of the page, or click the backtick key (`) to show the menu
- Click the padlock button next to Public IP to lock (Do not enable if using DHCP. Static addresses only)

![img](unlocked1.png)

- Public IP should now show padlock icon as locked as shown below

![img](locked1.png)

- Click the padlock again to remove the restriction at any point.

###### Alerts:
Once locked changing the Public IP of the node will result in an alert being generated as seen belo. At this point no data plane traffic will be allowed.

![img](alert-node-public-ip.png)