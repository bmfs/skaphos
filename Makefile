all: test lint

tidy:
	go mod tidy -v

build:
	go build -o bin/ ./...

test: build
	go test -cover -race ./...

test-coverage:
	go test ./... -race -coverprofile=coverage.txt && go tool cover -html=coverage.txt
