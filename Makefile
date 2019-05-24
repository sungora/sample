# Инициализация
# DIR_ROOT := $(realpath -m ..)
# APP := $(shell basename $(DIR))
DIR := $(realpath -m .)
APP := github.com/sungora/sample
DBMIG := "migration/postgres"
DBDSN := "host=localhost port=5432 user=postgres password=postgres dbname=sample sslmode=disable"

default: help

# Зависимости
dep:
	@cd $(DIR);	
	@if [ ! -f $(DIR)/go.mod ]; then \
		go mod init $(APP); \
	fi
	go mod tidy;
	go mod vendor;
.PHONY: dep

# Сборка
com:
#	@for dir in $(PROJECT_LIST); do \
#		cd $(DIR)/cmd/$${dir}; go build -i -mod vendor -o $(DIR)/bin/$${dir}; \
#	done
	@cd $(DIR);
	go build -i -mod vendor -o $(DIR)/bin/app $(DIR)/cmd/app;
.PHONY: com

# Свагер
swag:
	@cd $(DIR);
	swag i -g cmd/app/main.go;
.PHONY: swag

# Запуск в режиме разработки
run: com
	$(DIR)/bin/app -c config.yaml -migrate=false -log-db=true;
.PHONY: run

# Запуск в режиме отладки
runs: swag com
	$(DIR)/bin/app -c config.yaml;
.PHONY: runs

# Создание шаблона миграции
mig:
	@gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} create tpl;
.PHONY: mig

# Статус миграции
mig-st:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} status;
.PHONY: mig-st

# Миграция на одну позицию вниз
mig-down:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} down;
.PHONY: mig-down

# Миграция вверх до конца
mig-up:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} up;
.PHONY: mig-up

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    h			- Вывод этой документации"
	@echo "    dep			- Загрузка и обновление файла зависимостей (go.mod go.sum)"
	@echo "    run			- Запуск в режиме разработки с конфигом config.yaml"
	@echo "    runs		- Запуск в режиме отладки с конфигом config.yaml"
	@echo "    mig			- Создание шаблона миграции"
	@echo "    mig-st		- Статус миграции"
	@echo "    mig-down		- Миграция на одну позицию вниз"
	@echo "    mig-up		- Миграция вверх до конца"
	@echo "    swag		- Формирование swagger докеументации"
.PHONY: h
help: h
.PHONY: help

