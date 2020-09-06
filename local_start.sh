#!/usr/bin/env bash
select d in cmd/*/; do test -n "$d" && break; echo ">>> Invalid Selection"; done
PROJECT_NAME=$(echo $d | cut -d'/' -f1)
CURRENT_PATH=$PWD
cd $(echo $PROJECT_NAME/)
killall -9 $PROJECT_NAME
find $CURRENT_PATH -type f -name "*.go" | entr -r sh -c "clear && go build -o $CURRENT_PATH/bin/$PROJECT_NAME && $CURRENT_PATH/bin/$PROJECT_NAME"