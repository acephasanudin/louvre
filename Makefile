doc:
	swag fmt && swag init --parseDependency --parseInternal

dev:
	go run main.go start

test:
	@echo "\x1b[32;1m>>> running unit test and calculate coverage\x1b[0m"
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./... | grep -v mocks | tr '\n' ',')
	@go tool cover -html=coverage.txt

PROTO_DIR=proto
ADAPTER_DIR=app/adapter

buf-gen:
	@echo "Formatting code with buf..."
	@cd $(PROTO_DIR) && buf format -w .
	@echo "Generating Protobuf files..."
	@cd $(PROTO_DIR) && buf generate .

wire-gen:
	@echo "Generating Wire For Adapter..."
	@cd $(ADAPTER_DIR) && wire