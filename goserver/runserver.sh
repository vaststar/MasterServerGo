#!/bin/bash

cd "$(dirname "$0")"

echo "====start: cleaning up rundata===="
cd ..
mkdir -p rundata && rm -rf rundata/server_go && rm -rf rundata/config/*
cp -R ./goserver/config rundata/
echo "====done: cleaning up rundata===="

echo "====start: build goserver===="
cd goserver/src
go build main.go
mv main ../../rundata/server_go
cd ../../rundata
echo "====done: build goserver===="

echo "====start: run server===="
./server_go --config=./config/config.json