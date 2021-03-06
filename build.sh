#!/usr/bin/env bash
#/etc/crwords/crazy-words-backend/Dockerfile
git -C ./crazy-words-backend pull
PREV_VER=$CRW_TAG_VERSION
CUR_VER=$(($PREV_VER+1))
PREV_TAG=crw-$PREV_VER
CUR_TAG=crw-$CUR_VER
docker ps | grep crw | awk '{print $1}' | xargs docker stop
docker build --tag=$CUR_TAG --tag=crw:latest -f ./crazy-words-backend/Dockerfile
docker run -v /usr/share/dict:/dict -p 8099:8080 -d crw /app/main -words /dict/words -colors /app/list-of-colors -animals /app/animals-list
export CRW_TAG_VERSION=$CUR_VER