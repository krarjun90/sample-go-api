.PHONY: all
all: setup build migrate test

setup:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	go get -u github.com/krarjun90/go-bindata/...

build:
	dep ensure
	go build -o sample_go_api

start:
	./sample_go_api start

migrate:
	./sample_go_api migrate

rollback:
	./sample_go_api rollback

test:
	go test ./...