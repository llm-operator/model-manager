.PHONY: default
default: test

include common.mk

.PHONY: test
test: go-test-all

.PHONY: lint
lint: go-lint-all git-clean-check

.PHONY: generate
generate: buf-generate-all

.PHONY: build-server
build-server:
	go build -o ./bin/server ./server/cmd/

.PHONY: build-loader
build-loader:
	go build -o ./bin/loader ./loader/cmd/

.PHONY: build-docker-server
build-docker-server:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/model-manager-server:latest -f build/server/Dockerfile .

.PHONY: build-docker-loader
build-docker-loader:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/model-manager-loader:latest -f build/loader/Dockerfile .

.PHONY: build-docker-convert-gguf
build-docker-convert-gguf:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/experiments-convert_gguf:latest -f build/experiments/convert_gguf/Dockerfile build/experiments/convert_gguf/
