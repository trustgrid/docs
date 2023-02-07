---
title: Edge Node Behavior When Port 443 is blocked
---

## Symptoms

- Node will not update either automatically or when manually triggered

## Cause

- The Trustgrid edge node cannot connect to repo.trustgrid.io on port 443 to update packages

## Resolution

- Confirm the node can resolve repo.trustgrid.io
- Ensure the node can connect to repo.trustgrid.io:443
- [Verify no DPI-SSL or HTTPS proxy]({{<ref "ssl-tls-tampering">}}) is interfereing with the SSL certificate for repo.trustgrid.io
