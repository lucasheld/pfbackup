pfbackup
========

pfbackup backups pfSense configurations

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
```
