#!/bin/sh
git add -A
read tags
git commit -m $tags
git push origin devm:master
git tag -a $tags -m $tags
git push origin --tag $tags
sleep 30
git push origin --tag :$tags
git tag -d $tags