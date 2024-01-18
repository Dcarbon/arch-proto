# Images name: dcarbon/golang

FROM golang:1.21-alpine

RUN apk update && apk add --no-cache git ca-certificates alpine-sdk protobuf

WORKDIR /dcarbon
COPY . .

RUN echo  "Start download dependency" && \
    # go mod download && \
    go get -u -v google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get -u -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2  && \
    go get -u -v google.golang.org/protobuf/cmd/protoc-gen-go  && \
    go get -u -v google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -u -v github.com/mwitkow/go-proto-validators && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway  && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2  && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators && \
    echo "Load dependency successfully...!" 


ENV PATH=$PATH:$GOPATH/bin