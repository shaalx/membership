#!/bin/sh

git status
git add -A
read commits
git commit -m "$commits"

sleep 3