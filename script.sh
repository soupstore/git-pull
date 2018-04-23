#!/bin/sh

cd $TARGET

while true; do git pull https://$GIT_USER:$GIT_PASS@$REPO_SLUG.git; sleep 60; done
