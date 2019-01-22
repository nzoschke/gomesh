#!/bin/bash
set -x
STATUS=0

go version
diff -u <(echo -n) <(gofmt -d ./) || STATUS=$?
golint ./...   || STATUS=$?
go vet ./...   || STATUS=$?
go build ./... || STATUS=$?
go test ./...  || STATUS=$?

exit $STATUS