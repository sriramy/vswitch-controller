
BACKEND_GEN_DIR = backend

GO ?=go

BACKEND_SRC_DIRS = $(wildcard pkg/*)
BACKEND_ARCHIVES = $(patsubst pkg/%/%.go, $(BACKEND_GEN_DIR)/%.a, $(BACKEND_SRC_DIRS))

all: generate

$(BACKEND_GEN_DIR)/%.a: pkg/%/%.go
	@mkdir -p $(@D)
	$(GO) build -o $< -buildmode=c-archive $@/...

.PHONY: generate
generate: $(BACKEND_ARCHIVES)

.PHONY: clean
clean:
	rm -rf $(BACKEND_GEN_DIR)