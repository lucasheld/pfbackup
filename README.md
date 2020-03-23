pfbackup
========

pfbackup backups pfSense configurations

[![Build Status](https://travis-ci.org/lucasheld/pfbackup.svg?branch=master)](https://travis-ci.org/lucasheld/pfbackup)
[![Docker Build](https://img.shields.io/docker/build/lucasheld/pfbackup.svg)](https://hub.docker.com/r/lucasheld/pfbackup)
[![Docker Pulls](https://img.shields.io/docker/pulls/lucasheld/pfbackup.svg)](https://hub.docker.com/r/lucasheld/pfbackup)

## Compilation
```console
$ git clone https://github.com/lucasheld/pfbackup.git
$ cd pfbackup
$ go build
```

## Usage
```console
$ ./pfbackup --help
pfbackup backups pfSense configurations

Usage:
  pfbackup [flags]

Flags:
  -h, --help          help for pfbackup
      --no-verify     do not verify ssl certificate
      --pass string   pfSense password (required)
      --path string   path to output directory (default ".")
      --url string    pfSense url (required)
      --user string   pfSense username (required)
  -v, --version       Print the version number
```
