---
title: "Identity Providers"
date: 2023-1-12
weight: 9
---

{{% pageinfo %}}
The Trustgrid Portal can integrate with external identity providers. The identity providers added to the Portal may be used for ZTNA access and/or Portal authentication.
{{% /pageinfo %}}

To use an identity provider for Portal authentication, first configure the authentication domain.

![img](auth-domain.png)

This domain can be used to directly access the Portal, and unauthenticated users will be redirected to your chosen identity provider.

Four different types of identity providers are supported: AzureAD, GSuite, OpenID, and SAML.

{{< field-def "Type" >}}
the identity provider type.
{{< /field-def >}}
{{< field-def "Name" >}}
used inside the portal when associating an identity provider with a ZTNA application
{{< /field-def >}}
{{< field-def "Use for Portal Auth" >}}
whether or not this provider should be the default authentication provider. Only one provider can be used for Portal authentication.
{{< /field-def >}}
