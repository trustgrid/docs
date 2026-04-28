---
title: "Tool reference"
linkTitle: "Tool reference"
weight: 35
description: "Comprehensive reference for every public Trustgrid MCP tool."
---

This page documents every **public top-level MCP tool** currently exposed by the hosted Trustgrid MCP server.

- Sources inspected: `trustgrid/mcp/src/mcp/index.ts`, `trustgrid/mcp/src/mcp/node-tools.ts`, `trustgrid/mcp/src/mcp/scope-map.ts`, and `trustgrid/mcp/index.yaml`
- `switchOrg` is intentionally omitted because it is **stdio-only/internal** and is not part of the public hosted MCP surface
- helper functions available *inside* `codemode` JavaScript (for example `codemode.listNodes(...)`) are not separate MCP tools, so they are not listed as standalone rows here
- Sample interactions below are realistic illustrative snippets, not exhaustive schemas
- Public tool count from the current source: **50 total** (`4` codemode, `36` read, `10` node diagnostics)

{{% alert color="info" %}}
When the API spec declares an exact permission, this page lists it directly. When the spec does **not** declare a narrower per-endpoint permission, the table calls out the public MCP scope that exposes the tool instead of inventing precision that is not present in the source. For reference, the hosted `/mcp/read` scope bundle currently requires `alerts::read`, `audits::read:config`, `audits::read:node`, `domains::read`, `events::read`, `node-vpn::read`, `nodes::read`, `portal::access`, and `virtual-networks::read`.
{{% /alert %}}

## `codemode` tools

These are the four tools registered on `https://mcp.trustgrid.io/mcp/codemode`.

| Tool Name | Brief Description | Permissions Required | Sample Interaction |
| --- | --- | --- | --- |
| `search` | Search the Trustgrid API/function index and Trustgrid docs by keyword. | `/mcp/codemode` scope bundle (`portal::access` plus the codemode read scopes advertised by that endpoint). | Request: `{"query":"vpn routes"}`<br>Response: `{"content":[{"type":"text","text":"## API Functions\n\n- listNodeVpnRoutes ..."}]}` |
| `describe` | Return the full signature and docs for one codemode function. | `/mcp/codemode` scope bundle. | Request: `{"name":"listNodes"}`<br>Response: `{"content":[{"type":"text","text":"# listNodes\n\n**GET** /node ..."}]}` |
| `code` | Execute one async JavaScript function that calls `codemode.*` helpers and returns a shaped result. | `/mcp/codemode` scope bundle **plus** the permission(s) required by the underlying `codemode.*` calls. Example: `nodes::read` for `codemode.listNodes(...)`. | Request: `{"code":"async () => { const nodes = await codemode.listNodes({ limit: 2 }); return nodes.map(n => ({ name: n.name, online: n.online })); }"}`<br>Response: `{"content":[{"type":"text","text":"### Findings\n\n| name | online | ..."}]}` |
| `followUp` | Fetch the next page from a prior paginated `search` or `code` result using its cached `toolCallId`. | `/mcp/codemode` scope bundle. | Request: `{"toolCallId":"9c4d4b7e-..."}`<br>Response: `{"content":[{"type":"text","text":"## API Functions\n\n| Name | Method | ..."}]}` |

## `read` tools

These are the direct MCP tools registered on `https://mcp.trustgrid.io/mcp/read`.

