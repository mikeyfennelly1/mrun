#!/usr/bin/env bats

setup() {
    # Source only function definitions; the source guard prevents main from running.
    source "${BATS_TEST_DIRNAME}/../scripts/build.sh"
}

@test "build.sh can be sourced without error" {
    run bash -c "source \"${BATS_TEST_DIRNAME}/../scripts/build.sh\""
    [ "$status" -eq 0 ]
}

@test "INODE_CAPS exits successfully" {
    run INODE_CAPS
    [ "$status" -eq 0 ]
}

@test "INODE_CAPS outputs cap_linux_immutable" {
    run INODE_CAPS
    [ "$output" = "cap_linux_immutable" ]
}

@test "INODE_CAPS outputs exactly one capability" {
    run INODE_CAPS
    local count
    count=$(echo "$output" | wc -l | tr -d ' ')
    [ "$count" -eq 1 ]
}

@test "INODE_CAPS output has no trailing comma" {
    run INODE_CAPS
    [[ "$output" != *"," ]]
}
