.PHONY: install buildall runall runprod down test lint


test:
	go test ./..

lint:
	golangci-lint run
