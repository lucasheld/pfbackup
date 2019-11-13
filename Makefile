OUT_DIR := ./build/
OUT_NAME := pfbackup
OUT := ${OUT_DIR}${OUT_NAME}
VERSION := $(shell git describe --tags --abbrev=0)
PKG := github.com/lucasheld/pfbackup

help:
	@echo "make <target>"
	@echo " clean"
	@echo " build"
	@echo " build-all: includes build-linux, build-windows, build-osx"
	@echo " build-linux"
	@echo " build-windows"
	@echo " build-osx"

clean:
	@rm -fr ${OUT_DIR}

build:
	@go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}

build-all: build-linux build-windows build-osx

run-build = $(MAKE) FC=$(1) FFLAGS=$(2) PETSC_FFLAGS=$(3) TARGET=$@ LEXT="$(1)_$(UNAME)" -e syst

build-linux:
	@echo "building linux..."
	@GOOS=linux GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-linux-386 pfbackup.go
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-linux-amd64 pfbackup.go

build-windows:
	@echo "building windows..."
	@GOOS=windows GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-windows-386.exe pfbackup.go
	@GOOS=windows GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-windows-amd64.exe pfbackup.go

build-osx:
	@echo "building osx..."
	@GOOS=darwin GOARCH=386 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-darwin-386 pfbackup.go
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-X ${PKG}/version.Version=${VERSION}" -o ${OUT}-darwin-amd64 pfbackup.go
