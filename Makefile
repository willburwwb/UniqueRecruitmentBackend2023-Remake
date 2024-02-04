EXECUTABLE = main

.DEFAULT_GOAL := build

dep:
ifneq ($(GO_MODULE_STATE), on)
	go env -w GO111MODULE="on"
endif
ifeq ($(GO_PROXY), https://proxy.golang.org,direct)
	go env -w GOPROXY="https://goproxy.cn,direct"
endif
	go mod tidy

build: dep
	# swag init --parseDependency --parseInternal # gen swagger docs
	CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o ${EXECUTABLE}

server : build
	./${EXECUTABLE} server

proto:
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative --proto_path=. proto/open_platform/open.proto proto/sso/sso.proto

clean:
	rm -rf ${EXECUTABLE}