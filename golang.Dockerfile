# Images name: harbor.viet-tin.com/dcarbon/golang

FROM harbor.viet-tin.com/dcarbon/golang

WORKDIR /dcarbon/arch-proto

COPY . .

RUN ./scripts/build-go-grpc.sh