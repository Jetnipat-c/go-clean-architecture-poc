BIN := $(shell pwd)/bin
INTERNAL_PATH = $(shell pwd)/internal/
MIGRATIONS_PATH = $(INTERNAL_PATH)/db/schema/migrations

include .env

.PHONY: migration-create
migration-create:
	@echo "Creating migration"
	@goose -dir $(MIGRATIONS_PATH) create $(name) sql

.PHONY: migration-up
migration-up:
	@echo "Running migration"
	@goose -dir $(MIGRATIONS_PATH) postgres "postgresql://$(DB_USER):$(DB_PASS)@127.0.0.1:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

.PHONY: migration-down
migration-down:
	@echo "Rolling back migration"
	@goose -dir $(MIGRATIONS_PATH) postgres "postgresql://$(DB_USER):$(DB_PASS)@127.0.0.1:$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

