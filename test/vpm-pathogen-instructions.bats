#!/usr/bin/env bats
set -e

@test "vpm-pathogen-instructions runs successfully" {
  run vpm pathogen-instructions
  [ "$status" -eq 0 ]
}

@test "vpm-version ouputs version" {
  run vpm pathogen-instructions
  [ -n "$output" ]
}
