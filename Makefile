SHELL:=/bin/zsh #Default shell for the executing the shell commands
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
API_DIR:= "interview-accountapi/docker-compose.yml"
HTTP_CLIENT_LIB_DIR:= "client-library/gohttp"
GO_EXAMPLE := "client-library"
#.Phony targets
.PHONY: all startAPI stopAPI test coverage testAPICalls help

default: help
help:
	@echo  '  stopAPI     				- Runs unit tests on all the packages'
	@echo  '  all         				- Runs unit tests on all the packages'
	@echo  '  startAPI    				- Starts the form3 fake accounts API'
	@echo  '  coverage    				- Tests coverage for HTTP Client library'
	@echo  '  test       			 	- Runs unit test on HTTP Client library'
	@echo  '  testAPICalls       		- Makes API Calls to Fake Accounts API'

all: startAPI test coverage stopAPI
startAPI:
	@echo "${GREEN}############### Starting Form3 Fake Accounts API #####################"
	docker-compose -f ${API_DIR} up -d
stopAPI:
	@echo "${GREEN}############### Stopping Form3 Fake Accounts API #####################"
	docker-compose -f ${API_DIR} down
test:
	@echo "${GREEN}############# Running Unit Tests ##############"
	go clean -testcache
	pushd ${HTTP_CLIENT_LIB_DIR} && go test -v ./... && popd ${HTTP_CLIENT_LIB_DIR}
coverage:
	@echo "${GREEN}############ Running Test Coverage #################"
	go clean -testcache
	pushd ${HTTP_CLIENT_LIB_DIR}  && go test -v -cover ./... && popd ${HTTP_CLIENT_LIB_DIR}
testAPICalls:
	@echo "${GREEN}############ Running Test Coverage #################"
	pushd ${GO_EXAMPLE} && go run example.go && popd ${GO_EXAMPLE}

