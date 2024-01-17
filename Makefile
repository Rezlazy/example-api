BUILD_ENVPARMS:=CGO_ENABLED=0

BIN_DIR=`pwd`/bin
INTERNAL_PATH=$(CURDIR)/internal

MINIMOCK_BIN:=$(BIN_DIR)/minimock
GOOSE_BIN:=$(BIN_DIR)/goose

MIGRATIONS_DIR=migrations
BASE_GOOSE_COMMAND=$(GOOSE_BIN) -dir $(MIGRATIONS_DIR) postgres "postgresql://postgres:postgres@localhost:5432/"

OPENAPI_SPEC_PATH=./api/server/openapi.yaml
AUTOGEN_SERVER_DIR=./internal/api/http/server

OPENAPI_SPEC_CLIENT_PATH=./api/client/
AUTOGEN_CLIENT_DIR=./internal/client/

install_dependencies:
	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@v3.15.1
	GOBIN=$(BIN_DIR) go install github.com/ogen-go/ogen/cmd/ogen@v0.77.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.1.3

generate_openapi_server: install_dependencies
	PKG_NAME=`basename $(AUTOGEN_SERVER_DIR)`; \
	$(BIN_DIR)/ogen \
	--target $(AUTOGEN_SERVER_DIR) \
	-package $$PKG_NAME \
	--no-client --skip-unimplemented --allow-remote \
	$(OPENAPI_SPEC_PATH);

migration_create:
	$(BASE_GOOSE_COMMAND) create "$(name)" sql

migration_up:
	$(BASE_GOOSE_COMMAND) up

stop_dev:
	docker-compose down

run_dev: stop_dev
	docker-compose up


download_spec_client:
	curl "https://raw.githubusercontent.com/thatapicompany/apis/main/theCatAPI.com/thecatapi-oas.yaml" -o ${OPENAPI_SPEC_CLIENT_PATH}/cats/openapi.yaml

generate_openapi_client: install_dependencies
	$(BIN_DIR)/ogen -no-server -no-webhook-client --debug.noerr -no-webhook-server --skip-unimplemented --allow-remote \
	  -package dadata -target ./internal/client/cats/ ${OPENAPI_SPEC_CLIENT_PATH}/cats/openapi.yaml;