| Tool Name | Brief Description | Permissions Required | Sample Interaction |
| --- | --- | --- | --- |
| `get_cluster` | Get one cluster's detailed config and status. | `/mcp/read` scope bundle; this endpoint is included in the hosted read surface but the operation itself does not declare a narrower per-endpoint permission in the public spec. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io"}`<br>Response: `{"fqdn":"ha-east.example.trustgrid.io","health":"healthy","mode":"manualFailback"}` |
| `get_cluster_vpn_dns` | Get the DNS config for a cluster-attached VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `{"enabled":true,"upstream":[{"ip":"10.20.0.53","port":"53"}]}` |
| `get_cluster_vpn_network` | Get full inventory for one VPN network on a cluster. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `{"name":"corp","ip":"10.30.0.10/24","routes":[...],"services":[...]}` |
| `get_domain` | Get one domain's configuration. | `domains::read` | Request: `{"domainName":"prod"}`<br>Response: `{"name":"prod","gateway":{...},"thresholds":{...}}` |
| `get_network_group` | Get one network group inside a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp","groupName":"branch-sites"}`<br>Response: `{"name":"branch-sites","members":["hq","lab"]}` |
| `get_node` | Get one node by UID, FQDN, or node name. | `nodes::read` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"uid":"6d3...","name":"edge-01","online":true,"config":{...}}` |
| `get_node_by_fqdn` | Get one node directly by FQDN. | `nodes::read` | Request: `{"fqdn":"edge-01.example.trustgrid.io"}`<br>Response: `{"uid":"6d3...","fqdn":"edge-01.example.trustgrid.io","online":true}` |
| `get_node_vpn_dns` | Get the DNS config for a node-attached VPN network. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `{"enabled":true,"upstream":[{"ip":"10.20.0.53"}]}` |
| `get_node_vpn_network` | Get full inventory for one VPN network on a node. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `{"name":"corp","interfaces":[...],"routes":[...],"dns":{...}}` |
| `list_alerts_v2` | List recent org-wide alerts. | `alerts::read` | Request: `{"limit":2}`<br>Response: `[{"eventType":"NODE_DISCONNECTED","node":"edge-01"},{"eventType":"CERT_EXPIRING","node":"edge-02"}]` |
| `list_cluster_vpn_export_routes` | List dynamic export routes for a cluster VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `[{"networkCidr":"10.42.0.0/16","metric":100}]` |
| `list_cluster_vpn_import_routes` | List dynamic import routes for a cluster VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `[{"networkCidr":"172.16.20.0/24","metric":50}]` |
| `list_cluster_vpn_interfaces` | List interfaces attached to a cluster VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `[{"interfaceName":"branch-lan","inDefaultRoute":true}]` |
| `list_cluster_vpn_networks` | List VPN networks attached to a cluster. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io"}`<br>Response: `[{"name":"corp","ip":"10.30.0.10/24"},{"name":"dmz","ip":"10.31.0.10/24"}]` |
| `list_cluster_vpn_routes` | List static routes for a cluster VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `[{"networkCidr":"172.16.0.0/16","path":"vpn"}]` |
| `list_cluster_vpn_services` | List services published on a cluster VPN network. | `/mcp/read` scope bundle; the API spec does not declare a narrower per-endpoint permission here. | Request: `{"clusterFQDN":"ha-east.example.trustgrid.io","networkName":"corp"}`<br>Response: `[{"name":"erp","ip":"10.30.0.50","port":443}]` |
| `list_clusters` | List all clusters in the organization. | `/mcp/read` scope bundle; this endpoint is included in the hosted read surface but the operation itself does not declare a narrower per-endpoint permission in the public spec. | Request: `{}`<br>Response: `[{"fqdn":"ha-east.example.trustgrid.io","health":"healthy"},{"fqdn":"ha-west.example.trustgrid.io","health":"offline"}]` |
| `list_events` | List org-wide events in a time range with filters. | `events::read` | Request: `{"sTime":"2026-04-01T00:00:00.000Z","eTime":"2026-04-02T00:00:00.000Z","limit":2}`<br>Response: `[{"eventType":"NODE_CONNECTED","itemType":"Node"},{"eventType":"CONFIG_CHANGE","itemType":"Node"}]` |
| `list_network_access_policies` | List access policies for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"action":"allow","source":"10.0.0.0/8","dest":"10.30.0.0/24"}]` |
| `list_network_auth_groups` | List auth groups for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"name":"employees","groups":["okta-employees"]}]` |
| `list_network_groups` | List network groups for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"name":"branch-sites"},{"name":"datacenter"}]` |
| `list_network_objects` | List network objects for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"name":"hq","cidr":"10.10.0.0/16"},{"name":"lab","cidr":"10.20.0.0/16"}]` |
| `list_network_port_forwardings` | List port-forwarding rules for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"name":"rdp-jump","port":3389,"target":"10.30.0.25:3389"}]` |
| `list_network_routes` | List routes for a virtual network. | `virtual-networks::read` | Request: `{"domainName":"prod","networkName":"corp"}`<br>Response: `[{"networkCidr":"172.16.0.0/16","path":"vpn","metric":100}]` |
| `list_node_alerts_v2` | List alerts for a specific node. | `alerts::read` | Request: `{"nodeID":"edge-01","limit":2}`<br>Response: `[{"eventType":"NODE_DISCONNECTED"},{"eventType":"NODE_RECONNECTED"}]` |
| `list_node_events` | List events for one node in a time range. | `events::read` | Request: `{"nodeId":"edge-01","sTime":"2026-04-01T00:00:00.000Z","eTime":"2026-04-02T00:00:00.000Z"}`<br>Response: `[{"eventType":"NODE_CONNECTED","level":"INFO"}]` |
| `list_node_vpn_export_routes` | List dynamic export routes for a node VPN network. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `[{"networkCidr":"10.42.0.0/16","metric":100}]` |
| `list_node_vpn_import_routes` | List dynamic import routes for a node VPN network. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `[{"networkCidr":"172.16.20.0/24","metric":50}]` |
| `list_node_vpn_interfaces` | List interfaces attached to a node VPN network. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `[{"interfaceName":"corp-lan","outDefaultRoute":false}]` |
| `list_node_vpn_networks` | List VPN networks attached to a node. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01"}`<br>Response: `[{"name":"corp","ip":"10.30.0.21/24"},{"name":"dmz","ip":"10.31.0.21/24"}]` |
| `list_node_vpn_routes` | List static routes for a node VPN network. | `/mcp/read` scope bundle; only `list_node_vpn_services` declares a narrower `node-vpn::read` permission explicitly in the API spec. | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `[{"networkCidr":"172.16.0.0/16","path":"vpn"}]` |
| `list_node_vpn_services` | List services published on a node VPN network. | `node-vpn::read` | Request: `{"nodeID":"edge-01","networkName":"corp"}`<br>Response: `[{"name":"rdp","ip":"10.30.0.25","port":3389}]` |
| `list_nodes` | List nodes, with optional filtering and MCP-side projection helpers for `config`, `shadow`, and selected top-level fields. | `nodes::read` | Request: `{"limit":2,"name":true,"online":true}`<br>Response: `[{"name":"edge-01","online":true},{"name":"edge-02","online":false}]` |
| `list_virtual_networks` | List virtual networks in a domain. | `virtual-networks::read` | Request: `{"domainName":"prod"}`<br>Response: `[{"name":"corp","networkCidr":"10.30.0.0/24"},{"name":"dmz","networkCidr":"10.31.0.0/24"}]` |
| `tail_config_audit` | Tail real-time configuration audit entries. | `audits::read:config` | Request: `{"limit":2}`<br>Response: `[{"actor":"alice@example.com","action":"update","itemType":"Node"}]` |
| `tail_node_audit` | Tail real-time node audit entries. | `audits::read:node` | Request: `{"limit":2}`<br>Response: `[{"node":"edge-01","event":"service restarted"}]` |

