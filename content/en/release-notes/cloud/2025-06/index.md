---
title:  June 2025 Major Release Notes
linkTitle:  June 2025 Major Release
date: 2025-06-01
description: "June 2025 Major Cloud Release Notes"
type: docs
---

## VLAN Sub-Interface Improvements
- Adds SNMP support for VLAN interfaces, enhancing monitoring capabilities.
- Removes the node-level restriction on VLAN interface assignments.

## Infovisor Redesign
- Displays NIC IP addresses in the Infovisor for easier interface identification.
- Improves the Infovisor layout and structure for better usability.

## Cloud Integration Enhancements
- Exposes AWS interface "allowance_exceeded" flag in the cloud UI for better visibility into quota-related issues.
- Makes the AWS Route Table service available in the portal to provide more comprehensive cloud networking configuration.
- Removes the `next hop` field for interface routes on AWS and Azure nodes as both require the gateway for that network to handle traffic.

## Gateway and Data Plane Improvements 
- Move the below settings up to a header bar since these settings are not specific to the nodes gateway or client functions.
  - Enable UDP
  - Max Ingress
  - Max Egress
- Resolves an issue where column widths in the Data Plane panel could not be adjusted, improving UI flexibility and customization.
- Restores functionality for gateway diagnostics in the portal.

## Other Improvements and Fixes
- Corrects the export functionality for the node table, ensuring that Location and ISP data are accurately included in the exported CSV file.
- Renames "TCP Errors" to "TCP Stats" in the Nodes Overview for clarity.
- Enables sorting of domain routes by description to streamline navigation.
- Enhances the Flows Tool to allow specifying "Any" as a protocol option.
- Ensures that all permissions assigned to a user are displayed in the portal, aiding in effective access management.
- Updates repository connectivity status indicators to reflect real-time "Enabled" or "Disabled" states accurately.
- Adds support for applying route monitors at the domain level, improving route oversight.
- Fixes GChat event handling so new alerts are not immediately marked as resolved, improving alert visibility and reliability in incident response workflows.
- Renames local/remote peer fields in GRE Tunnel configurations to be more descriptive.
