BINARY_NAME=Go_Project


ifeq (${OS}, Windows_NT)
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
ifeq (${OS}, Windows_NT)
	go env -w GOARCH=amd64 GOOS=windows
	go build -o ${BINARY_NAME}-windows.exe main.go
else
ifeq ($(shell uname -s),Linux)
	go env -w GOARCH=amd64 GOOS=linux
	go build -o ${BINARY_NAME}-linux main.go
endif
ifeq ($(shell uname -s),Darwin)
	go env -w GOARCH=amd64 GOOS=darwin 
	go build -o ${BINARY_NAME}-darwin main.go
endif
endif
	@echo --------------${\n}
	@echo "Building Done"

run:
ifeq (${OS}, Windows_NT)
	./${BINARY_NAME}-windows.exe
else
ifeq ($(shell uname -s),Linux)
	./${BINARY_NAME}-linux
endif
ifeq ($(shell uname -s),Darwin)
	./${BINARY_NAME}-darwin
endif
endif

build_run: build run

clean:
	go clean
ifeq (${OS}, Windows_NT)
	${RM} ${BINARY_NAME}-windows.exe
else
ifeq ($(shell uname -s),Linux)
	${RM} ${BINARY_NAME}-linux
endif
ifeq ($(shell uname -s),Darwin)
	${RM} ${BINARY_NAME}-darwin
endif
endif
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