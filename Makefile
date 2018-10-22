CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

build: bins
	docker-compose build

dev: build
	docker-compose up

setup:
	go get -u github.com/golang/protobuf/protoc-gen-go
