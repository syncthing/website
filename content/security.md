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

Download the release (tar.gz file) and the checksum sha1sum.txt.asc file.

Example verifying release v0.14.11:

```
$ curl -sLO https://github.com/syncthing/syncthing/releases/download/v0.14.11/syncthing-linux-amd64-v0.14.11.tar.gz
$ curl -sLO https://github.com/syncthing/syncthing/releases/download/v0.14.11/sha1sum.txt.asc
```

Verify that the SHA1 checksum is correct for the release. Errors will be
printed for the release files you did not download - these can be ignored.
The important line is the one indicating the checksum is "OK" for the
downloaded release file.

```
$ sha1sum -c sha1sum.txt.asc
...
sha1sum: syncthing-linux-386-v0.14.11.tar.gz: No such file or directory
syncthing-linux-386-v0.14.11.tar.gz: FAILED open or read
syncthing-linux-amd64-v0.14.11.tar.gz: OK  <-- this one
sha1sum: syncthing-linux-armv5-v0.14.11.tar.gz: No such file or directory
syncthing-linux-armv5-v0.14.11.tar.gz: FAILED open or read
...
sha1sum: WARNING: 20 lines are improperly formatted
sha1sum: WARNING: 12 listed files could not be read
```

Import the old and new release keys (only necessary if you haven't done this previously).

```
$ curl -s https://syncthing.net/release-key.txt | gpg --import
```

Verify the signature on the checksum file. The import line is the "good signature" one.

```
$ gpg --verify sha1sum.txt.asc
gpg: Signature made Tue Nov 15 07:44:49 2016 CET
gpg:                using RSA key D26E6ED000654A3E
gpg: Good signature from "Syncthing Release Management <release@syncthing.net>"
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
```

