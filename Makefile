BINARY_NAME = goenvlist
BINARY_UNIX = $(BINARY_NAME)_unix
MAIN_PATH = ./main.go

.PHONY: all build clean help

all: build

build:
	go build -o $(BINARY_NAME) -v $(MAIN_PATH)

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

help:
	@echo "make - Build the binary"
	@echo "make build - Build the binary"
	@echo "make clean - Remove the binary files"
