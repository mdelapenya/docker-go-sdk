#!/bin/bash

# =============================================================================
# Common Shell Functions and Constants
# =============================================================================
# Description: Shared utilities and constants for Docker Go SDK release scripts
# 
# Environment Variables:
#   DRY_RUN          - Enable dry run mode (default: true)
#                      When true, commands are echoed instead of executed
#
# Functions:
#   curlGolangProxy  - Trigger Go proxy to fetch module (for publishing)
#   execute_or_echo  - Execute command or echo based on DRY_RUN setting
#   find_latest_tag  - Find latest tag for a given module
#   get_modules      - Get list of modules from go.work file
#   get_script_dir   - Get directory of the calling script
#
# Constants:
#   ROOT_DIR         - Root directory of the repository
#   CURRENT_DIR      - Current directory of the script
#   GITHUB_REPO      - GitHub repository identifier
#   DRY_RUN          - Dry run mode flag
#
# Usage:
#   source .github/scripts/common.sh
#
# =============================================================================

# Common constants and functions for release scripts

# Get the directory of the script that sources this file
get_script_dir() {
  cd "$( dirname "${BASH_SOURCE[1]}" )" && pwd
}

readonly CURRENT_DIR="$(get_script_dir)"
readonly ROOT_DIR="$(dirname $(dirname "${CURRENT_DIR}"))"
readonly GITHUB_REPO="github.com/docker/go-sdk"
readonly DRY_RUN="${DRY_RUN:-true}"

# This function is used to trigger the Go proxy to fetch the module.
# See https://pkg.go.dev/about#adding-a-package for more details.
function curlGolangProxy() {
  local module="${1}"
  local module_version="${2}"

  # e.g.:
  #   github.com/docker/go-sdk/client/v0.1.0-alpha001.info
  #   github.com/docker/go-sdk/client/v0.0.1.info
  #   github.com/docker/go-sdk/client/v0.1.0.info
  #   github.com/docker/go-sdk/client/v1.0.0.info
  if [[ "${DRY_RUN}" == "true" ]]; then
    echo "[DRY RUN] Would execute: curl https://proxy.golang.org/${GITHUB_REPO}/${module}/@v/${module_version}.info"
    return
  fi

  curl "https://proxy.golang.org/${GITHUB_REPO}/${module}/@v/${module_version}.info"
}


# Function to execute or echo commands based on DRY_RUN
execute_or_echo() {
  if [[ "$DRY_RUN" == "true" ]]; then
    echo "[DRY RUN] Would execute: $*"
  else
    "$@"
  fi
}

# Function to get modules from go.work
get_modules() {
  go work edit -json | jq -r '.Use[] | "\(.DiskPath | ltrimstr("./"))"' | tr '\n' ' ' && echo
}

# Function to find latest tag for a module
find_latest_tag() {
  local module="$1"
  git tag --list | grep -E "${module}/v[0-9]+\.[0-9]+\.[0-9]+.*" | sort -V | tail -n 1
}

# Portable in-place sed editing that works on both BSD (macOS) and GNU (Linux) sed
portable_sed() {
  local pattern="$1"
  local file="$2"
  
  if [[ "$DRY_RUN" == "true" ]]; then
    echo "[DRY RUN] Would execute: sed '$pattern' in $file"
    return
  fi
  
  # Detect sed version and use appropriate syntax
  if sed --version >/dev/null 2>&1; then
    # GNU sed (Linux)
    sed -i "$pattern" "$file"
  else
    # BSD sed (macOS)
    sed -i '' "$pattern" "$file"
  fi
}