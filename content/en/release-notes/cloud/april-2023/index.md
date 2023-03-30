---
title: April 2023 Release Notes
linkTitle: 'April 2023'
type: docs
date: 2023-04-04
resources:
   - src: FlowLogOrderOption.png
     params:
       byline: "Flow Log Order Selector"   

----
## Documentation 
Trustgrid is moving away from our documentation previously hosted on Confluence to docs-as-code system at https://docs.trustgrid.io. This should enable us to more tightly integrate the portal UI with related documentation. 

As a start, we've created a link called Documentation under the Management section of the main navigation bar that will open the docs in a new tab.

Additionally, we've created a section for release notes that will allow you to monitor release for [Control Plane](https://docs.trustgrid.io/release-notes/cloud/) and [Node](https://docs.trustgrid.io/release-notes/node/) releases via RSS feeds going forward. 

## Data Plane Stats
This release further improves on our recent addition of data plane telemetry information.
* Show data plane telemetry for offline nodes
* Show data plane telemetry for connections between public and private gateways

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
Our flow logs have historically been searched starting with the oldest in the time range first. This sometimes forced you to load more results multiple times to see the most current information. With this release we now can search either with oldest or newest first. By default the Flow Logs page will now show the last two hours with the most recent first. 

Additionally, in the Advanced Search there is a new option to select the `Flow Log Order` depending on your needs.


{{< tgimg src="FlowLogOrderOption.png" caption="Flow Log Order selector" width="50%">}}