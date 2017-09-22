#!/bin/bash
set -euo pipefail


function USAGE () {
    echo "postgres <cmd>"
    echo "  where cmd is:"
    echo "   - start"
    echo "   - stop"
    echo "   - restart"
    echo "   - wipe"
}

if [[ "$#" -ne "1" ]]; then
    USAGE
    exit 1
fi

function START() {
    mkdir -p $(pwd)/local/postgres

    docker run --name postgres \
        -p 5432:5432 \
        -e POSTGRES_PASSWORD="postgres" \
        -e POSTGRES_DBNAME="develop" \
        -e POSTGRES_DBUSER="admin" \
        -e POSTGRES_DBPASS="password" \
        -v $(pwd)/local/postgres:/var/lib/postgresql/data \
        -d \
        eeacms/postgres:9.6-2.0

}

function STOP () {
    docker rm -f postgres
}

function WIPE () {
    rm -rf $(pwd)/local/postgres
}

case $1 in
    wipe)
        set +e
        STOP
        set -e
        WIPE
        ;;
    stop)
        STOP
        ;;
    start)
        START
        ;;
    restart)
        STOP
        START
        ;;
    *)
        USAGE
        exit 1
        ;;
esac

# HOFSTADTER_BELOW

# HOFSTADTER_BELOW
