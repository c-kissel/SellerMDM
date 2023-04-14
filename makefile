GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=cmd/app/mdm.out

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: build
 
build:
	go build -o ${BINARY_NAME} cmd/app/main.go

api:
	oapi-codegen -package specs -config specs/config.yml specs/contract.openapi.yaml > specs/contract.gen.go
	
lint: 
	golangci-lint run

test:
	go test -v cmd/app/main.go
 
run:
	go build -o ${BINARY_NAME} cmd/app/main.go
	./${BINARY_NAME}

debug:
	go build -o ${BINARY_NAME} cmd/app/main.go
	./${BINARY_NAME} -c cmd/app/debug.yaml
 
clean:
	go clean
	rm ${BINARY_NAME}