#!/bin/bash

yarn --cwd ./ui build

bash build_docs.sh

# clean up the old build
rm -rf build

# build the application
mkdir -p build
go build -o build/

# COPY STATIC RESOURCES
## copy config files
cp -rf config build/
## copy ui resources
cp -rf ui/build build/ui
## copy docs
mkdir -p build/docs && cp -f docs/swagger* build/docs/