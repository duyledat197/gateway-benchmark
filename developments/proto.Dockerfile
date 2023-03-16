
FROM golang:1.19.1 AS protoc_gen_go

RUN apt update && apt install -y --no-install-recommends curl make git unzip apt-utils
ENV GO111MODULE=on
ENV PROTOC_VERSION=3.14.0
ENV GRPC_WEB_VERSION=1.2.1
ENV BUFBUILD_VERSION=0.24.0

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip
RUN unzip protoc-$PROTOC_VERSION-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.1.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/envoyproxy/protoc-gen-validate@v0.9.1

ENV MOD=$GOPATH/pkg/mod
RUN mv --force $MOD/github.com/grpc-ecosystem/grpc-gateway/v2@v2.1.0/third_party/googleapis/* /usr/local/include/
RUN mv $MOD/github.com/envoyproxy/protoc-gen-validate@v0.9.1/validate /usr/local/include/
RUN mv protoc3/include/google/protobuf /usr/local/include/google/

