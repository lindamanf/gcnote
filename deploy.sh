#!/bin/bash

docker-compose down
docker-compose build --no-cache
docker-compose up -d

WEBSERVER=gcnote_web_1
DBSERVER=gcnote_db_1

echo "--Section:start golang server"

# function dexec(){
#   COMMAND=$1
# 	docker exec -t $WEBSERVER bash -lc "cd /go/src/github.com/lindamanf/gcnote && ${COMMAND}"
# }
# dexec 'go build'
