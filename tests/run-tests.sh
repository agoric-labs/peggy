#!/bin/bash
TEST_TYPE=$1
set -eu

# Run test entry point script
docker exec -it peggy_test_instance /bin/sh -c "pushd /peggy/ && tests/container-scripts/integration-tests.sh 3 $TEST_TYPE"