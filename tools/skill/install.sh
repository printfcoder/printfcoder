#!/usr/bin/env bash
# Symlink every skill in tools/skill into the IDE skill directories.
# One source of truth; rerun after adding a new skill. No copies, no network.
set -euo pipefail

repo_root="$(cd "$(dirname "$0")/../.." && pwd)"
src_dir="$repo_root/tools/skill"

# IDE -> project-local skill directory
targets=(
  ".trae/skills"
  ".cursor/skills"
)

for skill_dir in "$src_dir"/*/; do
  skill="$(basename "$skill_dir")"
  for target in "${targets[@]}"; do
    dest="$repo_root/$target"
    mkdir -p "$dest"
    ln -sfn "../../tools/skill/$skill" "$dest/$skill"
    echo "linked $target/$skill -> tools/skill/$skill"
  done
done
