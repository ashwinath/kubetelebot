commit=$(shell git rev-parse HEAD)
all: build push

.PHONY: test
test:
	@go clean -testcache
	@go test -cover -race ./...

.PHONY: build
build:
	docker build -t $(REGISTRY)/kubetelebot:$(commit) -t $(REGISTRY)/kubetelebot:latest .

.PHONY: push
push:
	docker push $(REGISTRY)/kubetelebot:$(commit)
	docker push $(REGISTRY)/kubetelebot:latest
