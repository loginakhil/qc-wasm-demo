GO=tinygo build -target wasi

.PHONY: build
build:
	@mkdir -p "./builds"
	$(GO) -o ./builds/hello.wasm

.PHONY: node
node:
	npm install wasi
	node  --experimental-wasi-unstable-preview1 --no-warnings wasi-node.js

.PHONY: python
python:
	python -m pip install wasmtime
	python wasi-python.py

.PHONY: wasmedge
wasmedge:
	wasmedge builds/hello.wasm