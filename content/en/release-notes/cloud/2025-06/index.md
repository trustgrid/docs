---
title:  June 2025 Major Release Notes
linkTitle:  June 2025 Major Release
date: 2025-06-25
description: "June 2025 Major Cloud Release Notes"
type: docs
---

## VLAN Sub-Interface Improvements
- Adds SNMP support for VLAN interfaces, enhancing monitoring capabilities.
- Removes the node-level restriction on VLAN interface assignments.

## Infovisor Redesign
This release includes a significant redesign of the Infovisor, providing a more intuitive and user-friendly interface. Some key improvements include:
- Info is now grouped into multiple panels that can be collapsed or expanded, and even added/removed for easier navigation.
- All info fields now include a copy button when hovered over, allowing for easy copying of information. {{<tgimg src="copy-button.png" width="35%" caption="Copy Button">}}
- Displays NIC IP addresses in the Infovisor for easier interface identification.

## Cloud Integration Enhancements
- Adds new interface tools on AWS Nodes:
    - `AWS Route Tables` - This service will query the AWS API for all route tables associated with the interface and the associated routes and destinations.  This is useful for troubleshooting routes managed by clustered Trustgrid nodes.
    - `AWS Stats` - This service will retrieve <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-network-performance-ena.html#network-performance-metrics" target="_blank" rel="noopener"> metrics for the ENA Driver</a> providing the ability to check if interface limits are being exceeded.
- Removes the `next hop` field for interface routes on AWS and Azure nodes as both require the gateway for that network to handle traffic.

## Gateway and Data Plane Improvements 
- Move the below settings up to a header bar since these settings are not specific to the nodes gateway server or client functions.
  - Enable UDP
  - Max Ingress
  - Max Egress
- Resolves an issue where column widths in the Data Plane panel could not be adjusted, improving UI flexibility and customization.
- Restores functionality for gateway diagnostics in the portal.


## Observability
This release includes early access to the [Observability]({{<relref "/docs/observability">}}) feature, which allows for exporting data such as metrics and events to external systems using the OpenTelemetry (OTel) protocol. Trustgrid will use this early access stage to gather feedback and usage data to help determine if additional charges will apply in the future. Functionality and pricing are subject to change. Contact Trustgrid Support to enable this feature for your account.

Prelimiary testing has been done with [Splunk]({{<relref "/tutorials/observability-platforms/splunk">}}) and [Logstash]({{<relref "/tutorials/observability-platforms/logstash">}}) as exporters.

## Other Improvements and Fixes
- Generating a new API key now invalidates the previous key correctly
- Corrects the export functionality for the node table, ensuring that Location and ISP data are accurately included in the exported CSV file.
- Renames "TCP Errors" to "TCP Stats" in the Nodes Overview for clarity.
- Enables sorting of domain routes by description to streamline navigation.
- Enhances the interface `Flows` Tool to allow specifying "Any" as a protocol option.
- Ensures that all permissions assigned to a user are displayed in the portal, aiding in effective access management.
- Changes the Repo Connectivity status to "Healthy" or "Unhealthy" instead of "Enabled" or "Disabled" for better clarity.
- Adds support for applying route monitors at the domain level, improving route oversight.
- Fixes GChat event handling so new alerts are not immediately marked as resolved, improving alert visibility and reliability in incident response workflows.
- Renames local/remote peer fields in GRE Tunnel configurations to be more descriptive.
- Fixes an issue where IDP user/group sync could fail to refresh daily.
