version := 0.2.0
githash := $(shell git rev-parse --short HEAD)
buildtime := $(shell date -u '+%Y-%m-%d_%I:%M:%S%pm_%Z')
ldflags := "\
	-X github.com/spilliams/foo/internal/version.versionNumber=$(version)\
	-X github.com/spilliams/foo/internal/version.gitHash=$(githash)\
	-X github.com/spilliams/foo/internal/version.buildTime=$(buildtime)"

.PHONY: build
build:
	go build -ldflags $(ldflags) -o bin/foo main.go

.PHONY: install
install:
	go build -ldflags "${ldflags}" -o $$GOPATH/bin/foo main.go

.PHONY: test
test:
	go test ./...
