---
title: Request Trustgrid Trial Account
linkTitle: Request Trial
weight: 10
description: "Sign up for a free 30-day trial account to deploy agents and explore Trustgrid features. - 5 minutes"
---

The process below will guide you through requesting a free 30-day trial account. 

## High-Level Steps
- Sign up for a Trustgrid trial account
- Verify email
- Sign in

## Prerequisites
- Email address
- Chrome-based browser

## Step 1 - Create a Trustgrid account
1. Using Chrome or a similar browser navigate to https://portal.trustgrid.io/
1. Click the link at the bottom of the page to "Start a free trial" {{<tgimg src="start-trial.png" width="70%">}}
1. Provide the requested information. Note that the account will be created based on the provided Company Name.
{{<tgimg src="sign-up-form.png" alt="Sign up form" width="50%">}}
1. A verification email will be sent to the provided address. 
{{<tgimg src="email-verify.png" width="50%">}}
1. Check your email inbox for the verification email from Trustgrid. It may be in your spam folder. Click the link to verify your email address.
1. You will be redirected to the login page. Sign in using the credentials you provided during sign-up.
1. Upon login, you will be required to configure one-time password multifactor authentication (MFA). Use an app such as Google Authenticator or Authy to scan the QR code. Enter the returned one-time password when prompted.
{{<tgimg src="mfa-setup.png" width="40%">}}
1. You will be returned to the sign-in page to complete your login. Enter your email, password, and the one-time password from your authentication app.
{{<tgimg src="final-signin.png" width="40%">}}

Note: If you are not automatically redirected back to the Trustgrid portal just click the Portal link in the top right.
{{<tgimg src="auth-portal-link.png" width="35%" caption="Link back to Portal">}}


## Step 2 - Choose a domain
All Trustgrid accounts have a base domain that is used to create fully qualified domain names (FQDN) for resources such as nodes and clusters.  
1. On initial login, you will be prompted to provide the subdomain for your FQDN. This will be appended with `.trustgrid.io` automatically. {{<tgimg src="domain-setup.png" width="85%" caption="Prompt to create account domain">}}
    - This name must be unique among all Trustgrid customers, if you choose a domain name that is already taken you will be prompted to choose another.
    - The subdomain can only consist of lower-case letters, numbers, and hyphens(`-`). 

{{<alert color="warning">}}Once created the domain cannot be changed. Please choose carefully.{{</alert>}}


## Next Steps
You now have access to the Trustgrid portal and can begin [installing agents](/getting-started/trial/base-setup).
