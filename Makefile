OUT_DIR := ./build/
OUT_NAME := pfbackup
OUT := ${OUT_DIR}${OUT_NAME}
VERSION := $(shell git describe --tags --abbrev=0 || echo build-`date +%s`)
PKG := github.com/lucasheld/pfbackup

.PHONY: help
help:
	@echo "make <target>"
	@echo " clean"
	@echo " build"
	@echo " build-all: includes build-linux, build-windows, build-osx"
	@echo " build-linux"
	@echo " build-windows"
	@echo " build-osx"

.PHONY: clean
clean:
	@rm -fr ${OUT_DIR}

.PHONY: build
build:
	@go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}

.PHONY: build-all
build-all: build-linux build-windows build-osx

.PHONY: build-linux
build-linux:
	@echo "building linux..."
	@GOOS=linux GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-linux-386 pfbackup.go
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-linux-amd64 pfbackup.go

.PHONY: build-windows
build-windows:
	@echo "building windows..."
	@GOOS=windows GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-windows-386.exe pfbackup.go
	@GOOS=windows GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-windows-amd64.exe pfbackup.go

.PHONY: build-osx
build-osx:
	@echo "building osx..."
	@GOOS=darwin GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-darwin-386 pfbackup.go
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-${VERSION}-darwin-amd64 pfbackup.go
