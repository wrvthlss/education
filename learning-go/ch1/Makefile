.DEFAULT_GOAL := build

.PHONY:fmt vet build run all
fmt:
	go fmt hello.go

vet:
	go vet hello.go

build:
	go build hello.go

run:
	./hello

all: build run

clean:
	rm -f hello
	go clean -cache -testcache -modcache