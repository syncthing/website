---
title: Downloads
---

# Downloads

## Integrations

These are some popular and user friendly OS integrations, providing things like system tray icons, file browser integration, etc. These are good starting points if you are a new user unfamiliar with Syncthing, or not prone to loving the command line.

- **[SyncTrayzor](https://github.com/canton7/SyncTrayzor/releases/latest)**:
  Windows tray utility, filesystem watcher & launcher
- **[syncthing-macos](https://github.com/syncthing/syncthing-macos/releases/latest)**:
  macOS application bundle
- **[Syncthing-GTK](https://github.com/kozec/syncthing-gtk/releases/latest)**:
  cross-platform GUI wrapper

There's a wealth of further integrations of all kinds listed on the [community
contributions](https://docs.syncthing.net/users/contrib.html) page. Each
integration has their own issue tracker for integration-specific issues, but
discussion and assistance for all of them is welcome on the
[forum](https://forum.syncthing.net/).


## Android

The Android app is available on [Google Play](https://play.google.com/store/apps/details?id=com.nutomic.syncthingandroid) and [F-Droid](https://f-droid.org/packages/com.nutomic.syncthingandroid/).


## Base Syncthing

This is the basic Syncthing distribution, providing a command line / daemon like
executable and a web based user interface.

{{< release >}}

If you are unsure what to download and you're running on a normal computer,
please use the "64-bit (x86-64)" build for your operating system. If you're
running on an oddball system such as a NAS, please consult your vendor.


## Debian / Ubuntu Packages

You can choose between the "stable" (latest release) or "candidate" (earlier
release candidate) tracks. The stable channel is updated usually every first
Tuesday of the month.

```
# Add the "stable" channel to your APT sources:
echo "deb https://apt.syncthing.net/ syncthing stable" | sudo tee /etc/apt/sources.list.d/syncthing.list
```

The candidate channel is updated with release candidate builds, usually every
second Tuesday of the month. These predate the corresponding stable builds by
about three weeks.


```
# Add the "candidate" channel to your APT sources:
echo "deb https://apt.syncthing.net/ syncthing candidate" | sudo tee /etc/apt/sources.list.d/syncthing.list
```

Then proceed with the following steps to finish setting up the chosen track and
install Syncthing.

```
# Add the release PGP keys:
curl -s https://syncthing.net/release-key.txt | sudo apt-key add -

# Increase preference of Syncthing's packages ("pinning")
printf "Package: *\nPin: origin apt.syncthing.net\nPin-Priority: 990\n" | sudo tee /etc/apt/preferences.d/syncthing

# Update and install syncthing:
sudo apt-get update
sudo apt-get install syncthing
```

Depending on your distribution, you may see an error similar to the following
when running apt-get:

```
E: The method driver /usr/lib/apt/methods/https could not be found.
N: Is the package apt-transport-https installed?
E: Failed to fetch https://apt.syncthing.net/dists/syncthing/InRelease
```

If so, please install the apt-transport-https package and try again:

```
sudo apt-get install apt-transport-https
```

If you insist, you can also use the above URLs with http instead of https.

{{% sponsors %}}
