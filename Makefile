BINARY_NAME=Go_Project

build:
	@echo Building in Progress
	@echo --------------${\n}
	go env -w GOARCH=amd64 GOOS=darwin
	go build -o ${BINARY_NAME}-darwin main.go
	go env -w GOARCH=amd64 GOOS=linux
	go build -o ${BINARY_NAME}-linux main.go
	go env -w GOARCH=amd64 GOOS=windows
	go build -o ${BINARY_NAME}-windows.exe main.go
	@echo --------------${\n}
	@echo "Building Done"

run:
	./${BINARY_NAME}-windows

build_and_run: build run

clean:
	go clean
	del ${BINARY_NAME}-darwin
	del ${BINARY_NAME}-linux
	del ${BINARY_NAME}-windows.exe
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
	golangci-lint run --enable-all