---
title: Troubleshooting Menu
linkTitle: Troubleshooting
description: Additional tools for troubleshooting connectivity
---
Tools on this page allow for additional troubleshooting.
{{<tgimg src="troubleshooting-menu.png" caption="Troubleshooting Menu" width="70%">}}

## View Node Logs
This utility opens the Trustgrid service logs in a text viewer in follow mode so new lines will be displayed as written to the log file.  

To exit follow mode hit `control + C`.  You can then use arrow keys to move up and down through the log file.

To return to the Troubleshooting menu hit `control + C` (if you haven't already) and then input `:q` and hit return.

## Push Node Logs

This command will attempt to zip up logs and push them up to the Trustgrid control plane. If successful these will be available in the node's [debug logs]({{<ref "help-center/ops-logs/debug-logs">}}).
{{<alert color="warning">}} This requires a valid IP configuration with access to the Trustgrid control plane. Specifically port 443 access to `repo.trustgrid.io`.{{</alert>}}

## Force Upgrade
This command will attempt to upgrade the Trustgrid service to the latest published version. 
{{<alert color="warning">}} This requires a valid IP configuration with access to the Trustgrid control plane. Specifically port 443 access to `repo.trustgrid.io`.{{</alert>}}

## Restart Node Service
This will restart the Trustgrid service on the node. This is faster than a full reboot but may not clear any OS related issues.

## Reboot Node Hardware
This will perform a complete operating system reboot of the Trustgrid appliance.

## Network Tools Shell
This will open a shell terminal that allows for running a very limited number of commands to perform troubleshooting. 

{{<alert color="info">}}Input `exit` at the command line to exit the shell and return to the troubleshooting menu. {{</alert>}}

These commands include:
* `cat`
* `less`
* `groups`
* `ls`
* `nslookup`
    * use this to confirm DNS resolution is working with a command like `nslookup repo.trustgrid.io` 
    * or to specify a DNS server with a command like `nslookup repo.trustgrid.io 8.8.8.8` which will use Google's DNS server
* `ping`
* `telnet` - use this to confirm connectivity with a command like `telnet repo.trustgrid.io 443`
* `traceroute` 
* `mtr` - this is a more robust version of `traceroute` and can be used to confirm connectivity with a command like `mtr repo.trustgrid.io`
* `nc` 
    * use this to confirm connectivity with a command like `nc -vz repo.trustgrid.io 443` 
    * or confirm connectivity with a [specific gateway IP address and port]({{<ref "/docs/nodes/gateway#server-settings">}}) with a command like `nc -vz 35.171.100.16 8443"
* `openssl` - this can be use to confirm a [nothing is interfering with a valid TLS connection]({{<ref "/help-center/kb/startup-process/ssl-tls-tampering">}}) and view the certificates used with a command like `openssl s_client -connect repo.trustgrid.io:443 -showcerts | less` (use `:q` to exit)
* `curl`
* `ip` 
    * `ip link` or `ip l` lists the connection status of all interfaces
    * `ip address` or `ip a` lists the IP addresses assigned to all interfaces
    * append `show <interface name>` to either of the above to limit to a specific interface. e.g. `ip a show enp0s20f0`
    * `ip route` lists the current OS routing table
* `dig` for confirming DNS resolution
    * `dig repo.trustgrid.io +short` will provide just the IP address using the configured DNS servers
    * `dig @8.8.8.8 repo.trustgrid.io +short` will query Google's DNS server directly and provide the IP address
## ARPing
This command can be used to confirm layer 2 connectivity between the Trustgrid appliance and other devices on the same network.  For example, to confirm the default gateway IP address is reachable use a command like `arping -I eth0 192.168.127.1`. This will send an ARP request to the default gateway IP address and display the resulting MAC address.  

If you do not get a response or get an unexpected MAC address then layer 2 connectivity is not working. Verify the Trustgrid node is connected the intended network and/or VLAN. 

## Advanced Tools
{{<alert color="warning">}} The tools below should only be used if directed by Trustgrid support.{{</alert>}}
- Force Backup Mode
- Force Diagnostic Mode
- Exit from Backup/Diagnostics
- Update JVM Memory
- Update JVM Garbage Collection