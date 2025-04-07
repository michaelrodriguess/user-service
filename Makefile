APP_NAME := user-service
CMD_DIR := cmd/server

init:
	go mod init github.com/michaelrodriguess/$(APP_NAME)

tidy:
	go mod tidy

deps: tidy

build:
	go build -o $(APP_NAME) ./$(CMD_DIR)

run:
	go run ./$(CMD_DIR)/main.go

clean:
	rm -f $(APP_NAME)

reload: tidy run

.PHONY: init tidy build run clean deps
