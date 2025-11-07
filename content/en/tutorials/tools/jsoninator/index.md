---
title: JSONinator Walk Through
linkTitle: JSONinator
description: A step-by-step guide to using JSONinator for bulk configuration changes on Trustgrid nodes.
---

{{<alert color="warning">}}
**Notice:** JSONinator is provided as a convenience tool for Trustgrid customers. It is not part of the core Trustgrid platform and is not covered by Trustgrid’s standard service level agreements (SLAs) or product guarantees. Support for this tool is provided on a best-effort basis. Report any issues via the GitHub repository issues page.
{{</alert>}}

[JSONinator](https://github.com/trustgrid/jsoninator) is a command-line tool designed to make bulk changes to Trustgrid nodes matching a configured filter. This guide walks through its usage, safety considerations, and best practices.

## What is JSONinator?

JSONinator allows you to automate configuration changes across multiple nodes. It is especially useful for:
 - Applying the same configuration to many nodes
 - Ensuring consistency across environments
 - Reducing manual effort and errors

## Best Practices

- Always start with a dry-run and review the output.
- Test changes on a small group of nodes before scaling up.
- Keep your API keys secure and never share them.
- Save and version-control your plan files for audit and repeatability.

> **Warning:**
> - Always run JSONinator in dry-run mode first and thoroughly review the reports generated.
> - Apply changes to small groups of nodes and verify results before running against your entire fleet.

## Prerequisites

- Download the latest JSONinator release from [GitHub Releases](https://github.com/trustgrid/jsoninator/releases).
- The executable is **unsigned**. On Windows and macOS, you may need to allow the app to run:
  - **Windows:** Right-click the executable, select **Properties**, and check **Unblock** if present.
  - **macOS:** Use `xattr -d com.apple.quarantine ./jsoninator` in Terminal, or allow via **System Preferences > Security & Privacy** after the first launch attempt.
- Obtain Trustgrid API keys. See [API Access]({{<relref "docs/user-management/api-access/" >}}) for instructions.

## How JSONinator Works

JSONinator operates in three main stages defined in your plan file:

1. [**Input**](#input): The `input` section queries the Trustgrid API to fetch the current configuration data for entities, like nodes or clusters, that you want to modify. This typically retrieves a list of nodes or clusters and their existing settings, which become the starting point for your changes.

2. [**Pipeline**](#pipeline): The `pipeline` section processes each node or cluster's configuration using a series of processors:
    - <a href="https://github.com/trustgrid/jsoninator/blob/main/README.md#filter" target="_blank" rel="noopener"><strong>filter</strong></a>: Selects which nodes or clusters to include or exclude based on criteria (such as name prefix, suffix, or custom Go template queries).
    - <a href="https://github.com/trustgrid/jsoninator/blob/main/README.md#map" target="_blank" rel="noopener"><strong>map</strong></a>: Focuses on a nested field or sub-object within each node or cluster for further processing.
    - <a href="https://github.com/trustgrid/jsoninator/blob/main/README.md#transform" target="_blank" rel="noopener"><strong>transform</strong></a>: Modifies, adds, or removes fields in the configuration using templates (for example, enabling UDP or setting a default port).
    - <a href="https://github.com/trustgrid/jsoninator/blob/main/README.md#replace" target="_blank" rel="noopener"><strong>replace</strong></a>: Emits only the specified fields, allowing you to reshape the output as needed.
    This pipeline creates the new, intended configuration for each node or cluster.

3. [**Output**](#output): The `output` section defines where to send the processed (new) configuration. For each node or cluster, JSONinator posts the updated configuration to the specified HTTP endpoint (usually the Trustgrid API), using the method and headers you provide.

All actions and results are logged to report files in the "reports" subdirectory of your working directory.

By default, JSONinator runs in **dry-run** mode. No changes are made to nodes or clusters; instead, a report is generated showing what would change. To actually make changes, run with `-dryrun=false`.

## Plan YAML File Structure

A plan file defines what changes to make, which data to operate on, and where to send results. The main sections are:

- `input`: Where to fetch the data from (e.g., HTTP endpoint)
- `pipeline`: A list of processors to filter, map, or transform the data
- `output`: Where to send the processed data (e.g., HTTP endpoint)
- `report`: Output settings for generated reports

### Example plan file

{{<codeblock lang="yaml">}}
input:
  http:
    url: https://portal.trustgrid.io/api/node?projection[0]=uid&projection[1]=name&projection[2]=tags&projection[3][0]=config&projection[3][1]=gateway&projection[4]=fqdn
    headers:
      Authorization: "trustgrid-token ${TRUSTGRID_API_KEY_ID}:${TRUSTGRID_API_KEY_SECRET}"
      Accept: application/json

pipeline:
  processors:
    - filter:
        prefix:
          name: gw-
        query: |
          {{if (eq .type "Node")}}true{{else}}false{{end}}
    - map:
        field: config.gateway
    - transform:
        fields:
          udpEnabled: true
          udpPort: |
            {{if .udpPort}}{{.udpPort}}{{else if .port}}{{.port}}{{else}}8995{{end}}
          maxClientWriteMbps: |
            {{if or (eq .maxClientWriteMbps 0.0) (not .maxClientWriteMbps)}}nil{{else}}{{.maxClientWriteMbps}}{{end}}
          cert: |
            {{if .cert}}{{.cert}}{{else}}nil{{end}}

output:
  http:
    url: https://portal.trustgrid.io/api/node/{{.uid}}/config/gateway
    method: PUT
    status_codes: [200]
    headers:
      Authorization: "trustgrid-token ${TRUSTGRID_API_KEY_ID}:${TRUSTGRID_API_KEY_SECRET}"
      Content-Type: application/json

report:
  path: ./reports
{{</codeblock>}}

#### input
- **http**: Fetches data from an HTTP endpoint. You can specify `url` and optional `headers`.

#### pipeline
- **processors**: A list of steps to filter, map, or transform the data. Common processors include:
  - `filter`: Exclude items by prefix, suffix, or query (Go template)
  - `map`: Select a nested field for further processing
  - `transform`: Change or remove fields using templates
  - `replace`: Emit only specified fields using templates

#### output
- **http**: Sends each processed item to an HTTP endpoint. Specify `url`, `method`, `headers`, and allowed `status_codes`.

#### report
- **path**: Directory for generated report files (e.g., `changes.csv`, `filtered.csv`, `noops.csv`).

## Building a Plan YAML File

1. Identify the settings you want to change. Use the <a href="https://docs.trustgrid.io/api/" target="_blank" rel="noopener">Trustgrid API docs</a> to find field names and endpoints.
2. For practical discovery, use your browser's DevTools (<a href="https://developer.chrome.com/docs/devtools/open/" target="_blank" rel="noopener">Chrome DevTools instructions</a>) while editing a node in the Trustgrid portal. Watch network requests to see which fields are updated and their names.
3. If you want to filter by fields like `tags` or `name`, add those columns to the Nodes table in the portal and observe their API usage in DevTools.
4. Once you know the fields you wish to modify and/or filter with build the plan's input `url` to return those fields for processing. See the example plan and the <a href="https://github.com/trustgrid/jsoninator/blob/main/README.md" target="_blank" rel="noopener">JSONinator README</a> for details.

Building a plan file may require some iteration. Start simple, test in dry-run mode, utilize a test/dev accounts if available, start with a single node/cluster and expand as needed. Trustgrid can assist with plan file creation on a best-effort basis during our business hours; contact your Trustgrid support for help.

## Running JSONinator

Before running JSONinator, you must provide your Trustgrid API key ID and SECRET. The most secure way is to use your organization's existing secrets manager or environment variable management tool. If that is not available, set them as environment variables:

**macOS/Linux:**
```sh
export TRUSTGRID_API_KEY_ID="your_key_id_here"
export TRUSTGRID_API_KEY_SECRET="your_secret_here"
./jsoninator -plan=my-plan.yaml
# When finished, unset the variables:
unset TRUSTGRID_API_KEY_ID
unset TRUSTGRID_API_KEY_SECRET
```

**Windows PowerShell:**
```powershell
$env:TRUSTGRID_API_KEY_ID="your_key_id_here"
$env:TRUSTGRID_API_KEY_SECRET="your_secret_here"
./jsoninator -plan=my-plan.yaml
# When finished, remove the variables:
Remove-Item Env:TRUSTGRID_API_KEY_ID
Remove-Item Env:TRUSTGRID_API_KEY_SECRET
```

**Windows Command Prompt:**
```cmd
set TRUSTGRID_API_KEY_ID=your_key_id_here
set TRUSTGRID_API_KEY_SECRET=your_secret_here
jsoninator -plan=my-plan.yaml
REM When finished, unset the variables:
set TRUSTGRID_API_KEY_ID=
set TRUSTGRID_API_KEY_SECRET=
```

{{<alert color="warning">}}
**Security Reminder:**

Using `export` (macOS/Linux) or `set` (Windows Command Prompt) does not record your secrets in shell history. For best security, use your organization's secrets manager or environment variable management tool. Avoid putting secrets directly in your plan file or command line.
{{</alert>}}

1. Place your plan YAML file in a working directory.
2. Run JSONinator in dry-run mode (default):
  ```sh
  ./jsoninator -plan=my-plan.yaml
  ```
3. [Review the generated report files](#reviewing-reports) in the specified report path.
4. When satisfied, run with `-dryrun=false` to make changes:
  ```sh
  ./jsoninator -plan=my-plan.yaml -dryrun=false
  ```

## Reviewing Reports

After each run, review the report files in the `reports` directory to see which nodes would be (or were) changed. JSONinator creates a datestamped folder with three main files, each relating to nodes processed:

- **changes.csv**: Lists each node and field that was changed, showing the node name, field, previous value, and new value. For example:

  ```csv
  name,field,before,after
  mctest-edge3-2204.mctest.trustgrid.io,udpPort,<nil>,8443
  mctest-edge3-2204.mctest.trustgrid.io,maxClientWriteMbps,0,<no value>
  mctest-edge3-2204.mctest.trustgrid.io,udpEnabled,false,true
  ...
  ```

- **filtered.csv**: Lists nodes that were excluded by filters, with the node name and the filter reason. For example:

  ```csv
  name,filter
  aws-agent1.mctest.trustgrid.io,"field ""name"" does not have prefix ""gw-"""
  aws-agent2.mctest.trustgrid.io,"field ""name"" does not have prefix ""gw-"""
  ...
  ```

- **noops.csv**: Lists nodes that passed through the pipeline but had no changes made. For example:

  ```csv
  name
  mctest-esxedge5.mctest.trustgrid.io
  mctest-esxedge6.mctest.trustgrid.io
  ```

These reports provide details on actions taken, nodes affected, and any errors encountered. Reviewing them helps ensure your plan is working as intended before and after applying changes.

