#!/bin/bash

if [ $# == 0 ]; then
  echo "two few arguments"
  exit -1
fi

docker run --rm ethereum/client-go:alltools-v1.8.23 geth $@