---
title: Security
---

# Security


## Reporting Issues

**If you believe that you've found a Syncthing-related security vulnerability,
please report it by emailing security@syncthing.net.** Do not report it in the
open issue tracker. The [security team PGP key
(`B683AD7B76CAB013`)](/security-key.txt) can be used to send encrypted mail or
to verify responses received from that address.


## Release Signatures

The [release PGP key (`D26E6ED000654A3E`)](/release-key.txt) can be used to
verify the signatures on the official binary releases.


### Verifying a Release Signature

Download the release (tar.gz file) and the checksum sha256sum.txt.asc file.

Example verifying release v1.23.6:

```
$ curl -sLO https://github.com/syncthing/syncthing/releases/download/v1.23.6/syncthing-linux-amd64-v1.23.6.tar.gz
$ curl -sLO https://github.com/syncthing/syncthing/releases/download/v1.23.6/sha256sum.txt.asc
```

Verify that the SHA256 checksum is correct for the release. Errors will be
printed for the release files you did not download - these can be ignored.
The important line is the one indicating the checksum is "OK" for the
downloaded release file.

```
$ sha256sum -c sha256sum.txt.asc
...
sha256sum: syncthing-linux-386-v1.23.6.tar.gz: No such file or directory
syncthing-linux-386-v1.23.6.tar.gz: FAILED open or read
syncthing-linux-amd64-v1.23.6.tar.gz: OK  <-- this one
sha256sum: syncthing-linux-armv5-v1.23.6.tar.gz: No such file or directory
syncthing-linux-armv5-v1.23.6.tar.gz: FAILED open or read
...
sha256sum: WARNING: 14 lines are improperly formatted
sha256sum: WARNING: 35 listed files could not be read
```

Import the old and new release keys (only necessary if you haven't done this previously).

```
$ curl -s https://syncthing.net/release-key.txt | gpg --import
```

Verify the signature on the checksum file. The import line is the "good signature" one.

```
$ gpg --verify sha256sum.txt.asc
gpg: Signature made Mo 03 Jul 2023 10:09:30 UTC
gpg:                using RSA key D26E6ED000654A3E
gpg: Good signature from "Syncthing Release Management <release@syncthing.net>"
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
```

