#!/bin/sh
GOOS=js GOARCH=wasm go build -a -o wasm/main.wasm /go/src/golang/wasm/main.go
