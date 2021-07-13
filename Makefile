# Include .env if exists
-include .env

# -----------------------------------------------------------------------------
#                                 General
# -----------------------------------------------------------------------------

# Git related
COMMIT_HASH = $(shell git rev-parse --short HEAD)
TAG         = $(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD)

# GO related
GO      ?= go
LDFLAGS = -X "main.CommitHash=$(COMMIT_HASH)" -X "main.Tag=$(TAG)"

# Database related
MYSQL_CONN     = $(DB_USER):$(DB_PASS)@tcp($(DB_HOST))/$(DB_NAME)?parseTime=true
MIGRATIONS_DIR = resources/database/migrations

# -----------------------------------------------------------------------------
#                                 Migrations
# -----------------------------------------------------------------------------

# Dump the migration status for the current DB
# Usage: make migrations-status
migrations-status:
	@goose -dir "$(MIGRATIONS_DIR)" mysql "$(MYSQL_CONN)" status

# Creates new migration file in sequential order
# Usage: make migrations-create NAME=add_new_column
migrations-create:
	@goose -dir "$(MIGRATIONS_DIR)" create ${NAME} sql
	@goose -dir "$(MIGRATIONS_DIR)" mysql "$(MYSQL_CONN)" fix

# Migrate the DB to the most recent version available
# Usage: make migrations-up
migrations-up:
	@goose -dir "$(MIGRATIONS_DIR)" mysql "$(MYSQL_CONN)" up

# Migrate the DB up by 1
# Usage: make migrations-up-by-one
migrations-up-by-one:
	@goose -dir "$(MIGRATIONS_DIR)" mysql "$(MYSQL_CONN)" up-by-one

# Roll back the version by 1
# Usage: make migrations-down
migrations-down:
	@goose -dir "$(MIGRATIONS_DIR)" mysql "$(MYSQL_CONN)" down

# -----------------------------------------------------------------------------
#                                 Mocks
# -----------------------------------------------------------------------------

# Generate mocks for all interfaces
generate-mocks:
	$(GO) generate ./...

# -----------------------------------------------------------------------------
#                                 Tests
# -----------------------------------------------------------------------------

# Run tests
tests:
	@$(GO) test -count=1 -race -v -tags=unit --parallel 10 ./... -coverprofile=testdata/coverage.out

# View test coverage in browser
view-coverage:
	@$(GO) tool cover -html=testdata/coverage.out

# -----------------------------------------------------------------------------
#                                 Build
# -----------------------------------------------------------------------------

# Build api binary
build:
	$(GO) build -ldflags '$(LDFLAGS)' -o ./bin/main cmd/api/main.go
