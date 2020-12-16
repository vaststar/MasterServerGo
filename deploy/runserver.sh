#!/bin/bash

cd "$(dirname "$0")"

echo "cleaning up"
cd ..
rm -rf rundata/* && mkdir rundata
cp -R ./deploy/config rundata/config/

echo "build server"
cd "goserver"
go build main.go
mv main ../rundata/server_go
cd ../rundata

echo "run server"
./server_go --config=./config/config.json