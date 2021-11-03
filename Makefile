OUT_DIR := ./build
OUT_BIN_NAME := server

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	@go build -o $(OUT_DIR)/$(OUT_BIN_NAME) .

.PHONY: clean
clean:
	@rm -r $(OUT_DIR)
