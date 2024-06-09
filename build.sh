#!/bin/bash

# Read the version from the version file
VERSION=$(cat version.txt)

# Build the project with ldflags to set the version variable
go build -ldflags="-X 'github.com/lanytcc/spot/cmd.version=${VERSION}'" -o spotvm main.go
