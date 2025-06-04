---
categories: ["monitoring"]
tags: ["monitoring"]
title: "Observability"
date: 2025-06-03
weight: 4
---

{{% pageinfo %}}
Observability enables exporters for Metrics (stats such as CPU, Memory, Disk, bytes sent/rcv'd, ..etc), [Events]({{<ref "/docs/alarms/events" >}}), [Node Audits]({{<ref "/docs/nodes/shared/audits">}}), and [Configuration Changes]({{<ref "/docs/operations/changes" >}})
{{% /pageinfo %}}

Trustgrid currently supports an OpenTelemetry exporter for [Splunk]({{<ref "/tutorials/observability-platforms/splunk" >}}), as well as a generic HTTP exporter for platforms that accept either JSON or OTLP-encoded data, such as [Logstash]({{<ref "/tutorials/observability-platforms/logstash" >}}).

![img](observability-exporters.png)

#### Adding an Exporter

![img](create-exporter.png)

{{<fields>}}
{{<field "Type">}}
Splunk & HTTP are currently suppored
{{</field >}}
{{<field "Name">}}
Your desired exporter name
{{</field>}}
{{<field "Description">}}
A brief description for your exporter
{{</field>}}
{{<field "Instrumentation Types">}}
Categories of data exported by Trustgrid.
{{</field>}}
{{</fields>}}
{{<alert>}}The Exporter type cannot be changed after adding an exporter{{</alert>}}

#### Splunk Exporter Settings

![img](splunk-exporter.png)
{{<fields>}}
{{<field "Status">}}
Current exporter status.
{{</field>}}
{{<field "Name">}}
Current exporter name.
{{</field>}}
{{<field "Exporter Type">}}
Configured exporter type.
{{</field>}}
{{<field "Instrumentation Types">}}
Current categories of data exported by Trustgrid (e.g. Metrics).
{{</field>}}
{{<field "Endpoint">}}
The destination URL for the Splunk endpoint.
{{</field>}}
{{<field "Token">}}
The authentication token used to send data to the Splunk endpoint.
{{</field>}}
{{<field "Source">}}
Optional: A custom source label for events sent to Splunk.
{{</field>}}
{{<field "Source Type">}}
Optional: A Splunk sourcetype value to help categorize data.
{{</field>}}
{{<field "Index">}}
Optional: The name of the Splunk index to send data to.
{{</field>}}
{{<field "Enable Data Logging">}}
Enable sending logs (e.g. instrumentation other than metrics) to Splunk.
{{</field>}}
{{<field "Skip TLS Verification">}}
Disable TLS certificate validation for the Splunk endpoint.
{{</field>}}
{{<field "TLS Custom CA">}}
Optional: Provide a custom CA certificate (in PEM format) for TLS validation.
{{</field>}}
{{</fields>}}

#### HTTP Exporter Settings

![img](http-exporter.png)
{{<fields>}}
{{<field "Status">}}
Current exporter status.
{{</field>}}
{{<field "Name">}}
Current exporter name.
{{</field>}}
{{<field "Exporter Type">}}
Configured exporter type.
{{</field>}}
{{<field "Instrumentation Types">}}
Current categories of data exported by Trustgrid (e.g. Metrics).
{{</field>}}
{{<field "Endpoint">}}
Base URL for sending telemetry data via HTTP.
{{</field>}}
{{<field "Metrics Endpoint">}}
Optional override for the metrics-specific HTTP endpoint (e.g. https://hrl/v2/metrics).
{{</field>}}
{{<field "Logs Endpoint">}}
Optional override for the logs-specific HTTP endpoint (e.g. https://hrl/v2/logs).
{{</field>}}
{{<field "HTTP Timeout">}}
Maximum time to wait for an HTTP response (in seconds - Default: 30 seconds).
{{</field>}}
{{<field "Read Buffer Size">}}
Size of the read buffer used when receiving HTTP responses (Default: 0).
{{</field>}}
{{<field "Write Buffer Size">}}
Size of the write buffer used when sending HTTP requests (Default: 512 \* 1024).
{{</field>}}
{{<field "Encoding">}}
Specifies the telemetry data format. Options include JSON or OTEL.
{{</field>}}
{{<field "Compression">}}
Compression method for HTTP payloads Options include gzip, zstd, and none.
{{</field>}}
{{<field "Skip TLS Verification">}}
Disable TLS certificate validation.
{{</field>}}
{{<field "TLS Custom CA">}}
Optional: Provide a custom CA certificate (in PEM format) for TLS validation.
{{</field>}}
{{</fields>}}
