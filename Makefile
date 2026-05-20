.DEFAULT_GOAL := build 
.PHONY:fmt vet run
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
