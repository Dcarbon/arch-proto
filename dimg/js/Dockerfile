
FROM node:20.10-alpine
RUN apk add --no-cache git ca-certificates protobuf

RUN wget https://github.com/grpc/grpc-web/releases/download/1.5.0/protoc-gen-grpc-web-1.5.0-linux-x86_64 && \
    chmod +x protoc-gen-grpc-web-1.5.0-linux-x86_64 && \
    mv protoc-gen-grpc-web-1.5.0-linux-x86_64 /usr/bin/protoc-gen-grpc-web && \
    echo "Build done"

RUN npm i -g grpc-web