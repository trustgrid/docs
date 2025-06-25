---
title: Setting up Splunk for OpenTelemetry Export
linkTitle: Splunk
date: 2023-06-04
weight: 20
description: This guide walks you through configuring Splunk and the Trustgrid Portal to export OpenTelemetry (OTEL) data via the HTTP Event Collector (HEC) integration.
---

{{< alert title="Early Access Notice" color="warning" >}}
The Observability feature is currently in early access. We are actively gathering feedback and usage data to help determine if additional charges will apply in the future. Functionality and pricing are subject to change.  
<br><br>
To enable this feature for your account, please contact Trustgrid Support.
{{< /alert >}}

---

## Part 1: Configure Splunk to Receive OTEL Data

### Step 1: Enable HEC in Splunk

1. Log into your Splunk instance.
2. Navigate to: **Settings > Data Inputs**
3. Click **Add New** next to **HTTP Event Collector (HEC)**  
   {{<tgimg src="splunk-add-hec.png" width="95%" caption="Add HEC Input">}}

---

### Step 2: Create a New Input

1. Enter a **Name** for the input (e.g., `Trustgrid`)
2. Leave other options at their default values unless otherwise required.
3. Click **Next**  
   {{<tgimg src="splunk-add-input.png" width="95%" caption="HEC Input Config">}}

---

### Step 3: Create a New Index

1. On the **Input Settings** step, click **Create a new index**
2. Provide a name (e.g., `otel_metrics`) and set the **Index Data Type** to `Metrics` (Node Telemetry) or `Events` (Node Audits, Changes, Events)

   #### Metric index

   {{<tgimg src="splunk-metric-index.png" width="75%" caption="Create Metric Index">}}

   #### Event (Node Audits, Changes, Events) index

   {{<tgimg src="splunk-event-index.png" width="75%" caption="Create Event Index">}}

3. Save the new index
4. Ensure it's:
   - Added to the **Allowed Indexes**
   - Selected as the **Default Index**
     {{<tgimg src="splunk-index-settings.png" width="75%" caption="Index Settings">}}

---

### Step 4: Complete the Setup

1. Click **Review** and then **Done**
2. Copy the **Token Value** generated — you'll use this in the Trustgrid Portal setup  
   {{<tgimg src="splunk-token-success.png" width="75%" caption="Token Created">}}

---

## Part 2: Configure Trustgrid to Export OTEL Data

### Step 1: Access Exporter Settings

1. Log into the **Trustgrid Portal**
2. Navigate to: **Management > Observability**
3. Click **Add Exporter**

---

### Step 2: Configure Exporter

- **Type**: `splunk`
- **Name**: e.g., `splunk-prod-exporter`
- **Description**: Optional
- **Instrumentation Types** (select one or more):

  - `Metrics`
  - `Node Audits`
  - `Node Events`
  - `Changes`

{{<tgimg src="trustgrid-add-exporter-splunk.png" width="50%" caption="Token Created">}}

---

### Step 3: Provide Endpoint Details

- **Endpoint**: The full URL to your Splunk HEC endpoint, e.g. `https://your-splunk-host:8088`
- **Token**: Paste the token you copied earlier
- **Source** / **Source Type**: Optional — maps to Splunk's source fields
- **Index**: Optional - The name of the index created in Splunk (e.g., `metrics`)
- **Enable Data Logging** Optional - Used when sending log data like Node Audits, Events, and Changes to Splunk
- **TLS Custom CA**: Optional - Upload a certificate if using a custom CA
- **Skip TLS Verification**: Optional - Will skip TLS certificate verification (TLS will still be enabled)

{{<tgimg src="trustgrid-splunk-exporter-config.png" width="85%" caption="Splunk Exporter Config">}}

---
