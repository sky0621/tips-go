# Makefile for running sqldef (mysqldef) against the local MySQL container

DOCKER_COMPOSE ?= docker compose
DOCKER_COMPOSE_FILE ?= docker-compose.yml

SQLDEF_CMD ?= mysqldef
SQLDEF_SCHEMA ?= schema.sql
SQLDEF_DRY_RUN ?= false
SQLDEF_EXTRA_FLAGS ?=

SQLDEF_MYSQL_HOST ?= 127.0.0.1
SQLDEF_MYSQL_PORT ?= 63416
SQLDEF_MYSQL_USER ?= localuser
SQLDEF_MYSQL_PASSWORD ?= yuckyjuice
SQLDEF_MYSQL_DATABASE ?= db-for-datamodeling-study

ifeq ($(SQLDEF_DRY_RUN),true)
SQLDEF_EXTRA_FLAGS += --dry-run
endif

.PHONY: mysql-start mysql-stop sqldef

mysql-start:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d

mysql-stop:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

sqldef:
	$(SQLDEF_CMD) \
		--host $(SQLDEF_MYSQL_HOST) \
		--port $(SQLDEF_MYSQL_PORT) \
		--user $(SQLDEF_MYSQL_USER) \
		--password $(SQLDEF_MYSQL_PASSWORD) \
		$(SQLDEF_EXTRA_FLAGS) \
		--file $(SQLDEF_SCHEMA) \
		$(SQLDEF_MYSQL_DATABASE)
