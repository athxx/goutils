#!/bin/bash

DIR=$(dirname $(readlink -f $0))
PROJECT_BINARY_DIR=$1
cd ${PROJECT_BINARY_DIR}
VERSION=$(git rev-parse --short HEAD)
BUILD_MODE=$2
BUILD_AT=$(date +%s)
# chmod -R 755 $DIR/../
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

cd $DIR

rm -f $DIR/../bin/slm

# 判断路径是否test或者prd模式, 如果是则使用生产模式, 方法比较挫
# 建议后续使用gccgo进行打包, 体积相差20倍, 具体请看 https://go.dev/doc/install/gccgo
PARAM=""
if [[ $DIR =~ "/prd/" ]]; then
     PARAM="-ldflags \"-s -w\""
fi

echo "compile soulma backend..."
go build -a -o $DIR/../bin/slm -ldflags "-X 'svr/app/lib.version=$VERSION' -X 'svr/app/lib.builtAt=${BUILD_AT}' -X 'svr/app/lib.buildMode=${BUILD_MODE}'" ${PARAM}

# echo "compile [GO HOST SERVER] oxProc..."
# go build -mod=mod -ldflags "$PARAM" -o $DIR/../bin/oxProc $DIR/cmd/oxProc/

# echo "compile [GO HOST SERVER] oxFont..."
# go build -mod=mod -ldflags "$PARAM" -o $DIR/../bin/oxFont $DIR/cmd/oxFont/

# echo "compile [GO HOST APPLICATION] xkProc..."
# go build -mod=mod -ldflags "$PARAM" -o $DIR/../bin/xkProc $DIR/cmd/xkProc/

# echo "compile [GO HOST APPLICATION] xkFont..."
# go build -mod=mod -ldflags "$PARAM" -o $DIR/../bin/xkFont $DIR/cmd/xkFont/
