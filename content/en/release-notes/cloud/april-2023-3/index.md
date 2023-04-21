---
title: April 2023 Third Release Notes
linkTitle: 'April 2023 Third Release'
type: docs
date: 2023-04-30
---

## Change Audit Improvements
Several improvements where made around our change auditing system including:
* Tag changes are now Audited.
* Change records can now be replicated into a customer's AWS S3 bucket. This works much like the existing [S3 Flow Log Export]({{<ref "/docs/operations/flow-logs#flow-log-export">}}) and requires the same bucket policy and versioning settings.  To have this setup contact Trustgrid support. 
* Changes to Container settings now include the name of the container instead of its unique ID.
* Changes to a policy will now show what resources are covered by the policy.

## Flow Log Advanced Search
With this release you can now set the source or destination node to `local` for an Advanced Flow Log Search. This will cover flows that do not traverse the data plane, such as traffic that is forwarded between interfaces.

