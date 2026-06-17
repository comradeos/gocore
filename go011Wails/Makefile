APP_NAME := go011Wails
GO_BIN := $(shell go env GOPATH)/bin
WAILS := $(or $(shell command -v wails 2>/dev/null),$(GO_BIN)/wails)

.PHONY: help deps doctor dev build clean

help:
	@echo "Available targets:"
	@echo "  make deps    - install frontend dependencies"
	@echo "  make doctor  - run Wails environment checks"
	@echo "  make dev     - start the app in development mode"
	@echo "  make build   - build the desktop app"
	@echo "  make clean   - remove the packaged binary directory"

deps:
	cd frontend && npm install

doctor:
	"$(WAILS)" doctor

dev:
	"$(WAILS)" dev

build:
	"$(WAILS)" build -clean

clean:
	rm -rf build/bin
