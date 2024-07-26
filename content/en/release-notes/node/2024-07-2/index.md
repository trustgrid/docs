---
title: July 2024 Major Appliance Release Notes
linkTitle: July 2024 Major
type: docs
date: 2024-07-28
description: "Release notes for the July 2024 Major Trustgrid Appliance release"
---
{{< node-release package-version="TBD" core-version="TBD" release="n-2.20.0" >}}
## Configuration Retrieval Improvements
Appliances retrieve their configuration from the control plane services. Before this release, it was not obvious if a node was unable to pull down its configuration which could lead to unexpected behavior.  Two improvements have been made:
- A new event called `Configuration Update Failure` has been created that the appliance will send if it is unable to [connect to the configuration control plane API endpoint]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}). {{<tgimg src="config-update-fail.png" caption="Example Configuration Update Failure event">}}
- Appliances attempt to pull the latest configuration at startup. With previous versions, the appliance would attempt this multiple times before continuing the startup process. It was discovered that in some cases, such as if the connection timed out rather than being immediately rejected, this would cause the startup process to take more than 15 minutes. This release changes the flow so the node will attempt once and if that fails it will startup with the local configuration and continue attempting the retrieval in the background.

## Other Improvements and Fixes
- Resolves an issue where containers reported their running status incorrectly when Enabled and Stopped
- Nodes will no longer try to connect to gateways once they begin their shutdown/restart process. Previously this caused error messages in the logs.
- Resolves an issue on clustered nodes that lead to Netplan errors during shutdown
