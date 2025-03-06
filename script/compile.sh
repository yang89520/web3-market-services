#!/bin/bash

# 检查是否安装了 Go
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed" >&2
    exit 1
fi

# 确保 GOPATH 已设置
if [ -z "$GOPATH" ]; then
    export GOPATH=$HOME/go
    echo "GOPATH was not set, using default: $GOPATH"
fi

# 如果目录不存在 创建必要的目录
if [ ! -d "$GOPATH/bin" ]; then
    mkdir -p $GOPATH/bin
fi

# 如果不存在 protoc插件，则安装
if ! command -v protoc-gen-go &> /dev/null; then
    echo "Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    if [ $? -ne 0 ]; then
        echo "Failed to install protoc-gen-go" >&2
        exit 1
    fi
fi

# 如果不存在 protoc-gen-go-grpc 插件，则安装
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    if [ $? -ne 0 ]; then
        echo "Failed to install protoc-gen-go-grpc" >&2
        exit 1
    fi
fi

# 添加 GOBIN 到 PATH
export PATH="$GOPATH/bin:$PATH"

# 编译 proto 文件
echo "Compiling protobuf files..."
protoc -I ./ \
    --go_out=./ \
    --go-grpc_out=require_unimplemented_servers=false:. \
    proto/*.proto
if [ $? -ne 0 ]; then
    echo "Failed to compile protobuf files" >&2
    exit 1
else
    echo "Compilation completed successfully"
fi