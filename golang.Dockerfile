# Images name: dcarbon/golang

FROM dcarbon/golang

WORKDIR /dcarbon/arch-proto

COPY . .

RUN ./scripts/build-go-grpc.sh