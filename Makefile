COMPOSE_FILE := deploy/compose.yaml
ENV_FILE := .env

ifneq ($(wildcard $(ENV_FILE)),)
    include $(ENV_FILE)
    export
endif

PROJECT_NAME := solace
COMPOSE := docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) --env-file $(ENV_FILE)

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) restart

logs:
	$(COMPOSE) logs -f
