---
title: "Troubleshooting"
linkTitle: "Troubleshooting"
weight: 40
description: "Common MCP server setup issues and how to fix them."
---

## Common issues

### The server isn't appearing in my client

**Check that your client supports HTTP transport.** Not all MCP clients support Streamable HTTP — some only support stdio (local process) connections. The Trustgrid MCP server is HTTP-only. Verify your client's MCP documentation.

**Check the URL format.** The URL must include the path prefix:

```
https://mcp.trustgrid.io/mcp          ✓ correct
https://mcp.trustgrid.io              ✗ missing /mcp
https://mcp.trustgrid.io/mcp/         ✓ trailing slash is fine
```

**Check your domain.** If your organization has a tenant-specific endpoint (e.g., `mcp.acme.trustgrid.io`), use that rather than the default `mcp.trustgrid.io`.

---

### I'm getting 401 Unauthorized

**Missing or malformed header.** Confirm the `Authorization` header is present and uses the correct format:

```
# API token
Authorization: trustgrid-token YOUR_CLIENT_ID:YOUR_CLIENT_SECRET

# Bearer JWT
Authorization: Bearer YOUR_PORTAL_JWT
```

A common mistake is using `Bearer` with an API token: `Bearer clientId:secret` won't work. Use `trustgrid-token clientId:secret`.

**Expired JWT.** Portal JWTs expire with your browser session. Re-extract the `portaljwt` cookie or switch to an API token.

**OAuth loop.** If the server returns a `WWW-Authenticate` challenge and your client doesn't handle OAuth, it may retry indefinitely. Configure a static token in the headers instead.

---

### I'm getting 403 Forbidden

Your credentials are valid but lack permission for the requested operation. Check:

- The API key was generated for the correct user account
- The user account has appropriate permissions in the Trustgrid portal
- You're using the correct endpoint path — `/mcp/tools` requires node service scopes that `/mcp/read` does not request

---

### Tools work but return errors for specific nodes

**Node identifier.** Tool calls accept node UID, FQDN, or bare name. If you're using a bare name that isn't unique across your organization, the server may resolve to the wrong node. Use the UID or full FQDN for unambiguous lookups.

**Node is offline.** Most diagnostic tools require the node to be reachable. If the node is disconnected from the control plane, `test_*` tools will time out or return an error.

---

### Rate limit errors (429)

You're hitting the per-token request limit for the active group. Limits are:

- codemode: 10 requests / 60 seconds
- read: 30 requests / 60 seconds
- tools: 10 requests / 60 seconds

The response includes `Retry-After` header with the wait time in seconds. If you're hitting limits consistently, consider using the `read` group instead of `codemode` for read-heavy workflows, or split high-frequency queries across multiple credentials.

---

### OAuth won't complete / browser auth fails

**Check the portal URL.** OAuth flows redirect through `portal.trustgrid.io` (or your tenant portal URL). If your network blocks the portal, the OAuth flow will fail. Use a static API token instead.

**Check for cookie blockers.** Some browser extensions block third-party cookies and can interfere with OAuth callbacks. Try in a private window.

---

### The `code` tool says "not allowed" or isn't available

The `code` tool is only in the **codemode** group (`/mcp/codemode` or `/mcp`). If you're connected to `/mcp/read` or `/mcp/tools`, the code sandbox isn't available. Change your client config to point at `/mcp` or `/mcp/codemode`.

---

### `switchOrg` fails

`switchOrg` requires a Bearer JWT (portal JWT), not an API token. It is not available via `trustgrid-token clientId:secret` credentials. Log into the portal and use the `portaljwt` cookie value as your Bearer token.

---

## Submit feedback

Found a bug, missing tool, or unexpected behavior? Open an issue at:

**[github.com/trustgrid/mcp-features](https://github.com/trustgrid/mcp-features)**

Helpful things to include in your report:

- **Client name and version** — e.g., "Claude Code 1.2.3", "Cursor 0.45"
- **Tool group / endpoint URL** — `/mcp/codemode`, `/mcp/read`, etc.
- **The tool that failed** — the exact tool name (e.g., `test_tcp_connectivity`, `code`)
- **What you asked for** — the prompt or tool call that triggered the issue
- **What you got back** — the error message or unexpected response, including HTTP status codes if visible
- **What you expected** — a brief description of the intended behavior
- **Auth method** — OAuth, API token, or Bearer JWT (don't include the actual credential)

The more context you provide, the faster we can reproduce and fix it.
