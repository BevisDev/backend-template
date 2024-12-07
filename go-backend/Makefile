# name app build
APP_BUILD = go-backend

# main file
MAIN_DIR = src/main
MAIN_FILE = $(MAIN_DIR)/main.go

# Build the application
all: build

build: 
	@echo "Building..."
	@go build -o main $(MAIN_FILE)

# format the source code
fmt: 
	@echo "Formatting source code..."
	go fmt ./...
	
# run the application
run: 
	@echo "Running the application..."
	go run $(MAIN_FILE)