#!/usr/bin/env bash
set -euo pipefail
cargo build --release
echo "rust-crypto built."
