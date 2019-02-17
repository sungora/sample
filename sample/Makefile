# Инициализация
DIR := $(realpath -m .)
APP := $(shell basename $(DIR))
DIR_ROOT := $(realpath -m ..)

default: help

# Зависимости
dep:
	@if [ ! -f $(DIR)/go.mod ]; then \
		cd $(DIR); go mod init $(APP); \
	fi
	cd $(DIR); go get -u;
	cd $(DIR); go mod tidy;
	cd $(DIR); go mod vendor;
.PHONY: dep

# Сборка
build:
	cd $(DIR); go build -mod vendor -o $(DIR_ROOT)/bin/$(APP);
.PHONY: build

# Запуск
run:
	cd $(DIR); go mod vendor;
	cd $(DIR); go build -mod vendor -o $(DIR_ROOT)/bin/$(APP);
	$(DIR_ROOT)/bin/$(APP);
.PHONY: run

# Help
help:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep		             - Загрузка и обновление зависимостей"
	@echo "    build                 - Компиляция приложения"
	@echo "    run 		             - Запуск приложения"
.PHONY: help
