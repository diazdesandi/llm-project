PROJECT_NAME = llm-stack
COMPOSE_FILE = infra/compose/docker-compose.yml
COMPOSE_CMD = docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME)

up:
	$(COMPOSE_CMD) up -d

down:
	$(COMPOSE_CMD) down -v

logs:
	$(COMPOSE_CMD) logs -f

restart:
	$(COMPOSE_CMD) down -v
	$(COMPOSE_CMD) up --build