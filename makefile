include .env

.PHONY: migration migrate-up migrate-down migrate-status run

MIGRATIONS_PATH = migrate/migrations

# Create a new migration file (Up + Down sections)
migration:
	@read -p "Enter migration name: " name; \
	ts=$$(date +%Y%m%d%H%M%S); \
	file="$(MIGRATIONS_PATH)/$${ts}-$$name.sql"; \
	mkdir -p $(MIGRATIONS_PATH); \
	touch $$file; \
	echo "-- +migrate Up" >> $$file; \
	echo "-- write your UP migration SQL here" >> $$file; \
	echo "-- +migrate Down" >> $$file; \
	echo "-- write your DOWN migration SQL here" >> $$file; \
	echo "✅ Created migration file: $$file"

# Apply all up migrations
migrate-up:
	@go run migrate/main.go up

# Roll back last migration
migrate-down:
	@go run migrate/main.go down

# Show current migration status
migrate-status:
	@go run migrate/main.go status

# Run the app using Air
run:
	@echo "🚀 Starting server with Air..."
	@air
