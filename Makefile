NAME=ipleak
REPO=github.com/tutabeier/ipleak

BINARY=${NAME}
BINARY_SRC=$(REPO)/cmd/${NAME}

BUILD_DIR ?= build/
GO_LINKER_FLAGS=-ldflags="-s -w"

.PHONY: build

build:
	@mkdir -p ${BUILD_DIR}
	@echo "Building"
	@go build -o ${BUILD_DIR}/${BINARY} ${GO_LINKER_FLAGS} ${BINARY_SRC}

build-rasp:
	@mkdir -p ${BUILD_DIR}
	@echo "Building for RaspberryPI"
	@env GOOS=linux GOARCH=arm GOARM=5 go build -o ${BUILD_DIR}/${BINARY}rasp ${GO_LINKER_FLAGS} ${BINARY_SRC}
