#!/usr/bin/env bash

# This script is designed for the kyma-addons releasing process which is different from the Kyma releasing process.
set -e

# readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# source "${SCRIPT_DIR}"

for name in alpine-kubectl bootstrap bootstrap-helm golang; do make -C $name ci-release  ; done