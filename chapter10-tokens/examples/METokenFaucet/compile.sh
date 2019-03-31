#!/bin/bash

rm -rf contracts/build 
mkdir contracts/build
docker run --rm -v ${PWD}/contracts:/contracts --workdir /contracts ethereum/solc:0.5.6 --bin --optimize -o build/ METoken.sol