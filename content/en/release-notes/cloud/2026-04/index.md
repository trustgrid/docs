---
title: "April 2026 Cloud Release Notes"
linkTitle: "April 2026"
date: 2026-04-08
description: "April 2026 Cloud Release Notes"
type: docs
---

## April 8, 2026 - Major Release

### Observability HTTP Exporter Authentication
The [HTTP exporter]({{<relref "/docs/observability#http-exporter-settings">}}) now supports authentication when sending telemetry data to external endpoints. A new **Authentication** section has been added with three **Auth Type** options:
- **None** - No authentication (previous behavior)
- **Bearer Token** - Sends a bearer token in the `Authorization` header
- **Basic Auth** - Sends a username and password

### Other Improvements and Fixes
- <!-- TODO: add bullet items -->
