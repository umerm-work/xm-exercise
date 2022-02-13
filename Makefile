.PHONY: all
all: deps build

.PHONY: deps
deps:
	go mod tidy
	go mod download
	go mod vendor

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o bin/application .

.PHONY: run
run:
	go run main.go

