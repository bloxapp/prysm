FROM ubuntu:20.10

RUN  python --version
RUN  apt-get update && apt-get install -y curl wget telnet net-tools git gnupg procps
RUN  curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
RUN  mv bazel.gpg /etc/apt/trusted.gpg.d/
RUN  echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN  apt update && apt install -y bazel && apt update && apt install -y bazel-3.2.0
RUN  mkdir /data

RUN  git clone https://github.com/bloxapp/prysm.git
RUN  cd /prysm && bazel build //beacon-chain:beacon-chain

#ENTRYPOINT ["bazel", "run", "//beacon-chain:beacon-chain", "--", "--medalla", "--accept-terms-of-use"]


#ENTRYPOINT ["bazel", "run", "//beacon-chain:beacon-chain", "--", "--datadir=/data", "--http-web3provider=http://ethereum1-testnet:8545", "--rpc-host=0.0.0.0", "--rpc-port=4000", "--grpc-gateway-host=0.0.0.0", "--grpc-gateway-port=3500", "--monitoring-host=0.0.0.0", "--monitoring-port=6668", "--p2p-max-peers=500", "--block-batch-limit=512", "--medalla", "--accept-terms-of-use"]

EXPOSE 3500 4000 6668 13000
