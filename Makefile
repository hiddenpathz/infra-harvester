.PHONY: up down

# указываем на единый compose-файл
COMPOSE = docker compose -f docker-compose.yaml

# 1) Запустить ВСЁ (proxy + rabbit + kafka и т.д.)
up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down --remove-orphans