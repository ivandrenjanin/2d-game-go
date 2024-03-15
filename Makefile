.PHONY: build run

BUILD_DIR = ./bin

build:
	CGO_ENABLED=1 go build -ldflags="-w -s" -o $(BUILD_DIR)/2d-game-go ./main.go

run: build
	$(BUILD_DIR)/2d-game-go
