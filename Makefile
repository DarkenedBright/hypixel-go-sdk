SHELL := /bin/bash

.PHONY: generate

default: generate

generate:
	@echo -n "Running 'go mod tidy'..."
	@if go mod tidy; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

	@echo -n "Running 'openapi-generator generate'..."
	@openapi-generator generate -i ./swagger.json --git-repo-id hypixel-go-sdk --git-user-id DarkenedBright -g go -o . || { echo " FAILED"; exit 1; }
