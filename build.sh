#!/bin/bash

yarn --cwd ./ui build

# generate Swagger API documentation
swag init

# clean up the old build
rm -rf build

# build the application
mkdir -p build
go build -o build/

# COPY STATIC RESOURCES
## copy config files
cp -rf config build/
## copy ui resources
mkdir -p build/ui && cp -rf ui/build build/ui/
## copy docs
mkdir -p build/docs && cp -f docs/swagger* build/docs/