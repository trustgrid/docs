---
title: "Customer Responsibilities"
weight: 1
date: 2022-12-20
---

{{% pageinfo %}}
Trustgrid and our customers share the responsibilities of securing the Trustgrid network. Trustgrid assumes and automates a significant part of the overall security challenges. These challenges range from the cloud to the firmware on hardware appliances. Our customers are responsible for securing physical environments, identity and access management, and all security aspects of the environments where Trustgrid is installed.
{{% /pageinfo %}}

## Trustgrid Secures the Cloud

All cloud management components are secured and monitored by Trustgrid. This includes the hardware, software, and network for all cloud components including the Portal/Management API and other management tools. Customers are responsible for issuing credentials to the Portal for each authorized user and maintaining the security of those accounts.

## Network Security is a Shared Responsibility

Trustgrid is responsible for the security of all data plane traffic from the time it enters a Trustgrid [Node]({{< ref "/docs/concepts/node" >}}) until the time it exists a Trustgrid [Node]({{< ref "/docs/concepts/node" >}}). Customers are responsible for the security of all network traffic eggressing from any Trustgrid [Node]({{< ref "/docs/concepts/node" >}}).

## Trustgrid Teams with Customers for Edge Security

All virtual or hardware appliances are secured by Trustgrid. Trustgrid's ability to secure hardware is restricted to supported hardware appliances with specific exceptions. Customers are responsible for the security of the physical and logical environments into which Trustgrid is integrated.

## Secure Gateway Nodes

Trustgrid customer's should apply access controls to ingress traffic for Gateway Nodes. By restricting ingress access to Edge Node's IP ranges the system security is improved and many common attacks easily avoided.

## Encrypt Traffic with Customer Certificates

By default Trustgrid encrypts all traffic on the data and control planes with privately issued certificates from our PKI (Public Key Infrastructure). Customers may elect to provide their own certificates for data plane encryption ensuring Trustgrid maintains a zero-knowledge security posture of all data plane traffic.
