lint:
	@echo "=> Executing go lint"
	@golangci-lint run
test:
	@echo "=> Running tests"
	@go test ./... 
coverage:
	@echo "=> Running Coverage-Report"
	@echo "=> Creating coverage report\n"
	@go test -cover -coverprofile="coverage.out" ./...
	@go tool cover -html=./coverage.out
