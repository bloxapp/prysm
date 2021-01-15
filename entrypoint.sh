#!/bin/bash

cd /prysm

/usr/bin/bazel run //beacon-chain:beacon-chain -- --datadir=/data --http-web3provider=http://ethereum1-mainnet.blockchain:8545 --rpc-host=0.0.0.0 --rpc-port=14000 --grpc-gateway-host=0.0.0.0 --grpc-gateway-port=13500 --monitoring-host=0.0.0.0 --monitoring-port=16668 --p2p-max-peers=500 --mainnet --accept-terms-of-use --head-sync --verbosity=debug --p2p-host-dns=eth2-prysm-mainnet-ext.bloxinfra.com --p2p-tcp-port=15000 --p2p-udp-port=16000 --subscribe-all-subnets
