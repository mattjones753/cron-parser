build:
	go build -o bin/cron-parser cmd/cron-parser/main.go

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

test:
	go test ./...

e2e_test: build
	scripts/e2e_test.sh bin/cron-parser

all: deps lint test build e2e_test
