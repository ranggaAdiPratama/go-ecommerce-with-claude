DB_URL ?= postgres://rangga:mitsuha@localhost:5432/go?sslmode=disable

migrate-down:
	migrate -path ./migrations -database "${DB_URL}" down

migrate-up:
	migrate -path ./migrations -database "${DB_URL}" up

serve:
	go run ./cmd/api/main.go

sqlc:
	docker run --rm -v "/mnt/f/API/go/e_commerce:/src" -w /src sqlc/sqlc generate

tidy:
	go mod tidy

.PHONY: migrate-down migrate-up serve sqlc tidy