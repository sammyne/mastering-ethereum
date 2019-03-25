#!/bin/bash

docker run --rm -v ${PWD}/contracts:/contracts --workdir /contracts ethereum/solc:0.5.6 $@