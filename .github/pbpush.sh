#!/bin/bash
set -x

MSG=$(git log -1 --pretty=%B)
STATUS=0

# lint protos
prototool lint      proto     || STATUS=$?
prototool format -l proto     || STATUS=$?

# generate protos
prototool generate  proto     || STATUS=$?
prototool generate  proto_ext || STATUS=$?

# generate mocks
find gen -name 'mock_*.go' -delete
mockery -all -dir gen -inpkg  || STATUS=$?

# exit if any errors
[ $STATUS -eq 0 ] || exit $STATUS

# sync gen, proto, proto_ext folders
git clone https://nzoschke:${PUSH_TOKEN}@github.com/nzoschke/gomesh-interface.git && cd gomesh-interface
git checkout -b ${GITHUB_REF} origin/${GITHUB_REF}
git rm -r gen proto proto_ext
cp -r ../gen ../proto ../proto_ext .
git add -f .

# exit if no changes
[[ -z $(git status -uno --porcelain) ]] && echo "this branch is clean, no need to push..." && exit 0;

# push changes
git config user.email "gen@example.com"
git config user.name  "gen"
git commit -m "$MSG"
git push origin ${GITHUB_REF}