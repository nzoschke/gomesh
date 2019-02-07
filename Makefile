CMDS = $(shell find cmd/* -mindepth 1 -type d)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

clean:
	rm -rf bin/linux_amd64/*

dc-build:
	docker-compose -f config/docker/compose-gateway.yaml -f config/docker/compose-mesh.yaml -f config/docker/compose-service.yaml --project-directory . build

dc-down:
	docker-compose -f config/docker/compose-gateway.yaml -f config/docker/compose-mesh.yaml -f config/docker/compose-service.yaml --project-directory . down

dc-up-gateway:
	make -j bins
	docker-compose -f config/docker/compose-gateway.yaml --project-directory . up --abort-on-container-exit

dc-up-mesh:
	make -j bins
	docker-compose -f config/docker/compose-mesh.yaml --project-directory . up --abort-on-container-exit

dc-up-service:
	make -j bins
	docker-compose -f config/docker/compose-service.yaml --project-directory . up --abort-on-container-exit

workflow:
	docker build . -f .github/action/go/Dockerfile        -t action-go
	docker build . -f .github/action/yamllint/Dockerfile  -t action-yamllint
	docker build . -f .github/action/prototool/Dockerfile -t action-prototool

	docker run -v $(PWD):/github/workspace action-go        .github/golint.sh
	docker run -v $(PWD):/github/workspace action-yamllint  .github/yamllint.sh
	docker run -v $(PWD):/github/workspace action-prototool .github/pbpush.sh
