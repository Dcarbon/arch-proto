cd $( git rev-parse --show-toplevel )

OUT_DIR="."
OUT_GRPC="$OUT_DIR/pbg"
OUT_SWG="$OUT_DIR/swg"

rm -rf $OUT_DIR
mkdir -p $OUT_GRPC
mkdir -p $OUT_SWG


protoc -I proto -I proto/libs \
            --go_out=$OUT_DIR --go-grpc_out=$OUT_DIR  \
            --grpc-gateway_out $OUT_GRPC \
            --grpc-gateway_opt logtostderr=true \
            --grpc-gateway_opt paths=source_relative \
            --grpc-gateway_opt generate_unbound_methods=true \
            --openapiv2_out $OUT_SWG --openapiv2_opt use_go_templates=true \
            proto/*.proto

# protoc -I proto -I proto/libs \
#             --go_out=$OUT_DIR --go-grpc_out=$OUT_DIR  \
#             proto/*.proto


if [[ "$?" == "0" ]];then
    if [ ! -z "$GOPATH" ]; then
        echo "Add golang extension code "
        cp proto/extensions/go/*.ext $OUT_GRPC
    fi

    cd $OUT_GRPC
    # pwd 
    for f in *.go.ext; do
        mv "$f" "$( basename $f .ext )"
    done
else
    echo "Build error." 
fi 
