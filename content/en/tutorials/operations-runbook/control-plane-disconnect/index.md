---
Title: "Control Plane Disconnect"
Date: 2023-1-9
Tags: ["control plane disconnect", "help", "troubleshoot"]
---

> When the Control Plane is disconnected there is no way to utilize remote tools to resolve the issue so you will need to contact the End-user technical resource for the site to troubleshoot

When troubleshooting the control plane it is a good idea to familiarize yourself with the [Edge Node Startup Process]({{<relref "/help-center/kb/startup-process">}}). 

<table>
<thead>
<tr>
<th> Network Requirements </th>
</tr>
<tr>
<td>
In order to connect to the Trustgrid Control Plane, the following <strong>outbound traffic</strong> must be allowed from the node’s configured primary interface IP address <br/>
<ul>
    <li>TCP Port 443 and TCP/UDP 8443 to:</li>
        <ul>
            <li>35.171.100.16/28</li>
            <li>34.223.12.192/28</li>
        </ul>
    </li>
    <li>TCP/UDP Port 53 to the configured DNS servers. These DNS servers must be able to resolve DNS requests for the trustgrid.io domain</li>
</ul>
</table>

## Troubleshoot from the Node
1. Triage the total site connectivity to see if actions can be taken to restore functionality for the edge site while troubleshooting the specific node
1. Confirm with the site tech:
    1. There are no known power or internet issues at the site
    1. No changes have been made to any firewalls between the Trustgrid node and the internet (if applicable). To connect the Trustgrid node must have access to the Network Requirements defined above.
1. Have the site tech attempt to ping the inside interface IP address(es) to see if the device is showing as powered up and online. If the site is using a single-interface configuration this would be the Network Adapter 1 - WAN Interface IP(s) in the portal.
    1. If the ping is successful you have determined the device has power and that the operating system and Trustgrid software are running. In this case, you can focus on internet side issues.
    1. If the ping fails, work with the site tech to:
        1. Confirm the node is powered on
        1. Connect directly to the network of the inside interface and attempt ping from there. They should also connect directly to the inside interface and statically configure an IP in the same network.
1. Attempt power cycling the node by removing power and reconnecting for physical devices, or using the hypervisor management tools for virtual nodes.
1. Connect to the console of the device
    1. A normal node looks something like this: {{<tgimg src="normal-node.png" width="80%" caption="Normal node login screen">}} Work with the onsite tech to log in to the [Trustgrid Local Console Utility]({{<relref "/tutorials/local-console-utility#accessing-the-trustgrid-console-utility">}}).  This tool will display the connectivity status and allow you to alter the WAN/outside IP settings if needed. From the console, you can also use the [Network Tools Shell]({{<relref "/tutorials/local-console-utility/troubleshooting#network-tools-shell">}}) to do additional troubleshooting including:
        1. Use `ping` to confirm you can ping the WAN interface's default gateway IP.  Note that not all default gateways will respond to ping but most do.
        1. Use `dig @<dns server IP> zuul.trustgrid.io +short` to confirm the DNS server is resolving to an IP address. Replace `<dns server IP>` with the DNS server configured on the WAN interfaces. If this does not work, work with the site tech to confirm the DNS server is configured correctly.
        1. Use `openssl` to confirm the [TLS certificate chain to the Trustgrid Control Plane is valid]({{<relref "help-center/kb/startup-process/ssl-tls-tampering#verification">}}).  
    1. If you see a screen like below attempt rebooting the device to restore connectivity and contact Trustgrid support so we can investigate further. {{<tgimg src="node-panic.png" width="80%" caption="Kernel panic screen">}}

## Troubleshoot from an Independent Device in place of Node
If connectivity cannot be established from the node, you can attempt to connect from an independent device like a laptop plugged into the port the node was plugged into.

1. Disconnect the cable from the WAN/Outside port of the Trustgrid node and connect to a laptop NIC.  Statically assign the same **IP and DNS settings** that the Trustgrid node is using. 
1. Confirm DNS is functioning:
    1. Using `nslookup` (on Windows) or `dig` (Linux or MacOS) to confirm you can resolve zuul.trustgrid.io
    1. Open a browser and navigate to https://zuul.trustgrid.io:8443
        1. If the device can connect to that server and port you should see a warning like this because Trustgrid uses its own Certificate Authority (CA) {{<tgimg src="link.png" width="50%" caption="Certificate security warning">}}
        1. Click `Not Secure` and then click `Certificate (invalid)` to view the certificate chain. {{<tgimg src="certificate-invalid.png" width="50%" caption="Expanded certificate security warning">}}
        1. You should expect to see a chain like the below example: {{<tgimg src="chain.png" width="70%">}}If any different certificates or CAs indicate something like DPI-SSL/HTTPS Proxy is [interfering with the TLS Certificate Chain]({{<relref "help-center/kb/startup-process/ssl-tls-tampering#resolution">}}).  
        1. If the browser says it cannot connect this indicates a firewall or routing issue upstream.


## Troubleshoot from an Upstream Network Device
If the site contact has management access to the device between the node and the internet, like a firewall or router, you can attempt to capture relevant traffic to determine where the issue lies. Specifically, look for the following common issues:

{{<alert color="warning">}}When capturing traffic it is recommended to perform the capture on the interfaces **closest to the Trustgrid node**.  This ensures you see the complete impact of firewall, routing, and NAT rules.  
For example, if your firewall has a WAN interface connected to the internet and a LAN interface connected to the network the node is attached to, use the LAN interface for captures.  Capturing from the WAN interface would give an incomplete picture. A WAN interface may receive a SYN/ACK response, but due to the firewall configuration that packet may not be put on the LAN network as expected.{{</alert>}}

1. Blocked DNS Access - capture TCP & UDP on port 53 (DNS) to confirm DNS requests are being sent to the upstream DNS server and responses are being delivered to the node.
1. Confirm Control Plane Traffic - capture TCP port 8443 to the [Trustgrid Control Plane]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane" >}}) networks. Common issues seen include:
    - Connection attempts are being blocked (access rules) or responses are not being put back on the network containing then node (routing or NAT rules)
    - DPI-SSL or HTTPS [altering the TLS certificate chain]({{<relref "/help-center/kb/startup-process/ssl-tls-tampering#symptoms">}})

