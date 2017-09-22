#!/bin/bash
set -euo pipefail

DIR="$(dirname "${BASH_SOURCE[0]}")"

function USAGE () {
    echo " <cmd>"
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

CMD=${1}

# DslContext.name: server-db-common

$DIR/postgres.sh ${CMD}


# HOFSTADTER_BELOW

