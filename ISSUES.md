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

## 11. Registry only indexes amd64 / Docker schema 2 pushes — arm64 and OCI manifests silently fail (DOCS GAP, not a bug)

**Status: expected behavior, not a product bug. Needs to be documented in `/docs/repositories/`.**

- **What happens:** `docker push` from Apple Silicon (M-series) produces an OCI image manifest (`application/vnd.oci.image.manifest.v1+json`), typically arm64-only. The push completes successfully — HTTP 201, all blobs uploaded, digest returned. The image lands in the underlying registry storage and is retrievable by digest via `crane`. But the org's **tag index** never registers it, so:
  - Portal "Repositories → <name>" tag table shows the tag as missing.
  - Trustgrid API `/repositories/<namespace>/<image>` returns the unindexed tag as absent.
  - Node pulls from `repo.<env>.trustgrid.io/v2/.../manifests/<tag>` get a blank HTTP 401 (no `WWW-Authenticate` header), `RegistryMediator: Retries exhausted` in the node log, container stays `STOPPED`.
- **What works:** `docker push` from an amd64 Linux host produces `application/vnd.docker.distribution.manifest.v2+json` with `architecture: amd64`. The tag indexes correctly; portal lists it; node pulls and runs the container.
- **Reproduction:**
  - Pushed `nginx:amd64-test` from an amd64 Ubuntu test host (labinator `docs-edge-host-0512-2340`, x86_64). Manifest mediaType `vnd.docker.distribution.manifest.v2+json`, arch `amd64`. Portal nginx page lists `amd64-test` with digest. Container deploys → `RUNNING`, `curl http://node:8082/` → HTTP 200, 896 bytes.
  - Pushed `nginx:alpine-docker` from colima on Apple Silicon. Manifest mediaType `vnd.oci.image.manifest.v1+json`, arch `arm64`. Portal nginx page does NOT list `alpine-docker`. Container deploys → `STOPPED`, `RegistryMediator: Retries exhausted`.
  - Same node, same code path, only variable is the manifest format / arch.
- **Real product asks (smaller than the original "this is a bug" framing):**
  1. **Customer-facing error on push.** `docker push` of an unsupported manifest currently succeeds silently and leaves the customer with an "invisible" image. Either reject the push at the registry layer with a clear `MANIFEST_INVALID` error explaining the requirement, or accept and translate. Silent acceptance is the worst UX.
  2. **Node logging.** `RegistryMediator.getImageManifest` wraps the underlying HTTP error as `java.lang.Exception("Retries exhausted")`. Even now that we know the cause, the next instance of any registry failure will look identical. Surface the actual HTTP status + body + `WWW-Authenticate` header.
- **Docs fix (already wired up on this branch):**
  - `/docs/repositories/` callout: pushes must be linux/amd64 Docker schema 2. From Apple Silicon use `docker buildx build --platform linux/amd64 --push` or push from an amd64 host.
  - Quickstart tutorial: same callout above the `docker push` step.
  - Troubleshooting `RegistryMediator: Retries exhausted` entry: first thing to check is whether the image was pushed from arm64 or as multi-arch OCI.

## (Original #11) `docker.<env>.trustgrid.io` (documented push target) and `repo.<env>.trustgrid.io` (actual node pull target) are two different services

- **Where:** Trustgrid container registry / per-org image flow. Reproduced end-to-end on test tenant.
- **Customer-facing push target** (per `/docs/repositories/`): `docker.<env>.trustgrid.io`. `docker login -u trustgrid -p $JWT` against this host succeeds; JWT has `aud: docker.test.trustgrid.io`, `scopes: docker:*`. `docker push docker.test.trustgrid.io/nate.test.trustgrid.io/nginx:alpine-docker` completes successfully; image is visible via authenticated reads.
- **Node-facing pull target** (per node config): on the deployed node `docs-edge1-0512-2236`, `/var/lib/trustgrid/config/node-profile.json` and `node-info.json` both contain `"repo.uri": "https://repo.test.trustgrid.io/"`. Container deploy logs `Retries exhausted downloading https://repo.test.trustgrid.io/v2/nate.test.trustgrid.io/nginx/manifests/<tag>` (`/var/log/trustgrid/tg-default.log`). Same failure for both a multi-arch OCI index push (via crane) and a single-platform Docker schema 2 push (via docker push) — so it's not a manifest format issue.
- **Evidence they are different services:**
  - **IPs:** `docker.test` → `3.214.92.2`, `35.174.58.241`; `repo.test` → `18.205.177.116`.
  - **TLS chains:** `docker.test` → wildcard `*.test.trustgrid.io`, Amazon RSA 2048 M02 (public CA); `repo.test` → `CN=repo.test.trustgrid.io`, TrustGrid Intermediate CA (internal).
  - **Auth challenge on `/v2/`:** `docker.test` returns `401 WWW-Authenticate: Basic realm="https://docker.test.trustgrid.io",service="docker.test.trustgrid.io"`. `repo.test` returns `401` with **no `WWW-Authenticate` header at all** — confirmed from inside the node via `curl -v https://repo.test.trustgrid.io/v2/nate.test.trustgrid.io/nginx/manifests/alpine-docker` (TLS handshake succeeds, server responds `HTTP/2 401`, `content-length: 0`, no auth challenge headers). A spec-compliant Docker registry returns `WWW-Authenticate: Bearer realm="..."` so clients can fetch a token; this one does not. The node's Java HTTP client retries blindly and gives up — hence `RegistryMediator:111 - Retries exhausted` with a generic `java.lang.Exception` and no underlying cause.
