#!/bin/bash

manifest_dir=${1:-"../manifests/"}

for file in $manifest_dir/*; do
    echo -n "  - " >> ./configspec-head.yaml
    cat $file | base64 -w 0 >> ./configspec-head.yaml
    echo "" >> ./configspec-head.yaml
done
