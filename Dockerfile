# syntax=docker/dockerfile:1.0.0-experimental

FROM golang:1.13

ENV HOME /root
ENV PATH /go/bin:/src/bin:/usr/local/go/bin:$PATH
ENV USER "root"

WORKDIR /work

RUN apt-get update -y \
  && apt-get install -y --force-yes \
      jq \
      protobuf-compiler \
      less \
      curl \
      tree \
      pkg-config \
      libtool \
      autoconf \
      git

RUN go get github.com/openconfig/ygot/generator \
  && go get golang.org/x/tools/cmd/goimports \
  && go install github.com/openconfig/ygot/generator \
  && go install github.com/openconfig/ygot/proto_generator

CMD ["make", "build"]
