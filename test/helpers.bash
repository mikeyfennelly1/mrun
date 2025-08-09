#!/bin/bash

export MRUN=${MRUN:-mrun}

# default timeout for an mrun command
export MRUN_TIMEOUT=${MRUN:-600}

export MRUN_TEST_CACHE=${MRUN_TEST_CACHE:-/tmp}/mrun-test-cache-$(id -u)
mkdir -p ${MRUN_TEST_CACHE}

function mrun_basic_setup() {
    # Temporary subdirectory, in which tests can write whatever they like
    # and trust that it'll be deleted on cleanup.
    echo "mrun_basic_setup"
}