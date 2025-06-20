---
title: Setting up Logstash for OpenTelemetry Export
linkTitle: Logstash
date: 2023-06-04
weight: 20
---



This guide covers how to configure Logstash to receive OpenTelemetry data from a Trustgrid HTTP exporter.

## Step 1: Navigate to Exporters

1. In the Trustgrid Portal, go to **Observability > Exporters**.
2. Click **Add Exporter**.

## Step 2: Configure the Exporter

- **Type**: Select `http`.
- **Exporter Name**: Enter a descriptive name (e.g., `logstash`).
- **Description**: Optional.
- **Instrumentation Types**: Select one or more:
  - `Metrics`
  - `Node Audits`
  - `Node Events`
  - `Changes`

{{<tgimg src="trustgrid-logstash-add-exporter.png" width="50%" caption="Add Exporter">}}

- **Endpoint**: Set this to your Logstash endpoint (e.g., `http://logstash.internal:5044`).

> Trustgrid appends `/v1/metrics` or `/v1/logs` automatically depending on the instrumentation type. You can also explicitly set:
>
> - `Metrics Endpoint`: Overrides default for metrics
> - `Logs Endpoint`: Overrides default for logs (audits, events, changes)

{{<tgimg src="trustgrid-logstash-exporter-settings.png" width="75%" caption="Exporter Settings">}}

## Step 3: Configure Logstash

Here is an example `logstash.conf` that matches the Trustgrid export format:

```ruby
input {
  http {
    port => 5044
    codec => json
  }
}

output {
  stdout {
    codec => rubydebug
  }
}
```
