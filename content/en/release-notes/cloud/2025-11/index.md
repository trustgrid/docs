---
title:  November 2025 Minor Release Notes
linkTitle:  November 2025 Minor Release
date: 2025-11-06
description: "November 2025 Minor Cloud Release Notes"
type: docs
---

## Validation Fixes
* Portal now allows CIDRs for interface IPs to have a /31 subnet. Previously, broadcast and network address validations prevented this from saving.
* Fixes the `client` field validation for BGP configuration.

## Cluster Config Changes
* Cluster configuration sent to nodes (appliances) now includes a version counter managed in the control plane. This prevents split brain scenarios where a node that is unable to connect to the control plane could repeatedly attempt to become the active member, even when the connected member has been designated active. This functionality is utilized by nodes running the [November 2025 Minor Appliance Release](../node/2025-11/) release. 
