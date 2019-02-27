# Инициализация
# DIR_ROOT := $(realpath -m ..)
DIR := $(realpath -m .)
APP := $(shell basename $(DIR))
RUN = sample
PROJECT_LIST = sample pkg

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
	@for dir in $(PROJECT_LIST); do \
		cd $(DIR)/cmd/$${dir}; go build -i -mod vendor -o $(DIR)/bin/$${dir}; \
	done
.PHONY: build

# Запуск
run:
	cd $(DIR); go mod tidy;
	cd $(DIR); go mod vendor;
	cd $(DIR)/cmd/$(RUN); go build -i -mod vendor -o $(DIR)/bin/$(RUN);
	$(DIR)/bin/$(RUN);
.PHONY: run

# Help
help:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep		- Загрузка и обновление зависимостей"
	@echo "    build	- Компиляция приложения"
	@echo "    run		- Запуск приложения"
.PHONY: help
