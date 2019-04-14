# Инициализация
# DIR_ROOT := $(realpath -m ..)
# APP := $(shell basename $(DIR))
DIR := $(realpath -m .)
APP := github.com/sungora/sample

default: help

# Зависимости
dep:
	@if [ ! -f $(DIR)/go.mod ]; then \
		cd $(DIR); go mod init $(APP); \
	fi
	cd $(DIR); go mod tidy;
	cd $(DIR); go mod vendor;
.PHONY: dep

# Сборка
build:
#	@for dir in $(PROJECT_LIST); do \
#		cd $(DIR)/cmd/$${dir}; go build -i -mod vendor -o $(DIR)/bin/$${dir}; \
#	done
	cd $(DIR); go build -i -mod vendor -o $(DIR)/main;
.PHONY: build

# Запуск
run:
	cd $(DIR); go build -i -mod vendor -o $(DIR)/main;
	$(DIR)/main;
.PHONY: run

# Help
help:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep		- Загрузка и обновление зависимостей"
	@echo "    build	- Компиляция приложения"
	@echo "    run		- Компиляция и запуск приложения"
.PHONY: help
