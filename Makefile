CMDS = $(shell find cmd/{client,server} -type d -mindepth 1)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

clean:
	rm bin/linux_amd64/*

dc-up-gateway:
	make -j bins
	docker-compose -f config/docker/compose-mesh.yaml --project-directory . up

dc-up-mesh:
	make -j bins
	docker-compose -f config/docker/compose-mesh.yaml --project-directory . up
