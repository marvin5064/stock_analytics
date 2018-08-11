# Makefile
APPNAME =`basename ${PWD}`

.PHONY: build run test protos
build:
	@go build
run: build
	@./$(APPNAME)
test:
	@go test
dep:
	@dep ensure
	@dep ensure -update
protos:
	@protoc --proto_path=protobuf/schema \
			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:protobuf \
			--grpc-gateway_out=logtostderr=true:protobuf \
			protobuf/schema/stock/*.proto