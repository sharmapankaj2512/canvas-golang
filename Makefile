.PHONY : test test-unit test-regression

test: test-unit test-regression

test-unit:
	@go test ./... --tags=unit

test-regression:
	@go test ./... --tags=regression

test-progress:
	@go test ./... --tags=progress