#!/bin/bash
set -euo pipefail

pushd script
go run . > ../themes/default/layouts/partials/github-sponsors.html
popd

rel=$(curl -s https://api.github.com/repos/syncthing/syncthing/releases/latest \
	| grep tag_name \
	| awk '{print $2}' \
	| tr -d \",v)

echo "stable: $rel" > data/release.yaml
hugo
