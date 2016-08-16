
.PHONY:	run clean format build

build:	clean	format
	@go build

run:	build
	@./numbers

clean:
	@go clean

format:
	@go fmt

default: run
