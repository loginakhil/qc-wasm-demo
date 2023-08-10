DOCKER=podman
GO=tinygo build -target wasi

.PHONY: build
build:
	mkdir -p "./builds"
	$(GO) -o ./builds/hello.wasm
	$(DOCKER) build -t qc-wasm-demo --platform=wasi/wasm .

.PHONY: docker
docker:
	$(DOCKER) run --platform=wasi/wasm  localhost/qc-wasm-demo
