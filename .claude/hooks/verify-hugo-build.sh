#!/bin/bash
# Verify Hugo builds successfully after file changes.
# Used as a Claude Code PostToolUse hook for Edit/Write/MultiEdit operations.
# On failure, automatically reverts the file to its pre-edit state.

set -uo pipefail

# --- Environment setup ---
export PATH="$CLAUDE_PROJECT_DIR/node_modules/.bin:$PATH"

# --- Parse hook input ---
INPUT=$(cat)
TOOL_NAME=$(echo "$INPUT" | jq -r '.tool_name // empty')
FILE_PATH=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty')

# --- Early exits for irrelevant files ---
if [[ -z "$FILE_PATH" ]]; then
  exit 0
fi

case "$FILE_PATH" in
  *.md|*.toml|*.yaml|*.yml|*.html|*.json)
    ;;
  *)
    exit 0
    ;;
esac

# --- Revert helpers ---
revert_edit() {
  local file="$1"
  local new_string=$(echo "$INPUT" | jq -r '.tool_input.new_string // empty')
  local old_string=$(echo "$INPUT" | jq -r '.tool_input.old_string // empty')
  [[ -z "$new_string" || -z "$old_string" ]] && return 1

  python3 -c "
import sys
path, new, old = sys.argv[1], sys.argv[2], sys.argv[3]
with open(path, 'r') as f:
    content = f.read()
with open(path, 'w') as f:
    f.write(content.replace(new, old, 1))
" "$file" "$new_string" "$old_string" 2>/dev/null
}

revert_write() {
  local file="$1"
  git -C "$CLAUDE_PROJECT_DIR" checkout -- "$file" 2>/dev/null
}

revert_file() {
  case "$TOOL_NAME" in
    Edit)  revert_edit "$FILE_PATH" ;;
    Write) revert_write "$FILE_PATH" ;;
    *)     return 1 ;;
  esac
}

# --- Hugo build verification ---
OUTPUT=$(hugo --renderToMemory --source "$CLAUDE_PROJECT_DIR" 2>&1)

if [[ $? -ne 0 ]]; then
  echo "Hugo build FAILED after editing: $FILE_PATH" >&2
  echo "" >&2
  echo "$OUTPUT" | grep -i -E '(error|ERROR|fatal|FATAL)' >&2
  echo "" >&2

  if revert_file; then
    echo "The file has been automatically reverted to its pre-edit state." >&2
  else
    echo "Auto-revert failed. Please fix the issue manually." >&2
  fi
  exit 2
fi

exit 0
