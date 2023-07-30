build:
	docker compose up --build

up:
	docker compose up -d

run:
	go run ./cmd/server
