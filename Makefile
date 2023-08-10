.PHONY:deps
deps:
	sudo dnf -y install wasmedge golang vim crun-wasm tinygo nodejs binaryen bat
	python -m ensurepip && python -m pip install wasmtime