# name app build
APP_NAME = go-backend

# main file
MAIN_DIR = src/main
MAIN_FILE = $(MAIN_DIR)/main.go

# default build if not args in make
all: build

# Build the application
build: 
	@echo "Building the application..."
	@go build -o $(APP_NAME) $(MAIN_FILE)

# run the application
run: 
	@echo "Running the application..."
	@go run $(MAIN_FILE)