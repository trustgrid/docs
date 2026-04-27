---
title: "Authentication"
linkTitle: "Authentication"
weight: 10
description: "Authenticate to the Trustgrid MCP server with OAuth, JWTs, or API tokens."
---

Trustgrid MCP accepts two HTTP authorization schemes:

- `Authorization: Bearer <token>`
- `Authorization: trustgrid-token <client-id>:<client-secret>`

## OAuth

If your MCP client supports remote MCP OAuth, the simplest setup is to add the MCP URL and complete the browser login flow when prompted.

Use OAuth when:

- a human is using Claude, Cursor, VS Code, Codex, or a similar client
- you want the session to run as your own Trustgrid user
- you want token refresh handled by the client

Trustgrid's public MCP server uses remote Streamable HTTP, which means the client connects to a hosted HTTPS MCP endpoint instead of launching a local stdio process.

The hosted MCP server exposes the standard discovery and OAuth endpoints needed by compatible clients.

Use the environment-specific base URL that matches the MCP endpoint you configured:

| Environment | OAuth base URL | Discovery and OAuth endpoints |
| --- | --- | --- |
| Production | `https://mcp.trustgrid.io` | `https://mcp.trustgrid.io/.well-known/oauth-authorization-server`, `https://mcp.trustgrid.io/.well-known/oauth-protected-resource/mcp/...`, `https://mcp.trustgrid.io/oauth/authorize`, `https://mcp.trustgrid.io/oauth/token`, `https://mcp.trustgrid.io/oauth/register` |
| Stage | `https://mcp.stage.trustgrid.io` | `https://mcp.stage.trustgrid.io/.well-known/oauth-authorization-server`, `https://mcp.stage.trustgrid.io/.well-known/oauth-protected-resource/mcp/...`, `https://mcp.stage.trustgrid.io/oauth/authorize`, `https://mcp.stage.trustgrid.io/oauth/token`, `https://mcp.stage.trustgrid.io/oauth/register` |
| Test | `https://mcp.test.trustgrid.io` | `https://mcp.test.trustgrid.io/.well-known/oauth-authorization-server`, `https://mcp.test.trustgrid.io/.well-known/oauth-protected-resource/mcp/...`, `https://mcp.test.trustgrid.io/oauth/authorize`, `https://mcp.test.trustgrid.io/oauth/token`, `https://mcp.test.trustgrid.io/oauth/register` |
| Dev | `https://mcp.dev.trustgrid.io` | `https://mcp.dev.trustgrid.io/.well-known/oauth-authorization-server`, `https://mcp.dev.trustgrid.io/.well-known/oauth-protected-resource/mcp/...`, `https://mcp.dev.trustgrid.io/oauth/authorize`, `https://mcp.dev.trustgrid.io/oauth/token`, `https://mcp.dev.trustgrid.io/oauth/register` |

If you are using a non-production MCP endpoint, do not mix it with production OAuth endpoints. Keep the MCP URL, OAuth metadata URLs, and browser login flow on the same environment host.

## JWT Bearer tokens

If your client does not use OAuth, you can send a JWT directly:

```text
Authorization: Bearer <jwt>
```

JWT auth is useful when:

- your client supports custom HTTP headers but not MCP OAuth
- you need to test with a known user session

See [API Access]({{<relref "/docs/user-management/API-access" >}}) for general Trustgrid API credential guidance.

## API tokens

Trustgrid API tokens use the `trustgrid-token` scheme:

```text
Authorization: trustgrid-token <client-id>:<client-secret>
```

Use API tokens when:

- you are wiring Trustgrid MCP into automation
- your client supports static headers but not OAuth
- you want a non-human service identity

API tokens can be generated for users and [Service Users]({{<relref "/docs/user-management/service-users" >}}).

## Which should I use?

| Use case | Recommended auth |
| --- | --- |
| Interactive use in a client with MCP OAuth support | OAuth |
| Interactive use in a client without MCP OAuth support | JWT Bearer token |
| CI, agents, or long-lived automation | API token |

## Common auth failures

| Status | Meaning | What to do |
| --- | --- | --- |
| `401 Unauthorized` | Missing or invalid auth | Re-authenticate, refresh the header, or complete the OAuth flow |
| `403 Forbidden` | Authenticated, but your user or token lacks required permissions | Add the required Trustgrid policies or use a credential with broader access |
| `404 Session not found` | The client resumed a session against the wrong MCP scope path | Re-initialize against the same `/mcp/...` path you started with |

When OAuth is available, the server returns a `WWW-Authenticate` challenge that points the client at the protected-resource metadata URL for the requested MCP scope.