- **What this means:** unless the two services share backing storage (engineering needs to confirm), images customers push to `docker.<env>.trustgrid.io` per the docs **never become reachable from nodes**, because nodes are configured to pull from a different host. Even if they do share storage, the missing `WWW-Authenticate` header on `repo.<env>.trustgrid.io` is a real registry-config bug — the node has no way to authenticate, and there are no in-band hints (challenge header, status text, or response body) that explain the 401.
- **Suggested action:** (1) confirm whether `docker.<env>` and `repo.<env>` share backing storage. If yes, fix the `WWW-Authenticate` response on `repo.<env>` so registry clients (including the node's) can complete the auth handshake. If no, that is a major silent failure on customer image pushes — fix registry routing so node pulls land on the same backend as customer pushes. (2) regardless of the routing answer, the node's `RegistryMediator` / `HttpProcessor.downloadJson` swallows the underlying HTTP failure (status code, response body, registry error code) and rethrows as `java.lang.Exception("Retries exhausted")`. Add structured logging — the next time this bites, "401 from repo.test with no challenge" should appear directly in the log, not require an evening of probing to derive. (3) Document `repo.<env>.trustgrid.io` in `/docs/repositories/` as the node-side pull host, even if customers never interact with it directly — it's surfaced in `node-profile.json` and in support / troubleshooting situations.
- **Definitive proof — tag indexing is the bug, not registry routing:** Deployed four containers on the same node, same code path, varying only the image tag.
  - `:alpine` (pushed today via crane) → `STOPPED`, `RegistryMediator: Retries exhausted`
  - `:alpine-docker` (pushed today via `docker push` per docs) → `STOPPED`, same exception
  - `:alpine` again via portal Add Container modal → `STOPPED`, same exception
  - `:latest` (pushed by Nate in a prior session, **present in the portal Repositories tag list**) → **`RUNNING`**, `curl http://node:8081/` → HTTP 200, 615 bytes (nginx welcome page).
  Identical node, identical config, identical pull host. The only variable is whether the tag is registered in the org's repo index. The pre-existing indexed tag pulls fine; today's pushes don't. `crane ls` against `docker.test.<env>` confirms all three tags exist in the underlying registry storage — they just never got indexed at the portal/API/auth layer.
- **`/repositories` lists the repo, `/repositories/<name>` does NOT list its tags:** after pushing the new image `hello-world:v1`, `GET /repositories` includes `nate.test.trustgrid.io/hello-world` but `GET /repositories/nate.test.trustgrid.io/hello-world` returns `[]`. So the push pipeline updates the repo-level index but skips tag-level indexing entirely.
- **Verified portal vs API parity:** I created the same container three times (twice via API, once via the portal Add Container modal) on the upgraded single node. All three end in identical `Retries exhausted` `RegistryMediator` exceptions and `STOPPED` state. So this is not an API-only problem — customers using the portal hit the exact same wall.
- **Cross-reference:** prior bug WOR-9984 ("container repos: URI with multiple directories is broken", fixed in n-2.23.0, released 2025-08-25) produced the *exact same* Java stack trace text for a structurally different root cause (URL path-component bug for image names with extra `/`). The generic exception means any of these unrelated failures look identical from the log. Fix #2 above to make future debugging tractable.

## 10. Typo in node software hostname: `setup.trustrid.io`

- **Where:** Node-side Java log line: `ApiGatewayContext:209 - Using https host setup.trustrid.io`. Observed on `docs-edge1-0512-2236` (uid `3af13d4c-aadb-4936-93b1-334ff68a0a40`) running `1.5.20260512-2447` after fresh deploy + upgrade.
- **What:** The hostname `setup.trustrid.io` is missing the `g` — should presumably be `setup.trustgrid.io`. If the node software is actually attempting to resolve this typo'd hostname, the lookup will fail (and very likely *does* fail, because no `trustrid.io` registration is expected to mirror trustgrid's `setup` host).
- **Evidence:** Coincides on the same node with intermittent `Repo Connectivity failed` events and `CRITICAL Connection to Gateway=... Unauthorized` events around restart time. Repo connectivity recovered after the next reconnect, suggesting the misspelled call is non-fatal but contributes to startup noise and possibly to slow / blocked container image pulls.
- **Suggested action:** grep the node codebase for `trustrid` (note missing g) and fix the literal. May be a single character typo with surprising blast radius.

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

