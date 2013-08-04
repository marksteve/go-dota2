all: format

format:
	go fmt github.com/marksteve/go-dota2
	go fmt github.com/marksteve/go-dota2/dota2

example: format
	go run example.go

@PHONY: format
