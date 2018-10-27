CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

configs/sidecar.yaml: cmd/envoy-cfg/main.go
	go run $< /tmp/sidecar.yaml
	docker run \
		-v/tmp/sidecar.yaml:/tmp/sidecar.yaml \
		envoyproxy/envoy:latest \
		envoy -c /tmp/sidecar.yaml --mode validate
	mv /tmp/sidecar.yaml configs/sidecar.yaml

compose-discovery: bins
	docker-compose -f config/docker/compose-discovery.yaml -p gomesh build
	docker-compose -f config/docker/compose-discovery.yaml -p gomesh up

setup:
	go get -u github.com/golang/protobuf/protoc-gen-go
