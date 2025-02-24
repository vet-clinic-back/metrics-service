GOLANGCI_LINT = $(HOME)/bin/golangci-lint

.PHONY: update-swagger
update-swagger:
	swag init -g ./cmd/metrics-service/main.go -o docs


.PHONY: lint
lint:
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=golangci.yaml

.PHONY: lint-fast
lint-fast:
	$(GOLANGCI_LINT) run ./... --fast --config=golangci.yaml

.PHONY: env-test
env-test:
	export POSTGRES_HOST=localhost && \
	export POSTGRES_PORT=5432 && \
    export POSTGRES_USER=postgres && \
    export POSTGRES_PASSWORD=postgres && \
    export POSTGRES_DB=postgres && \
    echo "Environment variables set"