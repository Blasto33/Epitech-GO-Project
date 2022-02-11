BINARY_NAME=Go_Project
UNAME_S := $(shell uname -s)

ifdef OS
   RM = del /Q
   FixPath = $(subst /,\,$1)
else
   ifeq ($(shell uname), Linux)
      RM = rm -f
      FixPath = $1
   endif
endif

build:
	@echo Building in Progress
	@echo --------------${\n}
	ifeq ($(OS),Windows_NT)
		go env -w GOARCH=amd64 GOOS=windows
		go build -o ${BINARY_NAME}-windows.exe main.go
	endif
	ifeq ($(UNAME_S),Linux)
		go env -w GOARCH=amd64 GOOS=linux
		go build -o ${BINARY_NAME}-linux main.go
	endif
	ifeq ($(UNAME_S),Darwin)
		go env -w GOARCH=amd64 GOOS=darwin 
		go build -o ${BINARY_NAME}-darwin main.go
	endif
	@echo --------------${\n}
	@echo "Building Done"

run:
	./${BINARY_NAME}-linux

build_run: build run

clean:
	go clean
	${RM} ${BINARY_NAME}-darwin
	${RM} ${BINARY_NAME}-linux
	${RM} ${BINARY_NAME}-windows.exe
	@echo "Cleaning ${BINARY_NAME}"

test:
	@echo Making Test
	@echo --------------${\n}
	go test ./...
	@echo --------------${\n}
	@echo Test Done

test_coverage:
	@echo Making Test Coverage
	@echo --------------${\n}
	go test ./... -coverprofile=coverage.out
	@echo --------------${\n}
	@echo Test Coverage Done

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run

re:		clean build