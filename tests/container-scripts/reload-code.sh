#!/bin/bash
# Number of validators to start
NODES=$1
# what test to execute
TEST_TYPE=$2
set -eux

# Stop any currently running peggy and eth processes
pkill peggyd || true # allowed to fail
pkill geth || true # allowed to fail

# Wipe filesystem changes
for i in $(seq 1 $NODES);
do
    rm -rf "/validator$i"
done


# Do all the building before launching any chains.
cd /peggy/module/
make
make install
cd /peggy/orchestrator/test_runner
DEPLOY_CONTRACTS=1 TEST_TYPE=$TEST_TYPE NO_GAS_OPT=1 \
  PATH=$PATH:$HOME/.cargo/bin cargo build --release --bin test-runner

cd /peggy/
tests/container-scripts/setup-validators.sh $NODES
tests/container-scripts/run-testnet.sh $NODES

# deploy the ethereum contracts
pushd /peggy/orchestrator/test_runner
RUST_BACKTRACE=full RUST_LOG=INFO ../target/release/test-runner

# This keeps the script open to prevent Docker from stopping the container
# immediately if the nodes are killed by a different process
read -p "Press Return to Close..."
