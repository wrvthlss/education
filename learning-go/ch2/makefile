.DEFAULT_GOAL := build

.PHONY:fmt vet build run all
fmt:
	go fmt const.go

vet:
	go vet const.go

build:
	go build const.go

run:
	./const

all: build run

clean:
	rm -f const
	go clean -cache -testcache -modcache