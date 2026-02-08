#!/usr/bin/env bash
set -euo pipefail

GO=go
BUILD_FLAGS="-v -race"
DIST_DIR=build

BINS=(nmb)

for bin in "${BINS[@]}"; do
    (
        echo "$bin"
        $GO build $BUILD_FLAGS -o "$DIST_DIR/$bin" "./cmd/$bin"
    ) &
done

wait
echo "done"
