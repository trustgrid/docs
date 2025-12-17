---
title:  January 2026 Major Release Notes
linkTitle:  January 2026 Major Release
date: 2026-01-05
description: "January 2026 Major Cloud Release Notes"
type: docs
---
## Lifecycle State Attribute
A new `Lifecycle State` attribute has been added as a more formal and consistent way to manage resource states across Trustgrid. Historically this was managed via a [Production Status Tag]({{<relref "docs/nodes/shared/tags/prod-status-tag" >}}) or similar custom tags. The new `Lifecycle State` attribute provides a standardized approach to indicate the current state of a resource.

The `Lifecycle State` attribute can be set to one of the following values:
- `pre-production` - Indicates that the resource is in a pre-production state, such as deployment, configuration or testing.
- `production` - Indicates that the resource is in active production use.
- `maintenance` - Indicates that the resource is temporarily out of service for maintenance or updates.
- `decommissioned` - Indicates that the resource is no longer in use and is being retired.

The `Lifecycle State` attribute can be set and modified at the individual node (appliance or agent) from the [infovisor]({{<relref "docs/nodes/shared/infovisor#essentials">}}) section. It can also be set at the cluster level for appliances via the [cluster overview]({{<relref "docs/clusters" >}}) page, which will propagate the value to all appliances within the cluster.

Additionally, the `Lifecycle State` attribute is now available as a filter option in various sections of the Trustgrid UI, including:
- Nodes Table
- [Alarm Filters]({{<relref "docs/alarms/alarm-filters" >}})


## Events and Alarm Improvements
This release includes several enhancements to the Events and Alarms system to improve usability and functionality.

- The backend database for Events has been migrated to a faster solution that should allow for further improvements in the future.
- The Advanced Search option has been added that allows users to filter events by multiple criteria, including date range, event type, level, node, and text. The results can then be exported to CSV for further analysis.
- The Generic Webhook notification channel will now include the Alarm Filters `Description` data in a `notes:` field within the payload.
- The 'Domain' field is now included in the Events data for better context for customers with multiple accounts/organizations.
- Adds the ability to [send a test event]({{<relref "docs/alarms/channels#testing-channels" >}}) to a **Channel** to test it independent of any filter. 
-  `/v2/event` API endpoint now paginates results. Previously only the first page was returned regardless of the number of events.
- Resolves an issue where many events were being logged when an Alert Suppression window closed.


## Other Improvements and Fixes

- Resolves an issue where the `Path` field value disappered when editing VPN Dynamic and Static routes.
- Fixes an issue that caused flow logs sorted by Send/Receive bytes to not treat 0 bytes as the lowest value, improving sorting accuracy.
- Container Repository images can now include multiple directories. E.g. example.trustgrid.io/directory/image:tag
- Users will now receive a more helpful error message if their IdP credentials are invalid during SSO login attempts.
- Enabling UDP will now only warn about a restart being required if the node is enabled as a gateway. Edge nodes do not require a restart when enabling UDP.
- Virtual network routes now sort by `Destination CIDR`.

### Deprecations
- With this release ZTNA Applications and related pages have been removed. 
- VM Management functionality has been removed. 
