---
title: "Sniff Interface Traffic"
date: 2023-2-7
---

{{% pageinfo %}}
The sniff interface traffic tool performs a traffic capture on the selected node, which can help troubleshoot connectivity issues.
{{% /pageinfo %}}

## Summary

The Sniff Interface Traffic tool allows you to monitor traffic in real-time but there are certain instances where you need to be able to capture traffic to a file for further analysis. For example, intermittent issues may require leaving a capture running for an unknown period of time. Or, perhaps you need to review the application data within the packet payload.

The Traffic Capture is started with a standard tcpdump filter, just like the sniff interface tool, and max capture file size. Once the capture is stopped the captured data is written to .pcap files which can then be transferred for review.

{{<alert color="warning">}}The pcap files will include the data payload of any communication captured. If a non-encrypted protocol (such as HTTP or FTP) is used this could be visible in clear text.{{</alert>}}

Because of the risk of data being visible, there are several restrictions and controls in place:

1. The permission to run this tool (`node-advanced::service:tg-capture`) must be granted on a per-user basis. Not even administrators have this feature by default.
   ![img](pcap-permissions.png)
1. Trustgrid employees are specifically denied the ability to run this tool. Trustgrid employees are not permitted to view data transferred through our devices, so please do not send us the resulting pcap files.

## Accessing & Using Sniff Interface Traffic

1. Login to the Trustgrid portal and navigate to the Node from which you want to capture traffic.
1. Select `Interfaces` under the `Network` section.
1. Click the Sniff Traffic button

   ![img](network-tools.png)

1. Set the appropriate parameters.  
   {{<field-def "iface">}}The interface name{{</field-def>}}
   {{<field-def "filter">}}The tool utilizes [TCPDump filtering syntax](https://www.tcpdump.org/manpages/pcap-filter.7.html), which can help isolate interesting traffic. The below filter would show only ICMP and TCP traffic{{</field-def>}}
   ![img](filters.png)
1. Click `Execute` to view the captured traffic

![img](sniff-output.png)

## Transfer PCAP Files

Now that you have created the pcap files you will need to transfer them somewhere for analysis. The Trustgrid node has standard file transfer clients installed including scp, sftp, and ftp.

### On a Local Network

If the destination file server is on the local network of the node you will use the standard Terminal tool to run the file transfer commands.

1. Navigate to the node
1. Open a terminal from the toolbar at the top right

You will need to ensure that:

- The node has a valid route to the destination server's network
- Any firewall between the node and destination server allows the traffic
- The destination server is configured to allow the node IP to connect

Once you are in the terminal with the above configured an example scp command would look something like `scp ~/captures/capture-2021-07-16_17-59-23.pcap0 username@172.16.100.10:/captures/capture-2021-07-16_17-59-23.pcap`.

### Across the Virtual Network

If the destination file server is across a Trustgrid virtual network you will need to use the VPN Admin terminal to run the file transfer commands. This uses the nodes configured Virtual Management IP to communicate to remote resources.

1. Navigate to the node
1. Select the VPN panel. You may need to select the virtual network to use, if your device is connected to multiple.
1. Select the `Admin Terminal` button from the tools dropdown.

   ![img](admin-terminal.png)

You will need to ensure that:

- The node has a virtual management IP assigned. The tool will not launch without this.
- A route for the virtual management IP exists in the virtual network route table.
- On the destination node, that a 1:1 NAT exists for the destination file server assigning it to a virtual IP.
- A route that includes the destination file server's VIP exists in the network route table.
- A virtual network access policy exists that allows the source node's virtual management IP to connect to the destination file server's VIP on the required port for your protocol.
- The destination side node has a valid route to the destination server's network.
- Any firewall between the node and destination server allows the traffic.
- The destination is configured to allow the node IP to connect.

Once you are in the admin terminal with the above configured, an example SCP command would look like `scp ~/captures/capture-2021-07-16_17-59-23.pcap0 username@10.200.100.10:/captures/capture-2021-07-16_17-59-23.pcap`.

## Useful Filters

The below filters can be combined using "and" & "or" without quotes

- You can filter based on protocol such as "tcp", "udp", or "icmp" without quotes
- Best practice is to filter for what you want to see rather than filter out what you don't want to see
- To see only traffic for a specific host use "host x.x.x.x" with the IP address in place of the x.x.x.x
- To see only traffic for a specific port use "port XXXX" replacing the XXXX with the desired port number
- If capturing traffic on a clustered node, you should filter out the cluster port traffic. Typically port 9000 or 1975. E.g., `not port 9000`
- If capturing on the ETH0 - WAN Interface, you should filter out traffic to both Trustgrid's network (35.171.100.16/28) and your data plane gateways' network or public IP. You'll need to identify the actual network or IPs. E.g., `not net 35.171.100.16/28 and not net X.X.X.X/X` or `not net 35.171.100.16/28 and not host X.X.X.X`.
