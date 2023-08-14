# https://workers.wasmlabs.dev/
.PHONY: build
build:
	tinygo build -target wasi -o worker.wasm

.PHONY: deps
deps:
	curl -fsSL https://workers.wasmlabs.dev/install | bash -s -- --local

.PHONY: run
run:
	./wws .

.PHONY: test
test:
	 curl http://127.0.0.1:8080/worker