#!/bin/bash

env
git status
git branch

git config user.email "gen@example.com"
git config user.name  "gen"
git remote -v

git rm -rf .github/
git add gen/
git commit -m "gen ${GITHUB_SHA:0:7}"
git push -f origin ${GITHUB_REF}:${GITHUB_REF}-gen
