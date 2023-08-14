# https://developer.fermyon.com/spin/index

.PHONY: deps
deps:
	curl -fsSL https://developer.fermyon.com/downloads/install.sh | bash

.PHONY: setup
setup:
	./spin new http-go src

.PHONY: build
build:
	./spin build --from src

.PHONY: up
up:
	./spin up --from src