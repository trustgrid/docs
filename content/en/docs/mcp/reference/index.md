---
title: "Transport, endpoints, errors, and rate limits"
linkTitle: "Reference"
weight: 30
description: "Reference information for the public Trustgrid MCP server."
---

## Transport

The public Trustgrid MCP server supports **remote Streamable HTTP** endpoints served over **HTTPS**.

- Public transport: **remote Streamable HTTP over HTTPS**
- Primary endpoint: `https://mcp.trustgrid.io/mcp`
- Common recommended scoped endpoint: `https://mcp.trustgrid.io/mcp/codemode`

## Endpoint URLs

| Service | URL |
| --- | --- |
| MCP base URL | `https://mcp.trustgrid.io` |
| API base URL | `https://api.trustgrid.io` |
| Portal URL | `https://portal.trustgrid.io` |

## MCP paths

| Path | Meaning |
| --- | --- |
| `/mcp` | Bare MCP root. Public clients should generally use a scoped path instead. |
| `/mcp/codemode` | Compact agent-oriented tool surface |
| `/mcp/read` | Direct read tools |
| `/mcp/tools` | Node diagnostics |
| `/mcp/<combined-parts>` | Union of the requested groups |

## OAuth discovery endpoints

Compatible clients can use the hosted OAuth metadata endpoints:

- `/.well-known/oauth-authorization-server`
- `/.well-known/oauth-protected-resource/mcp/...`
- `/oauth/register`
- `/oauth/authorize`
- `/oauth/token`

## Error handling

### `401 Unauthorized`

Returned when the request is missing valid authorization. OAuth-capable clients receive a `WWW-Authenticate` challenge that points to the protected-resource metadata URL for the requested MCP scope.

### `403 Forbidden`

Returned by underlying Trustgrid APIs when the credential is valid but does not have the required Trustgrid permissions.

### `404 Not found`

Returned for invalid MCP scope paths or when a client tries to resume a session on a different scope path than the one it initialized against.

### `400 Bad Request`

The first MCP request must initialize the session correctly. If a client skips initialization or resumes a broken session, the server can require re-initialization.

### `500` / `502`

Server or upstream portal errors. Retry after a short backoff. If the issue persists, check Trustgrid status and support channels.

## Rate limits

Trustgrid MCP sits in front of Trustgrid APIs and may return `429 Too Many Requests` when platform or upstream limits are exceeded.

Clients should:

- retry with exponential backoff
- honor `Retry-After` if it is returned
- avoid firing many parallel tool calls against the same org unless needed
- prefer `codemode` workflows when possible so the model can batch related reads into fewer round trips

Trustgrid does not currently publish a separate fixed public MCP rate-limit table in these docs.

## Pagination

Some results are paginated.

- In `codemode`, paginated table results return a `toolCallId`
- Use `followUp` with that `toolCallId` to fetch the next page

## Auth header formats

```text
Authorization: Bearer <jwt>
Authorization: trustgrid-token <client-id>:<client-secret>
```
