#!/usr/bin/env bash

go install github.com/bokwoon95/wgo@latest

echo "Watching for changes to .go files to regenerate documentation..."

wgo -verbose -file=.go ./docs/generate.sh