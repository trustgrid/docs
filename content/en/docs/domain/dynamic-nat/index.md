---
Title: "Dynamic NAT (catch-all)"
Tags: ["dynamic-NAT", "Networking", "Networks"]
Date: 2023-1-9
---

{{% pageinfo %}}
In some scenarios it may be desirable to configure a dynamic NAT rule that translates unknown source IPs to an IP address on the Trustgrid virtual subnet. To configure dynamic NAT, you would create an inside NAT rule with a local CIDR of 0.0.0.0/0 and a /32 network CIDR.
{{% /pageinfo %}}

In this configuration, any traffic with an unknown source that is destined for the [Trustgrid network]({{< ref "docs/overview/networking" >}}) will be translated to the /32 that is specified in the dynamic NAT rule. More specific inside NAT rules will take priority.

### Example:

#### Edge Node Inside NATs

![img](edge-inside.png)

#### Gateway Node Inside NATs

![img](gateway-inside.png)

#### Ping from host 192.168.14.35 to 10.100.50.15
Local Capture:

![img](local-capture1.png)

Remote Capture:

![img](remote-capture1.png)

#### Ping from host 172.33.16.45 to 10.100.50.15
Local Capture:

![img](local-capture2.png)

Remote Capture:

![img](remote-capture2.png)