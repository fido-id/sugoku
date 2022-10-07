#!/bin/bash

echo -n "Running tests... "
while inotifywait -r . -e create,delete,modify; do
  echo "Running tests..."
  gotest -race -count=1 -v -coverpkg=./... -coverprofile=coverage.out ./...
done
