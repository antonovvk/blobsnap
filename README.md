BlobSnap
========

BlobSnap is a snapshot-based backup system built on top of [BlobStash](https://github.com/tsileo/blobstash), designed to provide "time machine" like features.

## Features

- Content addressed (with [BLAKE2b](https://blake2.net) as hashing algorithm), files are split into blobs, and retrieved by hash
- Incremental backups/snapshots thanks to data deduplication
- A special archive mode, for one-time backup/non-snapshotted backup, but still with dedup
- Server handles uploading/downloading blobs to/from different storage
- Client only query the server and send blobs to it (the client take care of chunking/building blobs)
- Read-only FUSE file system to navigate backups/snapshots
- Optional encryption (using [go.crypto/nacl secretbox](http://godoc.org/code.google.com/p/go.crypto/nacl))
- Take snapshot automatically every x minutes, using a separate client-side scheduler (provides Arq/time machine like backup)
- Possibility to incrementally archive blobs to AWS Glacier (with a recovery command-line tool)
- Strong test suite (unit tests + integration tests)
- Support for backing-up multiple hosts (you can force a different host to split backups into "different buckets")
- Backend routing, you can define rules to specify where blobs should be stored ("if-meta", "if-host-myhost"...) and setup custom context

Draws inspiration from [Camlistore](camlistore.org) and [bup](https://github.com/bup/bup) (files are split into multiple blobs using a rolling checksum).

## Components

### Fuse file system

**BlobFS** is the most convenient way to restore/navigate snapshots is the FUSE file system.

There is three magic directories at the root:

- **latest**: it contains the latest version of every snapshots/backups.
- **snapshots**: it let you navigate for every snapshots, you can see every versions.
- **at**: let access directories/files at a given time, it automatically retrieve the closest previous snapshots.

```console
$ blobstash mount /backups
2014/05/12 17:26:34 Mounting read-only filesystem on /backups
Ctrl+C to unmount.
```

```console
$ ls /backups
tomt0m
$ ls /backups/tomt0m
at  latest  snapshots
$ ls /backups/tomt0m/latest
writing
$ ls /backups/tomt0m/snapshots/writing
2014-05-11T11:01:07+02:00  2014-05-11T18:36:06+02:00  2014-05-12T17:25:47+02:00
$ ls /backups/tomt0m/at/writing/2014-05-12
file1  file2  file3
```
### Command-line client

**blobsnap** is the command-line client to perform/restore snapshots/backups.

```console
$ blobsnap put /path/to/dir/or/file
```

### Backup scheduler

The backup scheduler allows you to perform snapshots...

## Roadmap / Ideas

- an Android app to backup Android devices
- Follow .gitignore file
- Garbage collection (sparse files support for blob files)
- A web interface
- An S3-like HTTP API to store archive
- Fill an issue!

## Donate!

[![Flattr this git repo](http://api.flattr.com/button/flattr-badge-large.png)](https://flattr.com/submit/auto?user_id=tsileo&url=https%3A%2F%2Fgithub.com%2Ftsileo%2Fblobstash)

BTC 1HpHxwNUmXfrU9MR9WTj8Mpg1YUEry9MF4
