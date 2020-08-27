#!/bin/sh

echo "Executing docker build ..."
# docker build --tag scarrionv/simple-api:1.0
docker build . -t scarrionv/simple-api:1.0
echo "Executed docker build."

