#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo <&2 "useage: $(basename $0) version"
    exit
fi

docker build -t note-server:v$1 ../