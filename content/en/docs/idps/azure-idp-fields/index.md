---
categories: ["idps"]
title: "Azure AD"
linkTitle: "Azure AD"
---

#### Setup

Setting up the Azure AD integration requires the configuration of an Azure App registration.  See [Azure App registration configuration]({{< ref "#azure-app-registration-configuration" >}}).

![img](azure_ad_config.png) 

{{<fields>}}
{{<field "Issuer">}}
Specifies the Azure App OpenID metadata endpoint.
{{</field>}}
{{<field "Client ID">}}
Specifies the Azure App client ID.
{{</field>}}
{{<field "Secret">}}
Specifies the Azure App secret.
{{</field>}}
{{</fields>}}

{{<alert color="info">}}
Before being able to save the configuration, validation of the Issuer needs to be performed by pressing the `Test` button.
{{</alert>}}

Once the validation has been performed and the configuration saved, the sync section will be available.

![img](azure_ad_sync.png) 

- Sync all users - Syncs all available users along with all the groups associated to the users.
- Sync users from specific groups - Syncs all users from the selected groups.

{{<alert color="info">}}
- At any given point the sync selection can be changed and saved, and the sync will automatically start. 
{{</alert>}}

#### Azure App registration configuration

{{<alert color="info">}}
The following information is only partial Azure AD reference for configuring an App registration via the Microsoft Azure Portal.  Information presented in the following images might have changed.
{{</alert>}}

1. Create a new `App Registration`

![img](app_register.png)

![img](app_create.png) 

2. Once you’ve registered your application you need to make a note of the Application Client ID and Metadata endpoint that provides the Open ID metadata.  You can find it by click on Endpoints at the top of the application.  The Application Client ID will be used for the Client ID configuration parameter of the Identity Provider in Trustgrid Portal. The metadata endpoint goes in the Issuer configuration parameter
of the Identity Provider in Trustgrid Portal.  It should be something like https://login.microsoftonline.com/<app_id>/v2.0 Ignore what comes after the /.well-known part.

![img](app_endpoints.png) 

3. Configure the Authentication piece of the Azure app

![img](app_auth.png)

4. Add a platform by selecting the “Web” type and add the corresponding redirect uri and logout url, and select both Access Tokens and ID tokens checkboxes.

{{<fields>}}
{{<field "Trustgrid Redirect URL">}}
https://id.trustgrid.io/auth/openid/callback
{{</field>}}
{{<field "Trustgrid Logout URL">}}
https://id.trustgrid.io/logout
{{</field>}}
{{</fields>}}

5. You need to create a secret for the authentication piece.  You can select the expiration to be any of the selections.  Once you create the secret make sure to make a note of the “value” since that piece of information is what is going to be used under the Secret configuration parameter of the Identity Provider in Trustgrid Portal.

![img](app_secret.png)

6. Once you’ve created the client secret, navigate to “Token Configuration” and add the “email” optional claim

![img](app_token_config.png)

***

