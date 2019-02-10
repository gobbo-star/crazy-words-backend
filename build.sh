#!/usr/bin/env bash
#/etc/crwords/crazy-words-backend/Dockerfile
git pull
PREV_VER=$CRW_TAG_VERSION
CUR_VER=$(($PREV_VER+1))
PREV_TAG=crw-$PREV_VER
CUR_TAG=crw-$CUR_VER
docker ps | grep $PREV_TAG | awk '{print $1}' | xargs docker stop
docker build --tag=$CUR_TAG .
docker run -v /usr/share/dict:/dict -p 8099:8080 -d $PREV_TAG /app/main -words /dict/words -colors list-of-colors -animals animals-list
export CRW_TAG_VERSION=$CUR_VER