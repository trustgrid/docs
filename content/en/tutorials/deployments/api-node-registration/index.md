---
title: "Register a Node via the Trustgrid API"
linkTitle: "API Node Registration"
description: "How to use the Trustgrid REST API to generate node licenses and register appliance nodes programmatically"
tags: ["api", "node", "registration", "tutorial"]
---

The Trustgrid API exposes a license generation endpoint that allows you to register appliance nodes programmatically — without using the portal UI. This is useful for automated deployments, infrastructure-as-code workflows, and scripted provisioning pipelines.

## Overview

The traditional portal workflow for registering a new node appliance is:

1. Navigate to the **Nodes** page and click **Add Node**
2. Enter a node name and click **Create License**
3. Copy the generated license key and use it during node setup

The API-based registration replaces steps 1–2 with a single `GET /node/license` request and is otherwise identical: the returned license key is used the same way during node installation.

## Prerequisites

- A Trustgrid API key (client ID and client secret). See [API Access]({{<ref "/docs/user-management/API-access">}}) for instructions on generating one. For automated pipelines, consider using a [Service User]({{<ref "/docs/user-management/service-users">}}) credential instead of a personal API key.
- The API key's associated user must have permission to create node licenses. Users with the `builtin-tg-admin` policy have this access.
- A Trustgrid node running a supported image that is ready to be registered.
- `curl` and (for the verification step) [`jq`](https://jqlang.github.io/jq/) installed on the machine running these commands.

## API Reference

### Generate a Node License

Generates a license key that an appliance node uses to register with your Trustgrid organization.

| | |
|---|---|
| **Method** | `GET` |
| **Path** | `/node/license` |
| **Base URL** | `https://api.trustgrid.io` |
| **Tag** | Appliance |
| **Full spec** | [apidocs.trustgrid.io](https://apidocs.trustgrid.io) |

#### Request

**Authentication header**

```
Authorization: trustgrid-token <CLIENT-ID>:<CLIENT-SECRET>
```

**Query parameters**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | The name to assign to the new node. Must be unique within the organization. |

#### Response

| Status | Content-Type | Description |
|--------|--------------|-------------|
| `200 OK` | `text/plain` | The license key body — a multi-line plain-text string used during node registration. |
| `422 Unprocessable Entity` | `text/plain` | Validation error. The response body describes the problem (for example, a duplicate node name). |

## Tutorial: Registering a Node via the API

### Step 1 — Generate the License Key

Use `curl` to call the license endpoint with your API credentials and the desired node name. The `-f` flag causes `curl` to exit with a non-zero status on HTTP errors:

```bash
curl -sf \
  -H "Authorization: trustgrid-token YOUR-CLIENT-ID:YOUR-CLIENT-SECRET" \
  -G --data-urlencode "name=my-edge-node" \
  "https://api.trustgrid.io/node/license" \
  -o my-edge-node.lic
```

A successful response writes the license key to `my-edge-node.lic`. If the command fails (e.g., `curl` exits non-zero), check the HTTP response body for the error message — a `422` typically means the node name is already taken.

{{<alert color="info">}}
The node name provided in the `name` parameter becomes part of the FQDN used to identify the node in the portal. Choose a name that clearly identifies the node's role and location (e.g., `prod-gateway-us-east-1`). Node names must be unique within the organization.
{{</alert>}}

### Step 2 — Use the License During Node Setup

The license key generated in Step 1 is used exactly the same way as a license downloaded from the portal. Apply it during Trustgrid node installation per the relevant deployment guide:

- [Deploy to AWS]({{<ref "/tutorials/deployments/deploy-aws-ami">}})
- [Deploy to Azure]({{<ref "/tutorials/deployments/deploy-azure">}})
- [Deploy to GCP]({{<ref "/tutorials/deployments/deploy-gcp">}})
- [Deploy to vSphere]({{<ref "/tutorials/deployments/deploy-vsphere">}})

### Step 3 — Verify the Node Is Registered

Once the node has completed registration and connected to the Trustgrid control plane, you can confirm it is visible via the API. The example below uses `jq` to filter the node list:

```bash
curl -s \
  -H "Authorization: trustgrid-token YOUR-CLIENT-ID:YOUR-CLIENT-SECRET" \
  "https://api.trustgrid.io/node" | \
  jq '.[] | select(.name | startswith("my-edge-node")) | {name, uid, online, state}'
```

Example output for a successfully registered node:

```json
{
  "name": "my-edge-node.your-domain.trustgrid.io",
  "uid": "abc12345-...",
  "online": false,
  "state": "ACTIVE"
}
```

{{<alert color="info">}}
A newly registered node will show `"online": false` until it establishes its first connection to the Trustgrid control plane. `"state": "ACTIVE"` confirms the license is valid and the node is registered. The full node FQDN suffix (`.your-domain.trustgrid.io`) is specific to your organization.
{{</alert>}}

### Scripted Example

The following shell script combines Steps 1 and 2 into a reusable registration helper. It uses `--data-urlencode` to safely handle node names with spaces or special characters:

```bash
#!/usr/bin/env bash
# register-node.sh — generate a Trustgrid node license via the API
# Usage: ./register-node.sh <node-name> <output-file>

set -euo pipefail

NODE_NAME="${1:?Usage: $0 <node-name> <output-file>}"
OUTPUT_FILE="${2:?Usage: $0 <node-name> <output-file>}"

: "${TG_CLIENT_ID:?Set TG_CLIENT_ID environment variable}"
: "${TG_CLIENT_SECRET:?Set TG_CLIENT_SECRET environment variable}"

echo "Generating license for node: ${NODE_NAME}"

curl -sf \
  -H "Authorization: trustgrid-token ${TG_CLIENT_ID}:${TG_CLIENT_SECRET}" \
  -G --data-urlencode "name=${NODE_NAME}" \
  "https://api.trustgrid.io/node/license" \
  -o "${OUTPUT_FILE}"

echo "License saved to: ${OUTPUT_FILE}"
```

Run it by exporting your credentials as environment variables:

```bash
export TG_CLIENT_ID="your-client-id"
export TG_CLIENT_SECRET="your-client-secret"

./register-node.sh my-edge-node my-edge-node.lic
```

## Related Resources

- [API Access]({{<ref "/docs/user-management/API-access">}}) — how to generate and use API credentials
- [Service Users]({{<ref "/docs/user-management/service-users">}}) — machine-to-machine credentials for automated integrations
- [Remote Console Registration]({{<ref "/tutorials/local-console-utility/remote-registration">}}) — register a node interactively from the console (no API required)
- [Trustgrid API Documentation](https://apidocs.trustgrid.io) — full interactive API reference (Swagger UI)
