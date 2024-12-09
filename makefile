# Make makefile tolerant to spaces
.RECIPEPREFIX := $() $()

# Docker
.PHONY: docker-compose-down
docker-compose-down:
  docker-compose down --remove-orphans

.PHONY: docker-compose-up
docker-compose-up: docker-compose-down
  docker-compose up --build --remove-orphans

.PHONY: docker-compose-up-d
docker-compose-up-d: docker-compose-down
  docker-compose up -d --build --remove-orphans
