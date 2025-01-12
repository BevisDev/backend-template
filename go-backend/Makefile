# name app build
APP_NAME = gobe

# main file
MAIN_DIR = src/main
MAIN_FILE = $(MAIN_DIR)/main.go

# default build if not args in make
all: build

# Install and generate swag
swag_install:
	@echo "Install swag"
	@go install github.com/swaggo/swag/cmd/swag@latest

swag_gen:
	@echo "Generate swag"
	@swag init --generalInfo ./src/main/startup/run.go --output ./src/resources/swagger

# Build the application
build: 
	@echo "Building the application..."
	@go build -v -o $(APP_NAME) $(MAIN_FILE)

# run the application
run: 
	@echo "Running the application..."
	@go run $(MAIN_FILE)