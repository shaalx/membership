#!/bin/sh
for tag in $(git tag)
do
 echo $tag
 git push origin --tag :$tag
 git tag -d $tag
done
