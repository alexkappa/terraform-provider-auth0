#!/usr/bin/env bash

# This script maintains the text of a specific release from the CHANGELOG to be
# used as release notes.

set -e

if [[ ! -f CHANGELOG.md ]]; then
  echo "ERROR: CHANGELOG.md not found in pwd."
  echo "Please run this from the root of the terraform provider repository"
  exit 1
fi

version=$1

if [[ -z "$version" ]]; then
  echo "ERROR: version argument was not set."
  echo "Please run this with a version argument"
  exit 1
fi

awk -v ver="$version" '/## / { if (p) { exit }; if ($2 == ver) { p=1; next} } p' CHANGELOG.md
