.PHONY: clean build
APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) server.go

docker.run: docker.network docker.postgres docker.server

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.server.build:
	docker build -t pf-server .

docker.server: docker.server.build
	docker run --rm -d \
		--name pf-server \
		--network dev-network \
		-p 8080:8080 \
		pf-server

docker.postgres:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres -c log_statement=all -c log_duration=true

docker.stop: docker.stop.server docker.stop.postgres

docker.stop.server:
	docker stop pf-server

docker.stop.postgres:
	docker stop dev-postgres

sqlc.generate:
	go run github.com/kyleconroy/sqlc/cmd/sqlc@latest generate
