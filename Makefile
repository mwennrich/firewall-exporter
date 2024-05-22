GO111MODULE := on
DOCKER_TAG := $(or ${GIT_TAG_NAME}, latest)

.PHONY: all
all: test firewall-exporter

.PHONY: firewall-exporter
firewall-exporter:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/firewall-exporter *.go
	strip bin/firewall-exporter

.PHONY: test
test:
	go test -cover ./...

.PHONY: clean
clean:
	rm -f bin/*
