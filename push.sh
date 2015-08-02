#!/bin/sh
git status
git add -A
read commits
git commit -m "$commits"
git push origin master:master
