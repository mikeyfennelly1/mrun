#!/usr/bin/env bats

# bats test_tags=basic
@test "mrun version" {
    mrun version
    is "$output" "mrun version .*"               "'Version line' in output with version command"
}