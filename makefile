.PHONY: test

tidy:
	@go mod tidy > /dev/null 2>&1
test: tidy
	@go test -v ./...

