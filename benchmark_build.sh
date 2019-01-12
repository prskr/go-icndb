#!/usr/bin/env bash

echo "Starting benchmark with go build..."

start_time=$(date +%s%N)

go build cmd/icndb-server/main.go

end_time=$(date +%s%N)

diff=$(expr $(expr $end_time - $start_time) / 1000000)

echo "Took ${diff}ms to run go build"

echo "Starting benchmark with bazel build..."

start_time=$(date +%s%N)

bazel build //cmd/icndb-server

end_time=$(date +%s%N)

diff=$(expr $(expr $end_time - $start_time) / 1000000)

echo "Took ${diff}ms to run bazel build"