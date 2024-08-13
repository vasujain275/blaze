# Makefile

# Variables
BINARY_NAME=blaze
CMD_DIR=cmd/blaze
SOURCE=$(CMD_DIR)/main.go

# Default target: Build the CLI executable
build:
	@echo "Building the CLI executable..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(SOURCE)
	@echo "Build completed."

# Run the CLI with a sample command
run: build
	@echo "Running the CLI..."
	./$(BINARY_NAME) parse /path/to/file.torrent

# Clean up the build
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Clean up completed."

.PHONY: build run clean
