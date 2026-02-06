GO := go

DIST_DIR := build
BIN_DIRS := cmd/nmb

BINS := $(notdir $(BIN_DIRS))
OUTPUTS := $(addprefix $(DIST_DIR)/, $(BINS))

BUILD_FLAGS := -v -race
GOBUILD := $(GO) build $(BUILD_FLAGS)

_default: build

build: $(OUTPUTS)

$(DIST_DIR)/%: cmd/%/main.go
	@mkdir -p $(DIST_DIR)
	$(GOBUILD) -o "$@" "./cmd/$*/"

run:
ifndef BIN
	$(error You must specify a binary with BIN=<name>. Available: $(BINS))
endif
	@if [ ! -f "$(DIST_DIR)/$(BIN)" ]; then \
		echo "Binary $(DIST_DIR)/$(BIN) not found. Building it..."; \
		$(MAKE) $(DIST_DIR)/$(BIN); \
	fi
	@echo "Running $(BIN)..."
	@"./$(DIST_DIR)/$(BIN)"

clean:
	rm -rf $(DIST_DIR)

.PHONY: _default build run clean
