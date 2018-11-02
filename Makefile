CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

COMPOSE_CMD = docker-compose -p gomesh
COMPOSE_FILES = -f config/docker/compose-api.yaml \
	-f config/docker/compose-mesh.yaml \
	-f config/docker/compose-proxy.yaml

configs/sidecar.yaml: cmd/envoy-cfg/main.go
	go run $< /tmp/sidecar.yaml
	docker run \
		-v/tmp/sidecar.yaml:/tmp/sidecar.yaml \
		envoyproxy/envoy:latest \
		envoy -c /tmp/sidecar.yaml --mode validate
	# mv /tmp/sidecar.yaml configs/sidecar.yaml

compose-build:
	$(COMPOSE_CMD) $(COMPOSE_FILES) build

compose-down:
	$(COMPOSE_CMD) $(COMPOSE_FILES) down

compose-api:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-api.yaml   up --abort-on-container-exit

compose-mesh:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-mesh.yaml  up --abort-on-container-exit

compose-proxy:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-proxy.yaml up --abort-on-container-exit

generate:
	docker build -f config/docker/Dockerfile-prototool -t prototool .
	docker run -v $(PWD):/in prototool /bin/prototool.sh
	# FIXME: add to prototool package
	find gen/go/ -name 'mock*' | xargs rm
	mockery -all -dir gen/go -inpkg

.PHONY: vendor
vendor:
	git remote add -f -t master --no-tags protoc-gen-validate https://github.com/lyft/protoc-gen-validate.git || true
	git remote add -f -t master --no-tags grpc-gateway        https://github.com/grpc-ecosystem/grpc-gateway  || true
	git rm -rf vendor/
	git read-tree --prefix=vendor/github.com/lyft/protoc-gen-validate/validate/ -u protoc-gen-validate/master:validate
	git read-tree --prefix=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party -u grpc-gateway/master:third_party
