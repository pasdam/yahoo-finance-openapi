#!/bin/sh

set -e
set -o xtrace

SCRIPT_DIR="$(cd "$( dirname "$0" )" >/dev/null && pwd )"
PROJECT_ROOT=$(dirname $SCRIPT_DIR)

API_VERSION=$(sh $SCRIPT_DIR/api_version.sh)
LATEST_TAG=$(cd $SCRIPT_DIR && git fetch && git describe --tags $(git rev-list --tags --max-count=1) 2> /dev/null | sed "s/^v//")
if [ "$API_VERSION" = "$LATEST_TAG" ]; then
  1>&2 echo "Please bump the API version in the api.yml OpenApi spec"
  exit 1
fi
