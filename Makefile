.DEFAULT_GOAL=build


.PHONY: build
build:
	go build -o task ./cmd/app

fmt:
	@golangci-lint run ./...

clear:
	@rm task