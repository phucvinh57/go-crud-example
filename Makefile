# note: call scripts from /scripts

.PHONY: apidocs dev docker_dev clean sqlc migrate build

ENTRYPOINT := cmd/main.go
DOCKER_DEV := docker compose -f scripts/docker-compose.dev.yml
GEN_SWAGGER := swag init -d cmd -o api

bootstrap:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u
	$(DOCKER_DEV) up -d

apidocs:
	$(GEN_SWAGGER)

dev:
	$(GEN_SWAGGER) && go run $(ENTRYPOINT)

docker_dev:
	$(DOCKER_DEV) up -d

clean:
	$(DOCKER_DEV) down --volumes

sqlc:
	sqlc generate

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

build:
	go build ./cmd/main.go