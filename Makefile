SHELL := /bin/bash

.PHONY: generate

default: generate

generate:
	@echo -n "Removing old openapi directory..."
	@if rm -rf openapi; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi
	
	@echo -n "Running 'openapi-generator generate'..."
	@openapi-generator generate \
		--input-spec api/openapi-spec.yaml \
		--git-host github.com \
		--git-repo-id hypixel-go-sdk \
		--git-user-id DarkenedBright \
		--generator-name go \
		--type-mappings=integer=int64,number=float64 \
		--additional-properties=useOneOfDiscriminatorLookup=true,withGoMod=false \
		--inline-schema-options RESOLVE_INLINE_ENUMS=true \
		--output openapi \
		|| { echo " FAILED"; exit 1; }

	@echo -n "Removing generated test files..."
	@if rm -r openapi/test; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

	@echo -n "Running 'go mod tidy'..."
	@if go mod tidy; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi

	@echo -n "Running 'go fmt'..."
	@if go fmt ./...; then \
		echo " SUCCESS"; \
	else \
		echo " FAILED"; \
		exit 1; \
	fi
