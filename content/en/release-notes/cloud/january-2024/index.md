---
title: January 2024 Release Notes
linkTitle: 'January 2024 Release'
type: docs
date: 2024-01-21
description: "January 2024 cloud release with Flow Log and Portal Improvements"
---

## Default Authentication Provider Change
Trustgrid has previously utilized Auth0 as the authentication provider for customers who do not use a [custom identity provider (IdP)]({{<relref "/docs/idps">}}). With this release, Trustgrid has migrated to utilizing an internally managed authentication service to provide more control and stability.

{{<alert color="info">}}This change will have no impact on customers using their own Identity Providers (IdPs) for portal authentication{{</alert>}}

### User Migration Process
As part of this process users will be prompted to migrate their accounts from the previous Auth0 provider to the new internally managed authentication service. This will involve:
1. When existing users login to the portal they will be receive a prompt, like below, to recreate their account in the new system.{{<tgimg src= "migrate-account.png" width="50%" caption="Account migration prompt">}}
1. The users will then be prompted to provide their first and last name, and a new password.  The email address will be pre-filled. {{<tgimg src="migrate-user1.png" width="50%" caption="Prompt for password and additional details">}}
1. A verification email will then be sent to the invited user's email address. They will need to click the link in the email to verify their email and complete registration.
1. After verifying their email address the user will be prompted to configure Multi-Factor Authentication (MFA). Trustgrid recommends using a one-time password MFA such as Authy or Google Authenticator.
   1. Scan the QR code with your app. 
   1. Enter the passcode and click the `Submit` button. {{<tgimg src="/docs/user-management/users/sign-up-4.png" width="50%" >}}
1. The user is returned to the login screen. Login with the newly created email, password, and MFA code.  You will then be redirected back to the Trustgrid portal. 
   1. If you are not automatically redirected click the Portal link in the top right. {{<tgimg src="auth-portal-link.png" width="50%">}}

## Flow Logs Improvements

### Flow Log Performance
With this release, Trustgrid has changed out the backend system used for storage and querying of [flow logs]({{<relref "/help-center/flow-logs">}}). This new system provides significant performance improvements when viewing and exporting flow logs for nodes. While the initial default searches will sometimes take a few seconds longer, advanced searches are significantly faster than before.  Additionally, this change will enable future improvements for better reporting, analysis and visualization. 

### Flow Log Export TCP Flags
Prior to this release flow log TCP flags were represented in an aggregate hexadecimal number which required conversion to be useful. This release changes the export process to produce plain text values such as SYN, ACK, RST, FIN, etc. making analysis of TCP flags in exported logs much easier.

### Other Flow Log Table Improvements
- Start and Stop times now display the seconds in the portal
- Sent and Received Bytes fields now include commas between thousands for easier reading
- Columns resize and wrap as needed to prevent cutting off text

## Portal Improvements

### Favorite Shortcuts
A new favorites feature has been added to the portal navigation menu. Users can now favorite pages they access frequently for quicker access. 

Pages are added to favorites by:
1. Searching in the search bar at the top of the portal for a node, cluster, or page (such as Virtual Network) that is frequently accessed. 
1. Click on the star icon to the right. {{<tgimg src="favorite-star.png" width="60%" caption="Example of adding a favorite" alt="Search for 'Virtual Networks' with the star item selected on the right.">}}.
1. A link to the page will now appear at the top right of the portal no matter where you navigate. {{<tgimg src="favorite-example.png" width="35%" caption="Example of favorite link in menu" alt="Example of a favorite link for 'Virtual Networks' appearing at the top right of the portal.">}}

Removing a favorite page is the same process but unselecting the star icon instead.

### Removal of the Dashboard page
Based on feedback from users the Dashboard page has been removed and users now land on the Nodes table on login. You can also configure a [custom landing page for users if desired]({{<relref "/docs/user-management/users#change-a-user-landing-page">}})

### Advanced Node Options
A new [Advanced Options]({{<ref "/docs/nodes/appliances/advanced">}}) section has been added to nodes allowing the configuration of additional options like:
- JVM Memory settings
- Network flow timeout and time to live (TTL) and ARP intervals

### Clone Policy Option
To make it easier to [duplicate existing policies]({{<relref "/docs/user-management/policies#clone-existing-policy">}}) there is now a copy button to the far right of each policy in the Policies page. Clicking this button will create a duplicate of the policy that can be named and edited as desired. 
{{<tgimg src="clone-policy.png" width="80%">}}

## Other Issues Resolved
- Resolves an issue causing the [Gateway Clients]({{< relref "/docs/nodes/appliances/gateway/gateway-client">}}) page to load as a blank white page.
- Resolves an issue that caused some statistics on the Node Overview graphs to double the real value. This caused a spike in the graphs which was misleading.
- Resolves an issue with some of the "breadcrumb" links at the top of the page not working as expected. 
- Prevents external API dependencies, such as our status page, from delaying portal pages from loading.
