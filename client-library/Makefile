SHELL:=/bin/bash #Default shell for the executing the shell commands
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
#GO basic commands
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
#Params for make targets
TIMEOUT=60#Timeout for the terraform test.Default is 10m
PKGS=$()#List of packages to test with relative path e.g. ./s3 ./dynamodb
PACKAGES ?= $(shell go list ./...|sed 's/[_"]//g') #List of packages to test
#.Phony targets
.PHONY: all testpkgs test clean help

default: help
help:
	@echo  '  clean      - Remove most generated files after a test run with terratest'
	@echo  '  all        - Runs unit tests on all the packages'
	@echo  '  testpkgs   - Run unit tests on individual packages. Usage:${YELLOW}make testpkgs ARGS="-v" PKGS="./s3 ./dynamodb"'

all: test clean

clean:
	@echo '${GREEN}###############Cleaning temporary terraform files/dirs################'
	@for dir in $(PACKAGES); do \
		echo " ${YELLOW} Deleting $$dir/.terraform";rm -rf $$dir/.terraform;rm -rf $$dir/*.tfstate;rm -rf $$dir/.terraform.lock.hcl;rm -rf $$dir/terraform.tfstate.backup; \
	done

test:
	@echo "${GREEN}############Running Unit Tests##############"
	GO111MODULE=on find . -name go.mod -execdir $(GOTEST) -v -timeout $(TIMEOUT)m ./... \;|tee >(go-junit-report > report.xml)
	@echo "${GREEN}################Unit Testing Completed Successfully##############"

testpkgs:
	echo "$(ARGS)"
	echo "$(PKGS)"
	GO111MODULE=on $(GOTEST) -timeout $(TIMEOUT)m $(ARGS) $(PKGS)

fmt:
	$(GOFMT) ./...