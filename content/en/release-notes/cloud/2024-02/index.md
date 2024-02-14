---
title: February 2024 Release Notes - Minor
linkTitle: 'February 2024 Release'
type: docs
date: 2024-02-14
description: "February 2024 Cloud Release - Minor improvements and fixes"
---

## Flow Log Performance Improvements
This release further increases the performance of flow log search using our new database. It also resolves an issue that prevented all logs from being exported to CSV as expected. 

## Other Improvements
- The MTR interface in the portal now supports selecting a source IP from your interface and cluster IPs.
- The Reboot and Restart buttons in the portal will now be disabled if a node is offline. This is to make clear you cannot perform these actions on an offline node. 

## Other Fixes
- Resolves issue preventing adding Azure AD as an identity provider.
- Resolved an issue that could cause Order-based alerts not to be sent to configured channels.
- Adds input validation for AWS and Azure Route Table identifiers.
- [Debug Logs]({{<relref "/help-center/ops-logs/debug-logs">}}) tool now prevents being triggered twice and now supports viewing greater than 20 historic log archives.