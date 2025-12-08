#!/usr/bin/env bash

go install github.com/bokwoon95/wgo@latest

echo "Watching for .go file changes to regenerate documentation..."

wgo -verbose -file=.go -xdir examples \
  ./docs/generate.sh \
  :: go run ./docs/gen/main.go && echo "✔ Examples generated in ./examples/"
