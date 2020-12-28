#!/bin/bash
NODES=$1
TEST_TYPE=$2
set -eu

FILE=/contracts
if test -f "$FILE"; then
echo "Contracts already deployed, running tests"
else 
echo "Testnet is not started yet, please wait before running tests"
exit 0
fi 

set +e
killall -9 test-runner
set -e

pushd /peggy/orchestrator/test_runner
TEST_TYPE=$TEST_TYPE NO_GAS_OPT=1 \
  PATH=$PATH:$HOME/.cargo/bin cargo build --release --bin test-runner

RUST_BACKTRACE=full RUST_LOG=INFO ../target/release/test-runner
