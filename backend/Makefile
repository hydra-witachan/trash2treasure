include .env

APP_NAME := trash2treasure
MIGRATIONS_DIR := ./migrations

DBMATE_URL := mysql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=${DB_TIMEZONE}
DBMATE_CMD_PREFIX := dbmate --migrations-dir '${MIGRATIONS_DIR}'
DBMATE_CMD_WITH_URL_PREFIX := dbmate --url '${DBMATE_URL}' --migrations-dir '${MIGRATIONS_DIR}'

.PHONY: build
build:
	go build -v -o ./build/${APP_NAME} ./internal/*.go

.PHONY: start
start:
	./build/${APP_NAME}

.PHONY: migration-new
migration-new:
	${DBMATE_CMD_PREFIX} new ${name}

.PHONY: migration-status
migration-status:
	${DBMATE_CMD_WITH_URL_PREFIX} status

.PHONY: migration-up
migration-up:
	${DBMATE_CMD_WITH_URL_PREFIX} up

.PHONY: migration-down
migration-down:
	${DBMATE_CMD_WITH_URL_PREFIX} down