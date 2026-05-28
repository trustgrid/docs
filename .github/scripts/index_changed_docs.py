#!/usr/bin/env python3
"""Index changed markdown docs into Cloudflare Vectorize for the AI chatbot.

Reads a git diff --name-status output file, processes each changed/added/deleted
markdown file under content/en/, generates embeddings via Cloudflare Workers AI,
and upserts/deletes vectors in the tgcodemode-docs Vectorize index.
"""

import base64
import json
import os
import re
import sys
import urllib.request
from collections.abc import Iterable
from pathlib import Path

INDEX_NAME = "tgcodemode-docs"
EMBED_MODEL = "@cf/baai/bge-base-en-v1.5"
EMBED_BATCH = 50  # chunks per embedding call
UPSERT_BATCH = 50  # vectors per upsert call
DELETE_BATCH = 100  # ids per delete call
MAX_CHUNKS = 200  # generous upper bound on chunks per file
MAX_CHUNK_CHARS = 1800  # ~450 tokens at ~4 chars/token (bge-base-en is 512 max)


def cf_api(path: str, data: dict | None = None) -> dict:
    account_id = os.environ["CLOUDFLARE_ACCOUNT_ID"]
    api_token = os.environ["CLOUDFLARE_API_TOKEN"]
    url = f"https://api.cloudflare.com/client/v4/accounts/{account_id}{path}"
    body = json.dumps(data).encode() if data is not None else None
    req = urllib.request.Request(url, data=body, method="POST" if body else "GET")
    req.add_header("Authorization", f"Bearer {api_token}")
    if body:
        req.add_header("Content-Type", "application/json")
    with urllib.request.urlopen(req, timeout=120) as resp:
        return json.loads(resp.read())


def safe_id_prefix(filepath: str) -> str:
    """Create a URL-safe prefix for vector IDs from a file path."""
    # Strip content/en/ prefix for brevity
    rel = filepath.removeprefix("content/en/")
    # base64url encode so we avoid / and : issues
    return base64.urlsafe_b64encode(rel.encode()).decode().rstrip("=")


def strip_frontmatter(text: str) -> str:
    if text.startswith("---"):
        parts = text.split("---", 2)
        if len(parts) >= 3:
            return parts[2].strip()
    return text.strip()


def chunk_markdown(filepath: str, text: str) -> list[dict]:
    text = strip_frontmatter(text)
    if not text:
        return []

    # Extract title from first H1
    title = ""
    m = re.search(r"^#\s+(.+)$", text, re.MULTILINE)
    if m:
        title = m.group(1).strip()

    # Split on H2 sections
    sections = re.split(r"\n(?=##\s+)", text)
    chunks = []
    for sec in sections:
        sec = sec.strip()
        if not sec:
            continue
        # If a section is still too long, further split by paragraphs
        if len(sec) > MAX_CHUNK_CHARS:
            paras = sec.split("\n\n")
            buf = ""
            for para in paras:
                if len(buf) + len(para) + 2 > MAX_CHUNK_CHARS and buf:
                    chunks.append(buf.strip())
                    buf = para
                else:
                    buf = (buf + "\n\n" + para).strip() if buf else para
            if buf:
                chunks.append(buf.strip())
        else:
            chunks.append(sec)

    prefix = safe_id_prefix(filepath)
    result = []
    for idx, content in enumerate(chunks):
        # Build a human-readable excerpt for metadata
        excerpt = content.replace("\n", " ").strip()
        if len(excerpt) > 300:
            excerpt = excerpt[:297] + "..."
        result.append(
            {
                "id": f"{prefix}:{idx}",
                "text": content,
                "metadata": {
                    "path": filepath,
                    "title": title,
                    "excerpt": excerpt,
                    "chunk": idx,
                },
            }
        )
    return result


def embed(texts: list[str]) -> list[list[float]]:
    resp = cf_api(f"/ai/run/{EMBED_MODEL}", {"text": texts})
    if not resp.get("success"):
        raise RuntimeError(f"Embedding failed: {resp}")
    return resp["result"]["data"]


def delete_vectors(ids: list[str]) -> None:
    if not ids:
        return
    for i in range(0, len(ids), DELETE_BATCH):
        batch = ids[i : i + DELETE_BATCH]
        resp = cf_api(
            f"/vectorize/v2/indexes/{INDEX_NAME}/delete_by_ids", {"ids": batch}
        )
        if not resp.get("success"):
            raise RuntimeError(f"Delete failed: {resp}")
        print(f"  deleted {len(batch)} vectors")


def upsert_vectors(vectors: list[dict]) -> None:
    if not vectors:
        return
    for i in range(0, len(vectors), UPSERT_BATCH):
        batch = vectors[i : i + UPSERT_BATCH]
        resp = cf_api(f"/vectorize/v2/indexes/{INDEX_NAME}/upsert", {"vectors": batch})
        if not resp.get("success"):
            raise RuntimeError(f"Upsert failed: {resp}")
        print(f"  upserted {len(batch)} vectors")


def process_file(filepath: str, status: str) -> None:
    prefix = safe_id_prefix(filepath)
    # Delete old chunk IDs for this file (0..MAX_CHUNKS-1)
    old_ids = [f"{prefix}:{i}" for i in range(MAX_CHUNKS)]
    print(f"{status} {filepath}: cleaning old vectors")
    delete_vectors(old_ids)

    if status == "D":
        return

    content = Path(filepath).read_text(encoding="utf-8")
    chunks = chunk_markdown(filepath, content)
    if not chunks:
        print(f"  no chunks generated")
        return

    # Generate embeddings in batches
    all_embeddings = []
    texts = [c["text"] for c in chunks]
    for i in range(0, len(texts), EMBED_BATCH):
        batch = texts[i : i + EMBED_BATCH]
        embs = embed(batch)
        all_embeddings.extend(embs)
        print(
            f"  embedded batch {i // EMBED_BATCH + 1}/{(len(texts) - 1) // EMBED_BATCH + 1}"
        )

    vectors = []
    for chunk, values in zip(chunks, all_embeddings):
        vectors.append(
            {
                "id": chunk["id"],
                "values": values,
                "metadata": chunk["metadata"],
            }
        )

    print(f"  upserting {len(vectors)} vectors")
    upsert_vectors(vectors)


def parse_changed_files(lines: Iterable[str]) -> list[tuple[str, str]]:
    changed = []
    for raw_line in lines:
        line = raw_line.rstrip("\n")
        if not line:
            continue

        parts = line.split("\t")
        status = parts[0]

        if status.startswith("R"):
            if len(parts) != 3:
                continue
            old_path, new_path = parts[1], parts[2]
            for item_status, item_path in (("D", old_path), ("A", new_path)):
                if item_path.startswith("content/en/") and item_path.endswith(".md"):
                    changed.append((item_status, item_path))
            continue

        if len(parts) != 2:
            continue

        path = parts[1]
        if status not in {"A", "M", "D"}:
            continue
        if not path.startswith("content/en/") or not path.endswith(".md"):
            continue
        changed.append((status, path))

    return changed


def main() -> int:
    diff_file = sys.argv[1]
    with open(diff_file) as f:
        changed = parse_changed_files(f)

    if not changed:
        print("No markdown files changed.")
        return 0

    errors = 0
    for status, path in changed:
        try:
            process_file(path, status)
        except Exception as e:
            print(f"ERROR processing {path}: {e}")
            errors += 1

    if errors:
        print(f"\n{errors} file(s) failed to process.")
        return 1

    print("\nAll changed docs indexed successfully.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
