#!/usr/bin/env bash

BASEDIR=$(dirname $0)
WD=`pwd`
$BASEDIR/chk_run_docker_compose.sh
if [ "$?" == "1" ]; then
    exit 1
fi
HOST_SERVICES=""
if [ -z ${DOCKER_HOST+x} ]; then
    HOST_SERVICES="localhost"
else
    HOST_SERVICES=$(echo $DOCKER_HOST | sed 's/^.*:\/\///' | sed 's/:[0-9]*$//')
fi

DYNO="true" GAUTOCLOUD_HOST_SERVICES="$HOST_SERVICES" go run "$BASEDIR/doc-generator/doc-generator.go" > "$BASEDIR/../docs/connectors.md"
