# Makefile for t-cli

BINARY_NAME=t-cli

# Build the binary
build:
	@echo "Building..."
	go build -o ${BINARY_NAME} main.go
	@echo "Built as: ${BINARY_NAME}"

# Install to /usr/local/bin (Mac/Linux only)
install: build
	@echo "Installing to /usr/local/bin..."
	sudo mv ${BINARY_NAME} /usr/local/bin/
	@echo "Done! Type '${BINARY_NAME}' to run."

# Remove the binary
clean:
	@echo "Cleaning..."
	go clean
	rm -f ${BINARY_NAME}

# Run quickly (dev mode)
run:
	go run main.go

.PHONY: build install clean run
