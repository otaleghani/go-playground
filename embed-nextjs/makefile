# Define directories
UI_DIR := ui
DIST_DIR := dist

# Define binary name
GO_BINARY := swms

# Define Go and Node.js build parameters
GO_CMD := go
GO_BUILD := $(GO_CMD) build -o
NPM_INSTALL := npm install
NEXT_BUILD := npx next build

# Default target
all: clean install-deps build

# Ensure the latest versions of Go and Node.js are installed
install-deps:
	@echo "Ensure the latest versions of Go and Node.js are installed."
	# Uncomment and adjust these lines for your OS
	# curl -sL https://golang.org/dl/ | grep -Eo 'https://[a-zA-Z0-9./?=_%:-]*linux-amd64.tar.gz' | head -1 | xargs curl -O
	# sudo tar -C /usr/local -xzf go*.tar.gz
	# export PATH=$PATH:/usr/local/go/bin
	# curl -fsSL https://deb.nodesource.com/setup_14.x | sudo -E bash -
	# sudo apt-get install -y nodejs

# Clean target
clean:
	rm -rf $(DIST_DIR)
	rm -f $(GO_BINARY)

# Build target
build: build-ui build-backend package

# Build Next.js frontend
build-ui:
	mkdir -p $(DIST_DIR)/ui
	cp -r $(UI_DIR)/* $(DIST_DIR)/ui/
	cd $(DIST_DIR)/ui && $(NPM_INSTALL)
	cd $(DIST_DIR)/ui && $(NEXT_BUILD)

# Build Go backend
build-backend:
	$(GO_BUILD) $(DIST_DIR)/$(GO_BINARY)

# Package everything neatly
package:
	@echo "Packaging complete."

# Run target (example to show how to run the backend binary)
run: build
	./$(DIST_DIR)/$(GO_BINARY)

.PHONY: all clean install-deps build build-ui build-backend package run
