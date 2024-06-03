SERVICE=mc-insta-bot

DC=docker-compose -f docker-compose.yml -f docker-compose.override.yml

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build    - Build the Docker service"
	@echo "  make up       - Start the Docker service in detached mode (background)"
	@echo "  make down     - Stop the Docker service"
	@echo "  make logs     - Show logs of the Docker service"
	@echo "  make dev      - Start the Docker service in development mode with hot-reload"
	@echo "  make test     - Run Go tests inside the Docker container"
	@echo "  make clean    - Clean all containers, volumes, images, and orphaned containers"
	@echo "  make help     - Show this help message"

.PHONY: build
build:
	$(DC) build $(SERVICE)

.PHONY: up
up:
	$(DC) up -d $(SERVICE)

.PHONY: down
down:
	$(DC) down

.PHONY: logs
logs:
	$(DC) logs -f $(SERVICE)

.PHONY: dev
dev:
	$(DC) up --build

.PHONY: test
test:
	$(DC) run $(SERVICE) go test ./...

.PHONY: clean
clean:
	$(DC) down -v --rmi all --remove-orphans
