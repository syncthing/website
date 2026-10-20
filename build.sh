#!/bin/bash
set -euo pipefail

pushd script
go run ./sponsors > ../data/sponsors.yaml
go run ./release > ../data/release.yaml
popd

hugo --minify
