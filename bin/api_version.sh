#!/bin/sh

SCRIPT_DIR="$(cd "$( dirname "$0" )" >/dev/null && pwd )"
PROJECT_ROOT=$(dirname $SCRIPT_DIR)
cat $PROJECT_ROOT/api.yml | grep -e "^  version: " | sed "s/  version: //"
