---
title: Configuration Update Failure Troubleshooting
linkTitle: Config Update Failure
description: Troubleshooting nodes that are unable to pull their latest configurations
---
## Symptoms
- Changes made via the portal or api are not implemented on the actual node
- In versions after 20240719, the `Configuration Update Failure` event will be sent by appliance nodes {{<tgimg src="config-update-fail.png" caption="Example Configuration Update Failure event">}}

## Causes
Possible causes include:
- The node is unable to resolve Trustgrid.io DNS names, in this case specifically `gatekeeper.trustgrid.io`
- The node is unable to connect to the the [entire Trustgrid Control Plane networks]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}) on port 8443

## Troubleshooting
These steps should be performed via the Terminal or the [appliance console Network Tools Shell]({{<relref "/tutorials/local-console-utility/troubleshooting#network-tools-shell">}}).
1. Test DNS Resolution
    1. Determine the configured DNS server by viewing the WAN interface in the portal.  
    1. Use `dig @<DNS IP> gatekeeper.trustgrid.io +short` to confirm you get a response from each of your DNS servers. E.g. if the appliance is configured to use 8.8.8.8 and 1.1.1.1 use the commands `dig @8.8.8.8 gatekeeper.trustgrid.io +short` and `dig @1.1.1.1 gatekeeper.trustgrid.io +short`
    1. If the above does not resolve, make sure any firewall or access control list between the node and the WAN interface IP allows the IP to make DNS connections to the configured DNS servers on TCP & UDP port 53. If it successfully resolves, continue to the next step.
1. Confirm Network connectivity to the rest api endpoint with the command {{<codeblock>}}nc -vz gatekeeper.trustgrid.io 8443{{</codeblock>}}
    1. If this succeeds as shown in this image proceed to the next troubleshooting step. {{<tgimg src="nc-gatekeeper-success.png" width="75%">}}
    1. If this fails, confirm any firewall or access control list allows connectivity to [all required control plane networks]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}) on TCP port 8443.
1. Follow [these instructions to verify nothing is interfering with the TLS certificate chain.]({{<relref "/help-center/kb/startup-process/ssl-tls-tampering#verification">}})
1. Attempt to run the command: {{<codeblock>}}curl https://gatekeeper.trustgrid.io:8443/domain-info{{</codeblock>}}  The expected output is {{<codeblock>}}{"error":"node not authorized"}{{</codeblock>}}
    1. If you receive no response this indicates the connection is still failing.
    1. If you receive a different response Trustgrid support will likely need to assist in troubleshooting.