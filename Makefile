all: build
build: generate
	go build main.go
generate: tidy
	buf generate
tidy:
	go mod tidy
