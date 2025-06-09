# Creating variables for dir and name of binary
BINARY_NAME_base=numbers-to-words

BIN_DIR=bin

# Creating the required binary name for windows and other Operating Systems
ifeq ($(GOOS),windows)
	BINARY_NAME := $(BINARY_NAME_BASE).exe
else
	BINARY_NAME := $(BINARY_NAME_BASE)
endif

# Setting default make
all: test build

# Building binary
build:
	@echo "Building binary..."
	@go build -o $(BIN_DIR)/$(BINARY_NAME) ./src

# Running unit tests
test:
	@echo "Running tests..."
	@go test -v ./src/...

# Cleaning up the binary
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# Preventing conflicts
.PHONY: all build test clean