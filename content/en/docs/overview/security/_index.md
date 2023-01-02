---
title: "Trustgrid Security"
weight: 1
date: 2022-12-30
---

{{% pageinfo %}}
Trustgrid’s mission is to replace aging Edge connectivity solutions with innovative, software defined solutions that integrate security and compliance in a way impossible for legacy solutions, while improving efficiency in deployment and life-cycle management.
{{% /pageinfo %}}

### Authentication

Pre-shared keys (PSK) are the most common method for authenticated connectivity and present significant risk when implemented poorly. Certificate based authentication is difficult and requires advanced skill sets. Trustgrid is the ‘root of trust’ in a Public Key Infrastructure (PKI) built by our security experts to enable automated certificate management across Trustgrid networks. All devices enroll in the PKI and are managed centrally from Trustgrid’s cloud infrastructure. Certificates are issued per device and are used to authenticate, and for encryption of all traffic.

### Authorization

Central to the security of the Trustgrid network is an authorization model derived from Google’s [Beyond Corp](https://cloud.google.com/beyondcorp) (aka Zero Trust from Forrester) initiative. This places an implicit deny on all traffic that cannot be configured to allow all traffic. Many breaches have been caused by allowing all traffic to overcome the burden of proper security configuration.

Trustgrid uses a token to authorize communication between Trustgrid [nodes]({{< ref "/docs/node" >}}).

### TLS Encryption

All data on Trustgrid is encrypted from [node]({{< ref "/docs/node" >}}) to node using next generation encryption – TLS Mutual Authentication. The Internet Engineering Task Force (IETF) recommends TLS as the replacement to IPSec VPN.

### Automated Management

Trustgird offers tools for automating software updates at scale. A significant difficulty in securing traditional VPNs is the application of patches and updates to hardware appliances. Many vendors own and maintain hundreds or thousands of these devices and are simply unable to efficiently patch them. This leaves significant security vulnerabilities unaddressed in customer and vendor datacenters.

### Private Data Routing

Trustgrid enables private data connections routed on infrastructure owned by each vendor, not through centralized, multi-tenant gateways that touch hundreds of vendors’ connections concurrently.
