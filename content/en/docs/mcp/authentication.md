---
title: "Authentication"
linkTitle: "Authentication"
weight: 10
description: "How to authenticate to the Trustgrid MCP server using OAuth, API tokens, or portal JWTs."
---

The Trustgrid MCP server requires authentication for every request. There are three ways to provide credentials.

## API token (recommended)

An API token is a `clientId:clientSecret` pair tied to your Trustgrid user account. It carries the same permissions as your portal account.

**Generate a token:**

1. Log into the Trustgrid portal
2. Navigate to **User Management** → **API Access**
3. Click **Regenerate API keys**

> The client secret is only shown once at generation time. If you lose it, regenerate.

**Use the token:**

Pass the token as an HTTP `Authorization` header using the `trustgrid-token` scheme:

```
Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET
```

In most MCP client configs, this goes in a `headers` block alongside the server URL.

## OAuth 2.0

The server implements OAuth 2.0 with the [MCP authorization spec](https://modelcontextprotocol.io/specification/2025-06-18/basic/authorization). Clients that support the OAuth handshake — including Claude Desktop and Claude Code — will automatically open a browser-based login when you first connect. After you authorize, the client manages token refresh without further intervention.

No manual token setup is required for OAuth-capable clients. Just point the client at the MCP URL.

The OAuth authorization server metadata is available at:

```
https://mcp.<domain>.trustgrid.io/.well-known/oauth-authorization-server
```

## Bearer JWT

You can use the portal JWT from your active browser session as a Bearer token. This is handy for quick interactive use but not suitable for automation since JWTs expire with your session.

**Get the JWT:**

1. Log into the Trustgrid portal in your browser
2. Open DevTools → Application → Cookies
3. Copy the value of the `portaljwt` cookie

**Use the JWT:**

```
Authorization: Bearer YOUR_PORTAL_JWT
```

## Auth error handling

| Response | Meaning | Fix |
|---|---|---|
| `401 Unauthorized` | Missing or malformed `Authorization` header | Check that the header is present and correctly formatted |
| `401` with `WWW-Authenticate` header | Server is requesting OAuth | Your client should initiate the OAuth flow; if it doesn't, provide a static token instead |
| `403 Forbidden` | Token is valid but lacks required scope | The credential doesn't have permission for the requested operation — check API key permissions or generate a new key |
| `401` with expired message | Portal JWT has expired | Re-extract the `portaljwt` cookie or use an API token instead |

If you see a `WWW-Authenticate` challenge with a `resource_metadata` URL and your client doesn't handle OAuth, you need to provide a static API token or JWT directly in the config header.

## MSP / multi-org access

If you are an MSP managing multiple organizations, the `switchOrg` tool lets you change the active org context within a session. This requires a Bearer JWT (not an API token) and is only available in the codemode scope.

```
switchOrg({ orgId: "the-target-org-id" })
```
