---
title: Node Behavior When Port 443 is blocked
linkTitle: Port 443 blocked
description: Details the symptoms, potential causes, and resolution steps when port 443 to the Trustgrid Control Plane is blocked.
---

## Symptoms

- Node will not update either automatically or when manually triggered
- Node will not be able to pull down container images from the organization repository
- Node cannot send up debug logs

## Cause

- The Trustgrid node cannot connect to repo.trustgrid.io on port 443 to update packages

## Troubleshooting Connectivity
Either of the below methods can be used to verify connectivity:
- Use the [Interface TCP Port Test tool]({{<relref "/tutorials/interface-tools/tcp-port-test">}}). Make sure the WAN interfaces IP is the source, `repo.trustgrid.io` is the Host, and `443` is the target {{<tgimg src="tcp-port-test-success.png" width="80%" caption="Successful interface port test">}}
- From the terminal run the command {{<codeblock>}}nc -vz repo.trustgrid.io 443{{</codeblock>}} {{<tgimg src="terminal-tcp-success.png" width="80%" caption="Successful TCP connection test from terminal">}}

If the above tests are successful but the device still exhibits the symptoms listed above, the issue is likely that something is [interfering with the TLS certificate]({{<relref "/help-center/kb/startup-process/ssl-tls-tampering">}})

If unsuccessful:
1. First confirm `repo.trustgrid.io` is resolvable by the [configured DNS servers]({{<relref "/docs/nodes/interfaces#dns-servers-ip">}}). In this example, we will assume the DNS server is `8.8.8.8`, **replace this with the configured DNS server IPs**. From the terminal run the command: {{<codeblock>}}dig @8.8.8.8 repo.trustgrid.io{{</codeblock>}} {{<tgimg src="dns-dig-success.png" width="80%" caption="Successful DNS resolution. Note: returned IP address will vary.">}}
    - Repeat the above process with the second configured DNS server if available.
    - If either fails to confirm:
        - Any firewall rules are not blocking TCP and UDP port 53 to the configured DNS server
        - Confirm there is not an [interface route]({{<relref "/docs/nodes/interfaces#routes">}}) on any LAN interface for CIDR that includes the DNS server's IP. This will cause requests to route out the LAN interface instead of WAN and only after the Trustgrid service has started and brought up the LAN interfaces.
1. If DNS returns an IP address but the connection still fails, confirm any firewall between the node and the internet allows port `443` **and** port `8443` to the [Trustgrid control plane public IP ranges]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}).


## Resolution

- Confirm the node can resolve repo.trustgrid.io
- Ensure the node can connect to repo.trustgrid.io:443

