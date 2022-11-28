.PHONY: run runprod down test lint


run:
	swag init -o ./app/apiserver/docs -g ./app/apiserver/main.go
	go run ./app/apiserver

test:
	go test ./...

lint:
	golangci-lint run
