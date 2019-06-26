#!/usr/bin/env sh
set -e

TAR_FILE="/tmp/goreleaser.tar.gz"
RELEASES_URL="https://github.com/goreleaser/goreleaser/releases"
test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"
TAG=$1

last_version() {
  curl -sL -o /dev/null -w %{url_effective} "$RELEASES_URL/latest" |
    rev |
    cut -f1 -d'/'|
    rev
}

download() {
  test -z "$VERSION" && VERSION="$(last_version)"
  test -z "$VERSION" && {
    echo "Unable to get goreleaser version." >&2
    exit 1
  }
  rm -f "$TAR_FILE"
  curl -s -L -o "$TAR_FILE" \
    "$RELEASES_URL/download/$VERSION/goreleaser_$(uname -s)_$(uname -m).tar.gz"
}

extract() {
  tar -xf "$TAR_FILE" -C "$TMPDIR"
}

release_notes() {
  rm -f "${TMPDIR}/release-notes.md"
  bash scripts/release-notes.sh "$TAG" > "$TMPDIR/release-notes.md"
}

release() {
  "${TMPDIR}/goreleaser" --skip-validate --release-notes="${TMPDIR}/release-notes.md"
}

download
extract
release_notes
release
