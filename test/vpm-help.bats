#!/usr/bin/env bats
set -e

@test "vpm-help runs successfully" {
  run vpm --help
  [ "$status" -eq 0 ]
}

@test "vpm-help contains Usage" {
  run vpm --help
  [[ "$output" = *Usage* ]]
}

@test "vpm-help contains Options" {
  run vpm --help
  [[ "$output" = *Options* ]]
}

@test "vpm-help contains Commands" {
  run vpm --help
  [[ "$output" = *Commands* ]]
}
