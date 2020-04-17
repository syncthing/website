---
resources:
    src: stars.jpg
---


{{% jumbotron bgResource="stars.jpg" %}}
Syncthing is a **continuous file synchronization** program. It synchronizes
files between two or more computers in real time, safely protected from prying
eyes. Your data is your data alone and you deserve to choose where it is stored,
whether it is shared with some third party, and how it's transmitted over the
internet.

{{% center %}}
## Get Started
Grab one of the [downloads]({{< relref downloads >}}) and start syncing! \
Check out the [getting started
guide](https://docs.syncthing.net/intro/getting-started.html) for some tips
along the way.
{{% /center %}}
{{% /jumbotron %}}


{{% row %}}
{{% col %}}
### Private & Secure

- **Private.** None of your data is ever stored anywhere else other than on your
  computers. There is no central server that might be compromised, legally or
  illegally.

- **Encrypted.** All communication is secured using TLS. The encryption used
  includes perfect forward secrecy to prevent any eavesdropper from ever gaining
  access to your data.

- **Authenticated.** Every node is identified by a strong cryptographic
  certificate. Only nodes you have explicitly allowed can connect to your
  cluster.

  *If you have a security concern, please see [the security page]({{< ref security.md >}}) for details and contact information.*
{{% /col %}}

{{% col %}}
### Open

- **Open Protocol.** The protocol is a [documented
  specification](https://docs.syncthing.net/specs/bep-v1.html#bep-v1) — no
  hidden magic.

- **Open Source.** All source code is [available on
  GitHub](https://github.com/syncthing/syncthing) — what you see is what you
  get, there is no hidden funny business.

- **Open Development.** Any bugs found are [immediately visible](https://github.com/syncthing/syncthing/issues) for anyone to
  browse — no hidden flaws.

- **Open Discourse.** Development and usage is always [open for discussion](https://forum.syncthing.net/).
{{% /col %}}
{{% /row %}}


### Easy to Use

{{% row %}}
{{% col %}}
- **Powerful.** Synchronize as many folders as you need with different people or
  just between your own devices.

- **Portable.** Configure and monitor Syncthing via a responsive and powerful
  interface accessible via your browser. Works on Mac OS X, Windows, Linux,
  FreeBSD, Solaris and OpenBSD. Run it on your desktop computers and synchronize
  them with your server for backup.

- **Simple.** Syncthing doesn't need IP addresses or advanced configuration: it
  just works, over LAN and over the Internet. Every machine is identified by an
  ID. Give your ID to your friends, share a folder and watch: UPnP will do if
  you don't want to port forward or you don't know how.
{{% /col %}}

{{% col %}}
<img src="/img/screenshot.png" class="img img-fluid" alt="Screenshot of Syncthing Web interface">
{{% /col %}}

{{% /row %}}


{{% sponsor href="https://kastelo.net/" alt="Kastelo logo" id="kastelo-logo" %}}

Kastelo provides [commercial support](https://kastelo.net/stes/) for Syncthing
and sponsors Syncthing with development resources.

{{% /sponsor %}}