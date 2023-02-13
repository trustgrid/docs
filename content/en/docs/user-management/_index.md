---
Title: "User Management"
Date: 2022-12-30
weight: 21
---

{{% pageinfo %}}
User Management is used to securely control access to Trustgrid resources. Identities (users) are authenticated and granted permissions (via policies) to view or manage Trustgrid devices and applications.
{{% /pageinfo %}}

#### How It Works

1. Users' identities are authenticated via one of three mechanisms:
   a. Username, password and MFA via Trustgrid’s native authentication system
   b. Via a configured Identity Provider (IdP) using OpenID or SAML
   c. Trustgrid generated API token
2. Permissions are determined via one or more attached policies
3. These permissions are evaluated as the user makes requests against either via Trustgrid Portal or API
