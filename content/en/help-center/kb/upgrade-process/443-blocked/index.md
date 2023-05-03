---
title: Node Behavior When Port 443 is blocked
---

## Symptoms

- Node will not update either automatically or when manually triggered
- Node will not be able to pull down container images from the organization repository

## Cause

- The Trustgrid node cannot connect to repo.trustgrid.io on port 443 to update packages

## Resolution

- Confirm the node can resolve repo.trustgrid.io
- Ensure the node can connect to repo.trustgrid.io:443

