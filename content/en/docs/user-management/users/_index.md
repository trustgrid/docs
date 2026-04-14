---
Title: "Users"
---

{{<tgimg src="users.png" width="90%" caption="Users page">}}

The Users panel is the primary panel for managing user identities and the policies attached to them. All changes in this panel require either the [builtin-tg-acess-admin]({{<ref "/docs/user-management/policies#tg-builtin-access-admin">}}) policy or similar permissions.

## Adding or Inviting Users

Users can either be invited to the Trustgrid portal and utilize Trustgrid’s native authentication system, or be added and be authenticated by a customer configured Identity Provider (IdP).

### Inviting Users

For users that will utilize Trustgrid’s native authentication system, you will need to invite them with a valid email address.

1. Navigate to **User Management** → **Users**
1. Under **Actions** select **Invite User**
{{<tgimg src="invite.png" caption="Invite User action">}}
1. Enter the user’s valid email address and select the desired base policy (more can be attached later).
{{<tgimg src="send-invite.png" alt="Enter email and base policy">}}
1. Click **Send Invite**. You should see a confirmation that the invite was sent like the one below.
{{<tgimg src="invite-sent.png" alt="Invite sent confirmation" width="40%">}}
1. The user will receive an email with a link to the [Trustgrid Portal](https://portal.trustgrid.io) or similar. They will need to click said link.
{{<tgimg src="linky.png">}}
1. Once at the Portal, the user will need to click the **Start a free trial** option. **The user must use the same email address that was invited to get access to your account**.
{{<tgimg src="sign-up-1.png" caption="Click Start a free trial" width="50%">}}
1. The user will be prompted for their email, name, password, and company. **The user must use the same email address that was invited to get access to your account**.  
{{<tgimg src="sign-up-2.png" caption="Enter required information and click Sign Up" width="60%">}}
1. A verification email will then be sent to the invited user's email address. They will need to click the link in the email to verify their email and complete registration.
1. After verifying their email address the user will be prompted to configure Multi-Factor Authentication (MFA). Trustgrid recommends using a one-time password MFA such as Authy or Google Authenticator.
   1. Scan the QR code with your app. 
   1. Enter the passcode and click the **Submit** button. {{<tgimg src="sign-up-4.png" width="50%" >}}
1. The user is returned to the login screen. Login with the newly created email, password, and MFA code.  You will then be redirected back to the Trustgrid portal.

### Adding Users (with Identity Providers)

If your account has an [Identity Provider (IdP)]({{<ref "/docs/idps">}}) configured as a Portal Auth Provider, you use the Add User button to create an identity for them within Trustgrid.

{{<alert color="info">}} Some IdP’s allow for users to be synchronized automatically making this step unnecessary {{</alert>}}

1. Navigate to **User Management** → **Users**
1. Click the **Add User** button {{<tgimg src="add-user.png" width="30%">}}
1. Enter in the user’s email address. **This must be the same email address the IdP sends back to Trustgrid, if unsure consult with your IdP Admin**.
1. Select your Identity Provider (if more than one)
1. Select the desired base policy. More can be attached later.
1. Click **Save**
{{<tgimg src="save-new-user.png" caption="Add User Prompt" width="40%">}}
1. The user can then navigate to the portal [authentication domain]({{<ref "/docs/idps">}}) associated with the IdP. They will be redirected to the IdP page and required to enter their username, password and MFA (if configured), then automatically sent back to the Trustgrid portal.

## Manage User Policies/Permissions

### Attaching Policies 
To attach/detach policies attached to a user:

1. Under **User Management** → **Users** search for the target user and click their name.
1. To add a new policy:
   1. Click the **Attach Policy** button {{<tgimg src="attach-policy.png" width="40%">}}
   1. Search for the desired policy then select it. {{<tgimg src="choose-policy.png" width="60%">}}
   1. Click **Attach**

### Removing Policies 
To remove a policy from a user:
1. Under **User Management** → **Users** search for the target user and click their name.
1. Click **Detach** located to the right of policy you wish to remove. {{<tgimg src="detach-policy.png" width="70%">}}


### View Effective Permissions

To see what permissions a user currently has and what policy is providing that permission:
<ol>
<li> Under **User Management** → **Users** search for the target user and click their name. </li>
<li> In the right `Effective Permissions` pane, find the permission in question. These permissions are grouped by category. Each entry will show:
{{<tgimg src="effective-permissions.png" width="80%" caption="Example effective permissions">}}
<ol type="a">
   <li> If the permission is allowed ✅, explicitly denied ❌, or not defined (both icons gray) </li>
   <li> The action name </li>
   <li> A short description of what the permission allows</li>
   <li> what policy grants/denies the user the permission</li>
</ol>
</li>
</ol>

### Change a User Landing Page

The Landing Page allows you to designate where a user is directed within the portal on initial login. 

This is most useful for users that only need access to ZTNA Remote Access application at the `/apps` page or if you want them to automatically land on a specific page like `/nodes`. 

1. Under **User Management** → **Users** search for the target user and click their name.
1. Enter in the desired landing page path (e.g. `/apps`) and click **Save**
{{<tgimg src="change-landing.png" width="80%" caption="Change user landing page">}}


## Managing Group Membership


### View a User’s Group Membership

You can view all the Groups a user is a member of from the Groups panel. To change memberships you will need to use the **User Management** → **Groups** page.
To view:

1. Under User Management → User search for the target user and click their name.
1. Select the **Groups** panel on the left.
1. **Groups** will be listed in a table in the main panel.

{{<tgimg src="user-groups-table.png" caption="Example group membership table" width="80%">}}

### Adding a User to Groups
You can add a user to one or more group by:
1. Navigate to User Management > Users and click on the name of the user to add to groups.
1. Select the Groups panel from the left navigation bar.
1. Click the "Add to Group(s)" button.
1. From the prompt select the group or groups you want to add the user to. If there are many groups you can start typing the name to filter the options. 
{{<tgimg src="add-user-group-prompt.png" caption="Add Groups prompt" width="60%">}}
1. Click Save

### Removing a User from Groups
You can remove a user from a group by:
1. Select the radio checkbox next to the group you wish to remove.
1. From Actions, select Delete.
{{<tgimg src="user-group-delete.png" caption="Deleting a group membership" width="40%">}}
1. When prompted confirm you want to delete the user.

## View a User’s Associated Identity Provider (IdP)
If an Identity Provider is configured a user maybe associated with one or more IdP.  
To view:

1. Under **User Management** → **Users** search for the target user and click their name.
1. Select the **Identity Provider** panel on the left.
1. The identity Providers will be listed in a table in the main panel. If no IdP is listed this indicates the user is using the Trustgrid native authentication system.
{{<tgimg src="user-idp-table.png" caption="Example User Identity Provider table" width="80%">}}


