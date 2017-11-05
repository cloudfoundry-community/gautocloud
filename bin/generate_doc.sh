#!/usr/bin/env bash

BASEDIR=$(dirname $0)
WD=`pwd`
$BASEDIR/chk_run_docker_compose.sh
if [ "$?" == "1" ]; then
    exit 1
fi

DYNO="true" GAUTOCLOUD_HOST_SERVICES="localhost" go run "$BASEDIR/doc-generator/doc-generator.go" > "$BASEDIR/../docs/connectors.md"
