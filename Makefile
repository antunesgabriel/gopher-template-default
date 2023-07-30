build:
	docker compose up --build

up:
	docker compose up -d

run-server:
	go run ./cmd/server
