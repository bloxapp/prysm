#!/bin/bash

cd /prysm

/usr/bin/bazel run //beacon-chain:beacon-chain -- --datadir=/data --http-web3provider=http://ethereum1-testnet.blockchain:8545 --rpc-host=0.0.0.0 --rpc-port=4000 --grpc-gateway-host=0.0.0.0 --grpc-gateway-port=3500 --monitoring-host=0.0.0.0 --monitoring-port=6668 --p2p-max-peers=500 --pyrmont --accept-terms-of-use --head-sync --verbosity=debug --p2p-host-dns=eth2-prysm-ext.stage.bloxinfra.com --p2p-tcp-port=13000 --p2p-udp-port=12000
