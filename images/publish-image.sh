#!/usr/bin/env bash

set -e

usage () {
    echo "Usage: \$ ${BASH_SOURCE[1]} /path/to/image"
    exit 1
}

readonly SOURCES_DIR=$1

if [[ -z "${SOURCES_DIR}" ]]; then
    usage
fi

echo "using ${SOURCES_DIR} to build the image"

make -C "${SOURCES_DIR}" ci-release