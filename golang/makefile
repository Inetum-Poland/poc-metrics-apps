# Make makefile tolerant to spaces
.RECIPEPREFIX := $() $()

.PHONY: all
all: utils build run

.PHONY: utils
utils:
  go get -u -m all && \
  go mod tidy && \
  go fmt ./... && \
  go vet ./...

.PHONY: build
build: utils
  go build \
		-o bin/app

.PHONY: run
run: utils
  go run main.go

.PHONY: run.w
run.w: utils
  gow -l -e go,mod run main.go run

.PHONY: test
test: utils
  go test ./...
