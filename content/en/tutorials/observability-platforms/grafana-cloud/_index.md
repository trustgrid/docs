---
title: Setting up Grafana Cloud for OpenTelemetry Export
linkTitle: Grafana Cloud
date: 2024-01-01
weight: 30
description: >-
  This guide walks you through configuring Grafana Cloud and the Trustgrid Portal to export OpenTelemetry (OTEL) data.
---

{{< alert title="Early Access Notice" color="warning" >}}
The Observability feature is currently in early access. We are actively gathering feedback and usage data to help determine if additional charges will apply in the future. Functionality and pricing are subject to change.  
<br><br>
To enable this feature for your account, please contact Trustgrid Support.
{{< /alert >}}

This guide will walk you through configuring Grafana Cloud to receive OpenTelemetry data from the Trustgrid [Observability]({{<ref "/docs/observability" >}}) feature.

---

## Part 1: Configure Grafana Cloud to Receive OTEL Data

### Step 1: Add a New Connection

1. From the Grafana Cloud Navigation Menu, select **Connections** --> **Add new connection**

{{<tgimg src="grafana-connections-menu.png" width="80%" caption="Grafana Cloud Connections menu">}}

---

### Step 2: Search for OpenTelemetry (OTLP)

1. In the **Search Connections** box, search for `OTLP`
2. Select **OpenTelemetry (OTLP)** from the results

{{<tgimg src="grafana-search-otlp.png" width="80%" caption="Searching for OTLP in Add new connection">}}

---

### Step 3: Choose Your Instrumentation Method

1. Under **Choose your instrumentation method**, select **OpenTelemetry SDK**

{{<tgimg src="grafana-instrumentation-method.png" width="80%" caption="Selecting OpenTelemetry SDK as the instrumentation method">}}

---

### Step 4: Choose Your Language

1. Under **Choose your language**, select **Other**
2. Click **Next**

{{<tgimg src="grafana-choose-language.png" width="80%" caption="Selecting Other as the language">}}

---

### Step 5: Choose Your Infrastructure

1. Under **Choose your infrastructure**, select **Linux**
2. Click **Next**

{{<tgimg src="grafana-choose-infrastructure.png" width="80%" caption="Selecting Linux as the infrastructure">}}

---

### Step 6: Choose Your Instrumentation Method

1. Under **Choose your instrumentation method**, select **OpenTelemetry Collector**

{{<tgimg src="grafana-instrumentation-collector.png" width="80%" caption="Selecting OpenTelemetry Collector as the instrumentation method">}}

---

### Step 7: Create a Grafana Cloud Access Token

1. In the **Create a Grafana Cloud access token** box, enter a name for your new token
2. Click **Create token**

{{<tgimg src="grafana-create-token.png" width="80%" caption="Creating a Grafana Cloud access token">}}

---

### Step 8: Note Your Configuration Values

1. Scroll down to the section labeled **Configure the OpenTelemetry Collector**
2. Make note of the following values -- you will need them to configure your exporter in the Trustgrid portal:
   - `GRAFANA_CLOUD_OTLP_ENDPOINT`
   - `GRAFANA_CLOUD_INSTANCE_ID`
   - `GRAFANA_CLOUD_API_KEY`

{{<tgimg src="grafana-configure-collector.png" width="80%" caption="Configure the OpenTelemetry Collector section showing the required values">}}

---

### Step 9: Finish Setup (Optional)

1. Select **Next** until you reach the **Finish set up** screen. This will begin watching for OTEL data being sent to Grafana Cloud from your exporter.

{{<tgimg src="grafana-finish-setup.png" width="80%" caption="Grafana Cloud Finish set up screen checking for incoming OTEL data">}}

---

## Part 2: Configure the Trustgrid OTEL Exporter

### Step 10: Navigate to Observability and Create an Exporter

1. In the Trustgrid Portal, navigate to [Observability]({{<ref "/docs/observability" >}})
2. Click **+ Create Exporter**

{{<tgimg src="trustgrid-observability-exporters.png" width="80%" caption="Trustgrid Observability Exporters page">}}

---

### Step 11: Add Exporter Details

1. In the **Add Exporter** dialog:
   - Set the **Type** to `http`
   - Enter a **Name** for your exporter
   - Select the **Instrumentation Types** you want to send to Grafana Cloud
2. Click **Save**

{{<tgimg src="trustgrid-add-exporter.png" width="80%" caption="Add Exporter dialog">}}

---

### Step 12: Configure the Exporter

1. In the exporter configuration screen:
   - Set the **Endpoint** to the value from `GRAFANA_CLOUD_OTLP_ENDPOINT` noted in Step 8
   - Under **Authentication**, select **Basic Auth**
   - Set the **Username** to the value from `GRAFANA_CLOUD_INSTANCE_ID`
   - Set the **Password** to the value from `GRAFANA_CLOUD_API_KEY`
2. Click **Save**

{{<tgimg src="trustgrid-exporter-config.png" width="80%" caption="Trustgrid exporter configuration with Grafana Cloud values">}}

---

Your Grafana Cloud OpenTelemetry configuration is now complete. You should begin seeing the instrumentation types you selected flowing into Grafana Cloud.

---

## Sample Dashboard

Trustgrid provides a sample Grafana dashboard that visualizes key node metrics exported via OTEL. You can import it directly into your Grafana instance from the Grafana dashboard library:

**[Trustgrid Node Dashboard](https://grafana.com/grafana/dashboards/25256-trustgrid-node-dashboard/)**

To import it:

1. In Grafana, go to **Dashboards** --> **Import**
2. Enter the dashboard ID `25256` or paste the URL above
3. Click **Load**, select your Prometheus/Mimir datasource, and click **Import**
