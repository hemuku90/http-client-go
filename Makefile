SHELL:=/bin/zsh #Default shell for the executing the shell commands
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
API_DIR:= "interview-accountapi/docker-compose.yml"
HTTP_CLIENT_LIB_DIR:= "client-library"
#.Phony targets
.PHONY: all startAPI stopAPI test help

default: help
help:
	@echo  '  stopAPI     - Runs unit tests on all the packages'
	@echo  '  all        - Runs unit tests on all the packages'
	@echo  '  startAPI   - Starts the form3 fake accounts API'

all: startAPI test
startAPI:
	@echo "${GREEN}############### Starting Form3 Fake Accounts API #####################"
	docker-compose -f ${API_DIR} up -d
stopAPI:
	@echo "${GREEN}############### Stopping Form3 Fake Accounts API #####################"
	docker-compose -f ${API_DIR} down
test:
	@echo "${GREEN}############Running Unit Tests##############"
	cd client-library && go test -v ./...
