GO_BIN := $(shell go env GOPATH)/bin
WAILS_FROM_PATH := $(shell command -v wails 2>/dev/null)
WAILS_FROM_GOBIN := $(GO_BIN)/wails

ifeq ($(strip $(WAILS_FROM_PATH)),)
  ifeq ($(wildcard $(WAILS_FROM_GOBIN)),)
    WAILS :=
  else
    WAILS := $(WAILS_FROM_GOBIN)
  endif
else
  WAILS := $(WAILS_FROM_PATH)
endif

.PHONY: help check-wails deps doctor dev build clean

help:
	@echo "Available targets:"
	@echo "  make check-wails - verify that the Wails CLI is installed"
	@echo "  make deps    - install frontend dependencies"
	@echo "  make doctor  - run Wails environment checks"
	@echo "  make dev     - start the app in development mode"
	@echo "  make build   - build the desktop app"
	@echo "  make clean   - remove the packaged binary directory"

check-wails:
ifndef WAILS
	@echo "Wails CLI not found."
	@echo "Install it with:"
	@echo "  go install github.com/wailsapp/wails/v2/cmd/wails@latest"
	@echo "Then ensure your PATH includes: $(GO_BIN)"
	@exit 1
else
	@echo "Using Wails CLI: $(WAILS)"
endif

deps:
	cd frontend && npm install

doctor: check-wails
	"$(WAILS)" doctor

dev: check-wails
	"$(WAILS)" dev

build: check-wails
	"$(WAILS)" build -clean

clean:
	rm -rf build/bin
