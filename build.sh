#!/bin/bash
set -euo pipefail

pushd script
go run ./sponsors > ../themes/default/layouts/partials/github-sponsors.html
go run ./release > ../data/release.yaml
popd

hugo --minify
