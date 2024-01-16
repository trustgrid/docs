---
title: Trustgrid User Account Management
linkTitle: Account Management
description: Managing your password and MFA settings for your Trustgrid user account.
---
Below shows how to manage password and MFA settings for organizations relying on native Trustgrid user accounts, instead of a custom [Identity Provider]({{<relref "/docs/idps">}}).  These instructions assume you can authenticate successfully with your Trustgrid account. If not you will need to work with Trustgrid support.

## Change Password
To change your Trustgrid password:
1. Login to the Trustgrid portal using your current password and provide your MFA code
1. In another browser tab navigate to https://auth.trustgrid.io/account/change-password
1. Provide your current password, enter a new password twice, and click submit. {{<tgimg src="change-password.png" width="50%" caption="Change password prompt">}}

## Change MFA
To change your multi-factor authentication (MFA) application:
1. Login to the Trustgrid portal using your current password and provide your MFA code
1. In another browser tab navigate to https://auth.trustgrid.io/account/change-mfa
1. Scan the provided QR code with your new MFA application 
1. Enter the Authenticator Code provided by your new MFA application
1. Enter your current password and click submit. {{<tgimg src="change-mfa.png" width="50%" caption="Change MFA prompt">}}