---
Title: "Groups"
Date: 2022-12-30
---

![img](/docs/concepts/user-management/group-ss.png)

User Management Groups defined can be used to make Remote Access applications visible to users. 

There are two sources for Groups

#### Trustgrid-native Groups
These groups are created in the portal via the `Add Group` button.
Users can be added to the group by:

1. From the `Groups` page you can select the group, then click `Add Member`.
   
![img](/docs/concepts/user-management/add-to-group.png)

2. From the `user` section, search for and add the email address of the desired user.

![img](/docs/concepts/user-management/user-email.png)

3. Click `Save` and repeat process to add other users.

#### Identity Provider-Synchronized Groups
These groups are create based on Identity Provider (IdP) configuration for select IdPs (currently Gsuite and Azure OpenID). The IdP can be configured to sync all or select groups, and the sync process will update the group membership automatically.  

> if users are added to an IdP-synced group in the Trustgrid portal they will automatically be removed the next time they sync runs.

 


