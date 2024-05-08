.DEFAULT_GOAL=build


.PHONY: build
build:
	go build -o comp-club ./cmd/app