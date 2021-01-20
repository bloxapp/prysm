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
ARG RPC_PORT
ARG GRPC_GATEWAY_PORT
ARG MONITORING_PORT
ARG P2P_TCP_PORT
ARG P2P_UDP_PORT

ENV HTTP_WEB3_PROVIDER=$HTTP_WEB3_PROVIDER
ENV PRYSM_NETWORK=$PRYSM_NETWORK
ENV P2P_HOST_DNS=$P2P_HOST_DNS
ENV RPC_PORT=$RPC_PORT
ENV GRPC_GATEWAY_PORT=$GRPC_GATEWAY_PORT
ENV MONITORING_PORT=$MONITORING_PORT
ENV P2P_TCP_PORT=$P2P_TCP_PORT
ENV P2P_UDP_PORT=$P2P_UDP_PORT

RUN  cd /prysm && ls -lah /prysm && bazel build //beacon-chain:beacon-chain

EXPOSE 3500 4000 6668 13000 12000/UDP 15000 16000/UDP 16668 14000

CMD cd /prysm && /usr/bin/bazel run //beacon-chain:beacon-chain -- --datadir=/data --http-web3provider=$HTTP_WEB3_PROVIDER --rpc-host=0.0.0.0 --rpc-port=$RPC_PORT --grpc-gateway-host=0.0.0.0 --grpc-gateway-port=$GRPC_GATEWAY_PORT --monitoring-host=0.0.0.0 --monitoring-port=$MONITORING_PORT  --p2p-max-peers=100 --$PRYSM_NETWORK --accept-terms-of-use --verbosity=debug --p2p-host-dns=$P2P_HOST_DNS --p2p-tcp-port=$P2P_TCP_PORT --p2p-udp-port=$P2P_UDP_PORT --subscribe-all-subnets
