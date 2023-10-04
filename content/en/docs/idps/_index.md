---
title: "Identity Providers"
date: 2023-1-12
weight: 9
---

{{% pageinfo %}}
The Trustgrid Portal can integrate with external identity providers. The identity providers added to the Portal may be used for ZTNA access and/or Portal authentication.
{{% /pageinfo %}}

## Authentication Domain
To use an identity provider for first configure the authentication domain. This domain can be used to access the portal (if "Use for Portal Auth" is enabled) and is used in the configuration of at the IdP provider to forward users back to Trustgrid.
{{<tgimg src="auth-domain.png" caption="Example authentication domain configuration">}}


## Identify Providers
Each IdP has the following fields.  Beyond that the settings differ for each provider type.

{{<fields>}}
{{<field "Type" >}}
the identity provider type.
{{</field >}}
{{<field "Name" >}}
used inside the portal when associating an identity provider with a ZTNA application
{{</field >}}
{{<field "Use for Portal Auth" >}}
whether or not this provider should be the default authentication provider. Only one provider can be used for Portal authentication.
{{</field >}}
{{</fields>}}
