FROM ubuntu:20.10

RUN  apt-get update && apt-get install -y curl wget telnet net-tools git gnupg procps python3-pip
RUN  curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
RUN  mv bazel.gpg /etc/apt/trusted.gpg.d/
RUN  echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN  apt update && apt install -y bazel && apt update && apt install -y bazel-3.7.0
RUN  mkdir /data
RUN  mkdir /prysm
COPY . /prysm

RUN  cd /prysm && ls -lah /prysm && bazel build //beacon-chain:beacon-chain

EXPOSE 3500 4000 6668 13000

ENTRYPOINT ["bash", "/prysm/entrypoint.sh"]
