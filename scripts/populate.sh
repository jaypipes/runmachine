#!/usr/bin/env bash

DEBUG=${DEBUG:-0}
VERBOSE=${VERBOSE:-0}
VERSION=$(git describe --tags --always --dirty)
ROOT_DIR=$(cd $(dirname "$0")/.. && pwd)
SCRIPTS_DIR=$ROOT_DIR/scripts
LIB_DIR=$SCRIPTS_DIR/lib
FIXTURES_DIR=$ROOT_DIR/tests/data/

source $LIB_DIR/common

if debug_enabled; then
    set -o xtrace
fi

source $SCRIPTS_DIR/service-up.sh

for f in $FIXTURES_DIR/objects/partitions/*; do
    part_id=$( basename "$f")
    echo -n "creating partition '$part_id' ... "
    $SCRIPTS_DIR/runm.sh partition create -f tests/data/objects/partitions/$part_id
done

echo -n "creating provider definition  ... "
$SCRIPTS_DIR/runm.sh provider definition set -f tests/data/definitions/runm.provider.yaml

for f in $FIXTURES_DIR/objects/providers/*; do
    prov_id=$( basename "$f")
    echo -n "creating provider '$prov_id' ... "
    $SCRIPTS_DIR/runm.sh provider create -f tests/data/objects/providers/$prov_id
done
