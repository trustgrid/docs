---
title: April 2023 Release Notes
linkTitle: 'April 2023'
type: docs
date: 2023-04-04

---
## Documentation 
Trustgrid is moving away from our documentation previously hosted on Confluence to docs-as-code system at https://docs.trustgrid.io. This should enable us to more tightly integrate the portal UI with related documentation. 

As a start, we've created a link called Documentation under the Management section of the main navigation bar that will open the docs in a new tab.

Additionally, we've created a section for release notes that will allow you to monitor release for [Control Plane](https://docs.trustgrid.io/release-notes/cloud/) and [Node](https://docs.trustgrid.io/release-notes/node/) releases via RSS feeds going forward. 

### Swagger API Documentation Move
As part of this change we are moving our API documentation from https://portal.trustgrid.io/api-docs/ to https://docs.trustgrid.io/docs/api/. This will allow updating this documentation outside our normal release process. 


## Cluster Configuration
Prior to this release it was necessary to navigate to each member node and configure the [cluster heartbeat]({{<ref "/docs/nodes/cluster#heartbeat">}}) IP and port **after** it was added to the cluster.  In this release a prompt has been added to the ["Add Node" flow]({{<ref "/docs/clusters/manage-members#add-members">}}) that will allow you to configure these settings as part of adding the node to the cluster. 

## Data Plane Stats
This release further improves on our recent addition of data plane telemetry information.
* Show data plane telemetry for offline nodes
* Show data plane telemetry for connections between public and private gateways

## Virtual Network Route Table
Virtual Network Route Tables can now be sorted by Destination, Destination CIDR or Metric.  This can make it much easier to view and understand how routes are configured. 

{{<tgimg src="sorted-routes.png" width="80%" caption="Virtual Network Routes sorted by destination" alt="Table with multiple routes listing their Destination device, Destination CIDR, Metric and Description. Sorted by the destination device.">}}

## Container Import
### Disable Imported Containers
Previously, if you imported a container from another node or cluster it was created in the Enabled state. If the container was configured to run as a service this would cause the container to start immediately. This could create issues if additional changes were desired before the container attempts to start, such as creating volumes or adjusting environment variables. With this release the container is imported in a Disabled state and will need to be enabled once you are ready to use the container.

### Import Container Volumes
There is a new Action available under Container Management > Volumes to Import. This will enable you to import configured volumes configured on another node or cluster. This will simplify configuring containers that require several volumes to be defined.

(Note: Just like with container configuration import, the source must match the destination type. A node can import from another node and a cluster can import from another cluster)

{{< alert type="note" title="Note:" >}}
Just like with container configuration import, the source must match the destination type. A node can import from another node and a cluster can import from another cluster
{{< /alert >}}

## Flow Logs

### Flow Log Ordering
Our flow logs have historically been searched starting with the oldest in the time range first. This sometimes forced you to load more results multiple times to see the most current information. With this release we now can search either with oldest or newest first. By default the Flow Logs page will now show the last two hours with the most recent first. 

Additionally, in the Advanced Search there is a new option to select the `Flow Log Order` depending on your needs.


{{< tgimg src="FlowLogOrderOption.png" caption="Flow Log Order selector" width="50%">}}

### Flow Log S3 Export
Previously Trustgrid had a mechanism to export flow logs to an AWS S3 bucket on a nightly basis.  This was determined to be insufficient for our customers who wanted closer to real-time access to this data. With this release Trustgrid can configur near real-time S3 to S3 replication with our new [Flow Log Export]({{<ref "/docs/operations/flow-logs#flow-log-export">}}) system. 

### Flow CSV Export Date Format
Trustgrid has long supported the option to export your current view of the Flow Logs to a CSV file.  However, the format of dates was difficult to parse using most spreadsheet applications. With this release both the start time and end time (if available) is reported in a human readable format like `dd-mm-yyyy hh:mm:ss.000` and in Unix Timestamp Milliseconds (MS).  This should make analysis of the records much easier and more precise. 

{{<tgimg src="flow-log-csv-date.png" width="40%" caption="Example CSV export" alt="text CSV export example showing 'startTime' and 'startTimeMS' fields">}}


## Upgrade Improvements
### Multi-Node Upgrades
Upgrading multiple nodes used to require navigating to each node one at a time, clicking the upgrade button and then entering the node's name.   This prevented accidental upgrades but was difficult for managing large environments.  With this release users with the `nodes::service:node-upgrade` and `nodes::upgrade` permissions will have the ability to select multiple nodes from the Nodes table and then select "Upgrade" from the "Actions" menu.
{{<tgimg src="multi-node-upgrade.png" width="30%" caption="Multi-Node Upgrade Action" alt="Nodes table with two nodes selected, with the Actions dropdown showing the Upgrade option">}}

### Upgrade Status 
Starting with the node release dated `1.5.20220412-1238` nodes began reporting their upgrade status (In Process, Complete, Failed) to the Trustgrid cloud.  With this release this information and the date of the last change in status is displayed in the Info tab of nodes.
{{<tgimg src="upgrade-status.png" width="50%" caption="Upgrade Status section" alt="Table showing upgrade status of Complete and the completion time">}}


## Portal Access Permission
With this release there is a new permission `portal::access` that controls if a user is allowed (or denied) to authenticate into the Trustgrid management portal. This permission was automatically added to all [built-in Trustgrid Managed policies]({{<ref "docs/user-management/policies#trustgrid-managed-policy-descriptions" >}}), but any custom defined policies that are intended to be used on their own will need to add this permission if management portal access is desired.
{{<tgimg src="portal-access.png" width="30%" caption="Portal Access Permission" alt="Filtered list of permissions showing Portal section with portal::access allowed">}}
{{<alert color="warning">}} The `portal::access` must be applied with Resource scope of `*` to be effective. If using [resource scoped policies]({{<ref "docs/user-management/policies#resource-scoped-policies">}}) an additional policy will need to be created and attached to users to grant portal access. {{</alert>}}