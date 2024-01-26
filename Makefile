BUILD_FOLDER_NAME=build
BUILD_EXECUTABLE_NAME=server

default: run

.PHONY: run
run: build
	./$(BUILD_FOLDER_NAME)/$(BUILD_EXECUTABLE_NAME)

.PHONY: build
build:
	CGO_ENABLED=1 go build -C cmd/server -o ../../$(BUILD_FOLDER_NAME)/$(BUILD_EXECUTABLE_NAME)

.PHONY: clean
clean:
	rm -r data/
	rm -r build/
