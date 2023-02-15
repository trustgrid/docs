---
title: "Certificates"
date: 2023-1-12
weight: 3
---

{{% pageinfo %}}
Trustgrid allows importing of self-managed certificates that can encrypt traffic for both the Trustgrid mesh network (via the Gateway) and [ZTNA applications]({{<ref "/docs/applications">}}).
{{% /pageinfo %}}

![img](certificates.png)

#### Adding a Certificate

{{<fields>}}
{{<field "Certificate body">}}
The certificate body, in PEM format.
{{</field >}}
{{<field "Certificate chain">}}
The certificate authority's chain, in PEM format.
{{</field>}}
{{<field "Private key">}}
The certificate's private key, in PEM format.
{{</field>}}
{{</fields>}}

Once uploaded, the sensitive certificate details will not be available via the API, but the parsed information from the certificate, such as FQDN, the expiry date, and any warnings about the certificate, will be available in the table.
