.PHONY: build test

build:
	GOWORK=off go build -mod=readonly -o ksyun-slb.exe .

test:
	GOWORK=off go test -mod=readonly ./...
