#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo <&2 "useage: $(basename $0) id title content"
    exit
fi

curl -v http://localhost:8080/notes --include \
 --header "Content-Type: application/json" \
 --data "{\"id\": \"$1\", \"title\": \"$2\", \"content\": \"$3\"}"