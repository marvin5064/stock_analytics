# Makefile
APPNAME =`basename ${PWD}`

.PHONY: build run test
build:
	@go build
run: build
	@./$(APPNAME)
test:
	@go test