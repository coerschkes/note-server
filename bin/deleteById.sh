#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Missing parameter id"
    exit
fi

curl -X DELETE http://localhost:8080/notes/$1