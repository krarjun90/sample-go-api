.PHONY: all
all: setup build migrate test

setup:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	go get -u github.com/krarjun90/go-bindata/...
	go get -u github.com/vektra/mockery/.../

build-deps:
	dep ensure
	go generate ./...

build: build-deps
	go build -o sample_go_api

start: migrate
	./sample_go_api start

migrate: build-deps
	./sample_go_api migrate

rollback:
	./sample_go_api rollback

test: build-deps
	go test ./...