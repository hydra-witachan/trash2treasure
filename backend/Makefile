APP_NAME := trash2treasure
SOURCE_DIR := ./backend/
BUILD_DIR := ${SOURCE_DIR}/build

.PHONY: build
build:
	go build -v -o ${BUILD_DIR}/${APP_NAME} ${SOURCE_DIR}/*.go

.PHONY: start
start:
	${BUILD_DIR}/${APP_NAME}