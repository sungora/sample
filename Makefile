DIR := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

# Приложение
APP := "sample"

# Инициализация модуля
dep-init:
	@if [ ! -f $(DIR)/src//$(APP)/go.mod ]; then \
		cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go mod init $(APP); \
	fi
.PHONY: dep-init

# Зависимости
dep: dep-init
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go get;
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go mod tidy;
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go mod vendor;
.PHONY: dep

depup: dep-init
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go get -u;
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go mod tidy;
	cd $(DIR)/src/$(APP); GO111MODULE="on" GOPATH="$(DIR)" go mod vendor;
.PHONY: depup

# Сборка
build: 
	cd $(DIR)/src/$(APP); GO111MODULE="on" go build -o $(DIR)/bin/$(APP) -mod vendor;
.PHONY: build

# Test
test:
	@echo $(DIR);
.PHONY: test


