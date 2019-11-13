#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f src/base/gqlgen/generated.go \
    src/base/gqlgen/models/generated.go \
    src/base/gqlgen/resolvers/generated/resolver.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"