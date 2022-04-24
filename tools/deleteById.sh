#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo <&2 "useage: $(basename $0) id"
    exit
fi

curl -X DELETE http://localhost:8080/notes/$1