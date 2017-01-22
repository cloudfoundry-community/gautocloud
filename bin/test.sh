#!/usr/bin/env bash

echo "------------------"
echo "Running unit tests"
echo "------------------"
echo ""
BASEDIR=$(dirname $0)
go test -v "./$BASEDIR/../..." -args -ginkgo.randomizeAllSpecs -ginkgo.trace
code="$?"
if [ "$code" != "0" ]; then
    exit $code
fi
echo ""
echo "--------------------------"
echo ""

echo "--------------------------"
echo "Running integrations tests"
echo "--------------------------"
echo ""
$BASEDIR/chk_run_docker_compose.sh
if [ "$?" == "1" ]; then
    echo "Integration not ran."
    exit 0
fi
HOST_SERVICES=""
if [ -z ${DOCKER_HOST+x} ]; then
    HOST_SERVICES="localhost"
else
    HOST_SERVICES=$(echo $DOCKER_HOST | sed 's/^.*:\/\///' | sed 's/:[0-9]*$//')
fi
DYNO="true" GAUTOCLOUD_HOST_SERVICES="$HOST_SERVICES" go test -v "./$BASEDIR/../test-integration" -args -ginkgo.trace -ginkgo.randomizeAllSpecs
code="$?"
echo ""
echo "--------------------------"
echo ""

exit $code