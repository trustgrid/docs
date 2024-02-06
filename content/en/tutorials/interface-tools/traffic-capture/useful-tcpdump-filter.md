The below filters can be combined using "and" & "or" without quotes

- You can filter based on protocol such as "tcp", "udp", or "icmp" without quotes
- Best practice is to filter for what you want to see rather than filter out what you don't want to see
- To see only traffic for a specific host use "host x.x.x.x" with the IP address in place of the x.x.x.x
- To see only traffic for a specific port use "port XXXX" replacing the XXXX with the desired port number
- If capturing traffic on a clustered node, you should filter out the cluster port traffic. Typically port 9000 or 1975. E.g., `not port 9000`
- If capturing on the ETH0 - WAN Interface, you should filter out traffic to both Trustgrid's network (35.171.100.16/28) and your data plane gateways' network or public IP. You'll need to identify the actual network or IPs. E.g., `not net 35.171.100.16/28 and not net X.X.X.X/X` or `not net 35.171.100.16/28 and not host X.X.X.X`.
