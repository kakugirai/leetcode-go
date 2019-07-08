#!/usr/bin/env bash

for d in src/*; do
    if test -f "$d/solution_test.go"; then
        go test $d/solution.go $d/solution_test.go
    fi
done
