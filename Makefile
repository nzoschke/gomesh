CMDS = $(shell find cmd/{client,server} -type d -mindepth 1)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

clean:
	rm -rf bin/linux_amd64/*

dc-build:
	docker-compose -f config/docker/compose-mesh.yaml -f config/docker/compose-proxy.yaml --project-directory . build

dc-down:
	docker-compose -f config/docker/compose-mesh.yaml -f config/docker/compose-proxy.yaml --project-directory . down

dc-up-mesh:
	make -j bins
	docker-compose -f config/docker/compose-mesh.yaml --project-directory . up --abort-on-container-exit

dc-up-proxy:
	make -j bins
	docker-compose -f config/docker/compose-proxy.yaml --project-directory . up --abort-on-container-exit
