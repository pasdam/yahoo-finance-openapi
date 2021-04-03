#!/bin/sh

SCRIPT_DIR="$(cd "$( dirname "$0" )" >/dev/null && pwd )"
PROJECT_ROOT=$(dirname $SCRIPT_DIR)

mv $PROJECT_ROOT/gen/go/yf/go.* $PROJECT_ROOT/

OS="`uname`"
if [ "$OS" = "Linux" ]; then
  # Linux
  sed -i "s|/gen/go/yf||" $PROJECT_ROOT/go.mod

elif [ "$OS" = "Darwin" ]; then
  # Mac OSX
  sed -i '' "s|/gen/go/yf||" $PROJECT_ROOT/go.mod

else
  1>&2 echo "Unrecognized OS: $OSTYPE"
  exit 1
fi
