COMPOSE_FILE := "deploy/compose.yaml"
PROJECT_NAME := "solace"
ENV_FILE := ".env"

COMPOSE_CMD := "docker compose -f " + COMPOSE_FILE + " -p " + PROJECT_NAME + " --env-file " + ENV_FILE

up:
    {{ COMPOSE_CMD }} up -d

down:
    {{ COMPOSE_CMD }} down

restart:
    {{ COMPOSE_CMD }} restart

logs:
    {{ COMPOSE_CMD }} logs -f

status:
    {{ COMPOSE_CMD }} ps
