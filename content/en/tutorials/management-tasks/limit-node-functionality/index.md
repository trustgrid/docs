---
Title: "Limit Node Functionality to Current Public IP"
Date: 2023-01-09
Description: "This tutorial will show you how to lock a node to the current public IP address."
aliases: 
    - /tutorial/limit-node-functionality
---

{{% pageinfo %}}
This security feature allows restricting full node functionality to the current public IP address. If the public IP changes the data plane connectivity will cease to function and no data plane traffic will pass. It is the equivalent to disabling a node in the trustgrid portal. 
{{% /pageinfo %}}

{{<alert color="warning">}}This feature should not be used on networks where the Public IP is controlled via DHCP. It should only be used where the public address is statically assigned to the node and is not expected to change. {{</alert>}}
## Process to Restrict
1. Navigate to the node in the Trustgrid portal
1. Click the **Actions** button in the top right corner of the page and select **Lock Node IP** {{<tgimg src="lock-node-ip.png" caption="Lock Node IP" width="25%">}}

## Process to Unlock
1. Navigate to the node in the Trustgrid portal
1. Click the **Actions** button in the top right corner of the page and select **Unlock Node IP** {{<tgimg src="unlock-node-ip.png" caption="Lock Node IP" width="25%">}}

## Alerts
Once locked changing the Public IP of the node will result in an alert being generated as seen belo. At this point no data plane traffic will be allowed.
{{<tgimg src="alert-node-public-ip.png" caption="Alert generated when node attempts to connect from an unauthorized IP address.">}}
