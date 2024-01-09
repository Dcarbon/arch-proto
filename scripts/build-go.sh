cd $( git rev-parse --show-toplevel )

OUT_FOLDER=.
PACKAGE="$OUT_FOLDER/pb"

rm -rf $PACKAGE
mkdir -p $PACKAGE 

# Build  
protoc -I . --go_out=$OUT_FOLDER --go-grpc_out=$OUT_FOLDER  *.proto

if [[ "$?" == "0" ]];then
    echo "Build success"
else
    echo "Build error." 
fi 
