#!/bin/bash

docker-compose down
./build-docker.sh 1
docker-compose up -d 
