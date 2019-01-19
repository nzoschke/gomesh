CMDS = $(shell find cmd/* -mindepth 1 -type d)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

clean:
	rm -rf bin/linux_amd64/*

dc-build:
	docker-compose -f config/docker/compose-gateway.yaml -f config/docker/compose-mesh.yaml --project-directory . build

dc-down:
	docker-compose -f config/docker/compose-gateway.yaml -f config/docker/compose-mesh.yaml --project-directory . down

dc-up-gateway:
	make -j bins
	docker-compose -f config/docker/compose-gateway.yaml --project-directory . up --abort-on-container-exit

dc-up-mesh:
	make -j bins
	docker-compose -f config/docker/compose-mesh.yaml --project-directory . up --abort-on-container-exit

workflow:
	docker build . -f .github/action/make/Dockerfile -t make
	docker run -v $(PWD):/github/workspace make bins
	docker build . -f .github/action/yamllint/Dockerfile -t yamllint
	docker run -v $(PWD):/github/workspace yamllint -c /etc/yamllint.yaml config/*/*.yaml
	#docker build . -f .github/action/yamllint/Dockerfile -t yamllint