#!/usr/bin/env bash
hash docker-compose 2>/dev/null
if [ "$?" == "1" ]; then
    echo "You need docker compose and docker."
    exit 1
fi
BASEDIR=$(dirname $0)

to_wait="0"
while read -r line ; do
    echo "$line"
    echo $line | grep --quiet "up-to-date"
    exit_code="$?"
    if [ "$exit_code" != "0" ]; then
        to_wait="1"
    fi
done < <(docker-compose -f "$BASEDIR/../docker-compose-services/docker-compose.yml" up -d 2>&1)
if [ "$to_wait" == "1" ]; then
    echo ""
    echo "Ensure all docker start by waiting 30sec (wait only when containers are starting or updating)."
    sleep 30
    echo "Finished to wait"
fi
