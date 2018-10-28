CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

configs/sidecar.yaml: cmd/envoy-cfg/main.go
	go run $< /tmp/sidecar.yaml
	docker run \
		-v/tmp/sidecar.yaml:/tmp/sidecar.yaml \
		envoyproxy/envoy:latest \
		envoy -c /tmp/sidecar.yaml --mode validate
	# mv /tmp/sidecar.yaml configs/sidecar.yaml

compose-mesh: generate bins
	docker-compose -f config/docker/compose-mesh.yaml build
	docker-compose -f config/docker/compose-mesh.yaml up

compose-proxy: generate bins
	docker-compose -f config/docker/compose-proxy.yaml -p gomesh build
	docker-compose -f config/docker/compose-proxy.yaml -p gomesh up --abort-on-container-exit

generate:
	prototool generate
	bin/pbtool.sh

setup:
	go get -u github.com/golang/protobuf/protoc-gen-go
