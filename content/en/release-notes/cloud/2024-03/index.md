---
title: March 2024 Release Notes
linkTitle: March 2024 Release
date: 2024-03-20
description: "March 2024 cloud release"
type: docs
---

## Portal Notification Improvements
This release makes several changes to how a portal user is notified of new [events]({{<relref "/docs/alarms/events">}}) and commands (such as Restart and Reboot) related to the node being viewed. The toast pop-up notifications have a new look as shown below:
{{<tgimg src="new-notifications.png" width="50%" caption="Example notifications of a restart request and the resulting events">}}
The notifications will automatically close or can be dismissed by clicking the X at the top corner.  

This release also will notify if a Restart request fails to complete successfully by showing a Failed status toast and a message in the center of the screen. A future node appliance release will enable the same for Reboot requests. 

## Assign Policies via Groups
Historically policies had to be assigned to specific users. This release introduces the ability to [assign policies to groups]({{<relref "docs/user-management/groups#manage-group-permissions">}}). Policies assigned to a group will be inherited by all users in that group. This simplifies managing permissions when dealing with large numbers of users.

## User Account Management Improvement
We've added a Security menu option for users to manage their own account details like password and MFA settings. This link does not apply to users associated with an external Identity Provider (IdP). 
{{<tgimg src="/docs/user-management/users/account-mgmt/security.png" width="40%" caption="Security menu option">}}

## Other Issues and Improvements
- Disconnected peers in the Data Plane panel now show as a red X icon instead of a red circle. 
- Resolve an issue with [adding cluster members]({{<relref "docs/clusters/manage-members#add-member">}}). Prior to this release, if you select "Yes" to the prompt about configuring the heartbeat settings the portal failed to configure those settings. This release resolves that failure.
- Resolves an issue that caused flow log exports to not match the current filtered view.
- Resolves an issue where odd characters appeared in terminal windows for users running Windows-based operating systems.
- Resolves an issue that prevented deleting users associated with multiple IDPs
