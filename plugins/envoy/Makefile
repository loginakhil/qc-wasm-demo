.PHONY: start-server-a
start-server-a:
	python3 server.py 9091 "Hello AAA!"

.PHONY: start-server-b
start-server-b:
	python3 server.py 9092 "Hello BBB!"

.PHONY: start-envoy-proxy
start-envoy-proxy:
	envoy --config-path ./config.yml

.PHONY: build-envoy-plugin
build-envoy-plugin:
	tinygo build -o plugin.wasm -scheduler=none -target=wasi ./plugin.go

.PHONY: start-envoy-plugin-proxy
start-envoy-plugin-proxy:
	envoy --config-path ./config.wasm.yml

.PHONY: test-valid-pal-auth
test-valid-pal-auth:
	curl -i -H "x-pal-auth: racecar" http://localhost:8080/a

.PHONY: test-invalid-pal-auth
test-invalid-pal-auth:
	curl -i -H "x-pal-auth: sun" http://localhost:8080/a