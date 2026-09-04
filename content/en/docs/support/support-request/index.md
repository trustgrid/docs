---
title: Support Requests
linkTitle: Support Request
description: Submit and track support requests from within the portal
---

The **Support Requests** section lets you open a request with the [Trustgrid Support Team]({{<relref "help-center/trustgrid-support" >}}) and follow it through to resolution without leaving the portal.

Requests are listed in two columns:

- **Open Support Requests** - requests Trustgrid is still working.
- **Past Support Requests** - requests that have been closed.

Each entry shows the subject, its current status, and when it was last updated.

{{<alert color="info">}}Support requests are submitted with your user account email address, and you see only the requests you opened.{{</alert>}}

## Create a Request

Click **Create Request** to open the request form. Please provide as much detail as possible about the issue you are experiencing so our support team can help resolve it quickly.

{{<tgimg src="support-req.png" caption="Support Request form" width="60%">}}

{{<fields>}}
{{<field "Phone number">}}Provide the phone number you'd like to be contacted at regarding this issue.{{</field>}}
{{<field "Urgency">}}

- Normal - Trustgrid will respond during normal business hours (Mon-Fri 8:30am-5:30pm Central)
- Emergency - only for production impacting issues; Trustgrid will respond 24x7 within established SLA. This level of support is available to paying customers only.
  {{</field>}}
  {{<field "Problem Description">}} Please list details about the nature of the issue, any particular resources being impacted (e.g. a specific destination IP:port, source IP, etc.) and troubleshooting steps already taken.{{</field>}}
  {{<field "Impacted Nodes/Clusters">}}
  This field will allow you to search for and select the impacted nodes. If a large number of nodes is impacted please provide at least a few examples and clarify the impact in the [Problem Description](#problem-description) field.
  {{<tgimg src="support-node-list.png" width="60%">}}
  {{</field>}}
  {{</fields>}}

Click **Submit** to create the support request. It appears under **Open Support Requests** immediately.

## Reply to a Request

Select any request to open it. The conversation shows each message with its author and timestamp, along with status changes as they happen.

Type in the reply box and click **Send** to respond. Your reply is attributed to your portal user account and goes to the support engineer working the request.

Click **Mark Ticket Resolved** once the issue is addressed. The request moves to **Past Support Requests**.
