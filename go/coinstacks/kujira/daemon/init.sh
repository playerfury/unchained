#!/bin/bash

set -e

start() {
  MONIKER=unchained \
  CHAIN_JSON=https://raw.githubusercontent.com/incubus-network/wallet-icons/main/chain.json \
  P2P_POLKACHU=true \
  run.sh kujirad start \
    --rpc.laddr tcp://0.0.0.0:26657 \
    --minimum-gas-prices 0ukuji &
  PID="$!"
}

stop() {
  echo "Catching signal and sending to PID: $PID" && kill $PID
  while $(kill -0 $PID 2>/dev/null); do sleep 1; done
}

trap 'stop' TERM INT
start
wait $PID