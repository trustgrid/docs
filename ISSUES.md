# Container Overhaul — Gaps & Issues Found

Tracker for things discovered while writing/validating the container docs. Nothing here is filed yet — review and decide what becomes a ticket, what becomes a doc, what gets fixed in code.

## Format

- **Where:** path or system
- **What:** what's wrong / missing
- **Evidence:** what I saw / reproduced
- **Suggested action:** fix in code, file ticket, doc-only, ask Nate

---

## 1. trustgrid-test MCP `search` omits mutation endpoints

- **Where:** `mcp.test.trustgrid.io` (codemode MCP). Likely same for prod.
- **What:** `search` only returns `GET` endpoints — POST/PUT/DELETE for the same resource family are filtered out. Search for "create container", "post container", or "add container" never surfaces `PostV2NodeExecContainer`, even though the prod MCP search does return it.
- **Evidence:** Ran `search(query: "post container")` and `search(query: "create container")` against `trustgrid-test` — only GETs returned. The exact same query against `trustgrid-prod` (different org, same API surface) returns `PostV2NodeExecContainer`, `DeleteV2NodeExecContainerByContainerid`, `PutV2NodeExecContainerByContainerid`. Also confirmed `describe("PostV2NodeExecContainer")` returns "No function named ... available" on test, but works on prod.
- **Suggested action:** investigate why the test MCP's tool inventory drops mutations. May be a permissions filter (token scope?) or a deployment lag. Either way the agent UX is broken — you can read but not write, and you have to know the function name from elsewhere to call it.

## 9. Cluster-scope container does not appear in node-config

- **Where:** `POST /v2/cluster/{fqdn}/exec/container` then `POST /node/{id}/trigger/node-config`.
- **What:** Created `docs-nginx` at cluster scope (cluster `nate-edge-cls-0427-2048`, id `740fd2d7-...`). Cluster-scope GET returns the record. Node `node-config` trigger fires within seconds and still does not include the cluster container; it shows only the 3 node-scope records.
- **Suggested action:** confirm whether cluster→node config sync is async with measurable lag, broken on this node build, or requires a cluster-membership state that's missing (recall the 2nd member edge2's `cluster=None`).

## 8. PUT container config appears to append port mappings, not replace

- **Where:** `PUT /v2/node/.../exec/container/{id}/config` with `portMappings: [...]`.
- **What:** After two consecutive PUTs with the same single-entry `portMappings` array, the resulting `node-config` shows two identical entries. This suggests append-merge semantics. The API doc/signature implies set/replace.
- **Evidence:** node-config command[0] shows `portMappings: [ {…8080→80…}, {…8080→80…} ]`.
- **Suggested action:** decide on set-vs-merge semantics and document; today it's accidentally additive.

## 7. POST container creates duplicates by `name`

- **Where:** `POST /v2/node/.../exec/container`
- **What:** Three POSTs with `name: "docs-nginx"` all returned `{ok: true}` and yielded **three distinct container records** with three different UIDs but the same `name`. The portal-side affordance is by name, so duplicates are confusing.
- **Suggested action:** either 409 on name collision or upsert. Either is fine; today's behavior is the worst case.

## 6. All 4 nodes in qa-E2E-20260427204827 report `startup.error: true`

- **Where:** test-tenant nodes nate-gw1/2-0427-2048, nate-edge1/2-0427-2048 (the live labinator env).
- **What:** `shadow.reported.startup.error: true` on every node. Labinator's `check_connectivity` warns on all four. Containers configured via API don't actually run — I created a `docs-nginx` service container, port mapping applied, config visible in `node-config` trigger (versionCounter=2, exec cmd enabled=true, repo+tag valid since image was pushed mid-debug). Container still never started. From the edge test host, `nc 192.168.100.209 8080` → connection refused even after 60s.
- **Evidence:** described above; tested both with `latest` (tag absent → ok, never starts) and `alpine` (tag present after push → still never starts). `container-status` trigger returns `{}` and `container-logs` trigger returns `"Service local:container-logs has not registered an event handler"`.
- **Suggested action:** the env this came up on is a labinator-deployed L4-connector environment, not vanilla. Either (a) startup.error is benign and the container runtime is gated on something we're missing, or (b) it's actually broken on this node build. Worth poking at separately — either way, "look at startup.error" should be in the troubleshooting doc.

## 5. DELETE container schema rejects own response shape

- **Where:** `DELETE /v2/node/{nodeID}/exec/container/{containerID}`
- **What:** Returns `{ "error": ["commands[0] object contains unknown properties: keyProp"] }` — the delete handler appears to round-trip the container record through its own schema validator and reject an internal field (`keyProp`). The original container record is left in place; subsequent POST returns the **same** container ID (collision-by-name with the dead record).
- **Evidence:** Reproduced on `nate-edge1-0427-2048` container `31765b8e-e684-44a5-b90e-08a0bdb990fb`. DELETE → 4xx with the message above. List GET after delete still shows the record. POST of an identical-name container returns `ok` but the resulting record has the same UID, suggesting upsert-by-name semantics.
- **Suggested action:** fix the DELETE handler to skip its current schema check (or drop `keyProp` from internal records before validating). Also: POST should probably 409 on name collision rather than silently upserting.

## 4. Container event names not documented anywhere reachable

- **Where:** `POST /node/{nodeID}/trigger/{event}` for container-scoped events.
- **What:** The MCP server documents many node trigger events (node-config, vpn-routes, tg-ping, flows, etc.) but nothing container-specific. By probing I confirmed `container-logs` is a known event name (it returns `Service local:container-logs has not registered an event handler` when the service isn't running). `container-list`, `container-status`, `image-list`, `image-pull-status` etc. all hang or 404 with no help text. Whatever events the portal "container terminal" / "container logs" buttons use are undocumented and unsearchable from the API spec.
- **Suggested action:** add codemode wrappers like `test_container_status`, `test_container_logs`, `test_image_list` mirroring the existing `test_*` diagnostics. Document them as the canonical way to inspect container state.

## 3. Container port mapping: schema says `string`, API requires `number`

- **Where:** `PUT /v2/node/{nodeID}/exec/container/{containerID}/config` (port mappings) — and likely the cluster equivalent.
- **What:** codemode signature declares `hostPort?: string; containerPort?: string;` but submitting strings returns `{ "error": ["host port is required"] }`. Submitting integers (`hostPort: 8080`) works and the stored record echoes back integers.
- **Evidence:** Reproduced via curl PUT on `nate-edge1-0427-2048`. String → 400. Integer → 200 OK, GET returns `"hostPort": 8080`.
- **Suggested action:** fix the OpenAPI/codemode types to `number`, or accept string and coerce. Either way the type lies — and this will trip up SDK users and terraform generation.

## 2. ListNodes / GetCluster don't surface online/connection status

- **Where:** API (`GET /node`, `GET /cluster/{fqdn}`)
- **What:** No obvious way to filter "online nodes" or even know whether a node is currently connected without poking `POST /node/{id}/trigger/node-config?wait=1` and waiting for a timeout. `status`, `online`, `connection` fields all return null/absent on the list response. Cluster detail returns `nodes: []`.
- **Evidence:** `/node` and `/node?projection[][]=["status"]` both return null for status fields on all 19 nodes. Discovery of which nodes are reachable required hitting `node-config` per node with an 8s timeout.
- **Suggested action:** confirm whether there's a documented field/endpoint for connection state. If not, this is a gap — likely needs `online: bool` and `lastSeen` on the list response, or a dedicated endpoint. Worth a docs note either way ("how do I know if a node is online?").

