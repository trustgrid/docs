---
title: Trustgrid User Account Management
linkTitle: Account Management
description: Managing your password and MFA settings for your Trustgrid user account.
---
Below shows how to manage password and MFA settings for organizations relying on native Trustgrid user accounts, instead of a custom [Identity Provider]({{<relref "/docs/idps">}}).  These instructions assume you can authenticate successfully with your Trustgrid account. If not you will need to work with Trustgrid support.

## Change Password
To change your Trustgrid password:
1. Login to the Trustgrid portal using your current password and provide your MFA code
1. From your user profile dropdown, select "Security" {{<tgimg src="security.png" width="40%" caption="Security menu option">}}
1. Provide your current password, enter a new password twice, and click submit. {{<tgimg src="change-password.png" width="50%" caption="Change password prompt">}}

## Change MFA
To change your multi-factor authentication (MFA) application:
1. Login to the Trustgrid portal using your current password and provide your MFA code
1. From your user profile dropdown, select "Security" {{<tgimg src="security.png" width="40%" caption="Security menu option">}}
1. Scan the provided QR code with your new MFA application 
1. Enter the Authenticator Code provided by your new MFA application
1. Enter your current password and click submit. {{<tgimg src="change-mfa.png" width="50%" caption="Change MFA prompt">}}

{{<alert color="info">}}If you are unable to login to the Trustgrid portal you will need to work with Trustgrid support to regain access{{</alert>}}

## Reset MFA
Users with the `users::modify` permission, such as those with the `builtin-tg-access-admin` policy, can reset another user's MFA by following the steps below:
1. Login to the Trustgrid portal using your current password and provide your MFA code
1. Navigate to User Management > Users
1. Select the user you want to reset and select the "Reset MFA" option from the Actions drop down.
 {{<tgimg src="reset-mfa.png" width="50%" caption="Reset MFA prompt">}}
1. The user will be required to reconfigure MFA on their next login.

{{<alert color="info">}}This action only impacts MFA configured for users in the Trustgrid user database.  It does not impact MFA configured for users in your identity provider.{{</alert>}}