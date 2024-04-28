FROM golang:1.22

RUN mkdir /build
WORKDIR /build

RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler-grpc protobuf-compiler libprotobuf-dev && rm -rf /var/lib/apt/lists/*

RUN go mod init proto && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  && \
    go install github.com/envoyproxy/protoc-gen-validate@latest  && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

COPY proto /build/proto

WORKDIR /build/proto

CMD protoc \
   --validate_out=lang=go:../go \
   --go_out=../go \
   --go-grpc_out=../go \
   --grpc-gateway_out=../go \
   -I=. \
   -I=/usr/include \
  *.proto
