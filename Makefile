.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
DATABASE_URL=postgres://postgres:password@s3-postgres:5432/postgres?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./api/cmd/friends/main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	docker run -v $(PWD)/api/data/migrations:/migrations --network dev-network pg-migrate migrate -path=/migrations -database "$(DATABASE_URL)" up

migrate.down:
	docker run -v $(PWD)/api/data/migrations:/migrations --network dev-network pg-migrate migrate -path=/migrations -database "$(DATABASE_URL)" down --all

docker.run: docker.network docker.postgres docker.chi docker.build.migrate

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.chi.build:
	docker build -t chi .

docker.chi: docker.chi.build
	docker run --rm -d \
		--name s3-chi \
		--network dev-network \
		-p 5000:5000 \
		chi

docker.postgres:
	docker run --rm -d \
		--name s3-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${PWD}/pg_data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.build.migrate:
	docker build -t pg-migrate -f Dockerfile.migrate . --no-cache

docker.stop: docker.stop.chi docker.stop.postgres

docker.stop.chi:
	docker stop s3-chi

docker.stop.postgres:
	docker stop s3-postgres