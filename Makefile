THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build up start down logs ps db
help:
	make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
r:
	docker compose down && docker compose -f compose.yml up --build -d && docker compose logs --follow
restart:
	docker compose down && docker compose -f compose.yml up --build -d && docker compose logs --follow
b:
	docker compose -f compose.yml build $(c)
build:
	docker compose -f compose.yml build $(c)
up:
	docker compose -f compose.yml up -d $(c)
d:
	docker compose -f compose.yml down $(c)
down:
	docker compose -f compose.yml down $(c)
logs:
	docker compose -f compose.yml logs --tail=100 -f $(c)
ps:
	docker compose -f compose.yml ps

