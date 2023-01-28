#!/bin/sh

echo "Generating swaggo files"
swag init

echo "Building"
go build -o service_executable