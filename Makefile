PROJECT_NAME = llm-stack
COMPOSE_FILE = infra/compose/docker-compose.yml
BAKE_FILE = infra/compose/docker-bake.hcl
COMPOSE_CMD = docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME)

up:
	$(COMPOSE_CMD) up --build -d

down:
	$(COMPOSE_CMD) down -v

logs:
	$(COMPOSE_CMD) logs -f

build:
	docker buildx bake -f $(BAKE_FILE)
	
restart:
	$(MAKE) down
	docker buildx bake -f $(BAKE_FILE)
	$(MAKE) up

clean-restart:
	docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down -v --remove-orphans
	docker buildx bake --no-cache -f $(BAKE_FILE)
	docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d