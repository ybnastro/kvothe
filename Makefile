#!/bin/bash

build-app:
	@echo "Building core app..."
	@go build -o kvothe_app ./cmd/kvothe
	@echo "Done ✔"

clear-app:
	@echo "Removing core app..."
	rm -rf kvothe_app
	@echo "Done ✔"

build-cli:
	@echo "Building CLI app..."
	@go build -o kvothe_cli ./cmd/cli
	@echo "Done ✔"

clear-cli:
	@echo "Removing CLI app..."
	rm -rf kvothe_cli
	@echo "Done ✔"

build-mq:
	@echo "Building MQ app..."
	@go build -o kvothe_mq ./cmd/mq
	@echo "Done ✔"

clear-mq:
	@echo "Removing MQ app..."
	rm -rf kvothe_mq
	@echo "Done ✔"
