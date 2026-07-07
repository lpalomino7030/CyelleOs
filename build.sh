#!/usr/bin/env bash

set -e

PROFILE="profiles/cyelle"
WORKDIR="work"
OUTDIR="out"

sudo rm -rf "$WORKDIR" "$OUTDIR"

sudo mkarchiso \
    -v \
    -w "$WORKDIR" \
    -o "$OUTDIR" \
    "$PROFILE"