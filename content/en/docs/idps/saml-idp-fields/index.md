---
Title: "SAML"
Date: 2023-1-12
---

When configuring a SAML IDP, the provider will often provide a metadata XML file to simplify configuration. If provided, you can upload that on the SAML configuration page, which will automatically populate the fields necessary for SAML authentication.

{{<fields>}}
{{<field "Issuer" >}}
A URL that will point to the IDP, typically with some identifying information so the IDP can determine the source of the request. This will be provided by your SAML identity provider.
{{<field "Login URL" >}}
A URL that will initiate the login process. This will be provided by your SAML identity provider.
{{</field >}}
{{<field "Identity Provider Signing Certificate" >}}
A certificate used to authenticate requests between Trustgrid and the SAML IDP. This will be provided by your SAML identity provider.
{{</field >}}
{{</fields>}}
