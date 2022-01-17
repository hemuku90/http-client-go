SHELL:=/bin/bash #Default shell for the executing the shell commands
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
API_DOCKER_COMPOSE:= "interview-accountapi/docker-compose.yml"
HTTP_CLIENT_LIB_DOCKER_COMPOSE:= "client-library/gohttp/docker/docker-compose.yml"
GO_EXAMPLE := "client-library"
health = $(shell curl localhost:8080/v1/health|jq '.status')
#.Phony targets
.PHONY: all startAPI stopAPI test coverage testAPICalls help

default: help
help:
	@echo  '  stopAPI     				- Runs unit tests on all the packages'
	@echo  '  all         				- Runs unit tests on all the packages'
	@echo  '  startAPI    				- Starts the form3 fake accounts API'
	@echo  '  coverage    				- Tests coverage for HTTP Client library'
	@echo  '  unitTest       			- Runs unit test on HTTP Client library'
	@echo  '  testAPICalls       			- Makes API Calls to Fake Accounts API'

all: stopAPI startAPI unitTest coverage

startAPI:
	@echo "${GREEN}############### Starting Form3 Fake Accounts API #####################"
	@echo
	docker-compose -f ${API_DOCKER_COMPOSE} up -d

stopAPI:
	@echo "${GREEN}############### Stopping Form3 Fake Accounts API #####################"
	@echo
	docker-compose -f ${API_DOCKER_COMPOSE} down

unitTest:
	@echo "${GREEN}###################### Running Unit Tests ############################"
	@echo
	docker-compose -f ${HTTP_CLIENT_LIB_DOCKER_COMPOSE} run --rm unit

coverage:
	@echo "${GREEN}###################### Running Test Coverage #########################"
	@echo
	docker-compose -f ${HTTP_CLIENT_LIB_DOCKER_COMPOSE} run --rm coverage


testAPICalls:
	@echo "${GREEN}###################### Running test API Calls #########################"
	@echo
	docker-compose -f ${API_DOCKER_COMPOSE} up -d
	pushd ${GO_EXAMPLE} && go run example.go && popd ${GO_EXAMPLE}

