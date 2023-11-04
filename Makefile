include .env
export

MIGRATION_DIR=./migrations
MIGRATION_DSN="host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=$(POSTGRES_SSLMODE)"

define set_protocs
    TEMP_DIR=temp-$(1); \
    if [ ! -d "vendor-proto/$(4)" ]; then \
        git clone $(2) $$TEMP_DIR && \
        mkdir -p vendor-proto/$(4) && \
        cp -r $$TEMP_DIR/$(3)/* vendor-proto/$(4)/; \
    fi
endef

.PHONY: generate
generate:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 --proto_path vendor-proto \
		--go_out=pkg/note_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg/note_v1 --grpc-gateway_opt=paths=source_relative \
		--grpc-gateway_opt=logtostderr=true \
		--validate_out=lang=go,paths=source_relative:pkg/note_v1 \
		--swagger_out=allow_merge=true,merge_file_name=api:pkg/note_v1 \
		api/note_v1/note.proto

.PHONY: local-migration-status
local-migration-status:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir $(MIGRATION_DIR) postgres $(MIGRATION_DSN) down -v

.PHONY: vendor-proto
vendor-proto:
	$(call set_protocs,google/api,https://github.com/googleapis/googleapis,google/api,google/api)
	$(call set_protocs,envoyproxy/validate,https://github.com/envoyproxy/protoc-gen-validate,validate,envoyproxy/validate)
	$(call set_protocs,google/protobuf,https://github.com/protocolbuffers/protobuf,src/google/protobuf,google/protobuf)
	@rm -rf temp-*

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
