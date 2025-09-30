# Makefile for managing local sqldef sample environment

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

SQLDEF_PASSWORD_OPT := $(if $(strip $(SQLDEF_MYSQL_PASSWORD)),--password=$(SQLDEF_MYSQL_PASSWORD),)

.PHONY: mysql-start mysql-stop install-sqldef sqldef-dry-run sqldef-run go-run

mysql-start:
	@if docker ps -a --format '{{.Names}}' | grep -q '^$(MYSQL_CONTAINER_NAME)$$'; then \
		echo "$(MYSQL_CONTAINER_NAME) container already exists. Run 'make mysql-stop' first if you need a fresh instance."; \
		exit 1; \
	fi
	docker run -d --name $(MYSQL_CONTAINER_NAME) \
		-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
		-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
		-p $(MYSQL_PORT):3306 \
		$(MYSQL_IMAGE) \
		--default-authentication-plugin=mysql_native_password \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci

mysql-stop:
	@if docker ps -a --format '{{.Names}}' | grep -q '^$(MYSQL_CONTAINER_NAME)$$'; then \
		docker stop $(MYSQL_CONTAINER_NAME) >/dev/null; \
		docker rm $(MYSQL_CONTAINER_NAME) >/dev/null; \
	else \
		echo "$(MYSQL_CONTAINER_NAME) container not found"; \
	fi

install-sqldef:
	go install github.com/k0kubun/sqldef/cmd/mysqldef@latest

sqldef-dry-run:
	$(SQDEF_BIN) --host=$(SQLDEF_MYSQL_HOST) --port=$(SQLDEF_MYSQL_PORT) --user=$(SQLDEF_MYSQL_USER) $(SQLDEF_PASSWORD_OPT) --dry-run $(SQLDEF_MYSQL_DATABASE) < $(SCHEMA_FILE)

sqldef-run:
	$(SQDEF_BIN) --host=$(SQLDEF_MYSQL_HOST) --port=$(SQLDEF_MYSQL_PORT) --user=$(SQLDEF_MYSQL_USER) $(SQLDEF_PASSWORD_OPT) $(SQLDEF_MYSQL_DATABASE) < $(SCHEMA_FILE)

go-run:
	go run $(GO_MODULE)
