---
title: "July 2023 Appliance Release Notes"
linkTitle: "July 2023"
type: docs
date: 2023-07-13
---
{{< node-release package-version="1.5.20230712-1762" core-version="20230712-200223.7cced96" release="n-2.16.0" >}}


## New Event Types
Several new [event types]({{<ref "docs/alarms/event-types">}}) are generated with this version of the node appliance software.

### DNS Failure
If the node is unable to resolve DNS via its configured DNS servers it will generate a DNS Resolution error event. 
{{<tgimg src="dns-event.png" caption="Example DNS resolution events" width="70%">}}

### Interface Down Notification
This event is triggered when any configured interface's link state goes from UP to DOWN. 
{{<tgimg src="interface-event.png" caption="Example Interface health events" width="70%">}}

### Trustgrid Unhealthy Stop Notification
This event is triggered if the Trustgrid service fails to stop completely prior to the most recent start. This could indicate a hard power cycle event or something caused the Trustgrid service to halt unexpectedly.

{{<tgimg src="trustgrid-crash.png" caption="Example Node Stop Error" width="70%">}}

### Cluster Health Events
The following cluster health related events are now available and will now resolve themselves once the triggering condition clears.
- Network Health Check
- Data and Control Plane Disruption

## Virtual Network Port Forwards
Port forwards allow mapping a [layer 4 service]({{<ref "docs/nodes/shared/services">}}) to a virtual IP address in a Trustgrid [layer 3 virtual network]({{<ref "docs/domain/virtual-networks">}}). Prior to this release, port forwards could only be defined at the node or cluster level and this required an additional virtual network route to be defined to send the virtual IP traffic to the correct node or cluster.  

The new [Virtual Network Port Forward]({{<ref "docs/domain/virtual-networks/remote-port-forward">}}) feature combines the route and port forward configuration and consolidates the definition to the virtual network level. 

## Layer 4 Health Checks Deprecated
The [layer 4 service]({{<ref "docs/nodes/shared/services">}}) historically had an option to perform a health check where a connection attempt was made once per minute to the configured IP and port.  If this was configured on a cluster, the member nodes would mark themselves unhealthy if the connection failed. Since this was usually caused by an issue with the local network or target server, this frequently led to both nodes being marked as unhealthy and taking the cluster offline. As this feature was rarely used and caused more issues than it solved we have removed this it from this version of the node software. If this was previously configured on a service the setting will be ignored. 

## Support for requiring IMDSv2 on AWS EC2 Instances.
The AWS Instance Metadata Service (IMDS) provides information about the EC2 instance that the Trustgrid service queries to enable certain features.  Some potential security issues were identified in the IMDSv1 version of this service and AWS recommends [requiring IMDSv2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-existing-instances.html#modify-require-IMDSv2) going forward.  This version of the Trustgrid node software is confirmed to be fully compatible with IMDSv2-only instances.
