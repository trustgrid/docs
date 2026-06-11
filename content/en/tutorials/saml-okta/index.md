---
linkTitle: "Okta SAML"
title: "Enable SAML with Okta"
description: "Configure Okta as a SAML identity provider for the Trustgrid Portal."
tags:
  - SAML
  - Okta
  - authentication
---

This tutorial walks through configuring Okta as a SAML identity provider for the Trustgrid Portal using the self-service IDP configuration in **Organization Settings**.

## Prerequisites

- A Trustgrid organization with an [authentication domain]({{<relref "docs/idps#authentication-domain" >}}) configured.
- An Okta account with admin access.

## Part 1: Configure the IDP in the Trustgrid Portal

1. Navigate to **Organization Settings -> Identity Providers**.
2. Find your **Authentication Domain** on this page. If you haven't configured one yet, set it now. It must be in the form `<your-desired-name>.trustgrid.io`. Note this value; you will need it when configuring Okta.
3. Click **Add** and fill in:
   - **Type:** `SAML`
   - **Name:** a descriptive name, e.g. `Okta`
   - **Use for Portal Auth:** `Yes` (if using this provider for portal login)
4. Click **Save**. This opens the **Configuration** page. Leave it open; you will return to it in Part 3.

{{<tgimg src="idp-auth-domain.png" width="50%" caption="The Identity Providers page showing the Authentication Domain." >}}

## Part 2: Create the SAML App Integration in Okta

1. In the Okta admin console, go to **Applications -> Applications**.
2. Click **Create App Integration**.
3. Select **SAML 2.0**, then click **Next**.
4. On **General Settings**, enter an app name (e.g. `Trustgrid`), then click **Next**.
5. On **Configure SAML**, fill in:
   - **Single sign-on URL:** your Authentication Domain from Part 1 with `/saml` appended (e.g. `https://<subdomain>.trustgrid.io/saml`)
   - **Audience URI (SP Entity ID):** the same value as the Single sign-on URL
   - **Name ID format:** `EmailAddress`
   - **Application username:** `Email`
   - Leave all other fields at their defaults.
6. Click **Next**.
7. On the **Feedback** step, select **This is an internal app that we have created**, then click **Finish**.

{{<tgimg src="okta-configure-saml.png" width="50%" caption="The Configure SAML step with the four key fields filled in." >}}

## Part 3: Import the Okta Metadata into Trustgrid

After finishing, you land on the app's **Sign On** tab. Under **Metadata details**, find the **Metadata URL**.

{{<tgimg src="okta-sign-on.png" width="50%" caption="The Okta Sign On tab showing the Metadata URL and More details link." >}}

### Option A: XML upload (recommended)

1. Click **Copy** next to the Metadata URL, then open that URL in a new browser tab.
2. You will see raw XML. Save the file using **File -> Save Page As**, and make sure it is saved with a `.xml` extension (not `.html` or `.txt`).
3. On the Trustgrid IDP **Configuration** page, click **Choose File** and select the saved XML file.
4. The Issuer, Login URL, and Identity Provider Signing Certificate fields populate automatically.
5. Click **Save**.

### Option B: Manual entry

1. Click **More details** on the Okta **Sign On** tab.
2. Copy the following into the [SAML IDP fields]({{<relref "docs/idps/saml-idp-fields" >}}) in Trustgrid:
   - **Sign on URL** -> Login URL
   - **Issuer** -> Issuer
   - **Signing Certificate** -> Identity Provider Signing Certificate (paste the contents without the `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----` headers)
3. Click **Save**.

## Part 4: Assign Users in Okta

1. In the Okta app, go to the **Assignments** tab.
2. Click **Assign -> Assign to People** for individuals, or **Assign to Groups** for bulk assignment.
3. Assign all users who need Trustgrid portal access.

## Part 5: Add Users in the Trustgrid Portal

Follow the [Adding Users with Identity Providers]({{<relref "docs/user-management/users#adding-users-with-identity-providers" >}}) process:

- Use **Add User**, not **Invite User**.
- The email must exactly match what Okta sends in the SAML assertion.
- Select the Okta IDP when prompted.

## Part 6: Test the Login

Open an incognito or private browser window and navigate to your Authentication Domain URL. You should be redirected to Okta, authenticate, and land in the Trustgrid portal. Users can also launch Trustgrid from their Okta app tile.

If SAML fails for any reason, you can fall back to native authentication at `https://portal.trustgrid.io` using your Trustgrid portal credentials.
