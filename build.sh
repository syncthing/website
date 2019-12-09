#!/bin/bash
set -euo pipefail

rel=$(curl -s https://api.github.com/repos/syncthing/syncthing/releases/latest \
	| grep tag_name \
	| awk '{print $2}' \
	| tr -d \",v)

echo "stable: $rel" > data/release.yaml
hugo 

