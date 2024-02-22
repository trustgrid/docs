---
title: "October 2023 Appliance Release Notes"
linkTitle: "October 2023"
type: docs
date: 2023-10-02
---
{{< node-release package-version="1.5.20231003-1842" core-version="20231003-145438.e8103a2" release="n-2.17.0" >}}

## Support for running Ubuntu 22.04 LTS
This release adds support for running the Trustgrid node appliance on Ubuntu 22.04 LTS for new deployments on AWS, Azure and vSphere.  The 18.04 appliances will continue to be supported and security updates made available via the Trustgrid repository. 

{{<alert color="info">}}There are no plans to support upgrading OS on deployed devices at this time due to the risks involved. {{</alert  >}}

## (Preview Feature) Node Registration from the Trustgrid Console
This release includes a new feature that allows you to [register a node from the Trustgrid console]({{<ref "/tutorials/local-console-utility/remote-registration" >}}). This will allow a Trustgrid node running the a standard image to be registered with a Trustgrid account without needing to share a complete license token.  

The high-level process is:
1. A user with console access to a Trustgrid node logs in and initiates the registration process. This triggers a process that generates unique certificates for the node and generates a short activation code. The node will then begin checking with the Trustgrid control plane to see if it has been activated.
1. This code can then be shared with a user with access to the Trustgrid portal. That user can then enter the code in the Trustgrid portal via the **Activate Node** action.  After entering the code they are prompted to provide a name for the node and click **License Node**.
1. When the node detects its code has been activated it finalizes registration and then reboots.  The node is then available for use and configuration. 

This release also allows the user at the console to deregister a node. This will remove the node from the Trustgrid account and deactivate the license. 

{{<alert color="info">}} This feature is currently in preview and requires Trustgrid support to enable the feature for specific users. {{</alert  >}}

## Other Trustgrid Console Improvements
- The ability to run `traceroute`, `mtr`, `ip`, `curl`, `openssl` from the [Troubleshooting > Network Tools Shell]({{<ref "/tutorials/local-console-utility/troubleshooting#network-tools-shell">}})
- The ability to modify an interface's MTU settings from the [Advanced Network Configuration]({{<ref "/tutorials/local-console-utility/advanced-network-config#change-interface-mtu">}})

## New Tool for Monitoring Connectors
This release includes a new service that will allow you to see active [L4 Connector]({{<ref "/docs/nodes/shared/connectors">}}) sessions. This gives visibility into the active usage of any connector.


## Other Improvements and Fixes
- Audit messages for actions related to containers will now include the container name for improved visibility.
- IPSEC Tunnels will now display their up/down status on the Tunnels page for easier monitoring.
- Resolves an issue preventing AWS nodes from detecting their region correctly if they were set to require IMDSv2.
- Resolved issues that lead to false events for "packet(s) too large for UDP" and "Interface Down" notices during restart/reboot.
- Resolved an issue impacting proxy ARP responses on VLAN subinterfaces for outside NAT.