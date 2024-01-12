#!/bin/bash
archs="arm64 amd64"
oss="darwin linux windows"
for arch in $archs; do
    for os in $oss; do
        echo "Building $arch $os"
        GOOS=$os GOARCH=$arch go build -o build/$arch-$os/talkey ./cmd/...
    done
done