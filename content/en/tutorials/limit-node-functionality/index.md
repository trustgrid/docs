---
Title: "Limit Node Functionality to Current Public IP"
Date: 2023-01-09
---

{{% pageinfo %}}
This security feature allows restricting node functionality to the current public IP address. If the public IP changes the data plane connectivity will cease to function and no data plane traffic will pass. It is the equivalent to disabling a node in the trustgrid portal. This feature should not be used on networks where the Public IP is controlled via DHCP. It should only be used where the public address is statically assigned to the node and is not expected to change. 
{{% /pageinfo %}}

## Process to Restrict
1. Click `Info` in the top right corner of the page, or click the backtick key (`) to show the menu
1. Click the padlock button next to Public IP to lock (Do not enable if using DHCP. Static addresses only) {{<tgimg src="unlocked1.png" caption="Click the padlock to lock the node to the current public IP address.">}}
1. Public IP should now show padlock icon as locked as shown below. {{<tgimg src="locked1.png" caption="Public IP is now locked to the current public IP address.">}}

## Process to Unlock
1. Click `Info` in the top right corner of the page, or click the backtick key (`) to show the menu
1. Click the padlock next to "Locked IP" to unlock.

## Alerts
Once locked changing the Public IP of the node will result in an alert being generated as seen belo. At this point no data plane traffic will be allowed.
{{<tgimg src="alert-node-public-ip.png" caption="Alert generated when node attempts to connect from an unauthorized IP address.">}}
