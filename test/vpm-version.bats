#!/usr/bin/env bats
set -e

@test "vpm-version runs successfully" {
  run vpm version
  [ "$status" -eq 0 ]
}

@test "vpm-version ouputs version" {
  run vpm version
  [ -n "$output" ]
}
