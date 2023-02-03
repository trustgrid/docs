---
Title: "Authentication Audits"
Date: 2023-2-3
---

{{% pageinfo %}}
Trustgrid records information for authentication attempts into the control plane. This information can be used to troubleshoot authentication issues, or to monitor for suspicious activity.
{{% /pageinfo %}}

Authentication audits are available by navigating to `Operations`->`Authentication Audits`. Only the most recent 25 records are shown, but the last 30 days of audits can be downloaded by clicking the link at the bottom of the page.

{{<alert color="warning">}}
Successful logins from external identity providers are logged, but failed attempts are not. Check your IDP's documentation for information on how to view those.
{{</alert>}}
