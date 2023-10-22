.PHONY: build clean test fmt lint generate

build:
	go build -o github-project-automation cmd/main.go

clean:
	go clean
	rm -f github-project-automation

test:
	go test ./...

gen:
	go generate ./...

