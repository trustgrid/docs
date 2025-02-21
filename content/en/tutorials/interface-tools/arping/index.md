---
title: "ARP Ping Interface Tool"
linkTitle: "ARP Ping"
description: "Address Resolution Protocol (ARP) Ping"
---

The `arping` tool allows you to send an ARP request to a specific IP address on a local interface. This tool is useful for troubleshooting connectivity issues such as confirming if the interface is in the same layer 2 broadcast domain as the target IP address.

## Using ARP Ping

1. Login to the Trustgrid portal and navigate to the Node from which you want to send an ARP request.
1. Select `Interfaces` under the `Network` section. Then from the dropdown select the interface you want to send the ARP request from. 
1. Click the `ARP Ping` button. {{<tgimg src="arping-button.png" width="65%">}}
1. Confirm the interface ID and set the IP address to ping. {{<tgimg src="arping-ip.png" width="65%">}}
1. Click `Execute` to send the ARP request. 
1. The output will display the ARP response. 
    1. If an ARP response is returned you should see a single MAC address returned. If multiple MAC addresses are returned this indicates an IP conflict on the network.{{<tgimg src="arping-output.png" width="65%" caption="Example output of successful arping">}}
    1. If  no ARP response is returned you will see a `timeout` message. This indicates that the IP address is not in the same layer 2 broadcast domain as the interface. {{<tgimg src="arping-timeout.png" width="25%" caption="Example output of arping timeout">}}