## `tools` node diagnostics

These are the live node diagnostic tools registered on `https://mcp.trustgrid.io/mcp/tools`.

| Tool Name | Brief Description | Permissions Required | Sample Interaction |
| --- | --- | --- | --- |
| `tg_node_runtime_status` | Return live runtime/service status from the node itself. | `nodes::read`, `nodes::service:status`, `portal::access` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"services":{"trustgrid-node":"active"},"memory":{"usedPercent":61}}` |
| `tg_node_network_status` | Return live interface, link-state, and address info from the node. | `nodes::read`, `nodes::service:network-status`, `portal::access` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"interfaces":[{"name":"eth0","state":"UP","addresses":["10.0.1.10/24"]}]}` |
| `tg_node_dataplane_status` | Return current gateway-route/dataplane status from the node. | `nodes::read`, `nodes::service:gateway-routes`, `portal::access` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"routes":[{"gateway":"corp-gw","connected":true}]}` |
| `tg_node_errors_status` | Return current node errors, optionally only startup errors. | `nodes::read`, `nodes::service:errors`, `portal::access` | Request: `{"nodeID":"edge-01","startupOnly":true}`<br>Response: `{"errors":[]}` |
| `tg_node_tcp_test` | Test outbound TCP connectivity from a node to `host:port`. | `nodes::read`, `nodes::service:outbound-connection-check`, `portal::access` | Request: `{"nodeID":"edge-01","remoteAddress":"example.com:443"}`<br>Response: `{"success":true,"remoteAddress":"example.com:443"}` |
| `tg_node_repo_test` | Test whether the node can reach its configured package repositories. | `nodes::read`, `nodes::service:repo-connectivity`, `portal::access` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"repositories":[{"url":"https://repo.trustgrid.io","success":true}]}` |
| `tg_node_dns_test` | Test DNS resolution from the node against its configured DNS servers. | `nodes::read`, `nodes::service:dns-health`, `portal::access` | Request: `{"nodeID":"edge-01"}`<br>Response: `{"servers":[{"ip":"10.20.0.53","success":true,"latencyMs":12}]}` |
| `tg_node_gateway_latency_test` | Run a per-hop latency trace from the node to a named gateway. | `nodes::read`, `nodes::service:gateway-latency`, `portal::access` | Request: `{"nodeID":"edge-01","gateway":"corp-gw"}`<br>Response: `{"hops":[{"hop":1,"latencyMs":4.2},{"hop":2,"latencyMs":8.7}]}` |
| `tg_test_traffic` | Simulate a TCP packet through Trustgrid policy/routing logic and show where it would be allowed or blocked. | `nodes::read`, `nodes::service:packet-sim`, `portal::access` | Request: `{"nodeID":"edge-01","sourceIp":"10.10.1.25","destIp":"10.30.0.50","destPort":443}`<br>Response: `{"allowed":true,"matchedPolicy":"allow-hq-to-erp"}` |
| `tg_node_sniff_traffic` | Run a bounded packet capture on one node interface. | `nodes::read`, `nodes::service:tg-tcpdump`, `portal::access` | Request: `{"nodeID":"edge-01","iface":"eth0","filter":"tcp port 443","packetCount":5}`<br>Response: `{"output":"IP 10.10.1.25.51234 > 10.30.0.50.443: Flags [S] ..."}` |
