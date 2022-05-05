#!/bin/bash

manifest_dir=${1:-"../manifests/"}

for file in $manifest_dir/*; do
    echo -n "  - " >> ./head.yaml
    cat $file | base64 >> ./head.yaml
done
