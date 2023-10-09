include .env
export

MIGRATION_DIR=./migrations
MIGRATION_DSN="host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=$(POSTGRES_SSLMODE)"

PHONY: generate
generate:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
		--go_out=pkg/note_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
		api/note_v1/note.proto

PHONY: local-migration-status
local-migration-status:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) status -v

PHONY: local-migration-up
local-migration-up:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) up -v

PHONY: local-migration-down
local-migration-down:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) down -v