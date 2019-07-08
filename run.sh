#!/usr/bin/env bash

for d in src/*; do
    if test -f "$d/solution.go"; then
        go run $d/solution.go
    fi
done
