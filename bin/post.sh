#!/bin/bash

curl -v http://localhost:8080/notes --include \
 --header "Content-Type: application/json" \
 --data '{"id": 3, "title": "Test3", "content": "this is the third test"}'
