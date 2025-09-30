# Makefile for managing local sqldef sample environment

DOCKER_COMPOSE ?= docker compose
DOCKER_COMPOSE_FILE ?= docker-compose.yml

MYSQL_CONTAINER_NAME ?= sqldef-mysql
MYSQL_IMAGE ?= mysql:8.0
MYSQL_HOST ?= 127.0.0.1
MYSQL_PORT ?= 3306
MYSQL_ROOT_PASSWORD ?= secret
MYSQL_DATABASE ?= app_db

SQDEF_BIN ?= mysqldef
SCHEMA_FILE ?= sample/schema/schema.sql
GO_MODULE ?= ./sample

SQLDEF_MYSQL_HOST ?= $(MYSQL_HOST)
SQLDEF_MYSQL_PORT ?= $(MYSQL_PORT)
SQLDEF_MYSQL_USER ?= root
SQLDEF_MYSQL_PASSWORD ?= $(MYSQL_ROOT_PASSWORD)
SQLDEF_MYSQL_DATABASE ?= $(MYSQL_DATABASE)

export SQLDEF_MYSQL_HOST
export SQLDEF_MYSQL_PORT
export SQLDEF_MYSQL_USER
export SQLDEF_MYSQL_PASSWORD
export SQLDEF_MYSQL_DATABASE
export MYSQL_CONTAINER_NAME
export MYSQL_IMAGE
export MYSQL_PORT
export MYSQL_ROOT_PASSWORD
export MYSQL_DATABASE

SQLDEF_PASSWORD_OPT := $(if $(strip $(SQLDEF_MYSQL_PASSWORD)),--password=$(SQLDEF_MYSQL_PASSWORD),)

.PHONY: mysql-start mysql-stop install-sqldef sqldef-dry-run sqldef-run go-run

mysql-start:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d

mysql-stop:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

install-sqldef:
	go install github.com/k0kubun/sqldef/cmd/mysqldef@latest

sqldef-dry-run:
	$(SQDEF_BIN) --host=$(SQLDEF_MYSQL_HOST) --port=$(SQLDEF_MYSQL_PORT) --user=$(SQLDEF_MYSQL_USER) $(SQLDEF_PASSWORD_OPT) --dry-run $(SQLDEF_MYSQL_DATABASE) < $(SCHEMA_FILE)

sqldef-run:
	$(SQDEF_BIN) --host=$(SQLDEF_MYSQL_HOST) --port=$(SQLDEF_MYSQL_PORT) --user=$(SQLDEF_MYSQL_USER) $(SQLDEF_PASSWORD_OPT) $(SQLDEF_MYSQL_DATABASE) < $(SCHEMA_FILE)

go-run:
	go run $(GO_MODULE)
