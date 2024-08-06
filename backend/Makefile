COMPOSE=docker-compose
EXEC=$(COMPOSE) exec

# Run the backend
dev:
	$(COMPOSE) up

# Run the backend in the background
run:
	$(COMPOSE) up -d

# Stop the backend
stop:
	$(COMPOSE) down

# Run the tests

web:
	websocat -t ws://localhost:8080/ws