---
title: "TCP Port Test"
date: 2023-02-08
description: "Simple TCP port connectivity test tool on using local IPs"
---

{{% pageinfo %}}
The TCP Port Test tool initiates a TCP session with a target IP address and port to confirm fully layer 4 connectivity. This would be similar to using `netcat` or `telnet` to test connectivity to a TCP port.
{{% /pageinfo %}}

## Usage

1. Login to the Trustgrid portal and navigate to the Node from which you want to test connectivity.
1. Select **Interfaces** under the **Network** section.
1. Click the **Interface TCP Port Test** button {{<tgimg src="network-tools.png" width="85%" caption="Selecting TCP Port Test">}}
1. Update the host with the target IP address and the port with the target TCP port. Click **Execute** to test connectivity. {{<tgimg src="config.png" width="60%" caption="Configuring the TCP Port Test">}}
1. A new window will open with the results. If a new window does not open, check your browser's pop-up blocker settings.
   - A successful test will look like this: {{<tgimg src="success.png" width="85%" caption="Successful TCP Port Test result">}}
   - A failed connection will depend on the nature of the failure. Examples include
      - A connection actively refused (e.g. possibly firewall/ACL rejecting connections) {{<tgimg src="failed-refused.png" width="90%" caption="Connection refused result">}}
      - A connection timeout or no route message if there is no response (e.g. firewall/ACL dropping connection, bad destination IP, or no route to destination) {{<tgimg src="failed-noroute.png" width="90%" caption="No route to host result">}}
