# Make makefile tolerant to spaces
.RECIPEPREFIX := $() $()

.PHONY: all-golang
all-golang: utils-golang build-golang run-golang

.PHONY: build-golang
build-golang: utils-golang
  docker build -t golang-app golang

.PHONY: run-golang
run: utils-golang
  docker run -p 8080:8080 metrics-go-app

.PHONY: utils-golang
utils-golang:
  cd ./golang && \
  go mod tidy && \
  go fmt ./...&& \
  go vet ./...

.PHONY: docker-compose
docker-compose: utils-golang
  docker-compose down
  docker-compose up --build
