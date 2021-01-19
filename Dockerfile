FROM ubuntu:20.10

RUN  apt-get update && apt-get install -y curl wget telnet net-tools git gnupg procps python3-pip
RUN  curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
RUN  mv bazel.gpg /etc/apt/trusted.gpg.d/
RUN  echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN  apt update && apt install -y bazel && apt update && apt install -y bazel-3.7.0
RUN  mkdir /data
RUN  mkdir /prysm
COPY . /prysm

ARG HTTP_WEB3_PROVIDER
ARG PRYSM_NETWORK
ARG P2P_HOST_DNS


RUN  cd /prysm && ls -lah /prysm && bazel build //beacon-chain:beacon-chain

EXPOSE 3500 4000 6668 13000

RUN cd /prysm && /usr/bin/bazel run //beacon-chain:beacon-chain -- --datadir=/data --http-web3provider=$HTTP_WEB3_PROVIDER --rpc-host=0.0.0.0 --rpc-port=4000 --grpc-gateway-host=0.0.0.0 --grpc-gateway-port=3500 --monitoring-host=0.0.0.0 --monitoring-port=6668 --p2p-max-peers=500 --$PRYSM_NETWORK --accept-terms-of-use --head-sync --verbosity=debug --p2p-host-dns=$P2P_HOST_DNS --p2p-tcp-port=13000 --p2p-udp-port=12000 --subscribe-all-subnets
