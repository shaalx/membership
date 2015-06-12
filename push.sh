#!/bin/sh
git status
git add -A
read commits
git commit -m "$commits"
git push origi master:master
git push origin master:master
