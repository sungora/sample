# Инициализация
APP := "sample"
PROJECT_LIST = sample
DIR := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

default: help

# Зависимости
dep:
	@if [ ! -f $(DIR)/src/go.mod ]; then \
		cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod init $(APP); \
	fi
	cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod tidy;
	cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod vendor;
.PHONY: dep

dep-full:
	@if [ ! -f $(DIR)/src/go.mod ]; then \
		cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod init $(APP); \
	fi
	cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go get -u;
	cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod tidy;
	cd $(DIR)/src; GO111MODULE="on" GOPATH="$(DIR)" go mod vendor;
.PHONY: dep-full

# Сборка
build:
	@for dir in $(PROJECT_LIST); do \
		GOPATH="$(DIR)" go build -o $(DIR)/bin/$${dir} $${dir}; \
	done
.PHONY: build

# Help
help:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep                  - Загрузка зависимостей проекта"
	@echo "    dep-full             - Загрузка и одновление зависимостей проекта"
	@echo "    build                - Компиляция приложения"
.PHONY: help
