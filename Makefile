all: build
build:
	CGO_ENABLED=0 go build -o bin/xlsx2itop cmd/xlsx2itop/main.go
clean:
	rm -fr bin