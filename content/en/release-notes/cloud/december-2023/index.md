---
title: December 2023 Release Notes
linkTitle: 'December 2023 Release'
type: docs
date: 2023-12-10
description: "December cloud release with Flow Log and Portal Improvements"
---

Much of this release lays the groundwork needed for a new lighter-weight version of Node that will be made public next year. 

## Flow Logs Improvements

### Flow Log Performance Improvements
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

## Other Issues Resolved
- Resolves an issue causing the [Gateway Clients]({{< relref "/docs/nodes/gateway/gateway-client">}}) page to load as a blank white page.
- Resolves an issue that caused some statistics on the Node Overview graphs to double the real value. This caused a spike in the graphs which was misleading.
- Resolves an issue with some of the "breadcrumb" links at the top of the page not working as expected. 
- Prevents external API dependencies, such as our status page, from delaying portal pages from loading.
